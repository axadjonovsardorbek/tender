package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/axadjonovsardorbek/tender/pkg/models"
	token "github.com/axadjonovsardorbek/tender/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type AuthI interface {
	Register(context.Context, *models.RegisterReq) (*models.TokenRes, error)
	Login(context.Context, *models.LoginReq) (*models.TokenRes, error)
	GetProfile(context.Context, string) (*models.UserRes, error)
	UpdateProfile(context.Context, *models.UpdateReq) (*models.Void, error)
	DeleteProfile(context.Context, string) (*models.Void, error)
}

type AuthRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAuthRepo(db *sql.DB, rdb *redis.Client) *AuthRepo {
	return &AuthRepo{db: db, rdb: rdb}
}

func (r *AuthRepo) Register(ctx context.Context, req *models.RegisterReq) (*models.TokenRes, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO users(
		id,
		username,
		email,
		role,
		password
	) VALUES($1, $2, $3, $4, $5)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(query, id, req.Username, req.Email, req.Role, hashedPassword)

	if err != nil {
		fmt.Println("error while registering")
		return nil, err
	}

	tkn := token.GenerateJWTToken(id, req.Role)

	return &models.TokenRes{
		AccessToken:  tkn.AccessToken,
		RefreshToken: tkn.RefreshToken,
		Role:         req.Role,
		Id:           id,
	}, nil
}
func (r *AuthRepo) Login(ctx context.Context, req *models.LoginReq) (*models.TokenRes, error) {
	var id string
	var role string
	var password string

	query := `
	SELECT 
		id,
		role,
		password
	FROM
		users
	WHERE
		username = $1
	AND
		deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Username)

	err := row.Scan(
		&id,
		&role,
		&password,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		log.Println("Error while login: ", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	tkn := token.GenerateJWTToken(id, role)

	return &models.TokenRes{
		AccessToken:  tkn.AccessToken,
		RefreshToken: tkn.RefreshToken,
		Role:         role,
		Id:           id,
	}, nil
}

func (r *AuthRepo) GetProfile(ctx context.Context, id string) (*models.UserRes, error) {
	userData, err := r.rdb.Get(context.Background(), id).Result()
	if err == redis.Nil {
		user := models.UserRes{}
		query := `
	SELECT
		id,
		username,
		email,
		to_char(created_at, 'YYYY-MM-DD HH24:MI'),
		role
	FROM
		users
	WHERE
		id = $1
	AND
		deleted_at = 0
	`

		row := r.db.QueryRow(query, id)
		err := row.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.Role,
		)

		if err != nil {
			return nil, err
		}

		fmt.Println("Successfully got profile")

		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Printf("JSON marshalling error: %v", err)
			return nil, err
		}
		err = r.rdb.Set(context.Background(), id, jsonData, 5*time.Minute).Err()
		if err != nil {
			log.Printf("Redis set error: %v", err)
			return nil, err
		}

		return &user, nil
	} else if err != nil {
		log.Printf("Redis get error: %v", err)
		return nil, err
	}

	user := models.UserRes{}

	err = json.Unmarshal([]byte(userData), user)
	if err != nil {
		log.Printf("JSON unmarshalling error: %v", err)
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepo) UpdateProfile(ctx context.Context, req *models.UpdateReq) (*models.Void, error) {
	query := `
	UPDATE 
		users
	SET
	`

	var conditions []string
	var args []interface{}

	if req.Username != "" && req.Username != "string" {
		conditions = append(conditions, " username = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Username)
	}
	if req.Email != "" && req.Email != "string" {
		conditions = append(conditions, " email = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Email)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + "deleted_at = 0"

	args = append(args, req.Id)

	res, err := r.db.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s couldn't update", req.Id)
	}

	log.Println("Successfully updated user")

	return nil, nil
}

func (r *AuthRepo) DeleteProfile(ctx context.Context, id string) (*models.Void, error) {
	query := `
	UPDATE 
		users
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, id)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s not found", id)
	}

	return nil, nil
}

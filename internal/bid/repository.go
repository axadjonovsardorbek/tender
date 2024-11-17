package bid

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/axadjonovsardorbek/tender/pkg/models"
	"github.com/axadjonovsardorbek/tender/platform/websocket"
	"github.com/google/uuid"
)

type BidI interface {
	Create(context.Context, *models.CreateBidReq) (*models.Void, error)
	GetById(context.Context, string) (*models.BidRes, error)
	GetAll(context.Context, *models.GetAllBidReq) (*models.GetAllBidRes, error)
	Update(context.Context, *models.UpdateBidReq) (*models.Void, error)
	Delete(context.Context, *models.DeleteBidReq) (*models.Void, error)
}

type BidRepo struct {
	db *sql.DB
}

func NewBidRepo(db *sql.DB) *BidRepo {
	return &BidRepo{db: db}
}

func (r *BidRepo) Create(ctx context.Context, req *models.CreateBidReq) (*models.Void, error) {
	id := uuid.New().String()

	tr, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO bids(
		id,
		tender_id,
		contractor_id,
		price,
		delivery_time,
		comments
	) VALUES($1, $2, $3, $4, $5, $6)
	`

	_, err = tr.Exec(query, id, req.TenderId, req.ContractorId, req.Price, req.DeliveryTime, req.Comments)
	if err != nil {
		tr.Rollback()
		fmt.Println("error while creating bid")
		return nil, err
	}

	var user_id string
	query = `SELECT user_id FROM tenders WHERE id = $1`
	err = tr.QueryRow(query, req.TenderId).Scan(&user_id)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	message := fmt.Sprintf("price: %d, delivery_time: %d, comments: %s", req.Price, req.DeliveryTime, req.Comments)

	query = `INSERT INTO notifications(
				user_id,
				message,
				relation_id,
				type
			) VALUES($1, $2, $3, $4)`
	_, err = tr.Exec(query, user_id, message, id, "bid")
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	if err = tr.Commit(); err != nil {
		return nil, err
	}

	// Send the notification to WebSocket clients
	notification := map[string]interface{}{
		"user_id": user_id,
		"message": message,
		"type":    "bid",
	}
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}

	go func() {
		message := []byte(notificationJSON)
		websocket.BroadcastMessage(message)
	}()

	return &models.Void{}, nil
}

func (r *BidRepo) GetById(ctx context.Context, id string) (*models.BidRes, error) {
	bid := models.BidRes{}
	query := `
	SELECT
		id,
		tender_id,
		contractor_id,
		price,
		delivery_time,
		comments,
		status,
		to_char(created_at, 'YYYY-MM-DD HH24:MI')
	FROM
		bids
	WHERE
		id = $1
	AND
		deleted_at = 0
	`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&bid.Id,
		&bid.TenderId,
		&bid.ContractorId,
		&bid.Price,
		&bid.DeliveryTime,
		&bid.Comments,
		&bid.Status,
		&bid.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully got bid")

	return &bid, nil
}
func (r *BidRepo) GetAll(ctx context.Context, req *models.GetAllBidReq) (*models.GetAllBidRes, error) {
	bids := models.GetAllBidRes{}
	query := `
	SELECT
		COUNT(id) OVER () AS total_count,
		id,
		tender_id,
		contractor_id,
		price,
		delivery_time,
		comments,
		status,
		to_char(created_at, 'YYYY-MM-DD HH24:MI')
	FROM
		bids
	WHERE
		id = $1
	AND
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.TenderId != "" && req.TenderId != "string" {
		conditions = append(conditions, " tender_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.TenderId)
	}
	if req.ContractorId != "" && req.ContractorId != "string" {
		conditions = append(conditions, "contractor_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ContractorId)
	}
	if req.Price > 0 {
		conditions = append(conditions, "price <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Price)
	}
	if req.DeliveryTime > 0 {
		conditions = append(conditions, "delivery_time <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.DeliveryTime)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf("ORDER BY %s DESC", req.SortType)
	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err == sql.ErrNoRows {
		return &bids, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		bid := models.BidRes{}
		var count int64
		err := rows.Scan(
			&count,
			&bid.Id,
			&bid.TenderId,
			&bid.ContractorId,
			&bid.Price,
			&bid.DeliveryTime,
			&bid.Comments,
			&bid.Status,
			&bid.CreatedAt,
		)
		if err != nil {
			log.Println("error while scanning all bids: ", err)
			return nil, err
		}

		bids.Bids = append(bids.Bids, &bid)
		bids.TotalCount = count
	}

	fmt.Println("Successfully got bids")

	return &bids, nil
}
func (r *BidRepo) Update(ctx context.Context, req *models.UpdateBidReq) (*models.Void, error) {

	// get tender_id by bids id
	var tender_id string
	tenderIdQuery := `
		SELECT 
			b.tender_id
		FROM
			bids b
		JOIN
			tenders t
		ON
			t.id = b.tender_id
		AND
			t.deleted_at = 0
		AND
			t.status <> 'awarded'
		WHERE
			b.id = $1
		AND
			b.deleted_at = 0
		AND
			b.status = 'pending'
		`

	row := r.db.QueryRow(tenderIdQuery, req.Id)
	err := row.Scan(
		&tender_id,
	)

	if err != nil {
		return nil, err
	}

	if req.Status == "rejected" {
		query := `
		UPDATE 
			bids
		SET 
			status = $1
		WHERE 
			id = $2
		AND 
			deleted_at = 0
		AND
			status = 'pending'
		AND
			tender_id = $3
		`

		res, err := r.db.Exec(query, req.Status, req.Id, tender_id)

		if err != nil {
			return nil, err
		}

		if r, err := res.RowsAffected(); r == 0 {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("bid with id %s couldn't reject", req.Id)
		}

		return nil, nil
	} else if req.Status == "accepted" {
		tz, err := r.db.Begin()
		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				tz.Rollback()
				err = fmt.Errorf("panic occurred: %v", p)
			} else if err != nil {
				tz.Rollback()
			} else {
				err = tz.Commit()
			}
		}()

		// accept bid
		acceptQuery := `
		UPDATE 
			bids
		SET 
			status = $1
		WHERE 
			id = $2
		AND 
			deleted_at = 0
		AND
			status = 'pending'
		AND
			tender_id = $3
		`

		res, err := tz.Exec(acceptQuery, req.Status, req.Id, tender_id)

		if err != nil {
			return nil, err
		}

		if r, err := res.RowsAffected(); r == 0 {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("bid with id %s couldn't accept", req.Id)
		}

		// reject tender all bids besides accepted bid
		rejectQuery := `
		UPDATE
			bids
		SET 
			status = 'rejected'
		WHERE 
			deleted_at = 0
		AND
			status = 'pending'
		AND
			tender_id = $1
		`

		res, err = tz.Exec(rejectQuery, tender_id)

		if err != nil {
			return nil, err
		}

		if r, err := res.RowsAffected(); r == 0 {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("bid with id %s couldn't reject", req.Id)
		}

		// reject tender all bids besides accepted bid
		awardQuery := `
		UPDATE
			tenders
		SET 
			status = 'awarded'
		WHERE 
			deleted_at = 0
		AND
			status <> 'awarded'
		AND
			tender_id = $1
		`

		res, err = tz.Exec(awardQuery, tender_id)

		if err != nil {
			return nil, err
		}

		if r, err := res.RowsAffected(); r == 0 {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("tender with id %s couldn't awarded", req.Id)
		}
	}

	return nil, fmt.Errorf("invalid status")
}
func (r *BidRepo) Delete(ctx context.Context, req *models.DeleteBidReq) (*models.Void, error) {
	query := `
	UPDATE 
		bids
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND
		contractor_id = $2
	AND 
		deleted_at = 0
	AND
		status = 'pending'
	`

	res, err := r.db.Exec(query, req.Id, req.ContractorId)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bid with id %s couldn't delete", req.Id)
	}

	return nil, nil
}

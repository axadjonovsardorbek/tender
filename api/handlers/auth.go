package handlers

import (
	"context"
	"net/http"

	"github.com/axadjonovsardorbek/tender/pkg/models"
	hp "github.com/axadjonovsardorbek/tender/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register
// @Description Register
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param user body models.RegisterReq true  "Registration request"
// @Success 201 {object} models.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {

	var body models.RegisterReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		return
	}

	if body.Email == "" || body.Username == ""  {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username or email cannot be empty"})
		return
	}

	emailExists, err := h.Clients.Auth.IsEmailTaken(context.Background(), body.Email)
	if !emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}
	if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        hp.SmsSender(c, err, http.StatusInternalServerError)
        return
    }

	if !hp.IsValidEmail(body.Email) {
		c.JSON(400, gin.H{"message": "username or email cannot be empty"})
		return
	}

	res, err := h.Clients.Auth.Register(context.Background(), &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// Login godoc
// @Summary Login
// @Description Authenticate user with username and password
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param admin body models.LoginReq true "Login credentials"
// @Success 200 {object} models.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var body models.LoginReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		return
	}

	if body.Username == "" || body.Password == "" {
		c.JSON(400, gin.H{"message": "Username and password are required"})
		return
	}

	res, err := h.Clients.Auth.Login(context.Background(), &body)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else if err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Serverda xatolik"})
			hp.SmsSender(c, err, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, res)
}

// Profile godoc
// @Summary Profile
// @Description Get profile
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} models.UserRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /profile [get]
func (h *Handler) Profile(c *gin.Context) {
	user_id := hp.ClaimData(c, "user_id")
	if user_id == "" {
		return
	}

	res, err := h.Clients.Auth.GetProfile(context.Background(), user_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateProfile godoc
// @Summary UpdateProfile
// @Description Update profile
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param user body models.UpdateReq true  "Update request"
// @Success 200 {object} string "Updated profile"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /profile/update [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	user_id := hp.ClaimData(c, "user_id")
	if user_id == "" {
		return
	}

	var body models.UpdateReq

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		hp.SmsSender(c, err, http.StatusBadRequest)
		return
	}

	if (body.Email != "" && body.Email != "string") || !hp.IsValidEmail(body.Email){
		c.JSON(409, gin.H{"message": "Incorrect email"})
		return
	}

	_, err := h.Clients.Auth.UpdateProfile(context.Background(), &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated profile"})
}

// DeleteProfile godoc
// @Summary DeleteProfile
// @Description Delete profile
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} string "Deleted profile"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /profile/delete [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	user_id := hp.ClaimData(c, "user_id")
	if user_id == "" {
		return
	}

	_, err := h.Clients.Auth.DeleteProfile(context.Background(), user_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted profile"})
}

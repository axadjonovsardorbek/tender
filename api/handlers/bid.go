package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/axadjonovsardorbek/tender/pkg/models"
	hp "github.com/axadjonovsardorbek/tender/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// CreateBid godoc
// @Summary Create a new bid
// @Description Create a new bid with the provided details
// @Tags bid
// @Accept  json
// @Produce  json
// @Param bid body models.ApiCreateBidReq true "Bid Details"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /client/tenders/{id}/bids [post]
func (h *Handler) CreateBid(c *gin.Context) {
	user_id := hp.GetUserId(c)
	body := &models.ApiCreateBidReq{}
	req := &models.CreateBidReq{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing request body: ", err)
		hp.SmsSender(c, err, http.StatusBadRequest)
		return
	}

	req.ContractorId = user_id
	req.Comments = body.Comments
	req.DeliveryTime = body.DeliveryTime
	req.Price = body.Price
	req.TenderId = body.TenderId

	_, err = h.Clients.Bid.Create(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error creating bid: ", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	slog.Info("Bid created successfully")
	c.JSON(200, "Bid Crested Successfully")
}

// GetByIdBid godoc
// @Summary Get Bid by ID
// @Description Get Bid by their ID
// @Tags bid
// @Accept  json
// @Produce  json
// @Param id query string false "Bid ID"
// @Success 200 {object} models.BidRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /client/tenders/{id}/bids [get]
func (h *Handler) GetByIdBid(c *gin.Context) {
	bid_id := c.Query("id")

	res, err := h.Clients.Bid.GetById(context.Background(), bid_id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Bid not found or access denied"})
			hp.SmsSender(c, err, http.StatusNotFound)
			return
		}
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error getting Bid by ID: ", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	slog.Info("Bid retrieved successfully")
	c.JSON(200, res)
}

// UpdateBid godoc
// @Summary Update an Bid
// @Description Update an Bid's details
// @Tags bid
// @Accept  json
// @Produce  json
// @Param bid_id query string false "Bid ID"
// @Param bid body models.ApiUpdateBidReq true "Bid Update Details"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /client/tenders/{id}/award/{bid_id} [put]
func (h *Handler) UpdateBid(c *gin.Context) {
	reqBody := models.ApiUpdateBidReq{}

	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing request body: ", err)
		hp.SmsSender(c, err, http.StatusBadRequest)
		return
	}

	req := models.UpdateBidReq{
		Id:     c.Query("bid_id"),
		Status: reqBody.Status,
	}

	_, err = h.Clients.Bid.Update(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error updating Bid: ", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	slog.Info("Bid updated successfully")
	c.JSON(200, "Bid updated successfully")
}

// GetAllBids godoc
// @Summary Get all Bids
// @Description Get all Bids with optional filtering
// @Tags bid
// @Accept  json
// @Produce  json
// @Param id query string false "TenderId"
// @Param contractor_id query string false "ContractorId"
// @Param delivery_time query string false "DeliveryTime"
// @Param price query string false "Price"
// @Param sort_type query string false "SortType"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} models.GetAllBidRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /client/tenders/{id}/bid [get]
func (h *Handler) GetAllBids(c *gin.Context) {
	tender_id := c.Query("id")
	contractor_id := c.Query("contractor_id")
	delivery_time := c.Query("delivery_time")
	price := c.Query("price")
	sort_type := c.Query("sort_type")
	limit := c.Query("limit")
	offset := c.Query("offset")

	limitValue, offsetValue, err := parsePaginationParams(c, limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		slog.Error("Error parsing pagination parameters: ", err)
		hp.SmsSender(c, err, http.StatusBadRequest)
		return
	}

	var priceValue int64
	var timeValue int64
	if price != "" {
		parsedPrice, err := strconv.Atoi(price)
		if err != nil {
			slog.Error("Invalid limit value", err)
			c.JSON(400, gin.H{"error": "Invalid price value"})
		}
		priceValue = int64(parsedPrice)
	}

	if delivery_time != "" {
		parsedTime, err := strconv.Atoi(delivery_time)
		if err != nil {
			slog.Error("Invalid offset value", err)
			c.JSON(400, gin.H{"error": "Invalid deliver time value"})
		}
		timeValue = int64(parsedTime)
	}

	req := &models.GetAllBidReq{
		TenderId:     tender_id,
		ContractorId: contractor_id,
		Price:        priceValue,
		DeliveryTime: timeValue,
		SortType:     sort_type,
		Filter: models.Filter{
			Limit:  int(limitValue),
			Offset: int(offsetValue),
		},
	}

	res, err := h.Clients.Bid.GetAll(context.Background(), req)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Bid not found or access denied"})
			hp.SmsSender(c, err, http.StatusNotFound)
			return
		}
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error getting Bids: ", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	slog.Info("Bids retrieved successfully")
	c.JSON(200, res)
}

// DeleteBid godoc
// @Summary Delete an Bid
// @Description Delete an Bid by ID
// @Tags bid
// @Accept  json
// @Produce  json
// @Param id query string false "Bid ID"
// @Success 200 {string} string "Bid deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /client/bids/delete [delete]
func (h *Handler) DeleteBid(c *gin.Context) {
	Bid_id := c.Query("id")

	user_id := hp.GetUserId(c)

	_, err := h.Clients.Bid.Delete(context.Background(), &models.DeleteBidReq{Id: Bid_id, ContractorId: user_id})
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error deleting Bid by ID: ", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	slog.Info("Bid deleted successfully")
	c.JSON(200, "Bid deleted successfully")
}

func parsePaginationParams(c *gin.Context, limit, offset string) (int, int, error) {
	limitValue := 10
	offsetValue := 0

	if limit != "" {
		parsedLimit, err := strconv.Atoi(limit)
		if err != nil {
			slog.Error("Invalid limit value", err)
			c.JSON(400, gin.H{"error": "Invalid limit value"})
			return 0, 0, err
		}
		limitValue = parsedLimit
	}

	if offset != "" {
		parsedOffset, err := strconv.Atoi(offset)
		if err != nil {
			slog.Error("Invalid offset value", err)
			c.JSON(400, gin.H{"error": "Invalid offset value"})
			return 0, 0, err
		}
		offsetValue = parsedOffset
	}

	return limitValue, offsetValue, nil
}

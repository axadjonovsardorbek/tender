package models

type BidRes struct {
	Id           string `json:"id"`
	TenderId     string `json:"tender_id"`
	ContractorId string `json:"contractor_id"`
	Price        int64  `json:"price"`
	DeliveryTime int64  `json:"delivery_time"`
	Comments     string `json:"comments"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
}

type ApiCreateBidReq struct {
	TenderId     string `json:"tender_id"`
	ContractorId string `json:"contractor_id"`
	Price        int64  `json:"price"`
	DeliveryTime int64  `json:"delivery_time"`
	Comments     string `json:"comments"`
}

type CreateBidReq struct {
	TenderId     string `json:"tender_id"`
	ContractorId string `json:"contractor_id"`
	Price        int64  `json:"price"`
	DeliveryTime int64  `json:"delivery_time"`
	Comments     string `json:"comments"`
}

type GetAllBidReq struct {
	TenderId     string `json:"tender_id"`
	ContractorId string `json:"contractor_id"`
	DeliveryTime int64  `json:"delivery_time"`
	Price        int64  `json:"price"`
	SortType     string `json:"sort_type"`
	Filter       Filter `json:"filter"`
}

type GetAllBidRes struct {
	Bids       []*BidRes `json:"bids"`
	TotalCount int64     `json:"total_count"`
}

type UpdateBidReq struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type ApiUpdateBidReq struct {
	Status string `json:"status"`
}

type DeleteBidReq struct {
	Id           string `json:"id"`
	ContractorId string `json:"contractor_id"`
}

package models

type Tender struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Deadline    string `json:"deadline" binding:"required"`
	Budget      int64  `json:"budget" binding:"required,min=1"`
	FileUrl     string `json:"file_url"`
	Status      string `json:"status"`
	ClientID    string `json:"client_id" binding:"required"`
}

type CreateTenderReq struct {
	Title       string  `json:"title" bindeng:"required"`
	Description string  `json:"description"`
	Deadline    string  `json:"deadline"`
	Budget      float64 `json:"budget"`
	FileUrl     string  `json:"file_url"`
}

type UpdateStatus struct {
	Status string `json:"status" bindeng:"required"`
}

type UpdateTenderReq struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type GetAllTenderReq struct {
	UserId   string  `json:"user_id"`
	Title    string  `json:"title"`
	Deadline string  `json:"deadline"`
	Budget   float64 `json:"budget"`
	Status   string  `json:"status"`
	ClientID string  `json:"client_id"`
	Filter   Filter  `json:"filter"`
}

type GetAllTenderRes struct {
	Tenders    []*Tender `json:"tenders"`
	TotalCount int64     `json:"total_count"`
}

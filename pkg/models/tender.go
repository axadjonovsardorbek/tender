package models

type Tender struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    string  `json:"deadline"`
	Budget      float64 `json:"budget"`
	ClientID    string  `json:"client_id"`
	FileUrl     string  `json:"file_url"`
}

type CreateTenderReq struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    string  `json:"deadline"`
	Budget      float64 `json:"budget"`
	FileUrl     string  `json:"file_url"`
}

type GetAllTenderReq struct {
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

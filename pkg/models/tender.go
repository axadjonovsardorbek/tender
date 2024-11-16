package models

type Tender struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Deadline    string  `json:"deadline"`
	Budget      float64 `json:"budget"`
	Status      string  `json:"status"` // "open", "closed", "awarded"
	ClientID    int64   `json:"client_id"`
}

type GetAllTenderReq struct {
	Title    string  `json:"title"`
	Deadline string  `json:"deadline"`
	Budget   float64 `json:"budget"`
	Status   string  `json:"status"`
	ClientID int64   `json:"client_id"`
	Filter   Filter  `json:"filter"`
}

type GetAllTenderRes struct {
	Tenders    []*Tender `json:"tenders"`
	TotalCount int64     `json:"total_count"`
}

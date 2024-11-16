package models

type Void struct {
}

type ById struct {
	ID string `json:"id"`
}

type Filter struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

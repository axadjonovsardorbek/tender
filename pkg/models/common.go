package models

type Void struct {
}

type Filter struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

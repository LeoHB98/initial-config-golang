package models

type Log struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ReferenceKey string `json:"reference_key"`
	Field        string `json:"field"`
	Value        string `json:"value"`
	Errr         error  `json:"erro"`
}

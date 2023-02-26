package models

type Event struct {
	Type string      `json:"t"`
	Data interface{} `json:"d"`
}

package models

type Authorization struct {
	Token string `json:"token"`
}

type Authenticated struct {
	IsAuthenticated bool `json:"isAuthenticated"`
}

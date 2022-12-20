package models

type Login struct {
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type UserRes struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
	Tel        string `json:"tel"`
	Email      string `json:"email"`
}

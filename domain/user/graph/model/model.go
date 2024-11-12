package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewUser struct {
	Name string `json:"name"`
}

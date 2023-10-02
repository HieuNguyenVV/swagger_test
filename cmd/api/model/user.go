package model

type CreateUserRequest struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
} // @name CreateUserRequest

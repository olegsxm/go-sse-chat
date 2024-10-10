package models

//easyjson:json
type AuthResponse struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

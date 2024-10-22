package models

//easyjson:json
type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//easyjson:json
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

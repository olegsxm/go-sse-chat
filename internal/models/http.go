package models

type AuthRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

type CreatePrivateConversationRequest struct {
	Message     string `json:"message" validate:"required"`
	Participant string `json:"participant" validate:"required,uuid"`
}

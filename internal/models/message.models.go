package models

//easyjson:json
type MessageRequest struct {
	Message      string `json:"message"`
	Conversation string `json:"conversation"`
}

//easyjson:json
type MessageResponse struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
	CreatedAt string `json:"createdAt"`
}

//easyjson:json
type Message struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
	Read      bool   `json:"read"`
	Sender    User   `json:"sender"`
}

//easyjson:json
type Messages []Message

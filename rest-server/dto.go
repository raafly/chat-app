package restserver

import "time"

type UserReq struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password" valdate:"min=8"`
	CreatedAt time.Time `json:"created_at"`
}

type ContactResp struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Message struct {
	ID         string    `json:"id"`
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

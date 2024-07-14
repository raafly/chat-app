package restserver

import "time"

type UserDTO struct {
	Telp string `json:"telp"`
	Name string `json:"name"`
}

type OTP struct {
	OTP int `json:"otp"`
}

type ContactDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ContactID string `json:"contact_id"`
}

type MessageDTO struct {
	ID         string    `json:"id"`
	SenderID   int64     `json:"sender_id"`
	ReceiverID int64     `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

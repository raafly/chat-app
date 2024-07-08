package restserver

import "time"

type UserDTO struct {
	Telp      int64     `json:"telp"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ContactDTO struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"user_id"`
	ContactID int64     `json:"contact_id"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageDTO struct {
	ID         string    `json:"id"`
	SenderID   int64     `json:"sender_id"`
	ReceiverID int64     `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

package restserver

import "time"

type User struct {
	Telp      string
	Name      string
	OTP       int
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Contact struct {
	ID        int
	UserID    string `json:"user_id ,omitempty"`
	ContactID string
	CreatedAt time.Time `json:"created_at ,omitempty"`
	UpdatedAt time.Time `json:"updated_at ,omitempty"`
}

type Message struct {
	ID         int
	SenderID   string
	ReceiverID string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

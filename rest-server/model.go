package restserver

import "time"

type User struct {
	Telp      int64
	Name      string
	OTP       int
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Contact struct {
	ID        string
	UserID    int64
	ContactID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Message struct {
	ID         int
	SenderID   int64
	ReceiverID int64
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

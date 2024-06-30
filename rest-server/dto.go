package restserver

type UserReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ContactResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID         string `json:"id"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	Timestamp  int64  `json:"timestamp"`
}
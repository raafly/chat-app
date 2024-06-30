package restserver

type User struct {
	ID       string
	Name     string
	Password string
}

type Contact struct {
	ID     string
	Name   string
	UserID string
}

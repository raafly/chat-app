package restserver

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type AuthRepo interface {
	Create(ctx context.Context, dto *UserReq) error
	FindByEmail(ctx context.Context, name string) (*User, error)
	GetProfile(ctx context.Context, id string) (*User, error)
	GetContacts(ctx context.Context, id string) (*[]Contact, error)
	GetHistory(ctx context.Context, userIO string, contactID string) (*[]Message, error)
}

type AuthRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &AuthRepoImpl{db: db}
}

func (a *AuthRepoImpl) Create(ctx context.Context, dto *UserReq) error {
	_, err := a.db.ExecContext(ctx, "insert into user(id, username, password) values(?, ?)", dto.Name, dto.Password)
	if err != nil {
		return fmt.Errorf("err exec sql %w", err)
	}

	return nil
}

func (a *AuthRepoImpl) FindByEmail(ctx context.Context, name string) (*User, error) {
	var user User
	row := a.db.QueryRowContext(ctx, "select username, password from user where username = ?", name)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("err scan row %w", err)
	}

	return &user, nil
}

func (a *AuthRepoImpl) GetProfile(ctx context.Context, id string) (*User, error) {
	var user User
	row := a.db.QueryRowContext(ctx, "select username, password from user where id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("err scan row %w", err)
	}

	return &user, nil
}

func (a *AuthRepoImpl) GetContacts(ctx context.Context, id string) (*[]Contact, error) {
	rows, err := a.db.QueryContext(ctx, "select id, username, user_id from contact where user_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("err exec query %w", err)
	}

	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.Name, &contact.UserID)
		if err != nil {
			log.Printf("err scan %s", err)
		}
		contacts = append(contacts, contact)
	}

	return &contacts, nil
}

func (a *AuthRepoImpl) GetHistory(ctx context.Context, userIO string, contactID string) (*[]Message, error) {
	rows, err := a.db.QueryContext(ctx, "select content from messages where user_id = ? and contact_id = ?", userIO, contactID)
	if err != nil {
		return nil, fmt.Errorf("err exec sql %w", err)
	}

	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			log.Printf("err scan %s", err)
		}
		messages = append(messages, msg)
	}

	return &messages, nil
}

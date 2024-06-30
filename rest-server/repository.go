package restserver

import (
	"context"
	"database/sql"
	"fmt"
)

type AuthRepo interface {
	Create(ctx context.Context, dto *UserReq) error
	FindByEmail(ctx context.Context, dto *UserReq) (*User, error)
	GetProfile(ctx context.Context, id string) (*User, error)
	GetContacts(ctx context.Context, id string) (*[]Contact, error)
	GetHistory(ctx context.Context, userIO string, contactID string) (*[]Message, error)
}

type AuthRepoImpl struct {
	db *sql.DB
}

func (a *AuthRepoImpl) 	Create(ctx context.Context, dto *UserReq) error {
	_, err := a.db.ExecContext(ctx, "insert into user(id, name, password) values(?, ?)", dto.Name, dto.Password)
	if err != nil {
		return fmt.Errorf("err exec sql %w", err)
	}

	return nil
}

func (a *AuthRepoImpl) 	FindByEmail(ctx context.Context, dto *UserReq) (*User, error) {
	var user User
	row := a.db.QueryRowContext(ctx, "select name, password from user where name = ?", dto.Name)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("err scan row %w", err)
	}

	return &user, nil
}

func (a *AuthRepoImpl) GetProfile(ctx context.Context, id string) (*User, error) {
	var user User
	row := a.db.QueryRowContext(ctx, "select name, password from user where id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("err scan row %w", err)
	}

	return &user, nil
}

func (a *AuthRepoImpl) GetContacts(ctx context.Context, id string) (*[]Contact, error) {
	rows, err := a.db.QueryContext(ctx, "select id, name, user_id from contact where id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("err exec query %w", err)
	}

	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.Name, &contact.UserID)
		contacts = append(contacts, contact) 
	}

	return &contacts, nil
}

func (a *AuthRepoImpl) GetHistory(ctx context.Context, userIO string, contactID string) (*[]Message, error) {
	rows, err := a.db.QueryContext(ctx, "select message from message where user_id = ? and contact_id = ?", userIO, contactID)
	if err != nil {
		return nil, fmt.Errorf("err exec sql %w", err)
	}

	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.Timestamp)
		messages = append(messages, msg)
	}
	
	return &messages, nil
}

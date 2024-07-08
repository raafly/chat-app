package restserver

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/raafly/realtime-app/helper"
)

type AuthRepo interface {
	Create(ctx context.Context, dto *UserDTO) error
	FindByTelp(ctx context.Context, telp int64) (*User, error)
	GetContacts(ctx context.Context, contactID string) (*[]Contact, error)
	GetHistory(ctx context.Context, userIO string, contactID string) (*[]Message, error)
}

type AuthRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &AuthRepoImpl{db: db}
}

func (a *AuthRepoImpl) Create(ctx context.Context, dto *UserDTO) error {
	trx, err := a.db.Begin()
	if err != nil {
		_ = trx.Rollback()
		return helper.ErrInternalServerError()
	}
	
	_, err = trx.ExecContext(ctx, "insert into users(telp, username) values(?, ?)", dto.Telp, dto.Name)
	if err != nil {
		_ = trx.Rollback()
		return fmt.Errorf("err exec sql %w", err)
	}

	return trx.Commit()
}

func (a *AuthRepoImpl) 	FindByTelp(ctx context.Context, telp int64) (*User, error) {
	var user User
	row := a.db.QueryRowContext(ctx, "select telp, username, bio from users where telp = ?", telp)
	err := row.Scan(&user.Telp, &user.Name, &user.Bio)
	if err != nil {
		return nil, fmt.Errorf("err scan row %w", err)
	}

	return &user, nil
}

func (a *AuthRepoImpl) 	GetContacts(ctx context.Context, contactID string) (*[]Contact, error) {
	rows, err := a.db.QueryContext(ctx, "select id, username, user_id from contacts where contact_id = ?", contactID)
	if err != nil {
		return nil, fmt.Errorf("err exec query %w", err)
	}

	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.UserID, &contact.ContactID)
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

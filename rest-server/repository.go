package restserver

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/raafly/realtime-app/helper"
)

type AuthRepo interface {
	Create(ctx context.Context, dto *UserDTO) (string, error)
	FindByTelp(ctx context.Context, telp string) (*User, error)
	FindByOTP(ctx context.Context, telp, otp string) (string, error)
	GetContacts(ctx context.Context, contactID string) (*[]Contact, error)
	GetHistory(ctx context.Context, userIO string, contactID string) (*[]Message, error)
}

type AuthRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &AuthRepoImpl{
		db: db,
	}
}

func (a *AuthRepoImpl) Create(ctx context.Context, dto *UserDTO) (string, error) {
	otp := helper.RandomOTP()	

	_, err := a.db.ExecContext(ctx, "insert into users(telp, otp) values(?, ?)", dto.Telp, otp)
	if err != nil {
		log.Println(err)
		return "", helper.ErrBadRequest("duplicate nomer", nil)
	}

	return otp, nil
}

func (a *AuthRepoImpl) FindByTelp(ctx context.Context, telp string) (*User, error) {
	var user User
	row := a.db.QueryRowContext(ctx, "select telp, username, bio from users where telp = ?", telp)
	err := row.Scan(&user.Telp, &user.Name, &user.Bio)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("err scan row %w", err)
	}

	return &user, nil
}

func (a *AuthRepoImpl) FindByOTP(ctx context.Context, telp, otp string) (string, error) {
	var result string
	row := a.db.QueryRowContext(ctx, "select otp from users where telp = ?", telp)
	err := row.Scan(&result)
	if err != nil {
		return "", helper.ErrNotFound("otp not found", nil)
	}
	
	return result, nil
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

func (a *AuthRepoImpl) GetHistory(ctx context.Context, sender_id, receiver_id string) (*[]Message, error) {
	rows, err := a.db.QueryContext(ctx, "select id, sender_id, receiver_id, content, created_at from messages where sender_id = ? and receiver_id = ?", sender_id, receiver_id)
	if err != nil {
		log.Printf("err db %s", err.Error())
		return nil, helper.ErrNotFound("history nil", nil)
	}

	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			log.Printf("err scan %s", err)
			return nil, helper.ErrInternalServerError()
		}
		messages = append(messages, msg)
	}

	return &messages, nil
}
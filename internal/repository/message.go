package repository

import (
	"errors"
	"fmt"

	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

type messageRepository struct{}

func (r *messageRepository) CreateChat(message models.Message) (models.Chat, error) {
	chat := models.Chat{}

	tx, err := db.Begin()

	if err != nil {
		tx.Rollback()
		return chat, errors.New("creating chat db error")
	}

	row, err := db.Exec(`INSERT INTO chatS (name) values (null);`)
	if err != nil {
		tx.Rollback()
		return chat, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return chat, err
	}

	chat.Id = id

	_, err = db.Exec(`insert into chat_members ( user_id, chat_id) VALUES (?, ?)`, message.Sender, id)
	if err != nil {
		tx.Rollback()
		return chat, err
	}

	_, err = db.Exec(`insert into chat_members ( user_id, chat_id) VALUES (?, ?)`, message.Recipient, id)
	if err != nil {
		tx.Rollback()
		return chat, err
	}

	_ = tx.Commit()

	return chat, nil
}

func (r *messageRepository) FindChat(id int64) (models.Chat, error) {
	chat := models.Chat{}

	query := db.QueryRow("select id, name, created_at, updated_at from chats where id = ?", id)

	err := query.Scan(&chat.Id, &chat.Name, &chat.CreatedAt, &chat.UpdatedAt)

	if err != nil {
		return chat, err
	}

	return chat, nil
}

func (r *messageRepository) CreateMessage(message models.Message) (models.Message, error) {
	row, err := db.Exec(
		`INSERT INTO messages ( message, sender, recipient, chat_id) VALUES (?, ?, ?, ?)`,
		message.Message, message.Sender, message.Recipient, message.ChatId,
	)

	if err != nil {
		fmt.Println(err)
		return message, err
	}

	id, err := row.LastInsertId()

	if err != nil {
		return message, err
	}

	message.ID = id

	return message, nil
}

func (r *messageRepository) GetLastMessages(chatId, limit int) ([]models.Message, error) {
	messages := make([]models.Message, 0, 20)

	rows, err := db.Query("select id, message, sender, recipient, chat_id, created_at from messages where chat_id = ? order by created_at DESC limit ?", chatId, limit)
	defer rows.Close()

	if err != nil {
		return messages, err
	}

	for rows.Next() {
		message := models.Message{}

		e := rows.Scan(&message.ID, &message.Message, &message.Sender, &message.Recipient, &message.ChatId, &message.CreatedAt)

		if e != nil {
			fmt.Println(e)
		}

		messages = append(messages, message)
	}

	return messages, nil
}

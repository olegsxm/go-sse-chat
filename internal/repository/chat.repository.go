package repository

import (
	"database/sql"
	"log/slog"

	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

type ChatRepository struct{}

func (r ChatRepository) CreateConversation(from int64, to int64) (models.Conversation, error) {
	conversation := models.Conversation{}

	parts := make([]models.ConversationParticipants, 0, 2)

	parts = append(parts, models.ConversationParticipants{UserId: to})
	parts = append(parts, models.ConversationParticipants{UserId: from})

	tx, err := st.Sql().Begin()
	if err != nil {
		slog.Error(err.Error())
		err := tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
			return conversation, err
		}
	}

	res, err := tx.Exec(`insert into conversations (name) values (?)`, nil)
	if err != nil {
		slog.Error(err.Error())
		err := tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
			return conversation, err
		}
	}

	cid, err := res.LastInsertId()
	if err != nil {
		slog.Error(err.Error())
		err = tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
			return conversation, err
		}
	}

	conversation.ID = cid

	for _, part := range parts {
		_, err := tx.Exec("INSERT INTO conversation_participants (conversation_id, user_id) values (?, ?)", conversation.ID, part.UserId)
		if err != nil {
			slog.Error(err.Error())
			err := tx.Rollback()
			if err != nil {
				return conversation, err
			}
			return conversation, err
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error(err.Error())
		return conversation, err
	}

	return conversation, err
}

func (r ChatRepository) GetConversations() ([]models.ConversationDTO, error) {
	conversations := make([]models.ConversationDTO, 0, 128)

	tx, err := st.Sql().Begin()
	if err != nil {
		slog.Error(err.Error())
		err := tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
		}
	}
	rows, e := st.Sql().Query(`select * from conversations`)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(rows)

	if e != nil {
		return nil, e
	}

	for rows.Next() {
		c := models.Conversation{}
		rows.Scan(&c.ID, &c.Name)
		conversations = append(conversations, c.ToDTO())
	}

	return conversations, nil
}

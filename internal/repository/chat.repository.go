package repository

import (
	"database/sql"
	"errors"
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

func (r ChatRepository) GetConversations(uId int64) ([]models.ConversationDTO, error) {
	conversations := make([]models.ConversationDTO, 0, 128)

	tx, err := st.Sql().Begin()
	if err != nil {
		slog.Error(err.Error())
		err := tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
		}
	}

	cpRow, err := tx.Query(`select conversation_id from conversation_participants where user_id = ?`, uId)
	if err != nil {
		slog.Error(err.Error())
		err := tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
		}
	}
	defer func(cpRow *sql.Rows) {
		err := cpRow.Close()
		if err != nil {
			slog.Error(err.Error())
			_ = tx.Rollback()
		}
	}(cpRow)

	for cpRow.Next() {
		var dto models.ConversationDTO
		err := cpRow.Scan(&dto.ID)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		var message = models.Message{}

		mRow := tx.QueryRow(`select id, message, sender_id, conversation_id, created_at from messages where conversation_id = ? order by created_at DESC limit 1`, dto.ID)
		err = mRow.Scan(&message.ID, &message.Message, &message.SenderId, &message.ConversationId, &message.CreatedAt)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				_ = tx.Rollback()
				return nil, err
			}

			dto.Message = nil
		}

		if message.ID != 0 {
			var sender models.UserDTO
			sRow := tx.QueryRow(`select id, users.login from users where id = ?`, message.SenderId)

			err = sRow.Scan(&sender.ID, &sender.Login)

			message.Sender = &sender

			dto.Message = &message

			if dto.Name == "" {
				var id int64
				_ = tx.QueryRow(`select user_id from conversation_participants where conversation_id = ? and user_id != ?`, dto.ID, uId).Scan(&id)
				_ = tx.QueryRow(`select login from users where id = ?`, id).Scan(&dto.Name)
			}
		} else {
			var id int64
			_ = tx.QueryRow(`select user_id from conversation_participants where conversation_id = ? and user_id != ?`, dto.ID, uId).Scan(&id)
			_ = tx.QueryRow(`select login from users where id = ?`, id).Scan(&dto.Name)
		}

		conversations = append(conversations, dto)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return conversations, nil
}

func (r ChatRepository) CreateMessage(message models.Message) (int64, error) {
	res, e := st.Sql().Exec(`INSERT INTO messages ( message, sender_id, conversation_id, created_at) values ( ?, ?, ?, ?)`, message.Message, message.SenderId, message.ConversationId, message.CreatedAt)
	if e != nil {
		slog.Error(e.Error())
	}
	id, e := res.LastInsertId()
	if e != nil {
		slog.Error(e.Error())
	}

	return id, e
}

func (r ChatRepository) GetMessages(id int64, user int64) ([]models.Message, error) {
	messages := make([]models.Message, 0, 50)

	tx, err := st.Sql().Begin()
	if err != nil {
		slog.Error(err.Error())
		_ = tx.Rollback()
		return messages, err
	}

	rows, err := tx.Query(`SELECT id, message, sender_id, created_at from messages where messages.conversation_id = ? order by created_at limit 50`, id)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(rows)

	if err != nil {
		slog.Error(err.Error())
		return messages, err
	}

	for rows.Next() {
		var message models.Message
		e := rows.Scan(&message.ID, &message.Message, &message.SenderId, &message.CreatedAt)
		if e != nil {
			slog.Error(e.Error())
		}

		if message.SenderId != user {
			u := models.UserDTO{}
			er := tx.QueryRow(`select id, users.login from users where id = ?`, message.SenderId).Scan(&u.ID, &u.Login)
			if er != nil {
				slog.Error(er.Error())
			}

			message.Sender = &u
		}

		messages = append(messages, message)
	}

	_ = tx.Commit()

	return messages, nil
}

func (r ChatRepository) GetConversationsParticipants(id int64, excludeID int64) ([]int64, error) {
	ids := make([]int64, 0, 2)
	rows, err := st.Sql().Query("select user_id from conversation_participants where conversation_id = ? and user_id != ?", id, excludeID)
	defer rows.Close()
	if err != nil {
		slog.Error(err.Error())
		return ids, err
	}

	for rows.Next() {
		var id int64
		_ = rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
}

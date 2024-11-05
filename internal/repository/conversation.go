package repository

import (
	"chat/internal/models"
	"chat/internal/repository/queries"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type ConversationRepository struct {
	queries *queries.Queries
	pool    *pgxpool.Pool
}

func (r ConversationRepository) CreateConversation(ctx context.Context) (models.Conversation, error) {
	c, err := r.queries.CreateConversation(ctx)
	if err != nil {
		return models.Conversation{}, err
	}

	res := models.Conversation{
		Id:      c.ID,
		Name:    c.Name,
		Created: c.CreatedAt.Time,
	}

	return res, nil
}

func (r ConversationRepository) CreateConversationWithMessage(ctx context.Context, senderUid, participantUid uuid.UUID, message string) (models.Conversation, models.Message, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return models.Conversation{}, models.Message{}, err
	}

	qtx := r.queries.WithTx(tx)
	check, err := r.checkConversationParticipants(ctx, tx, senderUid, participantUid)
	if err != nil {
		_ = tx.Rollback(ctx)
		return models.Conversation{}, models.Message{}, err
	}

	if !check {
		_ = tx.Rollback(ctx)
		slog.Error("conversation already exist")
		return models.Conversation{}, models.Message{}, errors.New("conversation already exist")
	}

	conversation, err := qtx.CreateConversation(ctx)
	if err != nil {
		slog.Error("create conversation error", err.Error())
		_ = tx.Rollback(ctx)
		return models.Conversation{}, models.Message{}, err
	}

	err = qtx.CreateParticipant(ctx, queries.CreateParticipantParams{
		ConversationID: conversation.ID,
		UserID:         senderUid,
	})
	if err != nil {
		slog.Error("Create participant error:", err.Error())
		_ = tx.Rollback(ctx)
		return models.Conversation{}, models.Message{}, err
	}

	err = qtx.CreateParticipant(ctx, queries.CreateParticipantParams{
		ConversationID: conversation.ID,
		UserID:         participantUid,
	})
	if err != nil {
		slog.Error("Create participant error:", err.Error())
		_ = tx.Rollback(ctx)
		return models.Conversation{}, models.Message{}, err
	}

	m, err := qtx.CreateMessage(ctx, queries.CreateMessageParams{
		Message:        message,
		SenderID:       senderUid,
		ConversationID: conversation.ID,
	})
	if err != nil {
		slog.Error("Create message error:", err.Error())
		_ = tx.Rollback(ctx)
		return models.Conversation{}, models.Message{}, err
	}

	conv := models.Conversation{
		Id:      conversation.ID,
		Name:    conversation.Name,
		Private: *conversation.Private,
		Created: conversation.CreatedAt.Time,
	}

	msg := models.Message{
		Id:             m.ID,
		Message:        m.Message,
		ConversationId: conversation.ID,
		SenderId:       m.SenderID,
		CreatedAt:      m.CreatedAt.Time,
	}

	return conv, msg, tx.Commit(ctx)
}

func (r ConversationRepository) GetConversations(ctx context.Context) ([]models.Conversation, error) {
	c, err := r.queries.GetConversations(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]models.Conversation, len(c))

	for i, qc := range c {
		res[i] = models.Conversation{
			Id:      qc.ID,
			Name:    qc.Name,
			Created: qc.CreatedAt.Time,
		}
	}

	return res, nil
}

func (r ConversationRepository) GetConversationById(ctx context.Context, id string) (models.Conversation, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return models.Conversation{}, errors.New("invalid conversation id")
	}

	c, err := r.queries.GetConversation(ctx, uid)
	if err != nil {
		return models.Conversation{}, err
	}

	res := models.Conversation{
		Id:      c.ID,
		Name:    c.Name,
		Created: c.CreatedAt.Time,
	}

	return res, nil
}

func (r ConversationRepository) SetConversationName(ctx context.Context, id, name string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid conversation id")
	}

	err = r.queries.SetConversationName(ctx, queries.SetConversationNameParams{
		ID:   uid,
		Name: &name,
	})

	return err
}

func (r ConversationRepository) checkConversationParticipants(ctx context.Context, tx pgx.Tx, senderUid, participantUid uuid.UUID) (bool, error) {
	query := `
		SELECT conversation_id
		FROM participants p
		GROUP BY conversation_id
		HAVING array_agg(p.user_id ORDER BY p.user_id) = array [$1, $2]::uuid[];
	`

	rows, err := tx.Query(ctx, query, senderUid, participantUid)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if count > 0 {
			break
		}
		if err = rows.Err(); err != nil {
			slog.Error("rows error:", err.Error())
			break
		}

		count += 1
	}

	return count == 0, err
}

// NewConversationRepository Constructor
func NewConversationRepository(queries *queries.Queries, pool *pgxpool.Pool) ConversationRepository {
	return ConversationRepository{
		queries: queries,
		pool:    pool,
	}
}

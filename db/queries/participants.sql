-- name: CreateParticipant :exec
INSERT INTO participants (conversation_id, user_id)
VALUES ($1, $2);

-- name: GetConversationParticipants :many
SELECT users.* FROM participants
JOIN users
ON participants.conversation_id = $1 AND participants.user_id = users.id;
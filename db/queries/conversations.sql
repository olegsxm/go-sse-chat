-- name: GetConversations :many
SELECT *
FROM conversations;

-- name: GetConversation :one
SELECT *
from conversations
WHERE id = $1;

-- эта хня не работает
-- name: CheckConversationParticipants :many
SELECT conversation_id
from participants p
GROUP BY conversation_id
HAVING array_agg(p.user_id order by p.user_id) =
       array [$1]::uuid[];

-- name: CreateConversation :one
INSERT INTO conversations (name)
VALUES (null)
returning *;

-- name: SetConversationName :exec
UPDATE conversations
SET name = $2
WHERE id = $1;


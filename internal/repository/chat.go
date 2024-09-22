package repository

import (
	"fmt"

	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

const (
	pageLimit = 20

	get_chat_query = `
	SELECT c.id AS chat, users.login, c.updated_at FROM chat_members AS cm
         JOIN chats c ON cm.chat_id = c.id
         JOIN users ON users.id = c.id
         WHERE cm.user_id == ?
	ORDER BY c.updated_at DESC
	LIMIT ?
`
)

type chatRepository struct {
}

func (r *chatRepository) GetChats(userId int) []models.Chat {
	chats := make([]models.Chat, 0, pageLimit)

	query, err := db.Query(
		get_chat_query,
		userId,
		pageLimit)
	defer query.Close()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for query.Next() {
		chat := models.Chat{}

		err := query.Scan(&chat.Id, &chat.Name, &chat.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		chats = append(chats, chat)
	}

	return chats
}

func (r *chatRepository) GetChatsPage(lastChatId int) {

}

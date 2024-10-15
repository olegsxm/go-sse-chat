package repository

import (
	"database/sql"
	"log/slog"

	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

type UsersRepository struct{}

func (r UsersRepository) FindUsers(query string) ([]models.User, error) {
	users := make([]models.User, 0, 20)

	rows, err := st.Sql().Query(`SELECT id, login FROM users WHERE login LIKE ?`, "%"+query+"%")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			slog.Error("error while closing rows: ", err.Error())
		}
	}(rows)

	if err != nil {
		return users, nil
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Login)
		if err != nil {
			slog.Error("error while scanning row: ", err.Error(), " query: ", query)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

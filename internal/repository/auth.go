package repository

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

type authRepository struct {
}

func (r *authRepository) FindUserByLogin(login string) (models.User, error) {
	user := models.User{}
	row := db.QueryRow(`select id, login, password from users where login = ?`, login)

	row.Scan(&user.ID, &user.Login, &user.Password)

	return user, nil

}

func (r *authRepository) CreateUser(user models.User) (int64, error) {
	row, err := db.Exec(`insert into users(login, password) values(?, ?)`, user.Login, user.Password)

	if err != nil {
		log.Error(err)
		return 0, errors.New("error inserting new user")
	}

	return row.LastInsertId()
}

/** Move to user repository **/
func (r *authRepository) SearchUser(login string) ([]models.User, error) {
	query := "select id, login, password from users where login like '%" + login + "%'"
	rows, err := db.Query(query)
	defer rows.Close()

	users := make([]models.User, 0, 20)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			fmt.Println(err)
		}

		users = append(users, user)
	}

	return users, nil

}

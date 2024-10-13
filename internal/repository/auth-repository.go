package repository

import (
	"errors"
	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

type AuthRepository struct{}

func (r *AuthRepository) FindUserByLogin(login string) (models.User, error) {
	u := models.User{}

	rows, err := st.Sql().Query(
		`SELECT id, login, password FROM users WHERE login = $1 LIMIT 1`,
		login,
	)
	defer rows.Close()

	if err != nil {
		return models.User{}, err
	}

	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	if u.ID == 0 {
		return models.User{}, errors.New("user not found")
	}

	return u, nil
}

func (r *AuthRepository) CreateUser(login, password string) (int64, error) {
	query := `INSERT INTO users (login, password) VALUES ($1, $2)`
	res, e := st.Sql().Exec(query, login, password)

	if e != nil {
		return 0, e
	}

	return res.LastInsertId()
}

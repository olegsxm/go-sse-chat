package repository

type Repository struct {
	Auth *authRepository
}

func New() *Repository {
	return &Repository{
		Auth: &authRepository{},
	}
}

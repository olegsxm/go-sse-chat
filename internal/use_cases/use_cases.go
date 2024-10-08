package use_cases

type IRepository interface {
	Auth()
}

type UseCases struct{}

func (c *UseCases) Auth() {
}

func New(r IRepository) UseCases {
	return UseCases{}
}

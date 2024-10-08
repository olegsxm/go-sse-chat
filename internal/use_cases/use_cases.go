package use_cases

type IRepository interface {
	Auth()
}

type useCases struct{}

func (c *useCases) Auth() {
}

func New(r IRepository) useCases {
	return useCases{}
}

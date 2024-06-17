package services

type repository interface {
}

func NewInteractor(repo repository) *interactor {
	return &interactor{
		repo: repo,
	}
}

type interactor struct {
	repo repository
}

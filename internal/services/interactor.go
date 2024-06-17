package services

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
)

type repository interface {
	GetQuiz(ctx context.Context) (models.Quiz, error)
}

func NewInteractor(repo repository) *interactor {
	return &interactor{
		repo: repo,
	}
}

type interactor struct {
	repo repository
}

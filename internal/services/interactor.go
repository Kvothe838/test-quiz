package services

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
)

type repository interface {
	CreateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error
	DoesChoiceSelectionExist(ctx context.Context, selection models.ChoiceSelection) (bool, error)
	GetChoiceSelections(ctx context.Context) ([]models.ChoiceSelection, error)
	GetQuiz(ctx context.Context) (models.Quiz, error)
	SaveQuizSubmission(context.Context) (models.QuizSubmission, error)
	UpdateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error
	UpdateSubmission(ctx context.Context, submission models.QuizSubmission) error
}

func NewInteractor(repo repository) *interactor {
	return &interactor{
		repo: repo,
	}
}

type interactor struct {
	repo repository
}

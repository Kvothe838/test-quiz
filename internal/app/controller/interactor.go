package controller

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
)

type Interactor interface {
	ConfirmChoices(ctx context.Context) error
	GetQuiz(ctx context.Context) (models.Quiz, error)
	SelectChoice(ctx context.Context, selection models.ChoiceSelection) error
}

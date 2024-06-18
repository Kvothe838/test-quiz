package controller

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
)

type Interactor interface {
	CalcBetterThanPercentage(ctx context.Context, submissionID int) (int, error)
	GetQuiz(ctx context.Context) (models.Quiz, error)
	SelectChoice(ctx context.Context, selection models.ChoiceSelection) error
	SubmitQuiz(ctx context.Context) (models.QuizSubmission, error)
}

package controller

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
)

type Interactor interface {
	GetQuiz(ctx context.Context) (models.Quiz, error)
}

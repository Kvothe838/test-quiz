package memory

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
)

func (r *repository) GetQuiz(context.Context) (models.Quiz, error) {
	return r.quiz, nil
}

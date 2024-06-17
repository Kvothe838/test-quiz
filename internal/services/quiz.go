package services

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/pkg/errors"
)

func (in interactor) GetQuiz(ctx context.Context) (models.Quiz, error) {
	quiz, err := in.repo.GetQuiz(ctx)
	if err != nil {
		return models.Quiz{}, errors.Wrap(err, "could not get quiz")
	}

	return quiz, nil
}

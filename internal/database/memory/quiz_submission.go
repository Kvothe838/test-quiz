package memory

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	internalerrors "github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/pkg/errors"
)

func (r *repository) UpdateSubmission(ctx context.Context, submission models.QuizSubmission) error {
	for index, savedSubmission := range r.quizSubmissions {
		if savedSubmission.ID == submission.ID {
			savedSubmission.Selections = submission.Selections
			savedSubmission.HitsAmount = submission.HitsAmount
			r.quizSubmissions[index] = savedSubmission
		}
	}

	return nil
}

func (r *repository) SaveQuizSubmission(ctx context.Context) (models.QuizSubmission, error) {
	if r.submitted {
		return models.QuizSubmission{}, internalerrors.QuizAlreadySubmittedErr
	}

	selections, err := r.getChoiceSelections(ctx)
	if err != nil {
		return models.QuizSubmission{}, err
	}

	id := 1

	if len(r.quizSubmissions) > 0 {
		lastQuizSubmission := r.quizSubmissions[len(r.quizSubmissions)-1]
		id = lastQuizSubmission.ID + 1
	}

	submission := models.QuizSubmission{
		ID:         id,
		Selections: selections,
	}

	r.quizSubmissions = append(r.quizSubmissions, submission)

	r.submitted = true

	return submission, nil
}

func (r *repository) GetAllQuizSubmissions(ctx context.Context) ([]models.QuizSubmission, error) {
	return r.quizSubmissions, nil
}
func (r *repository) GetQuizSubmission(ctx context.Context, submissionID int) (models.QuizSubmission, error) {
	for _, submission := range r.quizSubmissions {
		if submission.ID == submissionID {
			return submission, nil
		}
	}

	return models.QuizSubmission{}, errors.New("quiz submission not found")
}

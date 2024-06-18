package memory

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
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
	selections, err := r.getChoiceSelections(ctx)
	if err != nil {
		return models.QuizSubmission{}, err
	}

	submission := models.QuizSubmission{
		Selections: selections,
	}

	r.quizSubmissions = append(r.quizSubmissions, submission)

	r.currentSelectionByQuestionID = make(map[int]models.ChoiceSelection)

	return submission, nil
}

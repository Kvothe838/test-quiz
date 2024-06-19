package services

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	internalerrors "github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/pkg/errors"
)

func (in *interactor) SubmitQuiz(ctx context.Context) (models.QuizSubmission, error) {
	quiz, err := in.repo.GetQuiz(ctx)
	if err != nil {
		return models.QuizSubmission{}, errors.Wrap(err, "could not get quiz")
	}

	missingChoiceSelections, err := in.getMissingChoiceSelections(ctx, quiz)
	if err != nil {
		return models.QuizSubmission{}, errors.Wrap(err, "could not get missing choice selections")
	}

	if len(missingChoiceSelections) != 0 {
		return models.QuizSubmission{}, internalerrors.MissingChoicesSelectionErr{QuestionIDs: missingChoiceSelections}
	}

	submission, err := in.repo.SaveQuizSubmission(ctx)
	if err != nil {
		return models.QuizSubmission{}, errors.Wrap(err, "could not save quiz submission")
	}

	hitsAmount := CalcHits(quiz.Questions, submission.Selections)

	submission.HitsAmount = hitsAmount

	err = in.repo.UpdateSubmission(ctx, submission)
	if err != nil {
		return models.QuizSubmission{}, errors.Wrap(err, "could not update submission")
	}

	return submission, nil
}

func CalcHits(questions []models.Question, selections []models.ChoiceSelection) int {
	hitsAmount := 0

	choiceIDByQuestionID := make(map[int]int)

	for _, selection := range selections {
		choiceIDByQuestionID[selection.QuestionID] = selection.ChoiceID
	}

	for _, question := range questions {
		selectedChoice := choiceIDByQuestionID[question.ID]
		var correctChoice int

		for _, choice := range question.Choices {
			if choice.IsCorrect {
				correctChoice = choice.ID
			}
		}

		if selectedChoice == correctChoice {
			hitsAmount++
		}
	}

	return hitsAmount
}

func (in *interactor) getMissingChoiceSelections(ctx context.Context, quiz models.Quiz) ([]int, error) {
	choiceSelections, err := in.repo.GetChoiceSelections(ctx)
	if err != nil {
		return nil, err
	}

	choiceSelectionIDByQuestionID := make(map[int]int)
	for _, choiceSelection := range choiceSelections {
		choiceSelectionIDByQuestionID[choiceSelection.QuestionID] = choiceSelection.ChoiceID
	}

	missingQuestionIDs := make([]int, 0)
	for _, question := range quiz.Questions {
		_, choiceExists := choiceSelectionIDByQuestionID[question.ID]

		if !choiceExists {
			missingQuestionIDs = append(missingQuestionIDs, question.ID)
		}
	}

	return missingQuestionIDs, nil
}

func (in *interactor) CalcBetterThanPercentage(ctx context.Context, submissionID int) (int, error) {
	submissions, err := in.repo.GetAllQuizSubmissions(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "could not get all quiz submissions")
	}

	submissionToCalc, err := in.repo.GetQuizSubmission(ctx, submissionID)
	if err != nil {
		return 0, errors.Wrapf(err, "could not get quiz submission for id %d", submissionID)
	}

	betterThanAmount := 0

	for _, submission := range submissions {
		if submission.ID == submissionID {
			continue
		}

		if submissionToCalc.HitsAmount > submission.HitsAmount {
			betterThanAmount++
		}
	}

	submissionsAmount := len(submissions) - 1
	betterThanPercentage := int(float64(betterThanAmount) / float64(submissionsAmount) * 100)

	return betterThanPercentage, nil
}

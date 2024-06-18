package services

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	internalerrors "github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/pkg/errors"
)

func (in *interactor) SubmitQuiz(ctx context.Context) (int, error) {
	missingChoiceSelections, err := in.getMissingChoiceSelections(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "could not get missing choice selections")
	}

	if len(missingChoiceSelections) != 0 {
		return 0, internalerrors.MissingChoicesSelectionErr{QuestionIDs: missingChoiceSelections}
	}

	submission, err := in.repo.SaveQuizSubmission(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "could not confirm choices")
	}

	quiz, err := in.repo.GetQuiz(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "could not get quiz to assert results")
	}

	hitsAmount := CalcHits(quiz.Questions, submission.Selections)

	submission.HitsAmount = hitsAmount

	err = in.repo.UpdateSubmission(ctx, submission)
	if err != nil {
		return 0, errors.Wrap(err, "could not update submission")
	}

	return hitsAmount, nil
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

func (in *interactor) getMissingChoiceSelections(ctx context.Context) ([]int, error) {
	choiceSelections, err := in.repo.GetChoiceSelections(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get choice selections")
	}

	quiz, err := in.repo.GetQuiz(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get quiz")
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

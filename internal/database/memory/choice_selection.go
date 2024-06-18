package memory

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/Kvothe838/fast-track-test-quiz/internal/pkg/errors"
	"github.com/samber/lo"
)

func (r *repository) CreateOptionSelection(ctx context.Context, selection models.ChoiceSelection) error {
	err := r.checkOptionAndQuestionExist(ctx, selection.QuestionID, selection.ChoiceID)
	if err != nil {
		return err
	}

	r.currentSelectionByQuestionID[selection.QuestionID] = selection

	return nil
}

func (r *repository) UpdateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error {
	err := r.checkOptionAndQuestionExist(ctx, selection.QuestionID, selection.ChoiceID)
	if err != nil {
		return err
	}

	currentSelection, exists := r.currentSelectionByQuestionID[selection.QuestionID]
	if !exists {
		return errors.SelectedOptionNotExistsErr
	}

	currentSelection.ChoiceID = selection.ChoiceID
	r.currentSelectionByQuestionID[selection.QuestionID] = selection

	return nil
}

func (r *repository) checkOptionAndQuestionExist(ctx context.Context, questionID, choiceID int) error {
	question, questionFound := lo.Find(r.quiz.Questions, func(question models.Question) bool {
		return question.ID == questionID
	})

	if !questionFound {
		return errors.QuestionNotExistsErr
	}

	choiceFound := lo.SomeBy(question.Choices, func(choice models.Choice) bool {
		return choice.ID == choiceID
	})

	if !choiceFound {
		return errors.ChoiceNotExistsErr
	}

	return nil
}

func (r *repository) DoesChoiceSelectionExist(ctx context.Context, selection models.ChoiceSelection) (bool, error) {
	_, exists := r.currentSelectionByQuestionID[selection.QuestionID]

	return exists, nil
}

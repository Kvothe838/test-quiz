package services

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/pkg/errors"
)

func (in *interactor) SelectChoice(ctx context.Context, selection models.ChoiceSelection) error {
	choiceExists, err := in.repo.DoesChoiceSelectionExist(ctx, selection)
	if err != nil {
		return errors.Wrap(err, "could not check if option selection exist")
	}

	if choiceExists {
		err = in.repo.UpdateChoiceSelection(ctx, selection)
		if err != nil {
			return errors.Wrap(err, "could not update choice selection")
		}
		return nil
	}

	err = in.repo.CreateChoiceSelection(ctx, selection)
	if err != nil {
		return errors.Wrap(err, "could not create option selection")
	}

	return nil
}

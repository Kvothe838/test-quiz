package services_tests

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/Kvothe838/fast-track-test-quiz/internal/services"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectChoice(t *testing.T) {
	tests := []struct {
		name           string
		selection      models.ChoiceSelection
		prepareRepo    func(*MockRepository)
		assertOnResult func(err error)
	}{
		{
			name: "success on SelectChoice with not existent choice selection",
			selection: models.ChoiceSelection{
				QuestionID: 1,
				ChoiceID:   1,
			},
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().DoesChoiceSelectionExist(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					false,
					nil,
				)

				mock.EXPECT().CreateChoiceSelection(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					nil,
				)
			},
			assertOnResult: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "success on SelectChoice with existent choice selection",
			selection: models.ChoiceSelection{
				QuestionID: 1,
				ChoiceID:   1,
			},
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().DoesChoiceSelectionExist(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					true,
					nil,
				)

				mock.EXPECT().UpdateChoiceSelection(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					nil,
				)
			},
			assertOnResult: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "error on DoesChoiceSelectionExist",
			selection: models.ChoiceSelection{
				QuestionID: 1,
				ChoiceID:   1,
			},
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().DoesChoiceSelectionExist(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					false,
					errors.New("test error"),
				)
			},
			assertOnResult: func(err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not check if option selection exist")
			},
		},
		{
			name: "error on CreateChoiceSelection",
			selection: models.ChoiceSelection{
				QuestionID: 1,
				ChoiceID:   1,
			},
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().DoesChoiceSelectionExist(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					false,
					nil,
				)

				mock.EXPECT().CreateChoiceSelection(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					errors.New("test error"),
				)
			},
			assertOnResult: func(err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not create choice selection")
			},
		},
		{
			name: "error on UpdateChoiceSelection",
			selection: models.ChoiceSelection{
				QuestionID: 1,
				ChoiceID:   1,
			},
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().DoesChoiceSelectionExist(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					true,
					nil,
				)

				mock.EXPECT().UpdateChoiceSelection(
					gomock.Any(),
					models.ChoiceSelection{
						QuestionID: 1,
						ChoiceID:   1,
					},
				).Return(
					errors.New("test error"),
				)
			},
			assertOnResult: func(err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not update choice selection")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := NewMockRepository(gomock.NewController(t))

			if test.prepareRepo != nil {
				test.prepareRepo(repo)
			}

			in := services.NewInteractor(repo)
			err := in.SelectChoice(context.Background(), test.selection)

			if test.assertOnResult != nil {
				test.assertOnResult(err)
			}
		})
	}
}

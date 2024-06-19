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

func TestSubmitQuiz(t *testing.T) {
	tests := []struct {
		name           string
		prepareRepo    func(*MockRepository)
		assertOnResult func(submission models.QuizSubmission, err error)
	}{
		{
			name: "success on SubmitQuiz without hits",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{},
					nil,
				)

				mock.EXPECT().GetChoiceSelections(
					gomock.Any(),
				).Return(
					[]models.ChoiceSelection{},
					nil,
				)

				mock.EXPECT().SaveQuizSubmission(
					gomock.Any(),
				).Return(
					models.QuizSubmission{
						ID: 1,
					},
					nil,
				)

				mock.EXPECT().UpdateSubmission(
					gomock.Any(),
					models.QuizSubmission{
						ID:         1,
						HitsAmount: 0,
					},
				).Return(
					nil,
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.NoError(t, err)
				assert.Equal(t, submission.HitsAmount, 0)
			},
		},
		{
			name: "success on SubmitQuiz with 2 hits",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{
						Questions: []models.Question{
							{
								ID: 1,
								Choices: []models.Choice{
									{
										ID:        1,
										IsCorrect: true,
									},
									{
										ID: 2,
									},
								},
							},
							{
								ID: 2,
								Choices: []models.Choice{
									{
										ID:        1,
										IsCorrect: true,
									},
									{
										ID: 2,
									},
								},
							},
							{
								ID: 3,
								Choices: []models.Choice{
									{
										ID:        1,
										IsCorrect: true,
									},
									{
										ID: 2,
									},
								},
							},
						},
					},
					nil,
				)

				selections := []models.ChoiceSelection{
					{
						QuestionID: 1,
						ChoiceID:   1,
					},
					{
						QuestionID: 2,
						ChoiceID:   1,
					},
					{
						QuestionID: 3,
						ChoiceID:   2,
					},
				}

				mock.EXPECT().GetChoiceSelections(
					gomock.Any(),
				).Return(
					selections,
					nil,
				)

				mock.EXPECT().SaveQuizSubmission(
					gomock.Any(),
				).Return(
					models.QuizSubmission{
						ID:         1,
						Selections: selections,
					},
					nil,
				)

				mock.EXPECT().UpdateSubmission(
					gomock.Any(),
					models.QuizSubmission{
						ID:         1,
						Selections: selections,
						HitsAmount: 2,
					},
				).Return(
					nil,
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.NoError(t, err)
				assert.Equal(t, submission.HitsAmount, 2)
			},
		},
		{
			name: "missing choice selections",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{
						Questions: []models.Question{
							{
								ID: 1,
								Choices: []models.Choice{
									{
										ID: 1,
									},
								},
							},
						},
					},
					nil,
				)

				mock.EXPECT().GetChoiceSelections(
					gomock.Any(),
				).Return(
					[]models.ChoiceSelection{},
					nil,
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "some questions have missing choice selection")
			},
		},
		{
			name: "error on GetQuiz",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{},
					errors.New("test error"),
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not get quiz")
			},
		},
		{
			name: "error on GetChoiceSelections",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{},
					nil,
				)

				mock.EXPECT().GetChoiceSelections(
					gomock.Any(),
				).Return(
					nil,
					errors.New("test error"),
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not get missing choice selections")
			},
		},
		{
			name: "error on SaveQuizSubmission",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{},
					nil,
				)

				mock.EXPECT().GetChoiceSelections(
					gomock.Any(),
				).Return(
					[]models.ChoiceSelection{},
					nil,
				)

				mock.EXPECT().SaveQuizSubmission(
					gomock.Any(),
				).Return(
					models.QuizSubmission{
						ID: 1,
					},
					errors.New("test error"),
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not save quiz submission")
			},
		},
		{
			name: "error on UpdateSubmission",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{},
					nil,
				)

				mock.EXPECT().GetChoiceSelections(
					gomock.Any(),
				).Return(
					[]models.ChoiceSelection{},
					nil,
				)

				mock.EXPECT().SaveQuizSubmission(
					gomock.Any(),
				).Return(
					models.QuizSubmission{
						ID: 1,
					},
					nil,
				)

				mock.EXPECT().UpdateSubmission(
					gomock.Any(),
					models.QuizSubmission{
						ID:         1,
						HitsAmount: 0,
					},
				).Return(
					errors.New("test error"),
				)
			},
			assertOnResult: func(submission models.QuizSubmission, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not update submission")
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
			submission, err := in.SubmitQuiz(context.Background())

			if test.assertOnResult != nil {
				test.assertOnResult(submission, err)
			}
		})
	}
}

func TestCalcBetterThanPercentage(t *testing.T) {
	tests := []struct {
		name           string
		submissionID   int
		prepareRepo    func(*MockRepository)
		assertOnResult func(percentage int, err error)
	}{
		{
			name:         "success with better than percentage calculated",
			submissionID: 1,
			prepareRepo: func(mock *MockRepository) {
				submissions := []models.QuizSubmission{
					{ID: 1, HitsAmount: 5},
					{ID: 2, HitsAmount: 3},
					{ID: 3, HitsAmount: 4},
				}
				mock.EXPECT().GetAllQuizSubmissions(
					gomock.Any(),
				).Return(
					submissions,
					nil,
				)

				mock.EXPECT().GetQuizSubmission(
					gomock.Any(),
					1,
				).Return(
					models.QuizSubmission{ID: 1, HitsAmount: 5},
					nil,
				)
			},
			assertOnResult: func(percentage int, err error) {
				assert.NoError(t, err)
				assert.Equal(t, percentage, 100)
			},
		},
		{
			name:         "success with no better than percentage calculated",
			submissionID: 2,
			prepareRepo: func(mock *MockRepository) {
				submissions := []models.QuizSubmission{
					{ID: 1, HitsAmount: 5},
					{ID: 2, HitsAmount: 3},
					{ID: 3, HitsAmount: 4},
				}
				mock.EXPECT().GetAllQuizSubmissions(
					gomock.Any(),
				).Return(
					submissions,
					nil,
				)

				mock.EXPECT().GetQuizSubmission(
					gomock.Any(),
					2,
				).Return(
					models.QuizSubmission{ID: 2, HitsAmount: 3},
					nil,
				)
			},
			assertOnResult: func(percentage int, err error) {
				assert.NoError(t, err)
				assert.Equal(t, percentage, 0)
			},
		},
		{
			name:         "error on GetAllQuizSubmissions",
			submissionID: 1,
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetAllQuizSubmissions(
					gomock.Any(),
				).Return(
					nil,
					errors.New("test error"),
				)
			},
			assertOnResult: func(percentage int, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not get all quiz submissions")
			},
		},
		{
			name:         "error on GetQuizSubmission",
			submissionID: 1,
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetAllQuizSubmissions(
					gomock.Any(),
				).Return(
					[]models.QuizSubmission{
						{ID: 1, HitsAmount: 5},
						{ID: 2, HitsAmount: 3},
					},
					nil,
				)

				mock.EXPECT().GetQuizSubmission(
					gomock.Any(),
					1,
				).Return(
					models.QuizSubmission{},
					errors.New("test error"),
				)
			},
			assertOnResult: func(percentage int, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not get quiz submission for id 1")
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
			percentage, err := in.CalcBetterThanPercentage(context.Background(), test.submissionID)

			if test.assertOnResult != nil {
				test.assertOnResult(percentage, err)
			}
		})
	}
}

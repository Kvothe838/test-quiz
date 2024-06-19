package services_tests

import (
	"context"
	"testing"

	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/Kvothe838/fast-track-test-quiz/internal/services"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetQuiz(t *testing.T) {
	tests := []struct {
		name           string
		prepareRepo    func(*MockRepository)
		assertOnResult func(quiz models.Quiz, err error)
	}{
		{
			name: "success on GetQuiz",
			prepareRepo: func(mock *MockRepository) {
				mock.EXPECT().GetQuiz(
					gomock.Any(),
				).Return(
					models.Quiz{
						Title:       "test title",
						Description: "test description",
						Questions:   nil,
					},
					nil,
				)
			},
			assertOnResult: func(quiz models.Quiz, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "test title", quiz.Title, "quiz must have 'test title' as title")
				assert.Equal(t, "test description", quiz.Description, "quiz must have 'test description' as description")
				assert.Nil(t, quiz.Questions, "questions must be nil")
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
			assertOnResult: func(quiz models.Quiz, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "test error")
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
			quiz, err := in.GetQuiz(context.Background())

			if test.assertOnResult != nil {
				test.assertOnResult(quiz, err)
			}
		})
	}
}

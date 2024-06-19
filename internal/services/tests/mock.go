package services_tests

import (
	"context"
	"github.com/Kvothe838/fast-track-test-quiz/internal/models"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryRecorder
}

func (m *MockRepository) CreateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) DoesChoiceSelectionExist(ctx context.Context, selection models.ChoiceSelection) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) GetChoiceSelections(ctx context.Context) ([]models.ChoiceSelection, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) GetAllQuizSubmissions(ctx context.Context) ([]models.QuizSubmission, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) GetQuizSubmission(ctx context.Context, submissionID int) (models.QuizSubmission, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) SaveQuizSubmission(ctx context.Context) (models.QuizSubmission, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) UpdateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) UpdateSubmission(ctx context.Context, submission models.QuizSubmission) error {
	//TODO implement me
	panic("implement me")
}

type MockRepositoryRecorder struct {
	mock *MockRepository
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	mock.recorder = &MockRepositoryRecorder{mock: mock}
	return mock
}

func (m *MockRepository) EXPECT() *MockRepositoryRecorder {
	return m.recorder
}

func (m *MockRepository) GetQuiz(ctx context.Context) (models.Quiz, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuiz", ctx)
	ret0, _ := ret[0].(models.Quiz)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRepositoryRecorder) GetQuiz(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuiz", reflect.TypeOf((*MockRepository)(nil).GetQuiz), ctx)
}

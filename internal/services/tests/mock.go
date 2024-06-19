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

func (m *MockRepository) CreateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChoiceSelection", ctx, selection)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockRepository) DoesChoiceSelectionExist(ctx context.Context, selection models.ChoiceSelection) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesChoiceSelectionExist", ctx, selection)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) GetChoiceSelections(ctx context.Context) ([]models.ChoiceSelection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChoiceSelections", ctx)
	ret0, _ := ret[0].([]models.ChoiceSelection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) GetAllQuizSubmissions(ctx context.Context) ([]models.QuizSubmission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllQuizSubmissions", ctx)
	ret0, _ := ret[0].([]models.QuizSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) GetQuizSubmission(ctx context.Context, submissionID int) (models.QuizSubmission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuizSubmission", ctx, submissionID)
	ret0, _ := ret[0].(models.QuizSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) SaveQuizSubmission(ctx context.Context) (models.QuizSubmission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveQuizSubmission", ctx)
	ret0, _ := ret[0].(models.QuizSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) UpdateChoiceSelection(ctx context.Context, selection models.ChoiceSelection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChoiceSelection", ctx, selection)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockRepository) UpdateSubmission(ctx context.Context, submission models.QuizSubmission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubmission", ctx, submission)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRepositoryRecorder) GetQuiz(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuiz", reflect.TypeOf((*MockRepository)(nil).GetQuiz), ctx)
}

func (mr *MockRepositoryRecorder) CreateChoiceSelection(ctx, selection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChoiceSelection", reflect.TypeOf((*MockRepository)(nil).CreateChoiceSelection), ctx, selection)
}

func (mr *MockRepositoryRecorder) DoesChoiceSelectionExist(ctx, selection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesChoiceSelectionExist", reflect.TypeOf((*MockRepository)(nil).DoesChoiceSelectionExist), ctx, selection)
}

func (mr *MockRepositoryRecorder) GetChoiceSelections(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChoiceSelections", reflect.TypeOf((*MockRepository)(nil).GetChoiceSelections), ctx)
}

func (mr *MockRepositoryRecorder) GetAllQuizSubmissions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllQuizSubmissions", reflect.TypeOf((*MockRepository)(nil).GetAllQuizSubmissions), ctx)
}

func (mr *MockRepositoryRecorder) GetQuizSubmission(ctx, submissionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuizSubmission", reflect.TypeOf((*MockRepository)(nil).GetQuizSubmission), ctx, submissionID)
}

func (mr *MockRepositoryRecorder) SaveQuizSubmission(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveQuizSubmission", reflect.TypeOf((*MockRepository)(nil).SaveQuizSubmission), ctx)
}

func (mr *MockRepositoryRecorder) UpdateChoiceSelection(ctx, selection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChoiceSelection", reflect.TypeOf((*MockRepository)(nil).UpdateChoiceSelection), ctx, selection)
}

func (mr *MockRepositoryRecorder) UpdateSubmission(ctx, submission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubmission", reflect.TypeOf((*MockRepository)(nil).UpdateSubmission), ctx, submission)
}

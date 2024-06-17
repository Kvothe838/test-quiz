package services_tests

import (
	"github.com/golang/mock/gomock"
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

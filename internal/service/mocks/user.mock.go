// Code generated by MockGen. DO NOT EDIT.
// Source: webook\internal\service\user.go

// Package svcmocks is a generated GoMock package.
package svcmocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/ecodeclub/webook/internal/domain"

	gomock "go.uber.org/mock/gomock"
)

// MockUserAndService is a mock of UserAndService interface.
type MockUserAndService struct {
	ctrl     *gomock.Controller
	recorder *MockUserAndServiceMockRecorder
}

// MockUserAndServiceMockRecorder is the mock recorder for MockUserAndService.
type MockUserAndServiceMockRecorder struct {
	mock *MockUserAndService
}

// NewMockUserAndService creates a new mock instance.
func NewMockUserAndService(ctrl *gomock.Controller) *MockUserAndService {
	mock := &MockUserAndService{ctrl: ctrl}
	mock.recorder = &MockUserAndServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAndService) EXPECT() *MockUserAndServiceMockRecorder {
	return m.recorder
}

// Signup mocks base method.
func (m *MockUserAndService) Signup(ctx context.Context, u *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signup indicates an expected call of Signup.
func (mr *MockUserAndServiceMockRecorder) Signup(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockUserAndService)(nil).Signup), ctx, u)
}

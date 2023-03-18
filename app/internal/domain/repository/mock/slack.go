// Code generated by MockGen. DO NOT EDIT.
// Source: slack.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/moneyforward/auriga/app/internal/model"
)

// MockSlackRepository is a mock of SlackRepository interface.
type MockSlackRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSlackRepositoryMockRecorder
}

// MockSlackRepositoryMockRecorder is the mock recorder for MockSlackRepository.
type MockSlackRepositoryMockRecorder struct {
	mock *MockSlackRepository
}

// NewMockSlackRepository creates a new mock instance.
func NewMockSlackRepository(ctrl *gomock.Controller) *MockSlackRepository {
	mock := &MockSlackRepository{ctrl: ctrl}
	mock.recorder = &MockSlackRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlackRepository) EXPECT() *MockSlackRepositoryMockRecorder {
	return m.recorder
}

// GetParentMessage mocks base method.
func (m *MockSlackRepository) GetParentMessage(ctx context.Context, channelID, ts string) (*model.SlackMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParentMessage", ctx, channelID, ts)
	ret0, _ := ret[0].(*model.SlackMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParentMessage indicates an expected call of GetParentMessage.
func (mr *MockSlackRepositoryMockRecorder) GetParentMessage(ctx, channelID, ts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentMessage", reflect.TypeOf((*MockSlackRepository)(nil).GetParentMessage), ctx, channelID, ts)
}

// GetReactions mocks base method.
func (m *MockSlackRepository) GetReactions(ctx context.Context, channelID, ts string, full bool) ([]*model.SlackReaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReactions", ctx, channelID, ts, full)
	ret0, _ := ret[0].([]*model.SlackReaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReactions indicates an expected call of GetReactions.
func (mr *MockSlackRepositoryMockRecorder) GetReactions(ctx, channelID, ts, full interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReactions", reflect.TypeOf((*MockSlackRepository)(nil).GetReactions), ctx, channelID, ts, full)
}

// ListUsersEmail mocks base method.
func (m *MockSlackRepository) ListUsersEmail(ctx context.Context, userID []string) ([]*model.SlackUserEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersEmail", ctx, userID)
	ret0, _ := ret[0].([]*model.SlackUserEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsersEmail indicates an expected call of ListUsersEmail.
func (mr *MockSlackRepositoryMockRecorder) ListUsersEmail(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersEmail", reflect.TypeOf((*MockSlackRepository)(nil).ListUsersEmail), ctx, userID)
}

// PostEphemeral mocks base method.
func (m *MockSlackRepository) PostEphemeral(ctx context.Context, channelID, message, ts, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostEphemeral", ctx, channelID, message, ts, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostEphemeral indicates an expected call of PostEphemeral.
func (mr *MockSlackRepositoryMockRecorder) PostEphemeral(ctx, channelID, message, ts, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostEphemeral", reflect.TypeOf((*MockSlackRepository)(nil).PostEphemeral), ctx, channelID, message, ts, userID)
}

// PostMessage mocks base method.
func (m *MockSlackRepository) PostMessage(ctx context.Context, channelID, message, ts string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMessage", ctx, channelID, message, ts)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockSlackRepositoryMockRecorder) PostMessage(ctx, channelID, message, ts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockSlackRepository)(nil).PostMessage), ctx, channelID, message, ts)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\rgoyal\GolandProjects\MinorProjectCli\LocalEyes\internal\interfaces\userRepoInterface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "localEyes/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// DeleteByUId mocks base method.
func (m *MockUserRepository) DeleteByUId(UId primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUId", UId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUId indicates an expected call of DeleteByUId.
func (mr *MockUserRepositoryMockRecorder) DeleteByUId(UId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUId", reflect.TypeOf((*MockUserRepository)(nil).DeleteByUId), UId)
}

// FindAdminByUsernamePassword mocks base method.
func (m *MockUserRepository) FindAdminByUsernamePassword(username, password string) (*models.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAdminByUsernamePassword", username, password)
	ret0, _ := ret[0].(*models.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAdminByUsernamePassword indicates an expected call of FindAdminByUsernamePassword.
func (mr *MockUserRepositoryMockRecorder) FindAdminByUsernamePassword(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAdminByUsernamePassword", reflect.TypeOf((*MockUserRepository)(nil).FindAdminByUsernamePassword), username, password)
}

// FindByUId mocks base method.
func (m *MockUserRepository) FindByUId(UId primitive.ObjectID) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUId", UId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUId indicates an expected call of FindByUId.
func (mr *MockUserRepositoryMockRecorder) FindByUId(UId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUId", reflect.TypeOf((*MockUserRepository)(nil).FindByUId), UId)
}

// FindByUsername mocks base method.
func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", username)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockUserRepositoryMockRecorder) FindByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockUserRepository)(nil).FindByUsername), username)
}

// FindByUsernamePassword mocks base method.
func (m *MockUserRepository) FindByUsernamePassword(username, password string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsernamePassword", username, password)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsernamePassword indicates an expected call of FindByUsernamePassword.
func (mr *MockUserRepositoryMockRecorder) FindByUsernamePassword(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsernamePassword", reflect.TypeOf((*MockUserRepository)(nil).FindByUsernamePassword), username, password)
}

// GetAllUsers mocks base method.
func (m *MockUserRepository) GetAllUsers() ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserRepositoryMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserRepository)(nil).GetAllUsers))
}

// UpdateActiveStatus mocks base method.
func (m *MockUserRepository) UpdateActiveStatus(UId primitive.ObjectID, status bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveStatus", UId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActiveStatus indicates an expected call of UpdateActiveStatus.
func (mr *MockUserRepositoryMockRecorder) UpdateActiveStatus(UId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveStatus", reflect.TypeOf((*MockUserRepository)(nil).UpdateActiveStatus), UId, status)
}

//Push Notification
func (m *MockUserRepository)PushNotification(UId primitive.ObjectID, title string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushNotification", UId, title)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActiveStatus indicates an expected call of UpdateActiveStatus.
func (mr *MockUserRepositoryMockRecorder)PushNotification(UId, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushNotification", reflect.TypeOf((*MockUserRepository)(nil).PushNotification), UId, title)
}

//Undo notification
func (m *MockUserRepository)ClearNotification(UId primitive.ObjectID) error{
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearNotification", UId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActiveStatus indicates an expected call of UpdateActiveStatus.
func (mr *MockUserRepositoryMockRecorder)ClearNotification(UId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearNotification", reflect.TypeOf((*MockUserRepository)(nil).ClearNotification), UId)
}

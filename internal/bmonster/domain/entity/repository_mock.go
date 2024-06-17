// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/bmonster/domain/entity/repository.go
//
// Generated by this command:
//
//	mockgen -source=./internal/bmonster/domain/entity/repository.go -destination=./internal/bmonster/domain/entity/repository_mock.go -package=entity
//

// Package entity is a generated GoMock package.
package entity

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPerformerRepository is a mock of PerformerRepository interface.
type MockPerformerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPerformerRepositoryMockRecorder
}

// MockPerformerRepositoryMockRecorder is the mock recorder for MockPerformerRepository.
type MockPerformerRepositoryMockRecorder struct {
	mock *MockPerformerRepository
}

// NewMockPerformerRepository creates a new mock instance.
func NewMockPerformerRepository(ctrl *gomock.Controller) *MockPerformerRepository {
	mock := &MockPerformerRepository{ctrl: ctrl}
	mock.recorder = &MockPerformerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerformerRepository) EXPECT() *MockPerformerRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockPerformerRepository) All() ([]Performer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]Performer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockPerformerRepositoryMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockPerformerRepository)(nil).All))
}

// MockScheduleRepository is a mock of ScheduleRepository interface.
type MockScheduleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleRepositoryMockRecorder
}

// MockScheduleRepositoryMockRecorder is the mock recorder for MockScheduleRepository.
type MockScheduleRepositoryMockRecorder struct {
	mock *MockScheduleRepository
}

// NewMockScheduleRepository creates a new mock instance.
func NewMockScheduleRepository(ctrl *gomock.Controller) *MockScheduleRepository {
	mock := &MockScheduleRepository{ctrl: ctrl}
	mock.recorder = &MockScheduleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleRepository) EXPECT() *MockScheduleRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockScheduleRepository) All() ([]Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockScheduleRepositoryMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockScheduleRepository)(nil).All))
}

// FilterByPerformer mocks base method.
func (m *MockScheduleRepository) FilterByPerformer(performer Performer) ([]Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterByPerformer", performer)
	ret0, _ := ret[0].([]Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterByPerformer indicates an expected call of FilterByPerformer.
func (mr *MockScheduleRepositoryMockRecorder) FilterByPerformer(performer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByPerformer", reflect.TypeOf((*MockScheduleRepository)(nil).FilterByPerformer), performer)
}

// Save mocks base method.
func (m *MockScheduleRepository) Save(schedules []Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", schedules)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockScheduleRepositoryMockRecorder) Save(schedules any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockScheduleRepository)(nil).Save), schedules)
}

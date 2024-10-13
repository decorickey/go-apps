// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=mock_repository.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	entity "github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockScrapingRepository is a mock of ScrapingRepository interface.
type MockScrapingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockScrapingRepositoryMockRecorder
}

// MockScrapingRepositoryMockRecorder is the mock recorder for MockScrapingRepository.
type MockScrapingRepositoryMockRecorder struct {
	mock *MockScrapingRepository
}

// NewMockScrapingRepository creates a new mock instance.
func NewMockScrapingRepository(ctrl *gomock.Controller) *MockScrapingRepository {
	mock := &MockScrapingRepository{ctrl: ctrl}
	mock.recorder = &MockScrapingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScrapingRepository) EXPECT() *MockScrapingRepositoryMockRecorder {
	return m.recorder
}

// FetchSchedulesByStudios mocks base method.
func (m *MockScrapingRepository) FetchSchedulesByStudios(arg0 []entity.Studio) ([]entity.Performer, []entity.Program, []entity.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSchedulesByStudios", arg0)
	ret0, _ := ret[0].([]entity.Performer)
	ret1, _ := ret[1].([]entity.Program)
	ret2, _ := ret[2].([]entity.Schedule)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// FetchSchedulesByStudios indicates an expected call of FetchSchedulesByStudios.
func (mr *MockScrapingRepositoryMockRecorder) FetchSchedulesByStudios(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSchedulesByStudios", reflect.TypeOf((*MockScrapingRepository)(nil).FetchSchedulesByStudios), arg0)
}

// FetchStudios mocks base method.
func (m *MockScrapingRepository) FetchStudios() ([]entity.Studio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStudios")
	ret0, _ := ret[0].([]entity.Studio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchStudios indicates an expected call of FetchStudios.
func (mr *MockScrapingRepositoryMockRecorder) FetchStudios() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStudios", reflect.TypeOf((*MockScrapingRepository)(nil).FetchStudios))
}

// MockStudioRepository is a mock of StudioRepository interface.
type MockStudioRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStudioRepositoryMockRecorder
}

// MockStudioRepositoryMockRecorder is the mock recorder for MockStudioRepository.
type MockStudioRepositoryMockRecorder struct {
	mock *MockStudioRepository
}

// NewMockStudioRepository creates a new mock instance.
func NewMockStudioRepository(ctrl *gomock.Controller) *MockStudioRepository {
	mock := &MockStudioRepository{ctrl: ctrl}
	mock.recorder = &MockStudioRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudioRepository) EXPECT() *MockStudioRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockStudioRepository) Save(arg0 []entity.Studio) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStudioRepositoryMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStudioRepository)(nil).Save), arg0)
}

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

// Save mocks base method.
func (m *MockPerformerRepository) Save(arg0 []entity.Performer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockPerformerRepositoryMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPerformerRepository)(nil).Save), arg0)
}

// MockProgramRepository is a mock of ProgramRepository interface.
type MockProgramRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProgramRepositoryMockRecorder
}

// MockProgramRepositoryMockRecorder is the mock recorder for MockProgramRepository.
type MockProgramRepositoryMockRecorder struct {
	mock *MockProgramRepository
}

// NewMockProgramRepository creates a new mock instance.
func NewMockProgramRepository(ctrl *gomock.Controller) *MockProgramRepository {
	mock := &MockProgramRepository{ctrl: ctrl}
	mock.recorder = &MockProgramRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProgramRepository) EXPECT() *MockProgramRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockProgramRepository) Save(arg0 []entity.Program) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockProgramRepositoryMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProgramRepository)(nil).Save), arg0)
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

// Save mocks base method.
func (m *MockScheduleRepository) Save(arg0 []entity.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockScheduleRepositoryMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockScheduleRepository)(nil).Save), arg0)
}

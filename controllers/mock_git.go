// Code generated by MockGen. DO NOT EDIT.
// Source: git.go

// Package controllers is a generated GoMock package.
package controllers

import (
	reflect "reflect"

	v5 "github.com/go-git/go-git/v5"
	config "github.com/go-git/go-git/v5/config"
	plumbing "github.com/go-git/go-git/v5/plumbing"
	gomock "github.com/golang/mock/gomock"
)

// MockRemoteClient is a mock of RemoteClient interface.
type MockRemoteClient struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteClientMockRecorder
}

// MockRemoteClientMockRecorder is the mock recorder for MockRemoteClient.
type MockRemoteClientMockRecorder struct {
	mock *MockRemoteClient
}

// NewMockRemoteClient creates a new mock instance.
func NewMockRemoteClient(ctrl *gomock.Controller) *MockRemoteClient {
	mock := &MockRemoteClient{ctrl: ctrl}
	mock.recorder = &MockRemoteClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteClient) EXPECT() *MockRemoteClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRemoteClient) List(o *v5.ListOptions) ([]*plumbing.Reference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", o)
	ret0, _ := ret[0].([]*plumbing.Reference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRemoteClientMockRecorder) List(o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRemoteClient)(nil).List), o)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// NewRemoteClient mocks base method.
func (m *MockClient) NewRemoteClient(c *config.RemoteConfig) RemoteClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoteClient", c)
	ret0, _ := ret[0].(RemoteClient)
	return ret0
}

// NewRemoteClient indicates an expected call of NewRemoteClient.
func (mr *MockClientMockRecorder) NewRemoteClient(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoteClient", reflect.TypeOf((*MockClient)(nil).NewRemoteClient), c)
}

// MockDriftWatcher is a mock of DriftWatcher interface.
type MockDriftWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDriftWatcherMockRecorder
}

// MockDriftWatcherMockRecorder is the mock recorder for MockDriftWatcher.
type MockDriftWatcherMockRecorder struct {
	mock *MockDriftWatcher
}

// NewMockDriftWatcher creates a new mock instance.
func NewMockDriftWatcher(ctrl *gomock.Controller) *MockDriftWatcher {
	mock := &MockDriftWatcher{ctrl: ctrl}
	mock.recorder = &MockDriftWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriftWatcher) EXPECT() *MockDriftWatcherMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockDriftWatcher) Add(name, namespace, origin, target string, interval int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", name, namespace, origin, target, interval)
}

// Add indicates an expected call of Add.
func (mr *MockDriftWatcherMockRecorder) Add(name, namespace, origin, target, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDriftWatcher)(nil).Add), name, namespace, origin, target, interval)
}

// Remove mocks base method.
func (m *MockDriftWatcher) Remove(name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockDriftWatcherMockRecorder) Remove(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDriftWatcher)(nil).Remove), name, namespace)
}

// Watch mocks base method.
func (m *MockDriftWatcher) Watch() chan<- interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(chan<- interface{})
	return ret0
}

// Watch indicates an expected call of Watch.
func (mr *MockDriftWatcherMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDriftWatcher)(nil).Watch))
}

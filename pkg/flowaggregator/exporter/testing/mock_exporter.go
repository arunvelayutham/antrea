// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/flowaggregator/exporter (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/flowaggregator/exporter/testing/mock_exporter.go -package testing antrea.io/antrea/pkg/flowaggregator/exporter Interface
//
// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	options "antrea.io/antrea/pkg/flowaggregator/options"
	entities "github.com/vmware/go-ipfix/pkg/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AddRecord mocks base method.
func (m *MockInterface) AddRecord(arg0 entities.Record, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRecord indicates an expected call of AddRecord.
func (mr *MockInterfaceMockRecorder) AddRecord(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecord", reflect.TypeOf((*MockInterface)(nil).AddRecord), arg0, arg1)
}

// Start mocks base method.
func (m *MockInterface) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockInterfaceMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInterface)(nil).Start))
}

// Stop mocks base method.
func (m *MockInterface) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockInterfaceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockInterface)(nil).Stop))
}

// UpdateOptions mocks base method.
func (m *MockInterface) UpdateOptions(arg0 *options.Options) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateOptions", arg0)
}

// UpdateOptions indicates an expected call of UpdateOptions.
func (mr *MockInterfaceMockRecorder) UpdateOptions(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOptions", reflect.TypeOf((*MockInterface)(nil).UpdateOptions), arg0)
}

// Code generated by mockery v2.9.4. DO NOT EDIT.

package ksqldb

import mock "github.com/stretchr/testify/mock"

// RespUnmarshaller is an autogenerated mock type for the RespUnmarshaller type
type RespUnmarshaller struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *RespUnmarshaller) Execute(_a0 []byte, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
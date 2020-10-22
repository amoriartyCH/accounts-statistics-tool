// Code generated by MockGen. DO NOT EDIT.
// Source: db/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/amoriartyCH/accounts-statistics-tool/models"
	gomock "github.com/golang/mock/gomock"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	reflect "reflect"
)

// MockTransactionClient is a mock of TransactionClient interface
type MockTransactionClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionClientMockRecorder
}

// MockTransactionClientMockRecorder is the mock recorder for MockTransactionClient
type MockTransactionClientMockRecorder struct {
	mock *MockTransactionClient
}

// NewMockTransactionClient creates a new mock instance
func NewMockTransactionClient(ctrl *gomock.Controller) *MockTransactionClient {
	mock := &MockTransactionClient{ctrl: ctrl}
	mock.recorder = &MockTransactionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionClient) EXPECT() *MockTransactionClientMockRecorder {
	return m.recorder
}

// GetAccountsTransactions mocks base method
func (m *MockTransactionClient) GetAccountsTransactions(dataDescription string) (*[]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsTransactions", dataDescription)
	ret0, _ := ret[0].(*[]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsTransactions indicates an expected call of GetAccountsTransactions
func (mr *MockTransactionClientMockRecorder) GetAccountsTransactions(dataDescription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsTransactions", reflect.TypeOf((*MockTransactionClient)(nil).GetAccountsTransactions), dataDescription)
}

// Shutdown mocks base method
func (m *MockTransactionClient) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockTransactionClientMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockTransactionClient)(nil).Shutdown))
}

// MockMongoDatabaseInterface is a mock of MongoDatabaseInterface interface
type MockMongoDatabaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDatabaseInterfaceMockRecorder
}

// MockMongoDatabaseInterfaceMockRecorder is the mock recorder for MockMongoDatabaseInterface
type MockMongoDatabaseInterfaceMockRecorder struct {
	mock *MockMongoDatabaseInterface
}

// NewMockMongoDatabaseInterface creates a new mock instance
func NewMockMongoDatabaseInterface(ctrl *gomock.Controller) *MockMongoDatabaseInterface {
	mock := &MockMongoDatabaseInterface{ctrl: ctrl}
	mock.recorder = &MockMongoDatabaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoDatabaseInterface) EXPECT() *MockMongoDatabaseInterfaceMockRecorder {
	return m.recorder
}

// Collection mocks base method
func (m *MockMongoDatabaseInterface) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Collection", varargs...)
	ret0, _ := ret[0].(*mongo.Collection)
	return ret0
}

// Collection indicates an expected call of Collection
func (mr *MockMongoDatabaseInterfaceMockRecorder) Collection(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockMongoDatabaseInterface)(nil).Collection), varargs...)
}
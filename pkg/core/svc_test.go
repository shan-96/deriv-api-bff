package core

import (
	"testing"

	"github.com/ksysoev/wasabi"
	"github.com/ksysoev/wasabi/mocks"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	mockCallsRepo := NewMockCallsRepo(t)
	mockDerivAPI := NewMockDerivAPI(t)
	mockConnRegistry := NewMockConnRegistry(t)

	svc := NewService(mockCallsRepo, mockDerivAPI, mockConnRegistry)

	assert.NotNil(t, svc)
	assert.Equal(t, mockCallsRepo, svc.ch)
	assert.Equal(t, mockDerivAPI, svc.be)
	assert.Equal(t, mockConnRegistry, svc.registry)
}

func TestService_PassThrough(t *testing.T) {
	mockCallsRepo := NewMockCallsRepo(t)
	mockDerivAPI := NewMockDerivAPI(t)
	mockConnRegistry := NewMockConnRegistry(t)

	svc := NewService(mockCallsRepo, mockDerivAPI, mockConnRegistry)

	mockConn := mocks.NewMockConnection(t)
	mockRequest := &Request{}

	conn := NewConnection(mockConn, func(_ string) {})

	mockConnRegistry.EXPECT().GetConnection(mockConn).Return(conn)
	mockDerivAPI.EXPECT().Handle(conn, mockRequest).Return(nil)

	err := svc.PassThrough(mockConn, mockRequest)

	assert.Nil(t, err)
}

func TestService_ProcessRequest(t *testing.T) {
	mockCallsRepo := NewMockCallsRepo(t)
	mockDerivAPI := NewMockDerivAPI(t)
	mockConnRegistry := NewMockConnRegistry(t)

	svc := NewService(mockCallsRepo, mockDerivAPI, mockConnRegistry)

	reqID := 1
	mockConn := mocks.NewMockConnection(t)
	mockRequest := &Request{
		Method: "testMethod",
		Params: map[string]any{"key": "value"},
		ID:     &reqID,
		data:   []byte(`{"req_id":1,"method":"testMethod","params":{"key":"value"}}`),
	}

	expectedResp := []byte(`{"echo":{"req_id":1,"method":"testMethod","params":{"key":"value"}},"req_id":1,"result":"success"}`)

	conn := NewConnection(mockConn, func(_ string) {})

	mockConnRegistry.EXPECT().GetConnection(mockConn).Return(conn)

	mockHandler := NewMockHandler(t)
	mockCallsRepo.EXPECT().GetCall("testMethod").Return(mockHandler)

	mockHandler.EXPECT().Handle(
		mock.Anything,
		mockRequest.Params,
		mock.Anything,
		mock.Anything,
	).Return(map[string]any{"result": "success"}, nil)

	mockConn.EXPECT().
		Send(wasabi.MsgTypeText, expectedResp).
		Return(nil)

	err := svc.ProcessRequest(mockConn, mockRequest)
	assert.Nil(t, err)
}

func TestService_ProcessRequest_UnsupportedMethod(t *testing.T) {
	mockCallsRepo := NewMockCallsRepo(t)
	mockDerivAPI := NewMockDerivAPI(t)
	mockConnRegistry := NewMockConnRegistry(t)

	svc := NewService(mockCallsRepo, mockDerivAPI, mockConnRegistry)

	mockConn := mocks.NewMockConnection(t)
	mockRequest := &Request{
		Method: "unsupportedMethod",
	}

	conn := NewConnection(mockConn, func(_ string) {})

	mockConnRegistry.EXPECT().GetConnection(mockConn).Return(conn)
	mockCallsRepo.EXPECT().GetCall("unsupportedMethod").Return(nil)

	err := svc.ProcessRequest(mockConn, mockRequest)
	assert.NotNil(t, err)
	assert.Equal(t, "unsupported method: unsupportedMethod", err.Error())
}

func TestService_ProcessRequest_HandlerError(t *testing.T) {
	mockCallsRepo := NewMockCallsRepo(t)
	mockDerivAPI := NewMockDerivAPI(t)
	mockConnRegistry := NewMockConnRegistry(t)

	svc := NewService(mockCallsRepo, mockDerivAPI, mockConnRegistry)

	mockConn := mocks.NewMockConnection(t)
	mockRequest := &Request{
		Method: "testMethod",
		Params: map[string]any{"key": "value"},
	}

	conn := NewConnection(mockConn, func(_ string) {})

	mockConnRegistry.EXPECT().GetConnection(mockConn).Return(conn)

	mockHandler := NewMockHandler(t)
	mockCallsRepo.EXPECT().GetCall("testMethod").Return(mockHandler)

	mockHandler.EXPECT().Handle(
		mock.Anything,
		mockRequest.Params,
		mock.Anything,
		mock.Anything,
	).Return(nil, assert.AnError)

	err := svc.ProcessRequest(mockConn, mockRequest)
	assert.ErrorIs(t, err, assert.AnError)
}

func TestService_ProcessRequest_APIError(t *testing.T) {
	mockCallsRepo := NewMockCallsRepo(t)
	mockDerivAPI := NewMockDerivAPI(t)
	mockConnRegistry := NewMockConnRegistry(t)

	svc := NewService(mockCallsRepo, mockDerivAPI, mockConnRegistry)

	mockConn := mocks.NewMockConnection(t)
	mockRequest := &Request{
		Method: "testMethod",
		Params: map[string]any{"key": "value"},
		data:   []byte(`{"method":"testMethod","params":{"key":"value"}}`),
	}

	expectedResp := []byte(`{"echo":{"method":"testMethod","params":{"key":"value"}},"error":{"code":"BadRequest","message":"Bad Request"}}`)

	conn := NewConnection(mockConn, func(_ string) {})

	mockConnRegistry.EXPECT().GetConnection(mockConn).Return(conn)

	mockHandler := NewMockHandler(t)
	mockCallsRepo.EXPECT().GetCall("testMethod").Return(mockHandler)

	apiErr := &APIError{Code: "BadRequest", Message: "Bad Request"}
	mockHandler.EXPECT().Handle(
		mock.Anything,
		mockRequest.Params,
		mock.Anything,
		mock.Anything,
	).Return(nil, apiErr)

	mockConn.EXPECT().Send(wasabi.MsgTypeText, expectedResp).Return(nil)

	err := svc.ProcessRequest(mockConn, mockRequest)
	assert.Nil(t, err)
}

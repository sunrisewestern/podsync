// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

package handler

import (
	gomock "github.com/golang/mock/gomock"
	patreon_go "github.com/mxpv/patreon-go"
	podcast "github.com/mxpv/podcast"
	api "github.com/mxpv/podsync/pkg/api"
	reflect "reflect"
)

// MockfeedService is a mock of feedService interface
type MockfeedService struct {
	ctrl     *gomock.Controller
	recorder *MockfeedServiceMockRecorder
}

// MockfeedServiceMockRecorder is the mock recorder for MockfeedService
type MockfeedServiceMockRecorder struct {
	mock *MockfeedService
}

// NewMockfeedService creates a new mock instance
func NewMockfeedService(ctrl *gomock.Controller) *MockfeedService {
	mock := &MockfeedService{ctrl: ctrl}
	mock.recorder = &MockfeedServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockfeedService) EXPECT() *MockfeedServiceMockRecorder {
	return _m.recorder
}

// CreateFeed mocks base method
func (_m *MockfeedService) CreateFeed(req *api.CreateFeedRequest, identity *api.Identity) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateFeed", req, identity)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeed indicates an expected call of CreateFeed
func (_mr *MockfeedServiceMockRecorder) CreateFeed(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateFeed", reflect.TypeOf((*MockfeedService)(nil).CreateFeed), arg0, arg1)
}

// BuildFeed mocks base method
func (_m *MockfeedService) BuildFeed(hashID string) (*podcast.Podcast, error) {
	ret := _m.ctrl.Call(_m, "BuildFeed", hashID)
	ret0, _ := ret[0].(*podcast.Podcast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildFeed indicates an expected call of BuildFeed
func (_mr *MockfeedServiceMockRecorder) BuildFeed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BuildFeed", reflect.TypeOf((*MockfeedService)(nil).BuildFeed), arg0)
}

// GetMetadata mocks base method
func (_m *MockfeedService) GetMetadata(hashId string) (*api.Metadata, error) {
	ret := _m.ctrl.Call(_m, "GetMetadata", hashId)
	ret0, _ := ret[0].(*api.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata
func (_mr *MockfeedServiceMockRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetMetadata", reflect.TypeOf((*MockfeedService)(nil).GetMetadata), arg0)
}

// Downgrade mocks base method
func (_m *MockfeedService) Downgrade(patronID string, featureLevel int) error {
	ret := _m.ctrl.Call(_m, "Downgrade", patronID, featureLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Downgrade indicates an expected call of Downgrade
func (_mr *MockfeedServiceMockRecorder) Downgrade(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Downgrade", reflect.TypeOf((*MockfeedService)(nil).Downgrade), arg0, arg1)
}

// MockpatreonService is a mock of patreonService interface
type MockpatreonService struct {
	ctrl     *gomock.Controller
	recorder *MockpatreonServiceMockRecorder
}

// MockpatreonServiceMockRecorder is the mock recorder for MockpatreonService
type MockpatreonServiceMockRecorder struct {
	mock *MockpatreonService
}

// NewMockpatreonService creates a new mock instance
func NewMockpatreonService(ctrl *gomock.Controller) *MockpatreonService {
	mock := &MockpatreonService{ctrl: ctrl}
	mock.recorder = &MockpatreonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockpatreonService) EXPECT() *MockpatreonServiceMockRecorder {
	return _m.recorder
}

// Hook mocks base method
func (_m *MockpatreonService) Hook(pledge *patreon_go.Pledge, event string) error {
	ret := _m.ctrl.Call(_m, "Hook", pledge, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Hook indicates an expected call of Hook
func (_mr *MockpatreonServiceMockRecorder) Hook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Hook", reflect.TypeOf((*MockpatreonService)(nil).Hook), arg0, arg1)
}

// GetFeatureLevelByID mocks base method
func (_m *MockpatreonService) GetFeatureLevelByID(patronID string) int {
	ret := _m.ctrl.Call(_m, "GetFeatureLevelByID", patronID)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetFeatureLevelByID indicates an expected call of GetFeatureLevelByID
func (_mr *MockpatreonServiceMockRecorder) GetFeatureLevelByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetFeatureLevelByID", reflect.TypeOf((*MockpatreonService)(nil).GetFeatureLevelByID), arg0)
}

// GetFeatureLevelFromAmount mocks base method
func (_m *MockpatreonService) GetFeatureLevelFromAmount(amount int) int {
	ret := _m.ctrl.Call(_m, "GetFeatureLevelFromAmount", amount)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetFeatureLevelFromAmount indicates an expected call of GetFeatureLevelFromAmount
func (_mr *MockpatreonServiceMockRecorder) GetFeatureLevelFromAmount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetFeatureLevelFromAmount", reflect.TypeOf((*MockpatreonService)(nil).GetFeatureLevelFromAmount), arg0)
}
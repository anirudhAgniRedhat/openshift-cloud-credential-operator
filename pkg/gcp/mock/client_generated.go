// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	iam "cloud.google.com/go/iam"
	storage "cloud.google.com/go/storage"
	context "context"
	gomock "github.com/golang/mock/gomock"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	iam0 "google.golang.org/api/iam/v1"
	admin "google.golang.org/genproto/googleapis/iam/admin/v1"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateServiceAccount mocks base method
func (m *MockClient) CreateServiceAccount(arg0 context.Context, arg1 *admin.CreateServiceAccountRequest) (*admin.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServiceAccount indicates an expected call of CreateServiceAccount
func (mr *MockClientMockRecorder) CreateServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceAccount", reflect.TypeOf((*MockClient)(nil).CreateServiceAccount), arg0, arg1)
}

// CreateServiceAccountKey mocks base method
func (m *MockClient) CreateServiceAccountKey(arg0 context.Context, arg1 *admin.CreateServiceAccountKeyRequest) (*admin.ServiceAccountKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceAccountKey", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServiceAccountKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServiceAccountKey indicates an expected call of CreateServiceAccountKey
func (mr *MockClientMockRecorder) CreateServiceAccountKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceAccountKey", reflect.TypeOf((*MockClient)(nil).CreateServiceAccountKey), arg0, arg1)
}

// DeleteServiceAccount mocks base method
func (m *MockClient) DeleteServiceAccount(arg0 context.Context, arg1 *admin.DeleteServiceAccountRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount
func (mr *MockClientMockRecorder) DeleteServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockClient)(nil).DeleteServiceAccount), arg0, arg1)
}

// DeleteServiceAccountKey mocks base method
func (m *MockClient) DeleteServiceAccountKey(arg0 context.Context, arg1 *admin.DeleteServiceAccountKeyRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccountKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccountKey indicates an expected call of DeleteServiceAccountKey
func (mr *MockClientMockRecorder) DeleteServiceAccountKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccountKey", reflect.TypeOf((*MockClient)(nil).DeleteServiceAccountKey), arg0, arg1)
}

// GetRole mocks base method
func (m *MockClient) GetRole(arg0 context.Context, arg1 *admin.GetRoleRequest) (*admin.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1)
	ret0, _ := ret[0].(*admin.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole
func (mr *MockClientMockRecorder) GetRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockClient)(nil).GetRole), arg0, arg1)
}

// GetServiceAccount mocks base method
func (m *MockClient) GetServiceAccount(arg0 context.Context, arg1 *admin.GetServiceAccountRequest) (*admin.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceAccount indicates an expected call of GetServiceAccount
func (mr *MockClientMockRecorder) GetServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceAccount", reflect.TypeOf((*MockClient)(nil).GetServiceAccount), arg0, arg1)
}

// ListServiceAccountKeys mocks base method
func (m *MockClient) ListServiceAccountKeys(arg0 context.Context, arg1 *admin.ListServiceAccountKeysRequest) (*admin.ListServiceAccountKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceAccountKeys", arg0, arg1)
	ret0, _ := ret[0].(*admin.ListServiceAccountKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServiceAccountKeys indicates an expected call of ListServiceAccountKeys
func (mr *MockClientMockRecorder) ListServiceAccountKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceAccountKeys", reflect.TypeOf((*MockClient)(nil).ListServiceAccountKeys), arg0, arg1)
}

// QueryTestablePermissions mocks base method
func (m *MockClient) QueryTestablePermissions(arg0 context.Context, arg1 *admin.QueryTestablePermissionsRequest) (*admin.QueryTestablePermissionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTestablePermissions", arg0, arg1)
	ret0, _ := ret[0].(*admin.QueryTestablePermissionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTestablePermissions indicates an expected call of QueryTestablePermissions
func (mr *MockClientMockRecorder) QueryTestablePermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTestablePermissions", reflect.TypeOf((*MockClient)(nil).QueryTestablePermissions), arg0, arg1)
}

// CreateWorkloadIdentityPool mocks base method
func (m *MockClient) CreateWorkloadIdentityPool(arg0 context.Context, arg1, arg2 string, arg3 *iam0.WorkloadIdentityPool) (*iam0.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkloadIdentityPool", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*iam0.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkloadIdentityPool indicates an expected call of CreateWorkloadIdentityPool
func (mr *MockClientMockRecorder) CreateWorkloadIdentityPool(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkloadIdentityPool", reflect.TypeOf((*MockClient)(nil).CreateWorkloadIdentityPool), arg0, arg1, arg2, arg3)
}

// CreateWorkloadIdentityProvider mocks base method
func (m *MockClient) CreateWorkloadIdentityProvider(arg0 context.Context, arg1, arg2 string, arg3 *iam0.WorkloadIdentityPoolProvider) (*iam0.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkloadIdentityProvider", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*iam0.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkloadIdentityProvider indicates an expected call of CreateWorkloadIdentityProvider
func (mr *MockClientMockRecorder) CreateWorkloadIdentityProvider(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkloadIdentityProvider", reflect.TypeOf((*MockClient)(nil).CreateWorkloadIdentityProvider), arg0, arg1, arg2, arg3)
}

// GetProjectIamPolicy mocks base method
func (m *MockClient) GetProjectIamPolicy(projectName string, request *cloudresourcemanager.GetIamPolicyRequest) (*cloudresourcemanager.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectIamPolicy", projectName, request)
	ret0, _ := ret[0].(*cloudresourcemanager.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectIamPolicy indicates an expected call of GetProjectIamPolicy
func (mr *MockClientMockRecorder) GetProjectIamPolicy(projectName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectIamPolicy", reflect.TypeOf((*MockClient)(nil).GetProjectIamPolicy), projectName, request)
}

// GetProjectName mocks base method
func (m *MockClient) GetProjectName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProjectName indicates an expected call of GetProjectName
func (mr *MockClientMockRecorder) GetProjectName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectName", reflect.TypeOf((*MockClient)(nil).GetProjectName))
}

// SetProjectIamPolicy mocks base method
func (m *MockClient) SetProjectIamPolicy(projectName string, request *cloudresourcemanager.SetIamPolicyRequest) (*cloudresourcemanager.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProjectIamPolicy", projectName, request)
	ret0, _ := ret[0].(*cloudresourcemanager.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetProjectIamPolicy indicates an expected call of SetProjectIamPolicy
func (mr *MockClientMockRecorder) SetProjectIamPolicy(projectName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProjectIamPolicy", reflect.TypeOf((*MockClient)(nil).SetProjectIamPolicy), projectName, request)
}

// TestIamPermissions mocks base method
func (m *MockClient) TestIamPermissions(arg0 string, arg1 *cloudresourcemanager.TestIamPermissionsRequest) (*cloudresourcemanager.TestIamPermissionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestIamPermissions", arg0, arg1)
	ret0, _ := ret[0].(*cloudresourcemanager.TestIamPermissionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TestIamPermissions indicates an expected call of TestIamPermissions
func (mr *MockClientMockRecorder) TestIamPermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestIamPermissions", reflect.TypeOf((*MockClient)(nil).TestIamPermissions), arg0, arg1)
}

// ListServicesEnabled mocks base method
func (m *MockClient) ListServicesEnabled() (map[string]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServicesEnabled")
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServicesEnabled indicates an expected call of ListServicesEnabled
func (mr *MockClientMockRecorder) ListServicesEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServicesEnabled", reflect.TypeOf((*MockClient)(nil).ListServicesEnabled))
}

// CreateBucket mocks base method
func (m *MockClient) CreateBucket(arg0 context.Context, arg1, arg2 string, arg3 *storage.BucketAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockClientMockRecorder) CreateBucket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockClient)(nil).CreateBucket), arg0, arg1, arg2, arg3)
}

// GetBucketPolicy mocks base method
func (m *MockClient) GetBucketPolicy(arg0 context.Context, arg1 string) (*iam.Policy3, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketPolicy", arg0, arg1)
	ret0, _ := ret[0].(*iam.Policy3)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketPolicy indicates an expected call of GetBucketPolicy
func (mr *MockClientMockRecorder) GetBucketPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketPolicy", reflect.TypeOf((*MockClient)(nil).GetBucketPolicy), arg0, arg1)
}

// SetBucketPolicy mocks base method
func (m *MockClient) SetBucketPolicy(arg0 context.Context, arg1 string, arg2 *iam.Policy3) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBucketPolicy", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBucketPolicy indicates an expected call of SetBucketPolicy
func (mr *MockClientMockRecorder) SetBucketPolicy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBucketPolicy", reflect.TypeOf((*MockClient)(nil).SetBucketPolicy), arg0, arg1, arg2)
}

// PutObject mocks base method
func (m *MockClient) PutObject(arg0 context.Context, arg1, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject
func (mr *MockClientMockRecorder) PutObject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockClient)(nil).PutObject), arg0, arg1, arg2, arg3)
}

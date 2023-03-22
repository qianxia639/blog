// Code generated by MockGen. DO NOT EDIT.
// Source: Blog/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "Blog/db/sqlc"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteBlog mocks base method.
func (m *MockStore) DeleteBlog(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlog indicates an expected call of DeleteBlog.
func (mr *MockStoreMockRecorder) DeleteBlog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlog", reflect.TypeOf((*MockStore)(nil).DeleteBlog), arg0, arg1)
}

// GetBlog mocks base method.
func (m *MockStore) GetBlog(arg0 context.Context, arg1 int64) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlog", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlog indicates an expected call of GetBlog.
func (mr *MockStoreMockRecorder) GetBlog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlog", reflect.TypeOf((*MockStore)(nil).GetBlog), arg0, arg1)
}

// GetChildComments mocks base method.
func (m *MockStore) GetChildComments(arg0 context.Context, arg1 db.GetChildCommentsParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildComments", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChildComments indicates an expected call of GetChildComments.
func (mr *MockStoreMockRecorder) GetChildComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildComments", reflect.TypeOf((*MockStore)(nil).GetChildComments), arg0, arg1)
}

// GetComments mocks base method.
func (m *MockStore) GetComments(arg0 context.Context, arg1 int64) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockStoreMockRecorder) GetComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockStore)(nil).GetComments), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// IncrViews mocks base method.
func (m *MockStore) IncrViews(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrViews", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrViews indicates an expected call of IncrViews.
func (mr *MockStoreMockRecorder) IncrViews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrViews", reflect.TypeOf((*MockStore)(nil).IncrViews), arg0, arg1)
}

// InsertBlog mocks base method.
func (m *MockStore) InsertBlog(arg0 context.Context, arg1 db.InsertBlogParams) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlog", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBlog indicates an expected call of InsertBlog.
func (mr *MockStoreMockRecorder) InsertBlog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlog", reflect.TypeOf((*MockStore)(nil).InsertBlog), arg0, arg1)
}

// InsertRequestLog mocks base method.
func (m *MockStore) InsertRequestLog(arg0 context.Context, arg1 db.InsertRequestLogParams) (db.RequestLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRequestLog", arg0, arg1)
	ret0, _ := ret[0].(db.RequestLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertRequestLog indicates an expected call of InsertRequestLog.
func (mr *MockStoreMockRecorder) InsertRequestLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRequestLog", reflect.TypeOf((*MockStore)(nil).InsertRequestLog), arg0, arg1)
}

// InsertType mocks base method.
func (m *MockStore) InsertType(arg0 context.Context, arg1 string) (db.Type, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertType", arg0, arg1)
	ret0, _ := ret[0].(db.Type)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertType indicates an expected call of InsertType.
func (mr *MockStoreMockRecorder) InsertType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertType", reflect.TypeOf((*MockStore)(nil).InsertType), arg0, arg1)
}

// ListBlogs mocks base method.
func (m *MockStore) ListBlogs(arg0 context.Context, arg1 db.ListBlogsParams) ([]db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlogs", arg0, arg1)
	ret0, _ := ret[0].([]db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlogs indicates an expected call of ListBlogs.
func (mr *MockStoreMockRecorder) ListBlogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlogs", reflect.TypeOf((*MockStore)(nil).ListBlogs), arg0, arg1)
}

// SearchBlog mocks base method.
func (m *MockStore) SearchBlog(arg0 context.Context, arg1 string) ([]db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchBlog", arg0, arg1)
	ret0, _ := ret[0].([]db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchBlog indicates an expected call of SearchBlog.
func (mr *MockStoreMockRecorder) SearchBlog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchBlog", reflect.TypeOf((*MockStore)(nil).SearchBlog), arg0, arg1)
}

// UpdateBlog mocks base method.
func (m *MockStore) UpdateBlog(arg0 context.Context, arg1 db.UpdateBlogParams) (db.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBlog", arg0, arg1)
	ret0, _ := ret[0].(db.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBlog indicates an expected call of UpdateBlog.
func (mr *MockStoreMockRecorder) UpdateBlog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlog", reflect.TypeOf((*MockStore)(nil).UpdateBlog), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

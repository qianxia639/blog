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

// CountArticle mocks base method.
func (m *MockStore) CountArticle(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountArticle", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountArticle indicates an expected call of CountArticle.
func (mr *MockStoreMockRecorder) CountArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountArticle", reflect.TypeOf((*MockStore)(nil).CountArticle), arg0)
}

// CreateCritique mocks base method.
func (m *MockStore) CreateCritique(arg0 context.Context, arg1 *db.CreateCritiqueParams) (db.Critique, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCritique", arg0, arg1)
	ret0, _ := ret[0].(db.Critique)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCritique indicates an expected call of CreateCritique.
func (mr *MockStoreMockRecorder) CreateCritique(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCritique", reflect.TypeOf((*MockStore)(nil).CreateCritique), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 *db.CreateUserParams) (db.User, error) {
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

// DeleteArticle mocks base method.
func (m *MockStore) DeleteArticle(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockStoreMockRecorder) DeleteArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockStore)(nil).DeleteArticle), arg0, arg1)
}

// GetArticle mocks base method.
func (m *MockStore) GetArticle(arg0 context.Context, arg1 int64) (db.GetArticleRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticle", arg0, arg1)
	ret0, _ := ret[0].(db.GetArticleRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockStoreMockRecorder) GetArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockStore)(nil).GetArticle), arg0, arg1)
}

// GetChildCritiques mocks base method.
func (m *MockStore) GetChildCritiques(arg0 context.Context, arg1 int64) ([]db.Critique, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildCritiques", arg0, arg1)
	ret0, _ := ret[0].([]db.Critique)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChildCritiques indicates an expected call of GetChildCritiques.
func (mr *MockStoreMockRecorder) GetChildCritiques(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildCritiques", reflect.TypeOf((*MockStore)(nil).GetChildCritiques), arg0, arg1)
}

// GetCritiques mocks base method.
func (m *MockStore) GetCritiques(arg0 context.Context, arg1 int64) ([]db.Critique, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCritiques", arg0, arg1)
	ret0, _ := ret[0].([]db.Critique)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCritiques indicates an expected call of GetCritiques.
func (mr *MockStoreMockRecorder) GetCritiques(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCritiques", reflect.TypeOf((*MockStore)(nil).GetCritiques), arg0, arg1)
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

// InsertArticle mocks base method.
func (m *MockStore) InsertArticle(arg0 context.Context, arg1 *db.InsertArticleParams) (db.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertArticle", arg0, arg1)
	ret0, _ := ret[0].(db.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertArticle indicates an expected call of InsertArticle.
func (mr *MockStoreMockRecorder) InsertArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertArticle", reflect.TypeOf((*MockStore)(nil).InsertArticle), arg0, arg1)
}

// ListArticles mocks base method.
func (m *MockStore) ListArticles(arg0 context.Context, arg1 *db.ListArticlesParams) ([]db.ListArticlesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListArticles", arg0, arg1)
	ret0, _ := ret[0].([]db.ListArticlesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArticles indicates an expected call of ListArticles.
func (mr *MockStoreMockRecorder) ListArticles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArticles", reflect.TypeOf((*MockStore)(nil).ListArticles), arg0, arg1)
}

// UpdateArticle mocks base method.
func (m *MockStore) UpdateArticle(arg0 context.Context, arg1 *db.UpdateArticleParams) (db.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", arg0, arg1)
	ret0, _ := ret[0].(db.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockStoreMockRecorder) UpdateArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockStore)(nil).UpdateArticle), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 *db.UpdateUserParams) (db.User, error) {
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

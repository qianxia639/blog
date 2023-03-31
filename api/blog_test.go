package api

import (
	mockdb "Blog/db/mock"
	db "Blog/db/sqlc"
	"Blog/token"
	"Blog/utils"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

type eqInsertBlogParamsMatcher struct {
	arg        db.InsertBlogParams
	created_at time.Time
}

func (e eqInsertBlogParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.InsertBlogParams)
	if !ok {
		return false
	}

	e.arg.CreatedAt = arg.CreatedAt

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqInsertBlogParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and created_at %v\n", e.arg, e.created_at)
}

func EqInsertBlogParams(arg db.InsertBlogParams, created_at time.Time) gomock.Matcher {
	return eqInsertBlogParamsMatcher{arg, created_at}
}

func TestInsertBlog(t *testing.T) {

	store := newTestDB(t)

	user, password := randomUser(t)

	arg := db.CreateUserParams{
		Username: user.Username,
		Password: password,
		Nickname: user.Nickname,
		Email:    user.Email,
	}

	user, err := store.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	title := utils.RandomString(6)
	content := fmt.Sprintf("%s-content", title)
	image := fmt.Sprintf("%s.jpg", title)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"owner_id": user.ID,
				"title":    title,
				"content":  content,
				"image":    image,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				createdAt := time.Now()
				arg := db.InsertBlogParams{
					OwnerID:   user.ID,
					Title:     title,
					Content:   content,
					Image:     image,
					CreatedAt: createdAt,
				}
				store.EXPECT().
					InsertBlog(gomock.Any(), EqInsertBlogParams(arg, createdAt)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "Internal Error",
			body: gin.H{
				"owner_id": user.ID,
				"title":    title,
				"content":  content,
				"image":    image,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				createdAt := time.Now()
				arg := db.InsertBlogParams{
					OwnerID:   user.ID,
					Title:     title,
					Content:   content,
					Image:     image,
					CreatedAt: createdAt,
				}
				store.EXPECT().
					InsertBlog(gomock.Any(), EqInsertBlogParams(arg, createdAt)).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "Duplicate Title",
			body: gin.H{
				"owner_id": user.ID,
				"title":    title,
				"content":  content,
				"image":    image,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertBlog(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Blog{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recoder.Code)
			},
		},
		{
			name: "Bad Request Error",
			body: gin.H{},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := "/blog"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.maker)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

func TestIncrViews(t *testing.T) {

	id := utils.RandomInt(1, 100)

	testCases := []struct {
		name          string
		blogId        int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1)
				store.EXPECT().
					IncrViews(gomock.Any(), gomock.Eq(id)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name:   "Not Found",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Blog{}, sql.ErrNoRows)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:   "Internal Error One",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "Internal Error Two",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(db.Blog{}, nil)
				store.EXPECT().
					IncrViews(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "Invalid ID",
			blogId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := fmt.Sprintf("/blog/incr/%d", tc.blogId)
			request, err := http.NewRequest(http.MethodPut, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

func TestDeleteBlog(t *testing.T) {

	id := utils.RandomInt(1, 100)

	testCases := []struct {
		name          string
		blogId        int64
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			blogId: id,
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, "user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1)
				store.EXPECT().
					DeleteBlog(gomock.Any(), gomock.Eq(id)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name:   "Not Found",
			blogId: id,
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, "user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Blog{}, sql.ErrNoRows)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:   "Internal Error One",
			blogId: id,
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, "user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "Internal Error Two",
			blogId: id,
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, "user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(db.Blog{}, nil)
				store.EXPECT().
					DeleteBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "Invalid ID",
			blogId: 0,
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, "user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := fmt.Sprintf("/blog/%d", tc.blogId)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.maker)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

type Page struct {
	PageNo   int32 `json:"page_no"`
	PageSize int32 `json:"page_size"`
}

func TestListBlog(t *testing.T) {

	page := Page{
		PageNo:   5,
		PageSize: 5,
	}

	testCases := []struct {
		name          string
		page          Page
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			page: page,
			buildStubs: func(store *mockdb.MockStore) {

				arg := db.ListBlogsParams{
					Limit:  page.PageSize,
					Offset: (page.PageNo - 1) * page.PageSize,
				}

				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Eq(arg)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "Internal Error",
			page: page,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListBlogsParams{
					Limit:  page.PageSize,
					Offset: (page.PageNo - 1) * page.PageSize,
				}
				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "Invalid Parameter",
			page: Page{
				PageNo:   0,
				PageSize: 0,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBlogs(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := fmt.Sprintf("/blog?page_no=%d&page_size=%d", tc.page.PageNo, tc.page.PageSize)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

func TestGetBlog(t *testing.T) {

	id := utils.RandomInt(1, 100)

	testCases := []struct {
		name          string
		blogId        int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name:   "Internal Error",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "Not Found",
			blogId: id,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Eq(id)).
					Times(1).
					Return(db.Blog{}, sql.ErrNoRows)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:   "Invalid Id",
			blogId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := fmt.Sprintf("/blog/%d", tc.blogId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

type eqUpdateBloggParamsMatcher struct {
	arg        db.UpdateBlogParams
	updated_at time.Time
}

func (e eqUpdateBloggParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.UpdateBlogParams)
	if !ok {
		return false
	}

	e.arg.UpdatedAt = arg.UpdatedAt

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqUpdateBloggParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and updated_at %v\n", e.arg, e.updated_at)
}

func EqUpdateBlogParams(arg db.UpdateBlogParams, updated_at time.Time) gomock.Matcher {
	return eqUpdateBloggParamsMatcher{arg, updated_at}
}

func TestUpdateBlog(t *testing.T) {

	store := newTestDB(t)

	user, password := randomUser(t)

	user, err := store.CreateUser(ctx, db.CreateUserParams{
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: password,
	})
	require.NoError(t, err)

	blog, err := store.InsertBlog(ctx, db.InsertBlogParams{
		OwnerID: user.ID,
		Title:   utils.RandomString(6),
		Content: utils.RandomString(50),
		Image:   fmt.Sprintf("%s.jpg", utils.RandomString(10)),
	})
	require.NoError(t, err)

	title := fmt.Sprintf("title-%s", time.Now().Format("2006-01-02 15:04:05"))
	content := fmt.Sprintf("content-%s", time.Now().Format("2006-01-02 15:04:05"))
	image := fmt.Sprintf("%s.jpg", time.Now().Format("2006-01-02 15:04:05"))

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "Internal Error",
			body: gin.H{
				"id": blog.ID,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateBlog(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Blog{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "Bad Request Error",
			body: gin.H{
				"id": 0,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "Duplicate Title",
			body: gin.H{
				"id":    blog.ID,
				"title": title,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateBlog(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Blog{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recoder.Code)
			},
		},
		{
			name: "Update Only Title",
			body: gin.H{
				"id":    blog.ID,
				"title": title,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				updatedAt := time.Now()
				arg := db.UpdateBlogParams{
					ID: blog.ID,
					Title: sql.NullString{
						String: title,
						Valid:  true,
					},
					UpdatedAt: updatedAt,
				}
				store.EXPECT().
					UpdateBlog(gomock.Any(), EqUpdateBlogParams(arg, updatedAt)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "Update Only Content",
			body: gin.H{
				"id":      blog.ID,
				"content": content,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				updatedAt := time.Now()
				arg := db.UpdateBlogParams{
					ID: blog.ID,
					Content: sql.NullString{
						String: content,
						Valid:  true,
					},
					UpdatedAt: updatedAt,
				}
				store.EXPECT().
					UpdateBlog(gomock.Any(), EqUpdateBlogParams(arg, updatedAt)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "Update Only Image",
			body: gin.H{
				"id":    blog.ID,
				"image": image,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				updatedAt := time.Now()
				arg := db.UpdateBlogParams{
					ID: blog.ID,
					Image: sql.NullString{
						String: image,
						Valid:  true,
					},
					UpdatedAt: updatedAt,
				}
				store.EXPECT().
					UpdateBlog(gomock.Any(), EqUpdateBlogParams(arg, updatedAt)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "Update Only All",
			body: gin.H{
				"id":      blog.ID,
				"title":   title,
				"image":   image,
				"content": content,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				updatedAt := time.Now()
				arg := db.UpdateBlogParams{
					ID: blog.ID,
					Title: sql.NullString{
						String: title,
						Valid:  true,
					},
					Content: sql.NullString{
						String: content,
						Valid:  true,
					},
					Image: sql.NullString{
						String: image,
						Valid:  true,
					},
					UpdatedAt: updatedAt,
				}
				store.EXPECT().
					UpdateBlog(gomock.Any(), EqUpdateBlogParams(arg, updatedAt)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := "/blog"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.maker)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

type Search struct {
	Title    string
	PageNo   int32
	PageSize int32
}

type eqSearchBloggParamsMatcher struct {
	arg   db.SearchBlogParams
	title string
}

func (e eqSearchBloggParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.SearchBlogParams)
	if !ok {
		return false
	}

	e.arg.Title = arg.Title

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqSearchBloggParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and titlee %v\n", e.arg, e.title)
}

func EqSearchBlogParams(arg db.SearchBlogParams, title string) gomock.Matcher {
	return eqSearchBloggParamsMatcher{arg, title}
}

func TestSearchBlog(t *testing.T) {

	search := Search{
		Title:    "s",
		PageNo:   1,
		PageSize: 5,
	}

	testCases := []struct {
		name          string
		search        Search
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			search: search,
			buildStubs: func(store *mockdb.MockStore) {
				title := "s"
				arg := db.SearchBlogParams{
					Title:  title,
					Limit:  search.PageSize,
					Offset: (search.PageNo - 1) * search.PageSize,
				}

				store.EXPECT().
					SearchBlog(gomock.Any(), EqSearchBlogParams(arg, title)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name:   "Internal Error",
			search: search,
			buildStubs: func(store *mockdb.MockStore) {
				title := "s"
				arg := db.SearchBlogParams{
					Title:  title,
					Limit:  search.PageSize,
					Offset: (search.PageNo - 1) * search.PageSize,
				}
				store.EXPECT().
					SearchBlog(gomock.Any(), EqSearchBlogParams(arg, title)).
					Times(1).
					Return([]db.SearchBlogRow{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "Invalid Parameter",
			search: Search{
				Title:    "s",
				PageNo:   0,
				PageSize: 0,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					SearchBlog(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recodre := httptest.NewRecorder()

			url := fmt.Sprintf("/blog/search?title=%s&page_no=%d&page_size=%d", tc.search.Title, tc.search.PageNo, tc.search.PageSize)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

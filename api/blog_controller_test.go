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
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

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

	typeName := fmt.Sprintf("%s-typeName", user.Username)

	ty, err := store.InsertType(context.Background(), typeName)
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
				"type_id":  ty.ID,
				"title":    title,
				"content":  content,
				"image":    image,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.InsertBlogParams{
					OwnerID: user.ID,
					TypeID:  ty.ID,
					Title:   title,
					Content: content,
					Image:   image,
				}
				store.EXPECT().
					InsertBlog(gomock.Any(), gomock.Eq(arg)).
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
				"type_id":  ty.ID,
				"title":    title,
				"content":  content,
				"image":    image,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.InsertBlogParams{
					OwnerID: user.ID,
					TypeID:  ty.ID,
					Title:   title,
					Content: content,
					Image:   image,
				}
				store.EXPECT().
					InsertBlog(gomock.Any(), gomock.Eq(arg)).
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
				"type_id":  ty.ID,
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

// func TestIncrViews(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:      "OK",
// 			body:      gin.H{},
// 			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetBlog(gomock.Any(), 0).
// 					Times(1).
// 					Return(db.Blog{}, nil)
// 				store.EXPECT().
// 					IncrViews(gomock.Any(), "").
// 					Times(1)
// 			},
// 			checkResponse: func(recoder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recoder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			server := newTestServer(t, store)
// 			recodre := httptest.NewRecorder()

// 			url := "/blog/incr"
// 			request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
// 			require.NoError(t, err)

// 			tc.setupAuth(t, request, server.maker)

// 			server.router.ServeHTTP(recodre, request)
// 			tc.checkResponse(recodre)
// 		})
// 	}
// }

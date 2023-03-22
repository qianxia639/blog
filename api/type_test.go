package api

import (
	mockdb "Blog/db/mock"
	db "Blog/db/sqlc"
	"Blog/token"
	"bytes"
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

func TestInsertType(t *testing.T) {

	user, _ := randomUser(t)

	typeName := fmt.Sprintf("%s-type", user.Username)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "Insert Success",
			body: gin.H{
				"type_name": typeName,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertType(gomock.Any(), gomock.Eq(typeName)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "Bad Request Error",
			body: gin.H{
				"type_name": "",
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertType(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "Internal Error",
			body: gin.H{
				"type_name": typeName,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertType(gomock.Any(), gomock.Eq(typeName)).
					Times(1).
					Return(db.Type{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "Duplicate TypeName",
			body: gin.H{
				"type_name": typeName,
			},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorizatin(t, req, tokenMaker, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					InsertType(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Type{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recoder.Code)
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

			url := "/type"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.maker)

			server.router.ServeHTTP(recodre, request)
			tc.checkResponse(recodre)
		})
	}
}

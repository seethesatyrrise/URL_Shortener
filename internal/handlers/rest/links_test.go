package rest

import (
	"URL_Shortener/internal/repo"
	"URL_Shortener/internal/service"
	mock_storage "URL_Shortener/internal/storage/mocks"
	"URL_Shortener/internal/utils"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetFullLink(t *testing.T) {
	type mockBehavior func(r *mock_storage.MockStorage, ctx context.Context, token string)
	tests := []struct {
		name                 string
		inputToken           string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "ok",
			inputToken: "rndomtoken",
			mockBehavior: func(s *mock_storage.MockStorage, ctx context.Context, token string) {
				s.EXPECT().GetLinkByToken(ctx, token).Return("full.link", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"link":"full.link"}`,
		},
		{
			name:       "no such token",
			inputToken: "NOLINKNOLI",
			mockBehavior: func(s *mock_storage.MockStorage, ctx context.Context, token string) {
				s.EXPECT().GetLinkByToken(ctx, token).Return("", utils.ErrNotFound)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `not found`,
		},
	}
	c := gomock.NewController(t)
	defer c.Finish()

	s := mock_storage.NewMockStorage(c)

	repo := repo.New(s)
	service := service.New(repo, s, utils.NewLogger())
	h := New(service)

	router := gin.Default()
	router.GET(":token", h.GetFullLink)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(s, context.Background(), tt.inputToken)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/"+tt.inputToken, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetToken(t *testing.T) {
	type mockBehavior func(r *mock_storage.MockStorage, ctx context.Context, link string, generatedToken string)
	tests := []struct {
		name                 string
		inputLink            string
		inputBody            string
		foundToken           string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "ok from storage",
			inputLink:  "test.link",
			inputBody:  `{"link":"test.link"}`,
			foundToken: "TESTTOKEN1",
			mockBehavior: func(s *mock_storage.MockStorage, ctx context.Context, link string, foundToken string) {
				s.EXPECT().TryGetTokenByLink(ctx, link).Return(foundToken, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"TESTTOKEN1"}`,
		},
		{
			name:       "bad req",
			inputLink:  "",
			inputBody:  `{"link":""}`,
			foundToken: "",
			mockBehavior: func(s *mock_storage.MockStorage, ctx context.Context, link string, foundToken string) {
				//s.EXPECT().TryGetTokenByLink(ctx, link).Return(foundToken, nil)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `bad request`,
		},
	}
	c := gomock.NewController(t)
	defer c.Finish()

	s := mock_storage.NewMockStorage(c)

	repo := repo.New(s)
	service := service.New(repo, s, utils.NewLogger())
	h := New(service)

	router := gin.Default()
	router.POST("/api/generate", h.GetToken)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(s, context.Background(), tt.inputLink, tt.foundToken)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/generate", bytes.NewBufferString(tt.inputBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}

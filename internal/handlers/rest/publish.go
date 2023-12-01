package rest

import (
	"URL_Shortener/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishError(c *gin.Context, err error) {
	code := getCode(err)
	c.Data(code, "text/html", []byte(err.Error()))
}

func PublishData(c *gin.Context, data []byte) {
	c.Data(http.StatusOK, "application/json", data)
}

func getCode(err error) int {
	switch err {
	case utils.ErrNotFound:
		return http.StatusNotFound
	case utils.ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

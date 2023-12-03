package rest

import (
	"URL_Shortener/internal/models"
	"URL_Shortener/internal/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFullLink(ctx *gin.Context) {
	var token string
	token = ctx.Param("token")

	fullLink, err := h.service.GetFullLink(ctx.Request.Context(), token)
	if err != nil {
		h.service.Logger.Error(err.Error())
		PublishError(ctx, err)
		return
	}

	data, err := json.Marshal(&models.Link{Link: fullLink})
	if err != nil {
		h.service.Logger.Error(err.Error())
		PublishError(ctx, err)
		return
	}

	PublishData(ctx, data)
	return
}

func (h *Handler) GetToken(ctx *gin.Context) {
	var link models.Link
	err := ctx.ShouldBind(&link)
	if err != nil {
		h.service.Logger.Error(err.Error())
		PublishError(ctx, utils.ErrBadRequest)
		return
	}

	if len(link.Link) < 1 {
		err = utils.ErrBadRequest
		h.service.Logger.Error(err.Error())
		PublishError(ctx, err)
		return
	}

	token, err := h.service.GetToken(ctx.Request.Context(), link.Link)
	if err != nil {
		h.service.Logger.Error(err.Error())
		PublishError(ctx, err)
		return
	}

	data, err := json.Marshal(&models.Token{Token: token})
	if err != nil {
		h.service.Logger.Error(err.Error())
		PublishError(ctx, err)
		return
	}

	PublishData(ctx, data)
	return
}

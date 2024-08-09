package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HandleRefreshToken	godoc
// @Summary      Làm mới access token
// @Description  Làm mới access token
// @Tags         App Token
// @Produce      json
// @Success      200 {object}  	common.OK
// @failure		 400 {object} 	common.Err
// @failure		 500 {object} 	common.Err
// @Router       /token/refresh [post]
// @Security     JWT
func (h *appTokenHandler) HandleRefreshToken(c *gin.Context) {
	var refreshToken = c.Request.Header["Authorization"][0]
	result, useCaseErr := h.refreshTokenUseCase.ExecRefreshToken(c.Request.Context(), strings.Split(refreshToken, " ")[1])
	if useCaseErr != nil {
		common.ResponseErr(c, http.StatusInternalServerError,
			common.Translate(useCaseErr.ErrCode()))
		return
	}

	common.SimpleResponseOK(
		c, http.StatusOK,
		result,
	)
}

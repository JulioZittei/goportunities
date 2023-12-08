package handler

import (
	"net/http"

	"github.com/JulioZittei/goportunities/error"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schema.Opening{}
	if err := db.First(&opening, id).Error;err != nil {
		err := &error.NotFoundOpeningError{
			Code: http.StatusNotFound,
			Context: ctx,
		}

		sendError(ctx, err.Code, err.Error(id))
		return
	}

	sendSuccess(ctx, http.StatusOK, opening)
}

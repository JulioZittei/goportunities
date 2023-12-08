package handler

import (
	"net/http"

	"github.com/JulioZittei/goportunities/error"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error;err != nil {
		err := &error.ListOpeningError{
			Code: http.StatusInternalServerError,
			Context: ctx,
		}

		sendError(ctx, err.Code, err.Error())
	}

	sendSuccess(ctx, http.StatusOK, openings)
}
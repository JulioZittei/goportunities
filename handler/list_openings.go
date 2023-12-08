package handler

import (
	"net/http"

	"github.com/JulioZittei/goportunities/error"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get openings
// @Description Get a list of job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} SuccessWithListResponse
// @Failure 400 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
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
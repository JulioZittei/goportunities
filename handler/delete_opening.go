package handler

import (
	"net/http"

	"github.com/JulioZittei/goportunities/error"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		err := &error.NotFoundOpeningError{
			Code: http.StatusNotFound,
			Context: ctx,
		}

		sendError(ctx, err.Code, err.Error(id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		err := &error.DeleteOpeningError{
			Code: http.StatusInternalServerError,
			Context: ctx,
		}
		sendError(ctx, err.Code, err.Error(id))
		return
	}

	sendSuccess(ctx,http.StatusNoContent, nil)
}

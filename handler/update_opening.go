package handler

import (
	"net/http"

	"github.com/JulioZittei/goportunities/error"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Param request body UpdateOpeningRequest true "Request body"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	var request map[string]interface{}

	if err := ctx.BindJSON(&request); err != nil {
		err := &error.BindJSONError{
			Code: http.StatusBadRequest,
			Context: ctx,
		}
		sendError(ctx, err.Code, err.Error())
		return
	}

	if len(request) == 0 {
		err := &error.BindJSONError{
			Code: http.StatusBadRequest,
			Context: ctx,
		}
		sendError(ctx, err.Code, err.Error())
		return
	}

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

	if err := db.Model(&opening).Updates(request).Error; err != nil {
		err := &error.UpdateOpeningError{
			Code: http.StatusInternalServerError,
			Context: ctx,
		}
		
		sendError(ctx, err.Code, err.Error(id))
		return
	}

	sendSuccess(ctx, http.StatusOK, opening)
}
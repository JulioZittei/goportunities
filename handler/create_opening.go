package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/JulioZittei/goportunities/error"
	"github.com/JulioZittei/goportunities/message"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	errorParams := []ErrorsParam{}
	request := CreateOpeningRequest{}
	lang, _ := ctx.Get("Accept-Language")

	if err := ctx.BindJSON(&request); err != nil {
		err := &error.BindJSONError{
			Code: http.StatusBadRequest,
			Context: ctx,
		}
		sendError(ctx, err.Code, err.Error())
		return
	}

	if err := validate.Struct(request); err != nil {
		logger.Errorf("validation error: %v", err.Error())

		for _, err := range err.(validator.ValidationErrors) {

			langTag := fmt.Sprint(lang)
			field := fmt.Sprint(err.Field())
			tag := fmt.Sprint(err.Tag())
			typ := fmt.Sprint(err.Type())
			paramValue := fmt.Sprint(err.Param())

			param := strings.ToLower(field)
			value := message.GetMessage(langTag, tag, typ, paramValue)

			errorParam := ErrorsParam{param, value}

			errorParams = append(errorParams, errorParam)
		}

		sendErrorParams(ctx, http.StatusUnprocessableEntity, &errorParams)
	}

	opening := schema.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err)

		err := &error.CreateOpeningError{
			Code: http.StatusInternalServerError,
			Context: ctx,
		}

		sendError(ctx, err.Code, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusCreated, opening)
}
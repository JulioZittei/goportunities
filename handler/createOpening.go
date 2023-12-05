package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/JulioZittei/goportunities/message"
	"github.com/JulioZittei/goportunities/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateOpeningHandler(ctx *gin.Context) {
	errorParams := []ErrorsParam{}
	request := CreateOpeningRequest{}
	lang, _ := ctx.Get("Accept-Language")


	ctx.BindJSON(&request)

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

		sendErrorParams(ctx, http.StatusUnprocessableEntity, errorParams)
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
		sendError(ctx, http.StatusInternalServerError, "Error while creating opening on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, opening)
}
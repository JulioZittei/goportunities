package handler

import (
	"fmt"
	"net/http"

	"github.com/JulioZittei/goportunities/message"
	"github.com/gin-gonic/gin"
)

type ErrorsParam struct {
	Param string `json:"param"`
	Message string `json:"message"`
}

func sendError(ctx *gin.Context, code int, detail string) {
	lang, _ := ctx.Get("Accept-Language")
	title := message.GetMessage(fmt.Sprint(lang), fmt.Sprint(code))

	ctx.Status(code)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"title": title,
		"detail": detail,
		"instance": ctx.Request.RequestURI,
	})
}

func sendErrorParams(ctx *gin.Context, code int, errorsParams *[]ErrorsParam) {
	lang, _ := ctx.Get("Accept-Language")
	title := message.GetMessage(fmt.Sprint(lang), fmt.Sprint(code), )

	ctx.Status(code)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"status": http.StatusText(code),
		"code": code,
		"title": title,
		"instance": ctx.Request.RequestURI,
		"invalid_params": errorsParams, 
	})
}

func sendSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.Status(code)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": "Success",
		"data": data,
	})
}
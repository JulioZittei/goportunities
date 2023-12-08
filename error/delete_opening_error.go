package error

import (
	"fmt"

	"github.com/JulioZittei/goportunities/message"
	"github.com/gin-gonic/gin"
)

type DeleteOpeningError struct {
	Code int
	Context *gin.Context
}

func (b *DeleteOpeningError) Error(id string) string {
	lang, _ := b.Context.Get("Accept-Language")
	langTag := fmt.Sprint(lang)

	return message.GetMessage(langTag, "DeleteOpeningError", id)
}
package error

import (
	"fmt"

	"github.com/JulioZittei/goportunities/message"
	"github.com/gin-gonic/gin"
)

type ListOpeningError struct {
	Code int
	Context *gin.Context
}

func (b *ListOpeningError) Error() string {
	lang, _ := b.Context.Get("Accept-Language")
	langTag := fmt.Sprint(lang)

	return message.GetMessage(langTag, "ListOpeningError", nil)
}
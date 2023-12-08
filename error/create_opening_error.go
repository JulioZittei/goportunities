package error

import (
	"fmt"

	"github.com/JulioZittei/goportunities/message"
	"github.com/gin-gonic/gin"
)

type CreateOpeningError struct {
	Code int
	Context *gin.Context
}

func (b *CreateOpeningError) Error() string {
	lang, _ := b.Context.Get("Accept-Language")
	langTag := fmt.Sprint(lang)

	return message.GetMessage(langTag, "CreateOpeningError", nil)
}
package router

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func getLocaleContext(ctx *gin.Context) {
	locale := os.Getenv("LANG")
	locale = strings.Split(locale, ".")[0]
	
	acceptLanguage := ctx.GetHeader("Accept-Language")

	if acceptLanguage != "" {
		ctx.Set("Accept-Language", acceptLanguage)
		ctx.Next()
	}
	languageTag, err := language.Parse(locale)
	if err != nil {
		logger.Errorf("error while parsing locale: %v", err)
	}

	ctx.Set("Accept-Language", languageTag.String())
	ctx.Next()
}
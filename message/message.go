package message

import (
	"fmt"
	"strings"
)

func getMessageBundle(lang string) map[string]string {
	lang = strings.ToLower(lang)

	switch lang {
	case "pt-br":
			return messagesPTBR
	case "en":
			return messagesEN
	default:
			return messagesEN
	}
}

func GetMessage(lang string, key string , args ...any) string {
	messages := getMessageBundle(lang)
	message := messages[key]

	count := strings.Count(message, "%s")

	if len(args) == 0 {
		return fmt.Sprint(message)
	}

	if len(args) > count {
		args = args[:count]
	} 

	return fmt.Sprintf(message, args...)
}
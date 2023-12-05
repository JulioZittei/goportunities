package message

var messagesEN = map[string]string{
	"required": "(type %s) is required",
	"gt":       "(type %s) must be greater than %s",
	"422": "Invalid request: verify params sent and try again",
	"400": "Invalid request: verify data sent and try again",
	"500": "An unexpected error occurred on the server",
}
package message

var messagesEN = map[string]string{
	"invalid_field": "is not expected",
	"invalid_type": "(type %s) is expected",
	"required": "(type %s) is required",
	"gt": "(type %s) must be greater than %s",
	"BindJSONError": "JSON malformed or missing required param",
	"CreateOpeningError": "Error while creating opening on database",
	"NotFoundOpeningError": "Opening with id: %s not found",
	"DeleteOpeningError": "Error while deleting opening with id: %s on database",
	"UpdateOpeningError": "Error while updating opening with id: %s on database",
	"ListOpeningError": "Error while listing openings on database",
	"422": "Invalid request: verify params sent and try again",
	"400": "Invalid request: verify data sent and try again",
	"404": "The requested resource was not found",
	"500": "An unexpected error occurred on the server",
}
package message

var messagesPTBR = map[string]string {
	"invalid_field": "não é esperado",
	"required": "(tipo %s) é obrigatório",
	"gt":       "(tipo %s) deve ser maior que %s",
	"BindJSONError": "JSON mal formado ou falta de parâmetro obrigatório",
	"CreateOpeningError": "Erro ao criar opening no banco de dados",
	"NotFoundOpeningError": "Opening com id: %s não encontrado",
	"DeleteOpeningError": "Erro ao deletar opening com id: %s no banco de dados",
	"UpdateOpeningError": "Erro ao atualizar opening com id: %s no banco de dados",
	"ListOpeningError": "Erro ao listar openings no banco de dados",
	"422": "Requisição inválida: verifique os parâmetros enviados e tente novamente",
	"400": "Requisição inválida: verifique os dados enviados e tente novamente",
	"404": "O recurso solicitado não foi encontrado",
	"500": "Um erro inesperado ocorreu no servidor",
}
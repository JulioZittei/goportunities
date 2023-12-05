package message

var messagesPTBR = map[string]string{
	"required": "(tipo %s) é obrigatório",
	"gt":       "(tipo %s) deve ser maior que %s",
	"422": "Requisição inválida: verifique os parâmetros enviados e tente novamente",
	"400": "Requisição inválida: verifique os dados enviados e tente novamente",
	"500": "Um erro inesperado ocorreu no servidor",
}
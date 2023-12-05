package handler

import (
	"github.com/JulioZittei/goportunities/config"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
	validate *validator.Validate
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
	validate = validator.New(validator.WithRequiredStructEnabled())
}


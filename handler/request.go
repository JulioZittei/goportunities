package handler

import (
	"fmt"
)

type CreateOpeningRequest struct {
	Role string `json:"role" validate:"required"`
	Company string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Remote *bool `json:"remote" validate:"required"`
	Link string `json:"link" validate:"required"`
	Salary int64 `json:"salary" validate:"required,gt=0"`
}


type UpdateOpeningRequest struct {
	Role string `json:"role" validate:"required"`
	Company string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Remote *bool `json:"remote" validate:"required"`
	Link string `json:"link" validate:"required"`
	Salary int64 `json:"salary" validate:"required,gt=0"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("query param: %s (type: %s) is required", name, typ)
}

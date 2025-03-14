package entity

import "github.com/go-playground/validator/v10"

// Определяем структуру Feedback с тегами validate
type Feedback struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	Email          string `json:"email" validate:"required,email"`
	Message        string `json:"message" validate:"required,min=10,max=1000"`
	Consent        bool   `json:"consent" validate:"required"`
	IdempotencyKey string `json:"idempotency_key" validate:"required,uuid"` // Пример с UUID
}

// Метод Validate для структуры Feedback
func (f *Feedback) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}

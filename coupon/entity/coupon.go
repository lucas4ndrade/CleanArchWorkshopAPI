package entity

import (
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/internal"
)

type Coupon struct {
	ID        string
	Code      string
	Value     float64
	CompanyID string
}

func (c *Coupon) Validate() (err error) {
	if c.Code == "" {
		return internal.NewValidationError("Código não pode ser vazio!!!")
	}

	if c.Value == float64(0) {
		return internal.NewValidationError("Valor do cupom não pode ser zero!!!")
	}

	return
}

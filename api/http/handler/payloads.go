package handler

import "github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"

type CouponCreatePayload struct {
	Code      string  `json:"code"`
	Value     float64 `json:"value"`
	CompanyID string  `json:"company_id"`
}

func (p *CouponCreatePayload) ToEntity() entity.Coupon {
	return entity.Coupon{
		Code:      p.Code,
		Value:     p.Value,
		CompanyID: p.CompanyID,
	}
}

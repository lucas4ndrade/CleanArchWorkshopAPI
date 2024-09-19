package presenters

import "github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"

type CouponPresenter struct {
	ID        string  `json:"id"`
	Code      string  `json:"code"`
	Value     float64 `json:"value"`
	CompanyID string  `json:"company_id"`
}

func NewCouponPresenter(c entity.Coupon) CouponPresenter {
	return CouponPresenter(c)
}

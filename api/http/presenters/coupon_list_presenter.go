package presenters

import "github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"

type CouponListPresenter []CouponPresenter

func NewCouponListPresenter(cs []entity.Coupon) (pr CouponListPresenter) {
	for _, c := range cs {
		pr = append(pr, NewCouponPresenter(c))
	}

	return
}

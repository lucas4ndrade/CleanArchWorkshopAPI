package handler

import (
	"net/url"

	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon"
)

func NewCouponListFinder(values url.Values) (f coupon.Finder) {
	f = coupon.Finder{}

	f.Code = values.Get("code")

	return
}

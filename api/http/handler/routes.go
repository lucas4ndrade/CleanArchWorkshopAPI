package handler

import (
	"github.com/go-chi/chi"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon"
)

func StartRoutes(
	uc coupon.UseCase,
) (router *chi.Mux) {
	router = chi.NewMux()

	router.Get("/coupons", listCoupons(uc))
	router.Post("/coupons", createCoupon(uc))
	router.Get("/companies/{id}/coupon_count", getCompanyCouponCount(uc))

	return
}

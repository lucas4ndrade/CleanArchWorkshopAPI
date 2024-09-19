package handler

import (
	"encoding/json"
	"net/http"

	"github.com/delivery-much/dm-go/render"
	"github.com/go-chi/chi"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/api/http/presenters"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/internal"
)

func listCoupons(uc coupon.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		f := NewCouponListFinder(r.URL.Query())
		couponList, err := uc.ListCoupons(ctx, f)
		if err != nil {
			render.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		p := presenters.NewCouponListPresenter(couponList)
		render.RespondJSON(w, http.StatusOK, p)
	}
}

func createCoupon(uc coupon.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var b CouponCreatePayload
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&b); err != nil {
			render.RespondError(w, http.StatusBadRequest, err)
			return
		}

		c := b.ToEntity()
		createdCoupon, err := uc.CreateCoupon(ctx, c)
		if err != nil {
			status := http.StatusInternalServerError
			if internal.IsValidationError(err) {
				status = http.StatusBadRequest
			}

			render.RespondError(w, status, err)
			return
		}

		p := presenters.NewCouponPresenter(createdCoupon)
		render.RespondJSON(w, http.StatusOK, p)
	}
}

func getCompanyCouponCount(uc coupon.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")

		res, err := uc.CountCompanyCoupons(ctx, id)
		if err != nil {
			render.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		p := presenters.NewCompanyCouponCountPresenter(res)
		render.RespondJSON(w, http.StatusOK, p)
	}
}

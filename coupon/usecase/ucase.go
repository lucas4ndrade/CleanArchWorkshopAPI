package usecase

import (
	"context"

	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"
)

type usecase struct {
	repo coupon.Repository
}

func NewUsecase(repo coupon.Repository) coupon.UseCase {
	return &usecase{
		repo,
	}
}

func (uc *usecase) ListCoupons(ctx context.Context, f coupon.Finder) ([]entity.Coupon, error) {
	return uc.repo.ListCoupons(ctx, f)
}

func (uc *usecase) CreateCoupon(ctx context.Context, c entity.Coupon) (createdCoupon entity.Coupon, err error) {
	if err = c.Validate(); err != nil {
		return
	}

	return uc.repo.CreateCoupon(ctx, c)
}

func (uc *usecase) CountCompanyCoupons(ctx context.Context, cID string) (int64, error) {
	return uc.repo.CountCompanyCoupons(ctx, cID)
}

package coupon

import (
	"context"

	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"
)

type UseCase interface {
	ListCoupons(ctx context.Context, f Finder) ([]entity.Coupon, error)
	CreateCoupon(ctx context.Context, c entity.Coupon) (entity.Coupon, error)
	CountCompanyCoupons(ctx context.Context, cID string) (int64, error)
}

type Repository interface {
	ListCoupons(ctx context.Context, f Finder) ([]entity.Coupon, error)
	CreateCoupon(ctx context.Context, c entity.Coupon) (entity.Coupon, error)
	CountCompanyCoupons(ctx context.Context, cID string) (int64, error)
}

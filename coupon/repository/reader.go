package repository

import (
	"context"

	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) ListCoupons(ctx context.Context, f coupon.Finder) (cs []entity.Coupon, err error) {
	query := mountListCouponQuery(f)

	cursor, err := r.client.
		Database(r.db).
		Collection(r.coll).
		Find(ctx, query)
	if err != nil {
		return
	}

	for cursor.Next(ctx) {
		var cm CouponModel
		err = cursor.Decode(&cm)
		if err != nil {
			return
		}
		cs = append(cs, UnmarshalCouponModel(cm))
	}

	return
}

func (r *repository) CountCompanyCoupons(ctx context.Context, cID string) (count int64, err error) {
	return r.client.
		Database(r.db).
		Collection(r.coll).
		CountDocuments(ctx, bson.M{
			"company_id": cID,
		})
}

func mountListCouponQuery(f coupon.Finder) (filter bson.M) {
	filter = bson.M{}

	if codeFilter := f.Code; codeFilter != "" {
		filter["code"] = codeFilter
	}

	return
}

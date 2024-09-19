package repository

import (
	"context"

	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) CreateCoupon(ctx context.Context, c entity.Coupon) (createdCoupon entity.Coupon, err error) {
	model := MarshalCouponModel(c)

	result, err := r.client.
		Database(r.db).
		Collection(r.coll).
		InsertOne(ctx, model)
	if err != nil {
		return
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	c.ID = insertedID.Hex()
	createdCoupon = c
	return
}

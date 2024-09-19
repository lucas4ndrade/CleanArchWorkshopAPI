package repository

import (
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	client *mongo.Client
	db     string
	coll   string
}

func NewRepository(
	client *mongo.Client,
	db string,
	coll string,
) coupon.Repository {
	return &repository{
		client,
		db,
		coll,
	}
}

type CouponModel struct {
	ID        string  `bson:"id,omitempty"`
	Code      string  `bson:"code,omitempty"`
	Value     float64 `bson:"value,omitempty"`
	CompanyID string  `bson:"company_id,omitempty"`
}

func MarshalCouponModel(c entity.Coupon) (cm CouponModel) {
	return CouponModel(c)
}

func UnmarshalCouponModel(cm CouponModel) (c entity.Coupon) {
	return entity.Coupon(cm)
}

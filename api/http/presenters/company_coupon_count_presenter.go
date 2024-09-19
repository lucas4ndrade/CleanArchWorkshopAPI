package presenters

type CompanyCouponCountPresenter struct {
	Total int64 `json:"total"`
}

func NewCompanyCouponCountPresenter(t int64) CompanyCouponCountPresenter {
	return CompanyCouponCountPresenter{
		Total: t,
	}
}

package delivery

import "golang-base/src/usecase"

type Delivery interface {
}

type delivery struct {
	usecase usecase.Usecase
}

func NewDelivery(usecase usecase.Usecase) Delivery {
	return &delivery{
		usecase: usecase,
	}
}

package usecase

import "golang-base/src/repository"

type Usecase interface {
}

type usecase struct {
	repo repository.Store
}

func NewUsecase(repo repository.Store) Usecase {
	return &usecase{
		repo: repo,
	}
}

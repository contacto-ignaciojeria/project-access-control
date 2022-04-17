package usecase

import (
	"context"
	"pac/src/user/domain"
)

type GettingStartedUseCase interface {
	Execute(context.Context,domain.Customer) (domain.Customer, error)
}

type GettingStarted struct {
}

func (u *GettingStarted) Execute(ctx context.Context,user domain.Customer) (domain.Customer,error){
	return domain.Customer{},nil
}
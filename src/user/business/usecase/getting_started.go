package usecase

import (
	"context"
	"pac/src/user/domain"
)

type GettingStartedUseCase interface {
	Execute(context.Context,domain.User) (domain.User, error)
}

type GettingStarted struct {
}

func (u *GettingStarted) Execute(ctx context.Context,user domain.User) (domain.User,error){
	return domain.User{},nil
}
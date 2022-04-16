package usecase

import "pac/internal/user/domain"

type IdentificationUseCase interface {
	Execute(domain.User) domain.User
}

type Identification struct {
}

func (u *Identification) Execute(user domain.User) domain.User{
	return domain.User{}
}
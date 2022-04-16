package outport

import (
	"context"
	"pac/src/user/domain"
)

type FindUserByEmail func(context.Context,domain.User) (domain.User,error)
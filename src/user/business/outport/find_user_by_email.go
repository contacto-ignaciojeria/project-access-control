package outport

import (
	"context"
	"pac/src/user/domain"
)

type FindUserByEmail func(context.Context,domain.Customer) (domain.Customer,error)
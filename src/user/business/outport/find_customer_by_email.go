package outport

import (
	"context"
	"pac/src/user/domain"
)

type FindCustomerByEmail func(context.Context,domain.Customer) (domain.Customer,error)
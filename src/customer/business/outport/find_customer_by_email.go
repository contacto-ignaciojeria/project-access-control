package outport

import (
	"context"
	"pac/src/customer/domain"
)

type FindCustomerByEmail func(context.Context,domain.Customer) (domain.Customer,error)
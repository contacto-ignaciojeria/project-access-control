package outport

import (
	"context"
	"pac/src/customer/domain"
)

type SaveCustomer func(context.Context, domain.Customer)  error
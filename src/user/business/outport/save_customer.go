package outport

import (
	"context"
	"pac/src/user/domain"
)

type SaveCustomer func(context.Context, domain.Customer)  error
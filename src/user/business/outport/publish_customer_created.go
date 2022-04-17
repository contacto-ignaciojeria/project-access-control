package outport

import (
	"context"
	"pac/src/user/domain"
)

type PublishCustomerCreated func(context.Context, domain.Customer) error
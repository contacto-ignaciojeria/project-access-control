package outport

import (
	"context"
	"pac/src/customer/domain"
) 

type PublishCustomerCreated func(context.Context, domain.Customer) error
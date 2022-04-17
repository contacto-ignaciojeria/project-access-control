package outport

import (
	"context"
	"pac/src/customer/domain"
)

type SendVerificationCode func(context.Context, domain.Customer) error
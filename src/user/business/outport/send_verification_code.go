package outport

import (
	"context"
	"pac/src/user/domain"
)

type SendVerificationCode func(context.Context, domain.Customer) error
package outport

import (
	"context"
	"pac/src/user/domain"
)

type PublishUserCreated func(context.Context, domain.User) error
package outport

import (
	"context"
	"pac/src/user/domain"
)

type SaveUser func(context.Context, domain.User)  error
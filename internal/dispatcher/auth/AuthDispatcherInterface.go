package authDispatcher

import "context"

type AuthDispatcherInterface interface {
	Check(ctx context.Context) (bool, error)
}

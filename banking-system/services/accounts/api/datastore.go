package api

import "context"

type Datastore interface {
	CreateAccount(ctx context.Context, info *AccountInfo) error
}

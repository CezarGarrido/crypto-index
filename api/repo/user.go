package repo

import "context"

type UserRepoInterface interface {
	Find(ctx context.Context, email, password string)
}

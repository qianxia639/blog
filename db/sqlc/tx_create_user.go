package db

import (
	"context"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg *CreateUserTxParams) (user User, err error) {
	err = store.execTx(ctx, func(q *Queries) error {
		user, err = q.CreateUser(ctx, &arg.CreateUserParams)
		if err != nil {
			return err
		}
		return arg.AfterCreate(user)
	})
	return
}

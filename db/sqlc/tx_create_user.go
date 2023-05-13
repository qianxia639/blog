package db

import (
	"context"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg *CreateUserTxParams) (User, error) {
	var user User
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		user, err = q.CreateUser(ctx, &arg.CreateUserParams)
		if err != nil {
			return err
		}
		return arg.AfterCreate(user)
	})
	return user, err
}

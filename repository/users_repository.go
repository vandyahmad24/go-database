package repository

import (
	"context"
	"github.com/vandyahmad24/go-database-mysql/model"
)

type UsersRepository interface {
	Insert(ctx context.Context, users model.Users)(model.Users, error)
	FindByID(ctx context.Context, id int32)(model.Users, error)
	FindAll(ctx context.Context)([]model.Users, error)
	DeleteByID(ctx context.Context, id int32)(error)
	//DeleteByID(ctx context.Context, id int32)(model.Users, error)
	//DeleteByID(ctx context.Context)(model.Users, error)
}
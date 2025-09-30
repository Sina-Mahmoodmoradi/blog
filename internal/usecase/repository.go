package usecase

import (
	"context"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)


type UserRepository interface{
	FindByEmail(ctx context.Context,email string) (*entity.User, error)
	FindByUsername(ctx context.Context,username string) (*entity.User, error)
	Save(ctx context.Context,user *entity.User) error
}
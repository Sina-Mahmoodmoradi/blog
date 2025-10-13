package usecase

import (
	"context"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)


type UserRepository interface{
	FindByEmail(ctx context.Context,email string) (*entity.User, error)
	FindByID(ctx context.Context,id uint) (*entity.User, error)
	FindByUsername(ctx context.Context,username string) (*entity.User, error)
	Save(ctx context.Context,user *entity.User) error
}

type PostRepository interface{
	Save(ctx context.Context,post *entity.Post) error
}
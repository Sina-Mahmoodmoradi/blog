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
	GetList(ctx context.Context ,AuthorID uint,offset ,limit int) ([]*entity.Post,error)
	Count(ctx context.Context,AuthorID uint)(int,error)
	GetById(ctx context.Context, id uint)(*entity.Post,error)
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id uint) error
}

type CommentRepository interface{
	Save(ctx context.Context, comment *entity.Comment)error
	GetList(ctx context.Context ,PostID uint,offset ,limit int) ([]*entity.Comment,error)
	Count(ctx context.Context,PostID uint)(int,error)
}
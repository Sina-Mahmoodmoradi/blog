package repository

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) usecase.PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Save(ctx context.Context,post *entity.Post) error{
	postModel := ToModelPost(post)
	if err:=r.db.WithContext(ctx).Create(postModel).Error;err!=nil{
		return fmt.Errorf("failed to create post %w",err)
	}

	post.ID = postModel.ID
	return nil
}

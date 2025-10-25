package repository

import (
	"context"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"gorm.io/gorm"
)




type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) usecase.CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Save(ctx context.Context,comment *entity.Comment) error{
	commentModel := ToModelComment(comment)
	if err:=r.db.WithContext(ctx).Create(commentModel).Error;err!=nil{
		return err
	}

	comment.ID = commentModel.ID
	return nil
}
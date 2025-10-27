package repository

import (
	"context"

	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
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


func (r *CommentRepository)GetList(ctx context.Context,postID uint,offset,limit int)([]*entity.Comment,error){
	var modelComments []models.Comment
	if err:=r.db.WithContext(ctx).Where("post_id = ?",postID).Offset(offset).Limit(limit).Order("created_at desc").Find(&modelComments).Error;err!=nil{
		return nil,err
	}
	comments := make([]*entity.Comment,0,len(modelComments))
	for _,comment := range modelComments{
		comments=append(comments, ToEntityComment(&comment))
	}
	return comments,nil
}

func (r *CommentRepository)Count(ctx context.Context,postID uint)(int,error){
	var count int64
	if err:= r.db.WithContext(ctx).Model(&models.Comment{}).Where("post_id = ?",postID).Count(&count).Error;err!=nil{
		return 0,err
	}
	return int(count),nil
}


func (r *CommentRepository)GetById(ctx context.Context, id uint)(*entity.Comment,error){
	var modelComment models.Comment
	if err:=r.db.WithContext(ctx).First(&modelComment,id).Error;err!=nil{
		return nil,err
	}

	return ToEntityComment(&modelComment),nil
}

func (r *CommentRepository)Update(ctx context.Context, comment *entity.Comment) error{
	return r.db.WithContext(ctx).Save(comment).Error
}


package repository

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
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

func (r *PostRepository)GetList(ctx context.Context ,AuthorID uint,offset ,limit int) ([]*entity.Post,error){
	var posts []models.Post
	if err:=r.db.WithContext(ctx).Where("author_id = ?", AuthorID).Order("created_at desc").Limit(limit).Offset(offset).Find(&posts).Error; err!=nil{
		return nil,fmt.Errorf("failed to get posts %w",err)
	}
	
	entityPosts := make([]*entity.Post,0,len(posts))
	for _,post:=range posts{
		entityPosts = append(entityPosts, ToEntityPost(&post))
	}

	return  entityPosts,nil
}


func (r *PostRepository)Count(ctx context.Context,AuthorID uint)(int,error){
	var count int64
	if err:=r.db.Model(&models.Post{}).Count(&count).Error;err!=nil{
		return 0,fmt.Errorf("failed to count posts %w",err)
	}

	return int(count),nil
}

func (r *PostRepository)GetById(ctx context.Context, id uint)(*entity.Post,error){
	var post models.Post
	if err:=r.db.WithContext(ctx).First(&post,id).Error; err!=nil{
		return nil,fmt.Errorf("failed to get post: %w",err)
	}

	return ToEntityPost(&post),nil
}


func (r *PostRepository)Update(ctx context.Context, post *entity.Post) error{
	return r.db.WithContext(ctx).Save(post).Error
}


func (r *PostRepository)Delete(ctx context.Context, id uint) error{
	return r.db.WithContext(ctx).Delete(&models.Post{},id).Error
}



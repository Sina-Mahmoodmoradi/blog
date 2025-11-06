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

func (r *PostRepository)GetList(ctx context.Context ,filter *usecase.PostFilter)([]*entity.Post,error){
	var posts []models.Post
	query := r.db.WithContext(ctx)

	if len(filter.Tags) > 0 {
		query = query.
		Joins("JOIN post_tags ON post_tags.post_id = posts.id").
		Joins("JOIN tags ON post_tags.tag_id = tags.id").
		Where("tags.name IN ?", filter.Tags).
		Group("posts.id")
	}

	if filter.AuthorID != nil {
		query = query.Where("author_id = ?", *filter.AuthorID)
	}

	if err := query.
	Order("posts.created_at desc").
	Limit(filter.Limit).
	Offset(filter.Offset).
	Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("failed to get posts %w", err)
	}
	
	entityPosts := make([]*entity.Post,0,len(posts))
	for _,post:=range posts{
		entityPosts = append(entityPosts, ToEntityPost(&post))
	}

	return  entityPosts,nil
}

// func (r *PostRepository)GetListByTags(ctx context.Context ,authorID *uint, tagNames []string,offset ,limit int) ([]*entity.Post,error){
// 	var posts []models.Post
// 	query:=r.db.WithContext(ctx).
// 	Joins("JOIN post_tags ON post_tags.post_id = posts.id").
// 	Joins("JOIN tags ON post_tags.tag_id = tags.id").
// 	Where("tags.name IN ?",tagNames).Group("posts.id")

// 	if authorID!=nil{
// 		query = query.Where("posts.author_id = ?",*authorID)
// 	}

// 	if err:=query.
// 	Order("posts.created_at desc").
// 	Limit(limit).
// 	Offset(offset).
// 	Find(&posts).Error; err!=nil{
// 		return nil,fmt.Errorf("failed to get posts %w",err)
// 	}
	
// 	entityPosts := make([]*entity.Post,0,len(posts))
// 	for _,post:=range posts{
// 		entityPosts = append(entityPosts, ToEntityPost(&post))
// 	}

// 	return  entityPosts,nil
// }


func (r *PostRepository)Count(ctx context.Context,filter *usecase.PostFilter)(int,error){
	var count int64
	query := r.db.Model(&models.Post{})

	if len(filter.Tags) > 0 {
		query = query.
		Joins("JOIN post_tags ON post_tags.post_id = posts.id").
		Joins("JOIN tags ON post_tags.tag_id = tags.id").
		Where("tags.name IN ?", filter.Tags).
		Distinct("posts.id")
	}

	if filter.AuthorID != nil {
		query = query.Where("author_id = ?", *filter.AuthorID)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count posts %w", err)	
	}

	return int(count),nil
}

// func (r *PostRepository)CountByTags(ctx context.Context,authorID *uint,tagNames []string)(int,error){
// 	var count int64
// 	query:= r.db.Model(&models.Post{}).
// 	Joins("JOIN post_tags ON post_tags.post_id = posts.id").
// 	Joins("JOIN tags ON post_tags.tag_id = tags.id").
// 	Where("tags.name IN ?",tagNames).Distinct("posts.id")

// 	if authorID!=nil{
// 		query = query.Where("posts.author_id = ?",*authorID)
// 	}

// 	if err:=query.Count(&count).Error;err!=nil{
// 		return 0,fmt.Errorf("failed to count posts %w",err)
// 	}

// 	return int(count),nil
// }

func (r *PostRepository)GetById(ctx context.Context, id uint)(*entity.Post,error){
	var post models.Post
	if err:=r.db.WithContext(ctx).First(&post,id).Error; err!=nil{
		return nil,fmt.Errorf("failed to get post: %w",err)
	}

	return ToEntityPost(&post),nil
}

func (r *PostRepository)GetByIdWithComments(ctx context.Context, id uint, limit int)(*entity.Post,error){
	var post models.Post
	if err:=r.db.WithContext(ctx).Preload("Comments",func (db *gorm.DB)*gorm.DB{
		return db.Order("created_at desc").Limit(limit)
	}).First(&post,id).Error; err!=nil{
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

func (r *PostRepository)AppendTags(ctx context.Context, post *entity.Post,tags []*entity.Tag) error{
	modelTags := make([]interface{},len(tags))
	for i,tag :=range tags{
		modelTags[i] = ToModelTag(tag)
	}
	modelPost := ToModelPost(post)

	if err:=r.db.WithContext(ctx).Model(modelPost).Association("Tags").Append(modelTags...);err!=nil{
		return fmt.Errorf("failed to append tags: %w",err)
	}
	return nil
}




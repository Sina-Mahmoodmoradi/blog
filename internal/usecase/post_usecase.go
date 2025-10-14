package usecase

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)




type PostUseCase struct{
	repo PostRepository
}

type CreatePostRequest struct{
	Title string
	Content string
	AuthorID uint
}

type GetPostsRequest struct{
	AuthorID uint
	Page int
	Limit int
}

type PaginatedPosts struct{
	Posts []entity.Post
	Total int
	Page int
	Limit int 
}


func NewPostUseCase(repo PostRepository) *PostUseCase{
	return &PostUseCase{
		repo: repo,
	}
}


func (u *PostUseCase)CreatePost(ctx context.Context,req *CreatePostRequest)(*entity.Post,error){
	post := &entity.Post{
		Title: req.Title,
		Content: req.Content,
		AuthorID: req.AuthorID,
	}

	if err:=u.repo.Save(ctx,post);err!=nil{
		return nil,fmt.Errorf("failed to create post: %w",err)
	}

	return post,nil
}


func (u *PostUseCase)GetAllPosts(ctx context.Context,req *GetPostsRequest)(*PaginatedPosts,error){
	offset:= (req.Page-1)*req.Limit
	posts, err:= u.repo.GetList(ctx,req.AuthorID,offset,req.Limit)
	if err!=nil{
		return nil,fmt.Errorf("failed to get posts:%w",err)
	}
	count,err := u.repo.Count(ctx,req.AuthorID)
	if err!=nil{
		return nil,fmt.Errorf("failed to get count of posts:%w",err)
	}

	return &PaginatedPosts{
		Posts: posts,
		Total: count,
		Page: req.Page,
		Limit: req.Limit,
	},nil
}




func (u *PostUseCase)GetPost(ctx context.Context,author_id, id uint)(*entity.Post,error){
	post,err := u.repo.GetById(ctx,id)
	if err!=nil{
		return nil,err
	}

	if post.AuthorID!=author_id{
		return nil,fmt.Errorf("post not found")
	}

	return post,nil
}

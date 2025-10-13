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

type CreatePostResponse struct{

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
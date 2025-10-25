package usecase

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)







type CommentUseCase struct {
	repo CommentRepository
	postRepo PostRepository
}

type CreateCommentRequest struct{
	AuthorID uint
	PostID uint
	Content string
}

func NewCommentUseCase(repo CommentRepository,postRepo PostRepository)*CommentUseCase{
	return &CommentUseCase{
		repo: repo,
		postRepo: postRepo,
	}
}


func (uc *CommentUseCase) CreateComment(ctx context.Context,req *CreateCommentRequest)(*entity.Comment,error){
	comment:= &entity.Comment{
		PostID: req.PostID,
		AuthorID: req.AuthorID,
		Content: req.Content,
	}

	//check if post exists!?
	if _,err:=uc.postRepo.GetById(ctx,req.PostID);err!=nil{
		return nil,fmt.Errorf("post does not exist: %w",err)
	}

	if err:=uc.repo.Save(ctx,comment);err!=nil{
		return nil,fmt.Errorf("failed to create comment:%w",err)
	}


	return comment,nil
}
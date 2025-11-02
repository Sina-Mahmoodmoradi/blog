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


type GetCommentsRequest struct{
	PostID uint
	Page int
	Limit int
}

type PaginatedComments struct{
	Comments []*entity.Comment
	Total int
	Page int
	Limit int
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


func (uc *CommentUseCase) GetAllComments(ctx context.Context,req *GetCommentsRequest )(*PaginatedComments,error){
	offset:= (req.Page-1)*req.Limit
	comments,err := uc.repo.GetList(ctx,req.PostID,offset,req.Limit)
	if err!=nil{
		return nil,fmt.Errorf("failed to get comment list:%w",err)
	}
	count,err := uc.repo.Count(ctx,req.PostID)
	if err!=nil{
		return nil,fmt.Errorf("failed to get count of comments:%w",err)
	}

	return &PaginatedComments{
		Comments: comments,
		Total: count,
		Page: req.Page,
		Limit: req.Limit,
	},nil

}


// func (uc *CommentUseCase) GetComment(ctx context.Context,commentID uint)(*entity.Comment,error){
// 	comment,err := uc.repo.GetById(ctx,commentID)
// 	if err!=nil{
// 		return nil,fmt.Errorf("failed to get comment: %w",err)
// 	}

// 	return comment,nil
// }



func (uc *CommentUseCase) UpdateComment(ctx context.Context,authorID,id uint,content string)(*entity.Comment,error){
	comment ,err := uc.repo.GetById(ctx,id)
	if err!=nil{
		return nil,fmt.Errorf("failed to get comment: %w",err)
	}

	if authorID!=comment.AuthorID{
		return nil,fmt.Errorf("unauthorized: cannot edit another users comment")
	}

	if content==""{
		return nil,fmt.Errorf("comment content cannot be empty")
	}
	comment.Content = content

	if err:=uc.repo.Update(ctx,comment);err!=nil{
		return nil,fmt.Errorf("failed to update comment: %w",err)
	}

	return comment,nil
}


func (u *CommentUseCase)DeleteComment(ctx context.Context, author_id, id uint)(error){
	comment,err := u.repo.GetById(ctx,id)
	if err!=nil{
		return err
	}

	if comment.AuthorID!=author_id{
		return fmt.Errorf("unauthorized: cannot delete another users comment")
	}

	if err:=u.repo.Delete(ctx,id);err!=nil{
		return err
	}

	return nil
}

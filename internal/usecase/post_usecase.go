package usecase

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)




type PostUseCase struct{
	repo PostRepository
	commentRepo CommentRepository
}

type CreatePostRequest struct{
	Title string
	Content string
	AuthorID uint
}

type GetPostsRequest struct{
	AuthorID *uint
	Page int
	Limit int
	TagNames  []string
}

type PaginatedPosts struct{
	Posts []*entity.Post
	Total int
	Page int
	Limit int 
}

type UpdatePostRequest struct{
	Title *string
	Content *string
}


func NewPostUseCase(repo PostRepository, commentRepo CommentRepository) *PostUseCase{
	return &PostUseCase{
		repo: repo,
		commentRepo: commentRepo,
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


	var (
		posts []*entity.Post
		count int 
		err error
	)
	switch {
	case len(req.TagNames) > 0:

		posts, err = u.repo.GetListByTags(ctx, req.AuthorID, req.TagNames, offset, req.Limit)
		if err != nil {
			return nil, fmt.Errorf("failed to get posts:%w", err)
		}
		count,err = u.repo.CountByTags(ctx,req.AuthorID,req.TagNames)
		if err!=nil{
			return nil,fmt.Errorf("failed to get count of posts:%w",err)
		}
	default:
		posts, err= u.repo.GetList(ctx,1,offset,req.Limit)
		if err!=nil{
			return nil,fmt.Errorf("failed to get posts:%w",err)
		}
		count,err = u.repo.Count(ctx,1)
		if err!=nil{
			return nil,fmt.Errorf("failed to get count of posts:%w",err)
		}
	}

	return &PaginatedPosts{
		Posts: posts,
		Total: count,
		Page: req.Page,
		Limit: req.Limit,
	},nil
}




func (u *PostUseCase)GetPost(ctx context.Context,author_id, id uint)(*entity.Post,bool,error){
	post,err := u.repo.GetByIdWithComments(ctx,id,5)
	if err!=nil{
		return nil,false,err
	}
	totalComments,err := u.commentRepo.Count(ctx,id)
	if err!=nil{
		return nil,false,err
	}
	hasMoreComments := totalComments > 5

	// if post.AuthorID!=author_id{
	// 	return nil,fmt.Errorf("post not found")
	// }

	return post,hasMoreComments,nil
}


func (u *PostUseCase)UpdatePost(ctx context.Context, author_id, id uint, req *UpdatePostRequest)(*entity.Post,error){
	post,err := u.repo.GetById(ctx,id)
	if err!=nil{
		return nil,err
	}

	if post.AuthorID!=author_id{
		return nil,fmt.Errorf("post not found")
	}

	if req.Title!=nil{
		post.Title = *req.Title
	}

	if req.Content!=nil{
		post.Content = *req.Content
	}

	if err:=u.repo.Update(ctx,post);err!=nil{
		return nil,fmt.Errorf("failed to update:%w",err)
	}

	return post,nil
}


func (u *PostUseCase)DeletePost(ctx context.Context, author_id, id uint)(error){
	post,err := u.repo.GetById(ctx,id)
	if err!=nil{
		return err
	}

	if post.AuthorID!=author_id{
		return fmt.Errorf("post not found")
	}

	if err:=u.repo.Delete(ctx,id);err!=nil{
		return err
	}

	return nil
}

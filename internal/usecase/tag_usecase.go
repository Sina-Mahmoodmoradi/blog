package usecase

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)





type TagUseCase struct {
	repo TagRepository
	postRepo PostRepository
}


func NewTagUseCase(repo TagRepository, postRepo PostRepository) *TagUseCase {
	return &TagUseCase{
		repo:    repo,
		postRepo: postRepo,
	}
}




func (uc *TagUseCase)AssignTagsToPost(ctx context.Context,userID,postId uint,tagNames []string)([]*entity.Tag,error){
	post,err := uc.postRepo.GetById(ctx,postId)
	if err!=nil{
		return nil,err
	}

	if post.AuthorID!=userID{
		return nil,fmt.Errorf("post does not exist")
	}

	tags,err:=uc.repo.GetOrCreateTags(ctx,tagNames)
	if err!=nil{
		return nil,err
	}

	if err:=uc.postRepo.AppendTags(ctx,post,tags);err!=nil{
		return nil,err
	}

	return tags,nil
}
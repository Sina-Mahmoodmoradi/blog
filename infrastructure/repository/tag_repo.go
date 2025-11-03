package repository

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"gorm.io/gorm"
)



type TagRepository struct{
	db *gorm.DB
}



func NewTagRepository(db *gorm.DB) usecase.TagRepository{
	return &TagRepository{
		db:db,
	}
}



func (r *TagRepository) GetOrCreateTags(ctx context.Context,tagNames []string)([]*entity.Tag,error){
	var existing []models.Tag 
	if err:=r.db.WithContext(ctx).Where("name IN ?",tagNames).Find(&existing).Error;err!=nil{
		return nil,fmt.Errorf("faild to get tags: %w",err)
	}

	tagMap := make(map[string]models.Tag)
	for _,tag:=range existing{
		tagMap[tag.Name] = tag
	}

	newTags := []models.Tag{}
	for _,t:= range tagNames{
		if _,exists:=tagMap[t];!exists{
			newTags = append(newTags, models.Tag{Name: t})
		}
	}

	if len(newTags)>0{
		if err:=r.db.WithContext(ctx).Create(&newTags).Error;err!=nil{
			return nil,fmt.Errorf("faild to create new tags: %w",err)
		}
		existing = append(existing, newTags...)
	}

	tags := make([]*entity.Tag,len(existing))
	for i,t :=range existing{
		tags[i] = ToEntityTag(&t)
	}

	return tags,nil

}
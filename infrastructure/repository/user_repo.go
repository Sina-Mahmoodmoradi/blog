package repository

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"gorm.io/gorm"
)



type UserRepository struct{
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) usecase.UserRepository{
	return &UserRepository{
		db: db,
	}
}


func (r *UserRepository)FindByEmail(ctx context.Context,email string)(*entity.User,error){
	user:= models.User{}
	if err:=r.db.WithContext(ctx).First(&user,"email = ?",email).Error;err!=nil{
		return nil,err
	}

	return ToEntity(&user),nil
}

	
func (r *UserRepository)FindByUsername(ctx context.Context,username string) (*entity.User, error){
	user:= models.User{}
	if err:=r.db.WithContext(ctx).First(&user,"username = ?",username).Error;err!=nil{
		return nil,err
	}

	return ToEntity(&user),nil
}

func (r *UserRepository)Save(ctx context.Context,user *entity.User) error{
	
	modelUser := ToModel(user)
	if err:=r.db.WithContext(ctx).Create(&modelUser).Error;err!=nil{
		return fmt.Errorf("failed to create user: %w",err)
	}

	user.ID = modelUser.ID
	return nil
}



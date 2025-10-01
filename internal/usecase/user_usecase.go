package usecase

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)


type UserUseCase struct{
	repo UserRepository
	hasher PasswordHasher 
}

type RegisterUserRequest struct{
	Username string
	Email string
	Password string 
}

type RegisterUserResponse struct{
	ID uint
	Username string
	Email string
}

func NewUserUseCase(repo UserRepository,hasher PasswordHasher) *UserUseCase{
	return &UserUseCase{
		repo: repo,
		hasher: hasher,
	}
}


func (u *UserUseCase) Register(ctx context.Context,req *RegisterUserRequest)(*RegisterUserResponse ,error){
	if existing,_:=u.repo.FindByUsername(ctx,req.Username);existing!=nil{
		return nil,fmt.Errorf("user with this username exists")
	}
	
	if existing,_:=u.repo.FindByEmail(ctx,req.Email);existing!=nil{
		return nil,fmt.Errorf("user with this email exists")
	}

	hashedPassword,err := u.hasher.Hash(req.Password)
	if err!=nil{
		return nil,fmt.Errorf("hash failed: %w",err)
	}

	user := &entity.User{
		Username: req.Username,
		Email: req.Email,
		PasswordHash: hashedPassword,
	}
	
	if err:=u.repo.Save(ctx,user);err!=nil{
		return nil,fmt.Errorf("failed to create user: %w",err)
	}

	return &RegisterUserResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
	},nil

}
package usecase

import (
	"context"
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)


type UserUseCase struct{
	Repo UserRepository
	Hasher PasswordHasher 
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
		Repo: repo,
		Hasher: hasher,
	}
}


func (u *UserUseCase) Register(ctx context.Context,req *RegisterUserRequest)(*RegisterUserResponse ,error){
	if existing,_:=u.Repo.FindByUsername(ctx,req.Username);existing!=nil{
		return nil,fmt.Errorf("user with this username exists")
	}
	
	if existing,_:=u.Repo.FindByEmail(ctx,req.Email);existing!=nil{
		return nil,fmt.Errorf("user with this email exists")
	}

	hashedPassword,err := u.Hasher.Hash(req.Password)
	if err!=nil{
		return nil,fmt.Errorf("hash failed: %w",err)
	}

	user := &entity.User{
		Username: req.Username,
		Email: req.Email,
		PasswordHash: hashedPassword,
	}
	
	if err:=u.Repo.Save(ctx,user);err!=nil{
		return nil,fmt.Errorf("failed to create user: %w",err)
	}

	return &RegisterUserResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
	},nil

}
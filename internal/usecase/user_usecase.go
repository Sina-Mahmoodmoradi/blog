package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)


type UserUseCase struct{
	repo UserRepository
	hasher PasswordHasher 
	tokenManager TokenManager
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


type LoginRequest struct{
	Identifier string
	Password string
}

type LoginResponse struct{
	Token string
}



func NewUserUseCase(repo UserRepository,hasher PasswordHasher,tokenManager TokenManager) *UserUseCase{
	return &UserUseCase{
		repo: repo,
		hasher: hasher,
		tokenManager: tokenManager,
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



func (u *UserUseCase)Login(ctx context.Context, req *LoginRequest) (*LoginResponse,error){
	var user *entity.User
	var err error
	if strings.Contains(req.Identifier,"@"){
		user,err = u.repo.FindByEmail(ctx,req.Identifier)
	}else{
		user,err = u.repo.FindByUsername(ctx,req.Identifier)
	}
	
	if err!=nil || user==nil{
		return nil,fmt.Errorf("invalid credential")
	}

	if !u.hasher.Compare(user.PasswordHash,req.Password){
		return nil,fmt.Errorf("invalid credential")
	}

	token,err := u.tokenManager.CreateToken(user.ID,time.Hour)
	if err!=nil{
		return nil,fmt.Errorf("failed to generate token: %w",err)
	}

	return &LoginResponse{Token: token},nil
}
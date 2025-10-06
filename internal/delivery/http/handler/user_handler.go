package handler

import (
	"net/http"

	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/middleware"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
)





type userHandler struct{
	useCase *usecase.UserUseCase
	tokenManager usecase.TokenManager
}


type RegisterRequest struct{
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterResponse struct{
	ID uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}
type CurrentUserResponse = RegisterResponse
type LoginRequest struct{
	Identifier string `json:"identifier" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct{
	Token string `json:"token"`
}

func NewUserHandler(useCase *usecase.UserUseCase,tokenManager usecase.TokenManager) *userHandler{
	return &userHandler{
		useCase: useCase,
		tokenManager: tokenManager,
	}
}

func (h *userHandler)RegisterRoutes(r *gin.Engine){
	r.POST("/register",h.Register)
	r.POST("/login",h.Login)


	auth:=r.Group("/auth")
	auth.Use(middleware.AuthMiddleware(h.tokenManager))
	{
		auth.GET("/me",h.getCurrentUser)
	}
}



func (h *userHandler)Register(c *gin.Context){
	var req RegisterRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	ctx:= c.Request.Context()
	ucResponse,err := h.useCase.Register(ctx,&usecase.RegisterUserRequest{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	})

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,RegisterResponse{
		ID: ucResponse.ID,
		Username: ucResponse.Username,
		Email: ucResponse.Email,
	})

}

func (h *userHandler)Login(c *gin.Context){
	ctx := c.Request.Context()
	var req LoginRequest

	if err:=c.ShouldBind(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	ucReq := usecase.LoginRequest{
		Identifier: req.Identifier,
		Password: req.Password,
	}
	ucRes,err := h.useCase.Login(ctx,&ucReq)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,LoginResponse{
		Token: ucRes.Token,
	})	
}


func (h *userHandler)getCurrentUser(c *gin.Context){
	ctx := c.Request.Context()

	userID,ok := c.Get("userID")
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
	}

	user,err := h.useCase.GetByID(ctx,userID.(uint))
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusOK,CurrentUserResponse{
		Username: user.Username,
		Email: user.Email,
		ID: user.ID,
	})

}

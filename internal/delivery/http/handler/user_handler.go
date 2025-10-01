package handler

import (
	"net/http"

	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
)





type userHandler struct{
	useCase *usecase.UserUseCase
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

func NewUserHandler(useCase *usecase.UserUseCase) *userHandler{
	return &userHandler{
		useCase: useCase,
	}
}

func (h *userHandler)RegisterRoutes(r *gin.Engine){
	r.POST("/register",h.Register)
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
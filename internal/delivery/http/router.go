package http

import (
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/passwordhasher"
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository"
	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/handler"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func SetupRouter(db *gorm.DB)*gin.Engine{
	r := gin.Default()

	//health check
	healthHandler := handler.NewHealthHandler(db)
	healthHandler.RegisterRoutes(r)

	//passwordHasher
	hasher := passwordhasher.NewBcryptHasher()

	//Repos
	userRepo := repository.NewUserRepository(db)

	//UseCase
	userUseCase := usecase.NewUserUseCase(userRepo,hasher)

	//handler
	handler := handler.NewUserHandler(userUseCase)
	handler.RegisterRoutes(r)
	
	return r
}
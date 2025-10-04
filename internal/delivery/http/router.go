package http

import (
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/passwordhasher"
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository"
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/tokenmanager"
	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/handler"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/Sina-Mahmoodmoradi/blog/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func SetupRouter(db *gorm.DB,cfg *config.Config)*gin.Engine{
	r := gin.Default()

	//health check
	healthHandler := handler.NewHealthHandler(db)
	healthHandler.RegisterRoutes(r)

	//passwordHasher
	hasher := passwordhasher.NewBcryptHasher()

	//tokenManager
	tokenManager:=tokenmanager.NewJWTTokenManager(cfg.JWTSecret)

	//Repos
	userRepo := repository.NewUserRepository(db)

	//UseCase
	userUseCase := usecase.NewUserUseCase(userRepo,hasher,tokenManager)

	//handler
	handler := handler.NewUserHandler(userUseCase)
	handler.RegisterRoutes(r)

	return r
}
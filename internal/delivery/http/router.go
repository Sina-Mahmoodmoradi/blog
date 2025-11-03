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
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	tagRepo := repository.NewTagRepository(db)

	//UseCase
	userUseCase := usecase.NewUserUseCase(userRepo,hasher,tokenManager)
	postUseCase := usecase.NewPostUseCase(postRepo,commentRepo)
	commentUseCase := usecase.NewCommentUseCase(commentRepo,postRepo)
	tagUseCase := usecase.NewTagUseCase(tagRepo,postRepo)


	//handler
	userHandler := handler.NewUserHandler(userUseCase,tokenManager)
	postHandler := handler.NewPostHandler(postUseCase,tokenManager)
	commentHandler := handler.NewCommentHandler(commentUseCase,tokenManager)
	tagHandler := handler.NewTagHandler(tagUseCase,tokenManager)

	//registering routes
	userHandler.RegisterRoutes(r)
	postHandler.RegisterRoutes(r)
	commentHandler.RegisterRoutes(r)
	tagHandler.RegisterRoutes(r)

	return r
}
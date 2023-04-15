package main

import (
	"final-project/controller"
	"final-project/middleware"
	"final-project/model"
	"final-project/repository"
	"final-project/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=rabbani11 dbname=golang-final sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Comment{}, model.Photo{}, model.SocialMedia{}, model.User{})

	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	socialMediaRepository := repository.NewSocialMediaRepository(db)

	photoController := controller.NewPhotoController(*photoRepository)
	commentController := controller.NewCommentController(*commentRepository)
	socialMediaController := controller.NewSocialMediaController(*socialMediaRepository)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	g := gin.Default()

	g.POST("/user/register", userController.Register)
	g.POST("/user/login", userController.Login)

	photoGroup := g.Group("/photo", middleware.AuthMiddleware)
	photoGroup.GET("/", photoController.GetAllPhoto)
	photoGroup.POST("/", photoController.CreatePhoto)
	photoGroup.DELETE("/:id", photoController.DeletePhoto)
	photoGroup.GET("/:id", photoController.GetPhotoById)
	photoGroup.PUT("/:id", photoController.UpdatePhoto)

	commentGroup := g.Group("/comment", middleware.AuthMiddleware)
	commentGroup.GET("/", commentController.GetAllComment)
	commentGroup.POST("/", commentController.CreateComment)
	commentGroup.DELETE("/:id", commentController.DeleteComment)
	commentGroup.GET("/:id", commentController.GetCommentById)
	commentGroup.PUT("/:id", commentController.UpdateComment)

	socialMediaGroup := g.Group("/socialMedia", middleware.AuthMiddleware)
	socialMediaGroup.GET("/", socialMediaController.GetAllSocialMedia)
	socialMediaGroup.POST("/", socialMediaController.CreateSocialMedia)
	socialMediaGroup.DELETE("/:id", socialMediaController.DeleteSocialMedia)
	socialMediaGroup.GET("/:id", socialMediaController.GetSocialMediaById)
	socialMediaGroup.PUT("/:id", socialMediaController.UpdateSocialMedia)

	g.Run(":8080")
}

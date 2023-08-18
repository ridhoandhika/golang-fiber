package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"golang-fiber/app"
	"golang-fiber/controller"
	"golang-fiber/exception"
	"golang-fiber/helper"
	"golang-fiber/middleware"
	"golang-fiber/repository"
	"golang-fiber/service"
)

func main() {
	//load environment
	envErr := godotenv.Load(".env")
	if envErr != nil {
		helper.PanicError(envErr)
	}
	// panggil connection db
	db := app.Database()

	//validator
	validate := validator.New()

	//repositor
	userRepository := repository.NewUserRepositoryImpl()

	//service
	userService := service.NewUserServiceImpl(
		userRepository,
		db,
		validate,
	)

	// controller
	userController := controller.NewUserControllerImpl(userService)

	// initialize router
	router := httprouter.New()

	//route
	//[USER]
	router.POST("/api/v1/user", userController.Create)
	router.POST("/api/v1/auth", userController.Auth)
	router.POST("/api/v1/refresh-token", userController.CreateWithRefreshToken)
	router.PUT("/api/v1/user/:user_id", userController.Update)
	router.DELETE("/api/v1/user/:user_id", userController.Delete)
	router.GET("/api/v1/user/:user_id", userController.FindById)
	router.GET("/api/v1/user", userController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}

package main

import (
	"net/http"

	"github.com/faridlan/daily/test/app"
	"github.com/faridlan/daily/test/controller"
	"github.com/faridlan/daily/test/helper"
	"github.com/faridlan/daily/test/repository"
	"github.com/faridlan/daily/test/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	db := app.NewConnection()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserRepositoryImpl(db, userRepo)
	userController := controller.NewUserController(userService)

	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:id", userController.Update)
	router.DELETE("/api/user/:id", userController.Delete)
	router.GET("/api/user/:id", userController.FindById)
	router.GET("/api/user/", userController.FindAll)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

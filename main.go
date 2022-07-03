package main

import (
	validator2 "github.com/go-playground/validator"
	"net/http"
	"simple-oauth-service/controller"
	"simple-oauth-service/database"
	"simple-oauth-service/helper"
	"simple-oauth-service/repository"
	"simple-oauth-service/router"
	"simple-oauth-service/service"
)

func main() {
	db := database.NewDB()
	rdb := database.NewRedisClient()
	validator := validator2.New()

	oauth2Repository := repository.NewOauth2Repository()
	oauthService := service.NewOauth2Service(oauth2Repository, db, rdb, validator)
	oauthController := controller.NewOauth2Controller(oauthService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validator)
	userController := controller.NewUserController(userService)

	userRoleRepository := repository.NewUserRoleRepository()
	userRoleService := service.NewUserRoleService(userRoleRepository, db, validator)
	userRoleController := controller.NewUserRoleController(userRoleService)

	dataScopeRepository := repository.NewDataScopeRepository()
	dataScopeService := service.NewDataScopeService(dataScopeRepository, db, validator)
	dataScopeController := controller.NewDataScopeController(dataScopeService)

	clientRepository := repository.NewClientRepository()
	clientService := service.NewClientService(clientRepository, db, validator)
	clientController := controller.NewClientController(clientService)

	r := router.NewRouter(oauthController, userController, userRoleController, dataScopeController, clientController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

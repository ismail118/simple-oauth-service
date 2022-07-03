package test

import (
	"fmt"
	validator2 "github.com/go-playground/validator"
	"net/http"
	"simple-oauth-service/controller"
	"simple-oauth-service/database"
	"simple-oauth-service/helper"
	"simple-oauth-service/repository"
	"simple-oauth-service/router"
	"simple-oauth-service/service"
	"testing"
)

func TestOauthController(t *testing.T) {
	db := database.NewDB()
	rdb := database.NewRedisClient()
	validator := validator2.New()
	oauth2RepositoryMock := repository.NewOauth2RepositoryMock()
	oauthService := service.NewOauth2Service(oauth2RepositoryMock, db, rdb, validator)
	oauthController := controller.NewOauth2Controller(oauthService)

	router := router.NewRouter(oauthController)
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("welcome test")
	}).Methods(http.MethodGet, http.MethodPost)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

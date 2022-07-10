package test

import (
	"database/sql"
	"fmt"
	validator2 "github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
	"net/http"
	"simple-oauth-service/controller"
	"simple-oauth-service/database"
	"simple-oauth-service/helper"
	"simple-oauth-service/repository"
	"simple-oauth-service/router"
	"simple-oauth-service/service"
	"testing"
	"time"
)

func NewDBTest() *sql.DB {
	db, err := sql.Open("mysql", "root:Colonelgila123@tcp(localhost:3306)/auth?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestSimpleOauthService(t *testing.T) {
	db := NewDBTest()
	rdb := database.NewRedisClient()
	validator := validator2.New()
	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"oauthserver99@gmail.com",
		"tvvlczggnxyvmham",
	)

	oauth2Repository := repository.NewOauth2Repository()
	oauthService := service.NewOauth2Service(oauth2Repository, db, rdb, validator)
	oauthController := controller.NewOauth2Controller(oauthService)

	userRoleRepository := repository.NewUserRoleRepository()
	userRoleService := service.NewUserRoleService(userRoleRepository, db, validator)
	userRoleController := controller.NewUserRoleController(userRoleService)

	dataScopeRepository := repository.NewDataScopeRepository()
	dataScopeService := service.NewDataScopeService(dataScopeRepository, db, validator)
	dataScopeController := controller.NewDataScopeController(dataScopeService)

	clientRepository := repository.NewClientRepository()
	clientService := service.NewClientService(clientRepository, db, validator)
	clientController := controller.NewClientController(clientService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, userRoleRepository, db, rdb, validator, dialer)
	userController := controller.NewUserController(userService)

	r := router.NewRouter(oauthController, userController, userRoleController, dataScopeController, clientController)

	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("welcome test")
	}).Methods(http.MethodGet, http.MethodPost)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

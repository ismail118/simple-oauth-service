package test

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	validator2 "github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
	"net/http"
	"net/url"
	"simple-oauth-service/controller"
	"simple-oauth-service/database"
	"simple-oauth-service/helper"
	"simple-oauth-service/repository"
	"simple-oauth-service/router"
	"simple-oauth-service/service"
	"testing"
	"time"
)

var State = "t3xegaMHDPYD5NwB"

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

	r.HandleFunc("/test/login", func(writer http.ResponseWriter, request *http.Request) {
		clientId := 1
		redirectUrl := "http://localhost:3000/test/callback"

		oauthUrl := fmt.Sprintf("http://localhost:3000/oauth/authorize?client_id=%d&redirect_url=%s&state=%s", clientId, redirectUrl, State)
		http.Redirect(writer, request, oauthUrl, http.StatusPermanentRedirect)
	}).Methods(http.MethodGet)

	r.HandleFunc("/test/callback", func(writer http.ResponseWriter, request *http.Request) {
		u, err := url.Parse(request.URL.String())
		helper.PanicIfError(err)

		code := u.Query()["code"][0]
		state := u.Query()["state"][0]
		if state != State {
			_, err2 := fmt.Fprintln(writer, "invalid state")
			helper.PanicIfError(err2)
			return
		}

		reqBody := map[string]interface{}{
			"code":          code,
			"client_id":     1,
			"client_secret": "ONRhfKsUOHoF8iV",
		}
		result, err := helper.SendPostHttpRequest("http://localhost:3000/oauth/access_token", reqBody)
		helper.PanicIfError(err)

		resBody := struct {
			Code   int    `json:"code"`
			Status string `json:"status"`
			Data   struct {
				AccessToken  string `json:"access_token"`
				RefreshToken string `json:"refresh_token"`
			} `json:"data"`
		}{}

		err = json.Unmarshal([]byte(result), &resBody)
		helper.PanicIfError(err)

		cookie := &http.Cookie{
			Name:  "jid",
			Value: resBody.Data.RefreshToken,
			Path:  "/",
		}

		helper.WriteToResponseBodyWithCookie(writer, cookie, resBody)
	}).Methods(http.MethodPost)

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

func TestError(t *testing.T) {
	err := http.ErrNoCookie
	fmt.Println(errors.Is(err, http.ErrNoCookie))
}

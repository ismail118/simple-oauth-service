package router

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"simple-oauth-service/controller"
	"simple-oauth-service/helper"
	"simple-oauth-service/middleware"
	"simple-oauth-service/templates"
)

func NewRouter(oauth2Controller controller.OAuth2Controller,
	userController controller.UserController,
	userRoleController controller.UserRoleController,
	dataScopeController controller.DataScopeController,
	clientController controller.ClientController) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/oauth/authorize", oauth2Controller.Authorize).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/oauth/login", oauth2Controller.Login).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/oauth/access_token", oauth2Controller.AccessToken).Methods(http.MethodPost)
	router.HandleFunc("/oauth/refresh_token", oauth2Controller.RefreshToken).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/oauth/revoke_refresh_token", oauth2Controller.RevokeRefreshToken).Methods(http.MethodPost)
	router.HandleFunc("/login", oauth2Controller.InternalLogin).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/api/user", userController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/user", userController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/user/{userId}", userController.FindById).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{userId}", userController.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/user/{userId}", userController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/user/change_password/{userId}", userController.ChangePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/user/validate", userController.Validate).Methods(http.MethodPost)

	router.HandleFunc("/api/user_role", userRoleController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/user_role", userRoleController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/user_role/{userRoleId}", userRoleController.FindById).Methods(http.MethodGet)
	router.HandleFunc("/api/user_role/{userRoleId}", userRoleController.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/user_role/{userRoleId}", userRoleController.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/data_scope", dataScopeController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/data_scope", dataScopeController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/data_scope/{dataScopeId}", dataScopeController.FindById).Methods(http.MethodGet)
	router.HandleFunc("/api/data_scope/{dataScopeId}", dataScopeController.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/data_scope/{dataScopeId}", dataScopeController.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/client", clientController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/client", clientController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/client/{clientId}", clientController.FindById).Methods(http.MethodGet)
	router.HandleFunc("/api/client/{clientId}", clientController.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/client/{clientId}", clientController.Delete).Methods(http.MethodDelete)

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.Use(middleware.PanicRecovery, middleware.AuthorizationHandler)

	return router
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	t, err2 := template.ParseFS(templates.Templates, "*.gohtml")
	helper.PanicIfError(err2)

	err2 = t.ExecuteTemplate(w, "page404.gohtml", nil)
	helper.PanicIfError(err2)
}

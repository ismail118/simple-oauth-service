package middleware

import (
	"github.com/go-playground/validator"
	"net/http"
	"simple-oauth-service/constanta"
	"simple-oauth-service/errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/response"
)

func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				ErrorHandler(w, r, err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if noFoundError(w, r, err) {
		return
	}
	if validationError(w, r, err) {
		return
	}
	if validationCustomError(w, r, err) {
		return
	}
	if unauthorizedError(w, r, err) {
		return
	}
	if forbiddenError(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: constanta.Status500,
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func noFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(errors.NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: constanta.Status404,
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: constanta.Status400,
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationCustomError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(errors.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: constanta.Status400,
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func unauthorizedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(errors.UnauthorizedError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: constanta.Status401,
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func forbiddenError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(errors.ForbiddenError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)

		webResponse := response.WebResponse{
			Code:   http.StatusForbidden,
			Status: constanta.Status403,
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

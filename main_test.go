package main

import (
	"net/http"
	"testing"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/kataras/iris/httptest"
)

func TestIrisThings(t *testing.T) {
	var app = NewApp()
	e := httptest.New(t, app, httptest.URL("http://localhost:3002"))

	result := e.POST("/administrator/login").
		WithJSON(&repository.LoginAdministratorParams{Email: "test@gmail.com", Password: "Testing"}).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object()

	result.ValueEqual("error", "Invalid email or password")
}

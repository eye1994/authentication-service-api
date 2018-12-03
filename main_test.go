package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"

	"gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("http://localhost:3002")

const administratorSchema = `{
  "title": "Administrator Schema",
  "type": "object",
  "properties": {
    "first_name": {
      "type": "string"
		},
		"last_name": {
      "type": "string"
		},
		"email": {
      "type": "string"
		},
		"id": {
      "type": "number"
    }
  },
  "required": ["first_name", "last_name", "email", "id"]
}`

func createAdministrator(email, password string) *repository.Administrator {
	passwordHash, _ := utils.GeneratePassword(password)

	administrator := repository.Administrator{
		Email:        email,
		PasswordHash: string(passwordHash),
		FirstName:    "Test",
		LastName:     "Testing",
	}

	repository.DB.Create(&administrator)
	return &administrator
}

func clearDatabase() {
	repository.DB.Exec("DELETE FROM administrators")
	repository.DB.Exec("DELETE FROM application_administrators")
	repository.DB.Exec("DELETE FROM users")
	repository.DB.Exec("DELETE FROM application_users")
	repository.DB.Exec("DELETE FROM applications")
}

func TestMain(m *testing.M) {
	app := EchoHandler()
	go app.Start(":3002")
	retCode := m.Run()
	app.Close()
	os.Exit(retCode)
}

func TestLoginFailed(t *testing.T) {
	clearDatabase()
	login := &repository.LoginAdministratorParams{
		Email:    "test@gmail.com",
		Password: "testiing",
	}

	test.Post("/administrator/login").
		JSON(login).
		Expect(t).
		Status(http.StatusUnauthorized).
		Type("json").
		JSON(map[string]string{"error": "Invalid email or password"}).
		Done()
}

func TestRegisterSuccess(t *testing.T) {
	clearDatabase()
	register := repository.RegisterAdministratorParams{
		Email:     "test@gmail.com",
		Password:  "testiing",
		FirstName: "Test",
		LastName:  "Testing",
	}

	test.Post("/administrator").
		JSON(register).
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		Done()
}

func TestLoginSuccess(t *testing.T) {
	clearDatabase()
	createAdministrator("test@gmail.com", "testing")

	login := &repository.LoginAdministratorParams{
		Email:    "test@gmail.com",
		Password: "testing",
	}

	test.Post("/administrator/login").
		JSON(login).
		Expect(t).
		Status(http.StatusOK).
		HeaderPresent("Authorization").
		Type("json").
		JSONSchema(administratorSchema).
		Done()
}

func TestProfile(t *testing.T) {
	clearDatabase()
	administrator := createAdministrator("test@gmail.com", "testing")
	token, _ := utils.SignJwt(int(administrator.ID))

	test.Get("/administrator/profile").
		SetHeader("Authorization", "Bearer "+token).
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(administratorSchema).
		Done()
}

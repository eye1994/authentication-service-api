package main

import (
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/eye1994/authentication-service-api/repository"
	"github.com/eye1994/authentication-service-api/utils"
	"github.com/google/uuid"

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

const applicationSchema = `{
  "title": "Administrator Schema",
  "type": "object",
  "properties": {
    "first_name": {
      "type": "string"
		},
		"url": {
      "type": "string"
		},
		"name": {
      "type": "string"
		},
		"id": {
      "type": "number"
    	},
		"public_token": {
      "type": "string"
		},
		"private_token": {
      "type": "string"
		}
  },
  "required": ["url", "name", "public_token"]
}`

const applicationsSchema = `{
  "title": "Administrator Schema",
  "type": "array",
  "properties": {
    "first_name": {
      "type": "string"
		},
		"url": {
      "type": "string"
		},
		"name": {
      "type": "string"
		},
		"id": {
      "type": "number"
    	},
		"public_token": {
      "type": "string"
		},
		"private_token": {
      "type": "string"
		}
  },
  "required": ["url", "name", "public_token"]
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

func createApplication(administrator *repository.Administrator) *repository.Application {
	publicToken, _ := uuid.NewUUID()
	privateToken, _ := uuid.NewUUID()
	application := repository.Application{
		URL:         "test",
		Name:        "testing",
		PublicToken: publicToken.String(),
		SecretToken: privateToken.String(),
	}

	repository.DB.Model(&administrator).Association("Applications").Append(&application)
	return &application
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

func TestCreateApplication(t *testing.T) {
	clearDatabase()
	administrator := createAdministrator("test@gmail.com", "testing")
	token, _ := utils.SignJwt(int(administrator.ID))

	createParams := repository.ApplicationParams{
		Name:             "Test",
		URL:              "http://testing.com",
		OpenRegistration: true,
	}

	test.Post("/application").
		JSON(createParams).
		SetHeader("Authorization", "Bearer "+token).
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(applicationSchema).
		Done()
}

func TestUpdateApplication(t *testing.T) {
	clearDatabase()
	administrator := createAdministrator("test@gmail.com", "testing")
	token, _ := utils.SignJwt(int(administrator.ID))

	updateParams := repository.ApplicationParams{
		Name:             "Testing",
		URL:              "http://test.com",
		OpenRegistration: true,
	}

	application := createApplication(administrator)
	url := "/application/" + strconv.Itoa(int(application.ID))

	test.Put(url).
		JSON(updateParams).
		SetHeader("Authorization", "Bearer "+token).
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(applicationSchema).
		Done()
}

func TestListApplication(t *testing.T) {
	clearDatabase()
	administrator := createAdministrator("test@gmail.com", "testing")
	token, _ := utils.SignJwt(int(administrator.ID))
	createApplication(administrator)

	test.Get("/application").
		SetHeader("Authorization", "Bearer "+token).
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(applicationsSchema).
		Done()
}

func TestDeleteApplication(t *testing.T) {
	clearDatabase()
	administrator := createAdministrator("test@gmail.com", "testing")
	token, _ := utils.SignJwt(int(administrator.ID))
	application := createApplication(administrator)
	url := "/application/" + strconv.Itoa(int(application.ID))

	test.Delete(url).
		SetHeader("Authorization", "Bearer "+token).
		Expect(t).
		Status(200).
		Done()
}

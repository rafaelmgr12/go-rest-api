package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelmgr12/go-rest-api/internal/database"
	"github.com/rafaelmgr12/go-rest-api/internal/models"
	"github.com/rafaelmgr12/go-rest-api/internal/utils"
	"github.com/steinfletcher/apitest"
)

func newApp() *fiber.App {
	var app *fiber.App = newApp()

	database.InitDatabase(utils.GetValue("DB_TEST_NAME"))

	return app
}

func getItem() models.Item {
	database.InitDatabase(utils.GetValue("DB_TEST_NAME"))
	item, err := database.SeedItem()
	if err != nil {
		panic(err)
	}

	return item
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		database.CleanSeeders()
	}
}

func TestSignup_Success(t *testing.T) {
	userData, err := utils.CreateFaker[models.User]()

	if err != nil {
		panic(err)
	}

	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    userData.Email,
		Password: userData.Password,
	}

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestSignup_ValidationFailed(t *testing.T) {
	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    "",
		Password: "",
	}

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	database.InitDatabase(utils.GetValue("DB_TEST_NAME"))
	user, err := database.SeedUser()
	if err != nil {
		panic(err)
	}

	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    "",
		Password: "",
	}

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *models.UserRequest = &models.UserRequest{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/v1/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

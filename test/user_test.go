package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	app := fiber.New()

	// Connect to test database
	db, err := sql.Open("postgres", "user=postgres password=password dbname=chat_app_test sslmode=disable")
	if err != nil {
		t.Errorf("Failed to connect to the test database: %v", err)
	}
	defer db.Close()

	// setup routing dan controller di sini

	// test case
	body := map[string]string{
		"email":    "test@email.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	// validasi response
	assert.Equal(t, http.StatusOK, resp.StatusCode, "status code harus sama dengan 200")
	responseData, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "User successfully registered", string(responseData), "response harus sama dengan 'User successfully registered'")

}

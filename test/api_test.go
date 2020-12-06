package test

import (
	"../api"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ID string

var postRequestBody = `{
		"firstName": "John",
		"lastName": "Smith",
		"age": 56,
		"address": "123 Main Street, New York, NY 10030"
	}`

var getResponseBody = `{"id":"%s","firstName":"John","lastName":"Smith","age":56,"address":"123 Main Street, New York, NY 10030"}`

var putRequestBody = `{
		"firstName": "James",
		"lastName": "Brown",
		"age": 56,
		"address": "123 Main Street, New York, NY 10030"
	}`

var putResponseBody = `{"id":"%s","firstName":"James","lastName":"Brown","age":56,"address":"123 Main Street, New York, NY 10030"}`

func TestGetAll(t *testing.T) {
	router := api.Router()

	request, _ := http.NewRequest("GET", "/peoples", nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, 200, responseRecorder.Code)
}

func TestPost(t *testing.T) {
	router := api.Router()

	body := bytes.NewBuffer([]byte(postRequestBody))
	request, _ := http.NewRequest(http.MethodPost, "/peoples", body)
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())

	ID = responseRecorder.Body.String()
}

func TestGet(t *testing.T) {
	router := api.Router()

	request, _ := http.NewRequest(http.MethodGet, "/peoples/"+ID, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, fmt.Sprintf(getResponseBody, ID), responseRecorder.Body.String())
}

func TestPut(t *testing.T) {
	router := api.Router()

	body := bytes.NewBuffer([]byte(putRequestBody))
	request, _ := http.NewRequest(http.MethodPut, "/peoples/"+ID, body)
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, fmt.Sprintf(putResponseBody, ID), responseRecorder.Body.String())
}

func TestDelete(t *testing.T) {
	router := api.Router()

	request, _ := http.NewRequest(http.MethodDelete, "/peoples/"+ID, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusNoContent, responseRecorder.Code)
}

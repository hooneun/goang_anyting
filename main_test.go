package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hooneun/golang_anyting/api"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	h, _ := api.API()
	router := setupRouter(h)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
}

func TestGetUserByIdRoute(t *testing.T) {
	h, _ := api.API()
	router := setupRouter(h)
	w := httptest.NewRecorder()
	id := 1
	req, _ := http.NewRequest("GET", "/user/"+strconv.Itoa(id), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.JSONEq(t, string(jsonUser), w.Body.String())
}

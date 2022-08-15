package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImagesByTopic(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/images", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

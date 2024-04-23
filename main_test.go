package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	//totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
}

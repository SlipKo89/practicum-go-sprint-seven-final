package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerForCorrectResponseAndBodyIsNotEmpty(t *testing.T) {
	//Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerForUnsupportCity(t *testing.T) {
	//Город, который передаётся в параметре city, не поддерживается.
	//Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.

	req := httptest.NewRequest("GET", "/cafe?count=2&city=Barnaul", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerForBigCount(t *testing.T) {
	//Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.

	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=999999", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

	assert.Equal(t, responseRecorder.Body.String(), "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент")
}

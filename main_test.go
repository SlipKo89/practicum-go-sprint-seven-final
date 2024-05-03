package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	//Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	//body_len := responseRecorder.Body.Len()
	//assert.Greater(t, body_len, 0)
	// Alternate variant
	assert.NotEmpty(t, responseRecorder.Body.Len())
}

func Test2(t *testing.T) {
	//Город, который передаётся в параметре city, не поддерживается.
	//Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.

	correct_reponse := "wrong city value"

	req := httptest.NewRequest("GET", "/cafe?count=2&city=Barnaul", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//fmt.Println(responseRecorder.Code)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	//fmt.Println(responseRecorder.Body.String())
	assert.Equal(t, responseRecorder.Body.String(), correct_reponse)
}

func Test3(t *testing.T) {
	//Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.

	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=999999", nil)
	req2 := httptest.NewRequest("GET", "/cafe?city=moscow&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	responseRecorder2 := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	response_with_big_count := responseRecorder.Body.String()

	handler.ServeHTTP(responseRecorder2, req2)

	response_with_true_count := responseRecorder2.Body.String()

	//fmt.Println(response_with_true_count, response_with_big_count)
	assert.Equal(t, response_with_big_count, response_with_true_count)
}

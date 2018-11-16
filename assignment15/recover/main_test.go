package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncThatPanics(t *testing.T) {
	assert.Panics(t, funcThatPanics, "The code is panicing")
}

func TestHello(t *testing.T) {
	//req, err := http.NewRequest("GET", "/", nil)
	//if err != nil {
	//	t.Fatal(err)
	//}

	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(hello)

	//handler.ServeHTTP(rr, req)

	rr, err := executeRequest("GET", "/", devMw(http.HandlerFunc(hello)))
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, rr.Code, http.StatusOK)

	expected := strings.Contains(rr.Body.String(), "<h1>Hello!</h1>")
	assert.Equalf(t, true, expected, "they should be equal")
}

func TestPanicAfterDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "/panic-after", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	panicAfterDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}
func TestDevMv(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	executeRequest("Get", "/panic", devMw(handler))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

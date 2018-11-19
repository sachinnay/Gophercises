package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	main()
	//os.Exit(m.Run())
	dashtest.ControlCoverage(m)
}
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
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("checking errror", err)
		}
		fmt.Println("checking errror")
	}()
	panicAfterDemo(rec, req)
	res := rec.Result()
	checkResponseCode(t, res.StatusCode, http.StatusInternalServerError)
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

func TestSourceCodeHandler_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()

	sourceCodeHandler(rec, req)
	checkResponseCode(t, http.StatusInternalServerError, rec.Code)
}
func TestSourceCodeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug?line=24&path=/home/gs-2008/go/src/github.com/sachinnay/Gophercises/assignment15/recover/main_test.go", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()

	sourceCodeHandler(rec, req)

	checkResponseCode(t, http.StatusOK, rec.Code)
}
func TestSourceCodeHandler_Failed(t *testing.T) {

	req, err := http.NewRequest("GET", "/debug?line=24&path=/home/gs-2008/go/src/github.com/sachinnay/Gophercises/assignment15/recover/main1.go", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()

	sourceCodeHandler(rec, req)

	checkResponseCode(t, http.StatusInternalServerError, rec.Code)
}

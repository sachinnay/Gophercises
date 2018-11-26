package handler

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment18/transform/primitive"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}

func TestHome(t *testing.T) {
	resp, err := executeRequest("GET", "/", GetMux())
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestUpload(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/birds.png")
	fmt.Println("======>", imgPath)
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	r, _ := http.NewRequest("POST", srv.URL+"/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := http.DefaultClient.Do(r)
	checkResponseCode(t, http.StatusOK, res.StatusCode)

}

// func TestUpload_ContentError(t *testing.T) {
// 	srv := httptest.NewServer(GetMux())
// 	defer srv.Close()
// 	h, _ := homedir.Dir()
// 	imgPath := filepath.Join(h, "img/birds.png")
// 	fmt.Println("======>", imgPath)
// 	file, err := os.Open(imgPath)
// 	if err != nil {
// 		t.Error("error in opening file")
// 	}
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormFile("image", file.Name())
// 	if err != nil {
// 		t.Error("error in copy")
// 	}
// 	_, err = io.Copy(part, file)
// 	if err != nil {
// 		t.Error("error in copy")
// 	}
// 	err = writer.Close()
// 	if err != nil {
// 		t.Error("error in close writer")
// 	}
// 	r, _ := http.NewRequest("POST", srv.URL+"/upload", body)
// 	r.Header.Set("Content-Type", writer.FormDataContentType())
// 	res, _ := http.DefaultClient.Do(r)
// 	checkResponseCode(t, http.StatusOK, res.StatusCode)

// }

func TestUpload_Error(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/birds.png")
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("test", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	r, _ := http.NewRequest("POST", srv.URL+"/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := http.DefaultClient.Do(r)
	checkResponseCode(t, http.StatusInternalServerError, res.StatusCode)

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

func TestTempFile(t *testing.T) {
	_, err := tempfile("/invalid/invalid", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

func TestModifyMode(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/birds.png?mode=3", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusOK, res.StatusCode)
}

func TestModify_ModeNegative(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/birds.png?mode=a", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestModify_ModeNegativeExt(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/test.txt?mode=2", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestModify_ModeNoShapes(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/birds.png?mode=3&n=5", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusOK, res.StatusCode)
}

func TestModifyMode_NegNoShapes(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/birds.png?mode=3&n=a", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestGenImage(t *testing.T) {
	rs := bytes.NewReader(nil)
	mode := primitive.ModeCombo
	_, err := genImage(rs, "txt", -1, mode)
	if err == nil {
		t.Error("Expected error but no error")
	}
}

func TestGenImages_Error(t *testing.T) {
	rs := bytes.NewReader(nil)
	opts := []genOpts{
		{N: -1, M: primitive.ModeCombo},
	}

	_, err := genImages(rs, "txt", opts...)
	if err == nil {
		t.Error("Expected error but no error")
	}
}
func TestRenderModeChoices_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rs := bytes.NewReader(nil)
	rec := httptest.NewRecorder()
	renderModeChoices(rec, req, rs, "txt")
	res := rec.Result()
	checkResponseCode(t, http.StatusInternalServerError, res.StatusCode)
}

func TestRenderNoShapes_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rs := bytes.NewReader(nil)
	rec := httptest.NewRecorder()
	renderNoShapesChoices(rec, req, rs, "txt", primitive.ModeCircle)
	res := rec.Result()
	checkResponseCode(t, http.StatusInternalServerError, res.StatusCode)
}

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// POST /api/email/send/<valid email>
// Body: <valid data>
func TestApiSendMail(t *testing.T) {
	router := Server()

	w := httptest.NewRecorder()
	data := "{\"subject\":\"demo\",\"body\":\"this is body\"}"
	req, _ := http.NewRequest("POST", "/api/email/send/example@gmail.com", strings.NewReader(data))
	router.ServeHTTP(w, req)
	gotStatusCode := w.Code
	wantStatusCode := 200
	if gotStatusCode != wantStatusCode {
		t.Errorf("Got Status Code = %d, Want Status Code = %d", gotStatusCode, wantStatusCode)
	}
	gotBody := string(w.Body.Bytes())
	wantBody := "{\"status\":\"success\"}"
	if gotBody != wantBody {
		t.Errorf("Got Body: '%s', Want Body: '%s'", gotBody, wantBody)
	}
}

// POST /api/email/send/<valid email>
// Body: <invalid data>
func TestApiSendMailInvalidBody(t *testing.T) {
	router := Server()

	w := httptest.NewRecorder()
	data := "{\"sub\":\"demo\",\"body\":\"this is body\"}"
	req, _ := http.NewRequest("POST", "/api/email/send/example@gmail.com", strings.NewReader(data))
	router.ServeHTTP(w, req)
	gotStatusCode := w.Code
	wantStatusCode := http.StatusBadRequest
	if gotStatusCode != wantStatusCode {
		t.Errorf("Got Status Code = %d, Want Status Code = %d", gotStatusCode, wantStatusCode)
	}
	gotBody := string(w.Body.Bytes())
	wantBody := "{\"status\":\"faild\"}"
	if gotBody != wantBody {
		t.Errorf("Got Body: '%s', Want Body: '%s'", gotBody, wantBody)
	}
}

// POST /api/email/send/<invalid email>
// Body: <valid data>
func TestApiSendMailInvalidToAddr(t *testing.T) {
	router := Server()

	w := httptest.NewRecorder()
	data := "{\"subject\":\"demo\",\"body\":\"this is body\"}"
	req, _ := http.NewRequest("POST", "/api/email/send/example@gmailcom", strings.NewReader(data))
	router.ServeHTTP(w, req)
	gotBody := string(w.Body.Bytes())
	wantBody := "{\"status\":\"faild\"}"
	if gotBody != wantBody {
		t.Errorf("Got Body: '%s', Want Body: '%s'", gotBody, wantBody)
	}
}

// POST /api/email/send/<invalid email>
// Body: <invalid data>
func TestApiSendMailInvalidToAddrInvalidBody(t *testing.T) {
	router := Server()

	w := httptest.NewRecorder()
	data := "{\"sub\":\"demo\",\"bod\":\"this is body\"}"
	req, _ := http.NewRequest("POST", "/api/email/send/example@gmailcom", strings.NewReader(data))
	router.ServeHTTP(w, req)
	gotBody := string(w.Body.Bytes())
	wantBody := "{\"status\":\"faild\"}"
	if gotBody != wantBody {
		t.Errorf("Got Body: '%s', Want Body: '%s'", gotBody, wantBody)
	}
}

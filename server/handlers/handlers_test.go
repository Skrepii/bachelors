package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "mike")
	form.Add("password", "wasowski")
	req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	b, _ := ioutil.ReadFile("./static/user.html")
	expected := string(b)

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestLoginHandler_fail(t *testing.T) {
	form := url.Values{}
	form.Add("username", "mike")
	form.Add("password", "rourke")
	req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	expectedBody := "Access Denied"
	if rr.Body.String() != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}

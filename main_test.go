package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "<h1>Profile list</h1>") {
		t.Errorf("Expected title in response, but not found")
	}
}

func TestFormHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/new", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(formHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "<form action=\"/save\" method=\"POST\">") {
		t.Errorf("Expected form in response, but not found")
	}
}

func TestSaveHandler(t *testing.T) {
	form := strings.NewReader("name=Diego&age=30&bio=Developer")
	req, err := http.NewRequest("POST", "/save", form)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(saveHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	if len(profiles) == 0 {
		t.Errorf("New profile was not saved.")
	}
}

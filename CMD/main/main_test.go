package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetBook_StatusOKAndBody(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	handleGetBook(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %d", rr.Result().StatusCode)
	}

	defer rr.Result().Body.Close()
	expected := "BOOK"
	b, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if got := string(b); got != expected {
		t.Errorf("Expected %s; got %s", expected, got)
	}
}

func handleGetBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("BOOK"))
}

func TestHandleGetBook_StatusMethodNotAllowed(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/books", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	handleGetBook(rr, req)

	if rr.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status MethodNotAllowed; got %d", rr.Result().StatusCode)
	}
}
func TestHandleGetBook_HTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetBook))
	defer server.Close()

	resp, err := http.Get(server.URL + "/books")
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := "BOOK"
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("Expected %s; got %s", expected, string(b))
	}
}

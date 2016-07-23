package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloRequest(t *testing.T) {
	router := InitializeRouter()

	s := httptest.NewServer(router)

	defer s.Close()

	r, _ := http.NewRequest("GET", s.URL+"/hello", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)
	got := string(body)
	expect := "I'm middleware\nHello"

	if got != expect {
		t.Errorf("Expected body to be: \n%s\n got:\n%s", expect, got)
	}
}

func TestHelloWorldRequest(t *testing.T) {
	router := InitializeRouter()

	s := httptest.NewServer(router)

	defer s.Close()

	r, _ := http.NewRequest("GET", s.URL+"/hello_world", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	body, _ := ioutil.ReadAll(w.Body)
	got := string(body)
	expect := "I'm middleware\nHello World\n"

	if got != expect {
		t.Errorf("Expected body to be: \n%s\n got:\n%s", expect, got)
	}
}

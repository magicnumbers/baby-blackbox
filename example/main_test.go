package main_test

import (
	"net/http"
	"os"
	"testing"

	blackbox "github.com/meowgorithm/baby-blackbox"

	// Import our main application. Depending on how your project is setup
	// you may need to do this differently. In many cases the common use case
	// is to do something along the lines of:
	//
	// import name "."
	//
	// ...which would normally be a bad circular import thing, but in the case
	// of testing it's fine.
	example "example"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCoolness(t *testing.T) {
	req := example.Response{}

	// Create an HTTP testing thing from our application's multiplexer
	mux := example.Init()
	api := blackbox.New(t, mux)

	// Perform a request
	api.Request("GET", "/", nil).
		OK().      // Assert that we want a `200 OK` response
		JSON(&req) // Unmarshal the response body into a struct

	if !req.Cool {
		t.Error("expected things to be cool, but they were not")
	}

	// In the next request we'll send this payload
	payload := struct {
		Message string `json:"message"`
	}{"Hello, there."}

	// Perform another request, this time with a payload, and expect to get
	// an HTTP `403 Forbidden` error.
	api.Request("POST", "/notcool", payload).
		Status(http.StatusForbidden)

}

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Very basis test
func TestAdd(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Errorf("Got Add(2, 2) = %d, expected 4", Add(2, 2))
	}
}

// Achieving better coverage with Table Driven Tests
type AddResult struct {
	x        int
	y        int
	expected int
}

var addResults = []AddResult{
	{1, 1, 2},
	{2, 2, 4},
}

func TestTableAdd(t *testing.T) {
	for _, test := range addResults {
		output := Add(test.x, test.y)
		if output != test.expected {
			t.Errorf("Got Add(%d, %d) = %d, expected %d", test.x, test.y, output, test.expected)
		}
	}
}

// Similar as above, but shortened code
func TestTableCalc(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{10, 20},
		{-1, 9},
		{0, 10},
		{990, 1000},
	}

	for _, test := range tests {
		if output := Calc(test.input); output != test.expected {
			t.Errorf("Got Calc(%d) = %d, expected %d", test.input, output, test.expected)
		}
	}
}

// Test to see if file contains expected content
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	if string(data) != "hello, world" {
		t.Errorf("Got '%s', expected 'hello, world'", string(data))
	}
}

func TestHttpHealthCheck(t *testing.T) {
	// Create a request to pass to our handler
	// Pass nil because we don't have any query parameters
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder which satisfies http.ResponseWriter to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HttpHealthCheck)

	// The handler satisfies http.Handler, so we can call its ServeHTTP method directly and pass in the Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check if the status code is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Got '%v', expected '%v'", status, http.StatusOK)
	}

	// Check if the response body is as expected
	expected := `{"status": ok}`
	if rr.Body.String() != expected {
		t.Errorf("Got '%v', expected '%v'", rr.Body.String(), expected)
	}
}

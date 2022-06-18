package main

import (
	"io/ioutil"
	"testing"
)

// Very basis test
func TestAdd(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Errorf("Add(2, 2) = %d, expected 4", Add(2, 2))
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
			t.Errorf("Add(%d, %d) = %d, expected %d", test.x, test.y, output, test.expected)
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
			t.Errorf("Calc(%d) = %d, expected %d", test.input, output, test.expected)
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
		t.Errorf("Expected 'hello, world', got '%s'", string(data))
	}
}

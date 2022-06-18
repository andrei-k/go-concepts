package main

import (
	"fmt"
	"io"
	"net/http"
)

func Calc(x int) (result int) {
	result = x + 10
	return result
}

func Add(x int, y int) int {
	return x + y
}

func HttpHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status": ok}`)
}

func main() {
	fmt.Println("Some testing examples")
}

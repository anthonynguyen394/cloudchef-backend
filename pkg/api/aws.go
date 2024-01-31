package api

import (
	"fmt"
	"net/http"
)

func AwsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
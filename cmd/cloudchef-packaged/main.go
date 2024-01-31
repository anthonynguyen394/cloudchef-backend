package main

import (
	"net/http"

	"github.com/anthonynguyen394/cloudchef-backend/pkg/api"
)

func main() {
	http.HandleFunc("/api/aws", api.AwsHandler)
	http.ListenAndServe(":8080", nil)
}

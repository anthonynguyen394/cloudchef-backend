package main

import (
	"net/http"

	"github.com/anthonynguyen394/cloudchef-backend/pkg/aws"
)

func main() {
	http.HandleFunc("/api/aws", aws.Handler)
	http.ListenAndServe(":8080", nil)
}


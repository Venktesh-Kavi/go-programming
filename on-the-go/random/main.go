package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://yubioci-dev-uls-gateway-api.go-yubi.in/apispec/UapiSpec.yaml")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer resp.Body.Close()
	buf := make([]byte, 8192)
	resp.Body.Read(buf)

	fmt.Printf("Response: %s\n", string(buf))
}

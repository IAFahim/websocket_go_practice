package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := router()
	fmt.Println("Listening on port 9806")
	http.ListenAndServe(":9806", mux)
}

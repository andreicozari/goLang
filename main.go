package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Boot the app")

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.ListenAndServe(port(), nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "Hello from a microservice written in GO lang")
}

func echo(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "text/plain")

	message := request.URL.Query()["message"][0]
	fmt.Fprintf(writer, message)
}

func port() string {
	port := ":" + os.Getenv("GO_SERVICE_PORT")

	if len(port) == 0 {
		port = ":8080"
	}

	fmt.Printf("The app is running on http://localhost" + port)
	return port
}

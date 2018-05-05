package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	port := getEnvOrDefault("PORT", "8080")
	log.Println("Listening on port", port)
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":"+port, nil)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	result := defaultValue
	val, ok := os.LookupEnv(key)
	if ok {
		result = val
	}
	return result
}

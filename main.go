package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO: globals are evil - there is a better way
var indexTemplate = GetIndexTemplate()
var hexTemplate = GetHexTemplate()

func myRouter(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "favicon") {
		log.Println(r.URL.Path)
	}
	if strings.ToLower(r.URL.Path) == "/hexcolors" {
		hexController(w, r)
	} else {
		indexTemplate.Execute(w, NoData{})
	}
}

func main() {
	port := getEnvOrDefault("PORT", "8080")
	log.Println("Listening on port", port)
	http.HandleFunc("/", myRouter)
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

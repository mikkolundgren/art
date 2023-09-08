package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mikkolundgren/art/asciiart"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    r.ParseMultipartForm(32 << 20)
    file, _, _ := r.FormFile("data")
	result, _ := asciiart.Draw(file)

    fmt.Fprint(w, result)
}

func main() {
    listenAddr := ":8080"
    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        listenAddr = ":" + val
    }
    http.HandleFunc("/api/AsciiArt", helloHandler)
    log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}
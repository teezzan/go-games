// main.go
package main

import (
    "embed"
    "log"
    "net/http"
    "path"
)

//go:embed static/*
var content embed.FS

func main() {
    // Serve files from the "static" folder
    fs := http.FileServer(http.FS(content))
    http.Handle("/games", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, path.Join("static", "index.html"))
            return
        }
        fs.ServeHTTP(w, r)
    }))
	// return hello world

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}))

    // Start the server
    log.Println("Starting server on :8090")
    if err := http.ListenAndServe(":8090", nil); err != nil {
        log.Fatal(err)
    }
}
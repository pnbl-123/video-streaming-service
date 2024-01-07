// display-service/main.go
package main

import (
    "net/http"
    "log"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./output.mp4"))) // Serve files from the current directory
    err := http.ListenAndServe(":8082", nil)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
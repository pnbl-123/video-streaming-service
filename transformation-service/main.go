package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "time"
)

func main() {
    ln, err := net.Listen("tcp", "localhost:8081")
    if err != nil {
        log.Fatal("Error setting up server:", err)
    }
    defer ln.Close()

    fmt.Println("Server is listening on localhost:8081...")

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal("Error accepting connection:", err)
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    timestamp := time.Now().Format("20060102150405")
    fileName := fmt.Sprintf("output.mp4", timestamp)

    file, err := os.Create(fileName)
    if err != nil {
        log.Fatal("Error creating file:", err)
    }
    defer file.Close()

    buf := make([]byte, 1024)

    for {
        n, err := conn.Read(buf)
        if err != nil && err != io.EOF {
            log.Fatal("Error reading data:", err)
        }

        if n == 0 {
            break
        }

        _, err = file.Write(buf[:n])
        if err != nil {
            log.Fatal("Error writing data:", err)
        }
    }

    fmt.Printf("Video data has been written to %s successfully.\n", fileName)
}
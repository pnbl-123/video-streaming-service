package main

import (
    "fmt"
    "io"
    "net"
    "os"
)

func main() {
    // Create a TCP server to receive data from the transformation service
    ln, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error setting up server:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Server is listening on localhost:8080...")

    for {
        // Accept a connection
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            return
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Open a file to write the video data
    file, err := os.Create("../assets/video.mp4")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    buf := make([]byte, 1024)

    for {
        // Read the data from the connection
        n, err := conn.Read(buf)
        if err != nil && err != io.EOF {
            fmt.Println("Error reading data:", err)
            return
        }

        if n == 0 {
            break
        }

        // Write the data to the file
        _, err = file.Write(buf[:n])
        if err != nil {
            fmt.Println("Error writing data:", err)
            return
        }
    }

    fmt.Println("Video data has been written to output.mp4 successfully.")
}
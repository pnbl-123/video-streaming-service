package utils

import (
    "fmt"
    "io"
    "net"
)

// EstablishConnection creates a TCP connection to the specified address.
func EstablishConnection(address string) (net.Conn, error) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return nil, fmt.Errorf("error establishing connection: %w", err)
    }

    return conn, nil
}

// SendData sends data over the specified connection.
func SendData(conn net.Conn, data []byte) error {
    _, err := conn.Write(data)
    if err != nil {
        return fmt.Errorf("error sending data: %w", err)
    }

    return nil
}

// ReceiveData receives data from the specified connection.
func ReceiveData(conn net.Conn, buf []byte) (int, error) {
    n, err := conn.Read(buf)
    if err != nil && err != io.EOF {
        return 0, fmt.Errorf("error receiving data: %w", err)
    }

    return n, nil
}
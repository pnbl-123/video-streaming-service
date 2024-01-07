package utils

import (
    "io/ioutil"
    "os"
)

// ReadVideoFile reads a video file and returns the data.
func ReadVideoFile(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    return data, nil
}
package main

import (
    "os"
    "io/ioutil"
    "fmt"
)

func getCSS() string {
    path := os.Getenv("GOPATH")
    file, err := ioutil.ReadFile(path + "/src/github.com/carljoshua/gitview/styles.css")

    if err != nil {
        fmt.Println("Unable to load styles")
        os.Exit(1)
    }
    return string(file)
}

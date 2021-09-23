package main

import (
    "fmt"
    "github.com/choryang/echo_example/myrouter"
)

func main() {
    fmt.Println("Welcome to the webserver")
    e := myrouter.New()

    e.Start(":8000")
}
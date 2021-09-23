package main

import (
    "fmt"
    "myrouter"
)

func main() {
    fmt.Println("Welcome to the webserver")
    e := myrouter.New()

    e.Start(":8000")
}
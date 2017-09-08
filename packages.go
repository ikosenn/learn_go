package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(60)
    fmt.Println("Random number is", rand.Intn(10))
}

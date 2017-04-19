package main

import (
    "fmt"
    "math/rand"
    "time"
)

func add(x int, y int) int {
    return x + y 
}

func main() {
    fmt.Println(add(41, 62))
    // rand.Seed(19990)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    fmt.Println(r.Intn(1000))
}
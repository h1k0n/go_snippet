package main

import "fmt"

func main() {
    ch:=make(chan bool)
    go func() {
        fmt.Println("do not block")
        close(ch)
    }()
    <-ch
}
        
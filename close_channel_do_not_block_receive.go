package main

import "fmt"

func main() {
    ch:=make(chan bool)
    go func() {
        fmt.Println("do not block")
        close(ch)
    }()    //这个goroutine不会block住，因为没有send或receive操作
    <-ch   //即便close了也不会阻塞
     fmt.Println("main")
}
        

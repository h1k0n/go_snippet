package main 

import "fmt"

func main() {
    // for i:=0;i<10;i++ {
    //     fmt.Println(i)
    // }
    
    // i:=0
    // for i<10 {
    //     fmt.Println(i)
    //     i++
    // }
    
    var i int
    for  {
        fmt.Println(i)
        i++
        if i>=10 {
            break
        }
    }
}
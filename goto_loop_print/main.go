package main 

import format "fmt"

func main()  {
    i:=0
    Label:{
        if i<10 {
            format.Println(i)
            i++
        }
        goto Label
    }
}
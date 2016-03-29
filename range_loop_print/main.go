package main 

import (
    "fmt"
)

func main()  {
    v:=[...]string{"a","b","c","d","e"}
    for index,value:=range v { //copy v get v1,v1's index and value
        fmt.Println(index,value)
    }
}
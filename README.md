# go_snippet
learning golang by program snippet


package main

import "fmt"

func main() {
	j:=&struct{abc int;def int;}{2,2}
	fmt.Println(j)
	
	a := &int(1)  //cannot take the address of int(1) int(1)是字面值，相当于不能&1
	fmt.Println(a)
}

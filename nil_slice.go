package main

import "fmt"

func main() {
	s :=[]int(nil) //var s []int //usually use in function []int when exception happen,return nill slice
	t :=[]int{} //t := make([]int,3)[3:] t:=make([]int,0  //represent empty data set, such as select return 
	fmt.Println("s ==", s," t == ",t)
	fmt.Println("len(s)=",len(s),"len(t)=",len(t))
	fmt.Println("cap(s)=",cap(s),"cap(t)=",cap(t))
	fmt.Println("s==nil ?",s==nil,"t==nil ?",t==nil)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}

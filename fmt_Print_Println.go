//Go语言-云动力
//fmt包中微小的区别，感叹作者的细心
package main

import "fmt"

func main() {
	var i = 42
	var s = "answer"
    fmt.Print("hello")
    fmt.Print('a')
    fmt.Print('余')
    fmt.Println('余')
    fmt.Printf("%c",97)
    fmt.Printf("%c",'a')
	fmt.Print(s, "is", i, 3.14, '\n', "\n")   //两个数字间自动空一格
    fmt.Println(s,"is",i,"\n",3.14,'\n')  //遇到”，“空一格，自动换行
}

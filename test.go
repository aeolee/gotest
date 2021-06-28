package main

import (
	"fmt"
)


func main()  {
	//A(1,2,3)
	//x()
	hello()
	test()
}

func test(){
	str := "hello world!"
	fmt.Printf("%s\n",str)
}

func A(a,b,c int){
	fmt.Print(a,b,c)
}
//输出不可见字符
func x(){
	fmt.Printf("%#v %#v %#v","	"," ","" +
		"")
}

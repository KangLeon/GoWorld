package main

import (
	"fmt"
)

//函数外面也可以定义变量
//包内变量不可以自动推导
var aa = 3
var ss = "kkk"
var bb= true

//还可以这样
var (
	aa1 = 3
	ss1 = "kkk"
	bb1 = true
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func  variableInitialValue()  {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction()  {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

//一种简写
func variableShorter() {
	a,b,c,s := 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,ss,bb)
	fmt.Println(aa1,ss1,bb1)
}

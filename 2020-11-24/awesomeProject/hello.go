package main

import (
	"fmt"
	"math"
	"math/cmplx"
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

//内建类型：原生支持负数类型
func euler() {
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1)
}

//强制类型转换
func triangle() {
	var a,b int = 3,4
	var c int
	//下面必须强制类型转换，负责会报错
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

//常量
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c= int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c)
}

//枚举
func enums()  {
	//普通枚举类型
	const (
		cpp = iota
		java
		python
		golang
	)

	//自增值的枚举
	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,ss,bb)
	fmt.Println(aa1,ss1,bb1)
	euler()
	consts()
	enums()
}

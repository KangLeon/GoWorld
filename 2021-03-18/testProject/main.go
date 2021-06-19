package main

import "fmt"

var (
	variables int
	slicex []int
	interfacex interface{}
)

func main() {
	test1()
	test2()
}

//变量
func test1()  {
	fmt.Println("Hello,world")

	var variables1 int
	variables2 := "自动推断"//简短声明只可以在函数内部使用，而不可以在函数外部使用

	var variables3 = 100
	var i = 100
	var j = 200
	println(i,j)
	i,j = j,i
	println(i,j)
	var variables4,variables5,variables6,variables7 = 1,"字符串","3.13",true

	fmt.Println(variables1,variables2,variables3,variables4,variables5,variables6,variables7)
}

//常量
func test2()  {
	const constVariables1 float64 = 3.1415926

	const constVariables2,constVariables3 = 100,"字符串"
	fmt.Println(constVariables1,constVariables2,constVariables3)

	const (
		iotaVariables1 = iota
		iotaVariables2 = iota
		iotaVariables3 = iota
	)

	const iotaVariables4 = iota
	fmt.Println(iotaVariables1,iotaVariables2,iotaVariables3,iotaVariables4)

	const (
		iotaVariables5 = iota
		iotaVariables6
		iotaVariables7
	)

	fmt.Println(iotaVariables5,iotaVariables6,iotaVariables7)

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday)

	const (
		iotaVariables8,iotaVariables9,iotaVariables10 = iota,iota,iota
	)

	const (
		iotaVariables11 = iota
		iotaVariables12 = "字符串"
		iotaVariables13 = iota
	)

	fmt.Println(iotaVariables8,iotaVariables9,iotaVariables10,iotaVariables11,iotaVariables12,iotaVariables13)
}
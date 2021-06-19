package main

import "fmt"

func main() {
	fmt.Println("Hello,world")

	var variables1 int
	variables2 := "自动推断"

	var variables3 = 100
	var variables4,variables5,variables6,variables7 = 1,"字符串","3.13",true

	fmt.Print(variables1,variables2,variables3,variables4,variables5,variables6,variables7)
}

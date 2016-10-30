package main

import "fmt"

//import "os"

func main() {

	//预定义
	//ota 比较特殊,可以被认为是一个可被编译器修改的常量,在每一个 const 关键字出现时被
	//重置为0,然后在下一个 const 出现之前,每出现一次 iota ,其所代表的数字会自动增1
	const (
		a = 1 << iota
		b = 1 << iota
	)
	fmt.Println(a) //1
	fmt.Println(b) //2

	//	_, _, la := niming()    只取最后一个返参
	//	fmt.Println(la)
	//	fmt.Println(os.Getenv("Home"))
}

func niming() (first, middle, last string) {
	return "young", "ming", "kai"
}

package main

import "fmt"

//import "strings"

func main() {
	//这样是错误的
	//var v1 int32
	//v2 := 64
	//强制转换
	//v1 = int32(v2)
	//fmt.Println(v1)
	//Go语言定义了两个类型 float32 和 float64 ,其中 float32 等价于C语言的 float 类型,
	//float64 等价于C语言的 double 类型
	//var value1 complex64 //复数

	//value1 = 3.2 + 12i
	//value3 := complex(3.2, 12)

	//字符串的操作
	//	str := "hello,wolrd"
	//	str1 := "youngkai"
	//	fmt.Println(len(str))
	//	fmt.Println(str[3])
	//	fmt.Println(str + str1)

	str := "hello,world"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	//数组
	//	[32]byte// 长度为32的数组,每个元素为一个字节
	//	[2*N] struct { x, y int32 } // 复杂类型数组
	//	[1000]*float64// 指针数组
	//	[3][5]int// 二维数组
	//	[2][2][2]float64// 等同于[2]([2]([2]float64))

	//Go语言还提供了一个关键字 range ,用于便捷地遍历容器中的元素。当然,数组也是 range的支持范围。上面的遍历过程可以简化为如下的写法:
	//	for i, v := range array {
	//		fmt.Println("Array element[", i, "]=", v)
	//	}

	//	arr := [5]int{1, 2, 2, 3, 4}
	//	for k, v := range arr {
	//		fmt.Print(k)
	//		fmt.Print(v)
	//		fmt.Println()
	//	}

	//数组切片-------类似于python的数据切片

	// 先定义一个数组
	//var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组创建一个数组切片
	//	var mySlice []int = myArray[:5]

	//	fmt.Println("Elements of myArray: ")

	//	for _, v := range myArray {

	//		fmt.Print(v, " ")

	//	}

	//	fmt.Println("\nElements of mySlice: ")

	//	for _, v := range mySlice {

	//		fmt.Print(v, " ")

	//	}

	//	fmt.Println()

	//Go语言提供的内置函数 make() 可以用于灵活地创建数组切片
	//mySlice1 := make([]int, 5) //创建一个初始元素个数为5的数组切片,元素初始值为0

	//mySlice2 := make([]int, 5, 10) //创建一个初始元素个数为5的数组切片,元素初始值为0,并预留10个元素的存储空间

	//mySlice3 := []int{1, 2, 3, 4, 5} //直接创建并初始化包含5个元素的数组切片

	//	for i := 0; i < len(mySlice3); i++ {
	//		fmt.Println(i)
	//	}

	//	for k, v := range mySlice3 {
	//		fmt.Print(k, "----", v) //字符串的连接
	//		fmt.Println()
	//	}

	//cap() 函数返回的是数组切片分配的空间大小,而 len() 函数返回的是数组切片中当前所存储的元素个数

	//fmt.Println(mySlice2)

	myslice := make([]int, 5, 10)
	fmt.Println(cap(myslice))
	fmt.Println(len(myslice))

}

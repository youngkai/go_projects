package main

import "fmt"
import "errors"

type student struct {
	name string
	id   int
}

func main() {
	//我们可以使用Go语言内置的函数 make() 来创建一个新 map
	//	person := make(map[string]student)
	//	person["1"] = student{"young", 12}
	//	person["2"] = student{"lee", 18}
	//	//Go语言提供了一个内置函数 delete() ,用于删除容器内的元素
	//	delete(person, "2")
	//	for _, v := range person {
	//		fmt.Println(v.name)
	//	}
	//fmt.Println(test(11))
	//	Num := 3
	//	switch {
	//	case 0 <= Num && Num <= 3:
	//		fmt.Printf("0-3")
	//	case 4 <= Num && Num <= 6:
	//		fmt.Printf("4-6")
	//	case 7 <= Num && Num <= 9:
	//		fmt.Printf("7-9")
	//	}
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func test(x int) int {
	if x == 10 {
		return 1
	} else {
		return x
	}
}

func add(num1 int, num2 int) (ret int, err error) {
	ret := num1 + num2
	if ret < 0 {
		err = errors.New("不可能小于0")
		return
	} else {
		return ret, nil
	}
}

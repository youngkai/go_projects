package main

import "fmt"
import "os"
import "io"

//import "errors"

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
	//	i := 0
	//HERE:
	//	fmt.Println(i)
	//	i++
	//	if i < 10 {
	//		goto HERE
	//	}

	//test1(1, 2, 3, 4)

	//test2(1, 2, "aaa")

	x := func(a, b int) int {
		return a + b
	}

	fmt.Println(x(1, 2))

	CopyFile("/home/youngk/go/bb.txt", "/home/youngk/workspace/aa.txt")

}

func test(x int) int {
	if x == 10 {
		return 1
	} else {
		return x
	}
}

//func add(num1 int, num2 int) (ret int, err error) {
//	ret := num1 + num2
//	if ret < 0 {
//		err = errors.New("不可能小于0")
//		return
//	} else {
//		return ret, nil
//	}
//}

//不定参数
func test1(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func test2(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println(v, "is a int")
		case string:
			fmt.Println(v, "is a string")
		default:
			fmt.Println(v, "unknow")
		}
	}
}

//defer关键字,一个函数中可以存在多个defer
func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	srcFile.Chmod()
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

//panic内置函数   Go语言引入了两个内置函数 panic() 和 recover() 以报告和处理运行时错误和程序中的错误场景
//当在一个函数执行过程中调用 panic() 函数时,正常的函数执行流程将立即终止,但函数中
//之前使用 defer 关键字延迟执行的语句将正常展开执行,之后该函数将返回到调用函数,并导致
//逐层向上执行 panic 流程,直至所属的goroutine中所有正在执行的函数被终止。错误信息将被报
//告,包括在调用 panic() 函数时传入的参数,这个过程称为错误处理流程。
//recover() 函数用于终止错误处理流程。一般情况下, recover() 应该在一个使用 defer
//关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复
//过程(使用 recover 关键字),会导致该goroutine所属的进程打印异常信息后直接退出。
//defer func() {
//if r := recover(); r != nil {
//log.Printf("Runtime error caught: %v", r)
//}
//}()

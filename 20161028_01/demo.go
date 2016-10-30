package main

import "fmt"

func main() {
	//exchange
	var i = 10
	var j = 100
	fmt.Println(i)
	fmt.Println(j)
	i, j = j, i
	fmt.Println(i)
	fmt.Println(j)
}

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func myfunc(a, b string) (c, d string) {
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(myfunc("helo", "world"))
}

package main

import "fmt"

// 浮点数

func main() {
	// math.MaxFloat32 // float32的最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
}

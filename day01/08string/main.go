package main

import (
	"fmt"
	"strings"
)

func main() {
	// \\转义为\
	path := "D:\\go\\src\\github.com"
	fmt.Println(path)

	// 多行的字符串
	s2 := `
		大家
		好
	`
	fmt.Println(s2)

	s3 := `D:\Go\src\github.com\Andrewcyh\day01`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	// 分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ss, "理想"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}

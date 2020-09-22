package main

import "fmt"

// byte和rune类型
// Go语言中为了处理非ASCII码类型的字符 定义了新的rune类型

func main() {
	s := "Hello玉泉"
	// len()求的是byte字节的数量
	n := len(s) // 求字符串s的长度，把长度保存到变量n中
	fmt.Println(n)

	// for i := 0; i < n; i++ { // 识别不了UTF-8
	// 	fmt.Printf("%c", s[i]) // %c：字符
	// }

	// for _, c := range s { // 从字符串中拿出具体的字符
	// 	fmt.Printf("%c", c) // %c：字符
	// }

	// 字符串修改
	s2 := "白萝卜"
	
}

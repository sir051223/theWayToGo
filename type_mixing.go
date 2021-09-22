package main

func main() {
	var a int
	var b int32
	a = 15
	print(a)
	// b = a + a // 编译错误 cannot use a + a (type int) as type int32 in assignment
	b = b + 5
	print(b)
}

package bitoperation

//位运算交换
func Swap(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func Swap1(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 按位取反
//取反逻辑
//转二进制
//计算补码(正数补码等于本身,负数补码在原码基础上符号位不变，按位取反，末位加1)
//按位取反
//转位原码(正数就为本身，负数符号位不变按位取反，末位加1)
//https://www.cnblogs.com/shy1766IT/p/6184874.html
func nagation(a int) int {
	return ^a + 1
}

//与
func and(a, b int) int {
	return a & b
}

//或
func or(a, b int) int{
	return a | b
}

//异或
func xor(a, b int) int {
	return a ^ b
}



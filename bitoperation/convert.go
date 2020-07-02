package bitoperation

import (
	"fmt"
	"strconv"
)

// 十进制转2进制
func TenTo2(i int64) string{
	return strconv.FormatInt(i, 2)
}

func TenTo2_1(i int64) string {
	return fmt.Sprintf("%b", i)
}

// 十进制转8进制
func TenTo8(i int64) string{
	return strconv.FormatInt(i, 8)
}

func TenTo8_1(i int64) string{
	return fmt.Sprintf("%o", i)
}

// 十进制转16进制
func TenTo16(i int64) string{
	return strconv.FormatInt(i, 16)
}

func TenTo16_1(i int64) string {
	return fmt.Sprintf("%x", i)
}

//任意进制转10进制
func XTo10(x string, base int) (int64, error) {
	return strconv.ParseInt(x, base, 64)
}
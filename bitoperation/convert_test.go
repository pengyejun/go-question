package bitoperation

import (
	"fmt"
	"testing"
)

func TestTenTo2(t *testing.T) {
	var i int64 = 20
	x := TenTo2(i)
	x1 := TenTo2_1(i)
	if x != x1 || x != "10100" {
		fmt.Printf("TenTo2 value:%s\n", x)
		fmt.Printf("TenTo2_1 value:%s\n", x1)
		t.Fatal("TenTo2Error")
	}
}

func TestTenTo8(t *testing.T) {
	var i int64 = 20
	x := TenTo8(i)
	x1 := TenTo8_1(i)
	if x != x1 || x != "24" {
		fmt.Printf("TenTo8 value:%s\n", x)
		fmt.Printf("TenTo8_1 value:%s\n", x1)
		t.Fatal("TenTo8Error")
	}
}

func TestTenTo16(t *testing.T) {
	var i int64 = 20
	x := TenTo16(i)
	x1 := TenTo16_1(i)
	if x != x1 || x != "14" {
		fmt.Printf("TenTo16 value:%s\n", x)
		fmt.Printf("TenTo16_1 value:%s\n", x1)
		t.Fatal("TenTo16Error")
	}
}

func TestXTo10(t *testing.T) {
	x, err := XTo10("20", 16)
	if err != nil || x != 32 {
		fmt.Printf("XTo10 value:%d\n", x)
		t.Fatal("TestXTo10Error")
	}
}


func BenchmarkTenTo2(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		TenTo2(int64(i))
	}
}

func BenchmarkTenTo2_1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		TenTo2_1(int64(i))
	}
}

func BenchmarkTenTo8(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		TenTo8(int64(i))
	}
}

func BenchmarkTenTo8_1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		TenTo8_1(int64(i))
	}
}


func BenchmarkXTo10(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		XTo10(string(i), 16)
	}
}

func TestSwap(t *testing.T) {
	b, a := Swap(10, 20)
	if b != 20 || a != 10 {
		t.Fatal("Swap Error")
	}
}

func BenchmarkSwap(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		Swap(10, 16)
	}
}

func BenchmarkSwap1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		Swap1(10, 16)
	}
}
package main

import (
	"testing"
)

// go语言：表格驱动测试
/*
	TestAdd
	1. 文件名必须为*_test结尾
	2. 基础测试函数必须Test*开头（基础测试) testing.T
	3. 性能测试函数必须BenchMark*开头(性能测试) testing.B
*/
func TestAdd(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{3, 5, 8},
		{7, 8, 15},
		{10, 11, 21},
	}
	for _, tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("测试add(%d, %d)方法: 结果值：%d; 期望值：%d", tt.a, tt.b, tt.c, actual)
		}
	}
}
func BenchmarkAdd(b *testing.B) {
	e, f, g := 7, 8, 15
	for v := 0; v < b.N; v++ {
		if actual := add(e, f); actual != g {
			b.Errorf("测试add(%d, %d)方法: 结果值：%d; 期望值：%d", e, f, g, actual)
		}
	}
}
func add(a, b int) int {
	return a + b
}

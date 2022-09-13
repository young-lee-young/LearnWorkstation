package main

import "testing"

/**
	测试

	文件以_test.go结尾
	方法以Test开始
	参数固定 t *testing.T

	命令行运行测试
	go test _test.go
	go test -coverprofile=test.out
	go tool cover -html=test.out

	代码覆盖率
	go test -bench .
	go test -bench . -cpuprofile cpu.out
	go tool pprof cpu.out
 */
func TestTriangle(t *testing.T)  {
	tests := []struct {a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c{
			t.Errorf("calcTriangle(%d, %d) got %d expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	num1, num2, num3 := 12, 35, 37

	for i := 0; i < b.N; i ++ {
		actual := calcTriangle(num1, num2)
		if actual != num3 {
			b.Errorf("calcTriangle(%d, %d) got %d expected %d", num1, num2, actual, num3)
		}
	}
}

package main

import "testing"


 //测试加法
func TestAdd(t *testing.T)  {
	tests := []struct{a, b, c int} {
		{1, 2, 3}, // 1+2 = 3
		{3, 5, 6},
		{7, 8, 9},
	}

	for _, tt := range tests{
		if res := add(tt.a, tt.b); res != tt.c {
			t.Errorf("error； got %d; exption", res)
		}
	}
}

//测试减法
func TestReduce(t *testing.T)  {
	tests := []struct{a, b, c int} {
		{10, 4, 6},	// 10 - 4 ==6
		{9, 4, 5},
		{9, 5, 1},
	}
	for _, tt :=  range tests{
		res := reduce(tt.a, tt.b)
		if res != tt.c {
			t.Errorf("div(%d, %d) error: got %d; exption %d", tt.a, tt.b, tt.c, res)
		}
	}
}

/**
性能测试
 */
func BenchmarkAdd(bb *testing.B)  {
	a, b, c := 1, 2, 3
	bb.ResetTimer() // 忽略准备数据的时间
	for i :=0; i<bb.N; i++ {
		if res := add(a, b); res != c {
			bb.Errorf("error； got %d; exption", res)
		}
	}
}
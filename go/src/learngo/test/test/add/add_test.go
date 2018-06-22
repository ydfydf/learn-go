package main

import "testing"

func TestAddFunc(t *testing.T) { //测试函数必须以Test为前缀才能运行
	tests := []struct{
		a, b, c int
	}{
		{1,2,3},
		{2,3,5},
		{12,234,0},
		{1111,23234,0},
	}

	for _, tt := range tests {
		if actual := Add(tt.a,tt.b); actual != tt.c {
			t.Errorf("add error,need:%d,given:%d\n",actual,tt.c)
		}
	}
}

func BenchmarkAdd(b *testing.B){ //性能测试函数必须以Benchmark为前缀才能运行
	x , y , z := 1111 , 23234 ,24345 //挑选一组最复杂或者最难的用于性能测试
	for i := 0 ; i < b.N ;i ++ { //做性能测试需要做很多次，具体次数是由算法决定，即b.N来决定
		actual := Add(x,y)
		if actual != z {
			b.Errorf("add error,need:%d,given:%d\n",actual,z)
		}
		//输出：2000000000	         0.26 ns/op，表示运行2000000000次，每次0.26纳秒
	}


}
package ltest

import (
	"fmt"
	"testing"
)

func TestT2(t *testing.T)  {
	var (
		a = 1
		b = 2
		expected = 3
	)
	ret := T2(a, b)
	if ret != expected {
		t.Errorf("result error")
	}
	type ins struct{
		a int
		b int
	}

	var tests = []struct{
		ins
		ret int
	} {
		{ins{1, 2},3},
		{ins{4, 2},6},
		{ins{1, 2},3},
		{ins{3, 2},5},
		{ins{1, 2},3},
		{ins{1, 10},11},
		{ins{10, 2},12},
		{ins{7, 2},9},
	}
	for _, tt := range tests {
		actual := T2(tt.a, tt.b)
		if actual != tt.ret {
			t.Errorf("error")
		}
	}
}

func BenchmarkT2(b *testing.B) {
	for i :=0; i< b.N; i++ {
		fmt.Sprintf("hello")
	}
}



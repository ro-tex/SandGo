package main

import (
	"testing"
)

// result helps us prevent the compiler from removing our functions
var result interface{}

// BenchmarkRands is a table test that covers both `fr` and `rr`
func BenchmarkRands(b *testing.B) {
	tests := []struct {
		name string
		fn   func() uint64
	}{
		{"FR", fr},
		{"RR", rr},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			var n uint64
			for i := 0; i < b.N; i++ {
				n = tt.fn()
			}
			result = n
		})
	}
}

func BenchmarkFR(b *testing.B) {
	var n uint64
	for i := 0; i < b.N; i++ {
		n = fr()
	}
	result = n
}

func BenchmarkRR(b *testing.B) {
	var n uint64
	for i := 0; i < b.N; i++ {
		n = rr()
	}
	result = n
}

func BenchmarkS_Default(b *testing.B) {
	ch := make(chan struct{})
	for i := 0; i < b.N; i++ {
		s(ch)
	}
}

func BenchmarkS_ChanMsg(b *testing.B) {
	ch := make(chan struct{})
	go func() {
		for {
			ch <- struct{}{}
		}
	}()
	for i := 0; i < b.N; i++ {
		s(ch)
	}
}

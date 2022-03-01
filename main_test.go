package main_test

import (
	"testing"

	a "github.com/trrrrrys/ambiguousud"
)

func TestAmbiguousConvert(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect interface{}
	}{
		{
			"yyyymmdd hh:mm:ss to unixtimestamp",
			"20200101 00:00:00",
			1577804400,
		},
		{
			"yyyymmdd hh:mm:ss to unixtimestamp",
			"20200101 00:00:00",
			1577804400,
		},
		{
			"yyyymmdd hh:mm:ss to unixtimestamp",
			"20200101 00:00:00",
			1577804400,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// get value
			got, err := a.AmbiguousConvert(tt.in)
			if err != nil {
				t.Fatalf("error: %+v", err)
			}
			if got != tt.expect {
				t.Errorf("expect: %v, got: %v", tt.expect, got)
			}
		})
	}
}

// Benchmark

func BenchmarkAmbiguousConvert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.AmbiguousConvert("20200101 00:00:00")
	}
}

func BenchmarkAmbiguousConvert2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.AmbiguousConvert("2020-01-01 00:00:00")
	}
}

func BenchmarkAmbiguousConvert3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.AmbiguousConvert("2020/01/01 00:00:00")
	}
}

func BenchmarkAmbiguousConvert4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.AmbiguousConvert("20200101")
	}
}

func BenchmarkAmbiguousConvert5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.AmbiguousConvert("2020-01-01")
	}
}

func BenchmarkAmbiguousConvert6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.AmbiguousConvert("2020/01/01")
	}
}

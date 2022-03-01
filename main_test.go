package main_test

import (
	"reflect"
	"testing"

	a "github.com/trrrrrys/ambiguousud"
)

func TestAmbiguousConvert(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect any
	}{
		{
			"yyyymmdd hh:mm:ss to unixtimestamp",
			"20200101 00:00:00",
			int64(1577804400),
		},
		{
			"yyyy/mm/dd hh:mm:ss to unixtimestamp",
			"2020/01/01 00:00:00",
			int64(1577804400),
		},
		{
			"yyyy-mm-dd hh:mm:ss to unixtimestamp",
			"2020-01-01 00:00:00",
			int64(1577804400),
		},
		{
			"ut to yyyy-mm-dd hh:mm:ss",
			"1577804400",
			"2020-01-01 00:00:00",
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
				t.Log("expect: ", reflect.TypeOf(tt.expect))
				t.Log("got: ", reflect.TypeOf(got))
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

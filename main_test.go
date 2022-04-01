package main

import (
	"testing"
)

func BenchmarkGenerateBwmarrinSnowflake(b *testing.B) {

	bw := NewBW()

	for i := 0; i < b.N; i++ {
		bw.Generate()
	}
}

func BenchmarkGenerateFishTennisSnowflake(b *testing.B) {

	ft := NewFt()

	for i := 0; i < b.N; i++ {
		ft.NextId()
	}
}

func BenchmarkGenerateBwmarrinSnowflakeParallel(b *testing.B) {

	bw := NewBW()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bw.Generate()
		}
	})

}

func BenchmarkGenerateFishTennisSnowflakeParallel(b *testing.B) {

	ft := NewFt()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ft.NextId()
		}
	})

}

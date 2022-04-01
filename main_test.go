package main

import (
	"testing"
)

func BenchmarkBwmarrinSnowflake(b *testing.B) {

	bw := NewBW()

	for i := 0; i < b.N; i++ {
		bw.Generate()
	}
}

func BenchmarkFishTennisSnowflake(b *testing.B) {

	ft := NewFt()

	for i := 0; i < b.N; i++ {
		ft.NextId()
	}
}

func BenchmarkBwmarrinSnowflakeParallel(b *testing.B) {

	bw := NewBW()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bw.Generate()
		}
	})

}

func BenchmarkFishTennisSnowflakeParallel(b *testing.B) {

	ft := NewFt()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ft.NextId()
		}
	})

}

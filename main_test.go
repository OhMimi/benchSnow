package main

import "testing"

func BenchmarkGenerateBwmarrinSnowflake(b *testing.B) {

	bw := NewBW()

	for i := 0; i < b.N; i++ {
		bw.Generate().Int64()
	}
}

func BenchmarkGenerateFishTennisSnowflake(b *testing.B) {

	ft := NewFt()

	for i := 0; i < b.N; i++ {
		ft.NextId()
	}
}

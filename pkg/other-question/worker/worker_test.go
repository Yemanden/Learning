package worker

import (
	"context"
	"testing"
)

const (
	benchmarkWorkTest1Name  = "One worker"
	benchmarkWorkTest1Input = 1

	benchmarkWorkTest2Name  = "Two workers"
	benchmarkWorkTest2Input = 2

	benchmarkWorkTest3Name  = "Three workers"
	benchmarkWorkTest3Input = 3

	benchmarkWorkTest4Name  = "Four workers"
	benchmarkWorkTest4Input = 4

	benchmarkWorkTest5Name  = "Five workers"
	benchmarkWorkTest5Input = 5

	benchmarkWorkTest6Name  = "Six workers"
	benchmarkWorkTest6Input = 6

	benchmarkWorkTest7Name  = "Seven workers"
	benchmarkWorkTest7Input = 7

	benchmarkWorkTest8Name  = "Eight workers"
	benchmarkWorkTest8Input = 8
)

// BenchmarkWork ...
func BenchmarkWork(b *testing.B) {
	b.Run(benchmarkWorkTest1Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest1Input)
	})

	b.Run(benchmarkWorkTest2Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest2Input)
	})

	b.Run(benchmarkWorkTest3Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest3Input)
	})

	b.Run(benchmarkWorkTest4Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest4Input)
	})

	b.Run(benchmarkWorkTest5Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest5Input)
	})

	b.Run(benchmarkWorkTest6Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest6Input)
	})

	b.Run(benchmarkWorkTest7Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest7Input)
	})

	b.Run(benchmarkWorkTest8Name, func(b *testing.B) {
		foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		foreman.Work(ctx, benchmarkWorkTest8Input)
	})
}

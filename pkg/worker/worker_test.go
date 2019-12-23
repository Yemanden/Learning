package worker

import (
	"context"
	"testing"
)

const (
	BenchmarkWorkTest1Name  = "One workers"
	BenchmarkWorkTest1Input = 1
	BenchmarkWorkTest2Name  = "Two workers"
	BenchmarkWorkTest2Input = 2
	BenchmarkWorkTest3Name  = "Three workers"
	BenchmarkWorkTest3Input = 3
	BenchmarkWorkTest4Name  = "Four workers"
	BenchmarkWorkTest4Input = 4
	BenchmarkWorkTest5Name  = "Five workers"
	BenchmarkWorkTest5Input = 5
	BenchmarkWorkTest6Name  = "Six workers"
	BenchmarkWorkTest6Input = 6
	BenchmarkWorkTest7Name  = "Seven workers"
	BenchmarkWorkTest7Input = 7
	BenchmarkWorkTest8Name  = "Eight workers"
	BenchmarkWorkTest8Input = 8
)

// BenchmarkWork проверяет среднее значение скорости выполнения функции Work
// при значениях аргументов от 1 до 10
func BenchmarkWork(b *testing.B) {
	b.Run(BenchmarkWorkTest1Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest1Input)
	})
	b.Run(BenchmarkWorkTest2Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest2Input)
	})
	b.Run(BenchmarkWorkTest3Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest3Input)
	})
	b.Run(BenchmarkWorkTest4Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest4Input)
	})
	b.Run(BenchmarkWorkTest5Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest5Input)
	})
	b.Run(BenchmarkWorkTest6Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest6Input)
	})
	b.Run(BenchmarkWorkTest7Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest7Input)
	})
	b.Run(BenchmarkWorkTest8Name, func(b *testing.B) {
		Foreman := NewForeman()
		ctx := context.Background()
		b.ResetTimer()
		Foreman.Work(ctx, BenchmarkWorkTest8Input)
	})
}

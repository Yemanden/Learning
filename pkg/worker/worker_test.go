package worker

import "testing"

// BenchmarkWork проверяет среднее значение скорости выполнения функции Work
// при значениях аргументов от 1 до 10
func BenchmarkWork(b *testing.B) {
	for i := 1; i <= 10; i++ {
		Work(uint(i))
	}
}

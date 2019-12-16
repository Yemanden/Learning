package Worker

import "testing"

// BenchmarkWork проверяет среднее значение скорости выполнение функции Work
// при значениях аргументов от 1 до 10
func BenchmarkWork(b *testing.B) {
	for i := 1; i <= 10; i++ {
		Work(1)
	}
}

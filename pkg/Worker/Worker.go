package Worker

import (
	"math/rand"
	"time"
)

// Work принимает количество работников и возвращает отчёт о проделанной ими работе
// В данном контексте работа - возведение в квадрат случайного числа от -10 до 10
func Work(workerCount uint) []int {
	tasks := make(chan int, workerCount)
	results := make(chan int, workerCount)

	for i := uint(0); i < workerCount; i++ {
		go workerUsing(tasks, results)
	}

	time.Sleep(10000) // Формирование задач для работяг

	// В канал с задачами помещаем "сформированные" задачи
	for i := uint(0); i < workerCount; i++ {
		tasks <- rand.Intn(20) - 10 // Задача - случайное число от -10 до 10
	}

	// Результаты всех работ записываются в report и возвращаются функцией
	report := make([]int, workerCount)
	for i := uint(0); i < workerCount; i++ {
		report[i] = <-results
	}
	return report
}

// Функция выполнения работы (возведения числа в квадрат с задержкой)
// На вход принимает канал с задачами для чтения и канал с результатами для записи
func workerUsing(tasks <-chan int, results chan<- int) {
	task := <-tasks
	result := task * task
	workDuration := time.Duration(1000 + rand.Intn(3000-1000))
	time.Sleep(time.Millisecond * workDuration)
	results <- result
}

package worker

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Worker is interface, contains Work method.
type Worker interface {
	Work(context.Context, uint)
}

type foreman struct {
}

// NewForeman is constructor, returns "foreman" object that implements the interface Worker
func NewForeman() Worker {
	return new(foreman)
}

// Функция выполнения работы (возведения числа в квадрат с задержкой)
// На вход принимает канал с задачами для чтения и канал с результатами для записи
func workerUsing(ctx context.Context, report chan<- int) {
	var task int
	task = ctx.Value("task").(int)
	done := make(chan struct{})
	go func() {
		workDuration := time.Duration(1000 + rand.Intn(3000-1000))
		time.Sleep(time.Millisecond * workDuration)
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return
	case <-done:
		result := task * task
		report <- result
	}
}

// Work generates "workerCount" workers to complete the task. Who is faster
func (f *foreman) Work(ctx context.Context, workerCount uint) {
	var task = 33
	cnlCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	newCtx := context.WithValue(cnlCtx, "task", task)
	report := make(chan int)

	for i := uint(0); i < workerCount; i++ {
		go workerUsing(newCtx, report)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Cancel")
	case output := <-report:
		fmt.Println(output)
	}
}

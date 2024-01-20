package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime/pprof"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	ID      int
	Content string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		processTime := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(processTime) // Simulating task processing time
		fmt.Printf("Worker %d completed task %d in %v\n", id, task.ID, processTime)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5
	numTasks := 10

	tasks := make(chan Task, numTasks)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		labels := pprof.Labels("worker", strconv.Itoa(i))
		pprof.Do(context.Background(), labels, func(_ context.Context) {
			go worker(i, tasks, &wg)
		})
	}

	// Add tasks to the queue
	for j := 1; j <= numTasks; j++ {
		wg.Add(1)
		tasks <- Task{ID: j, Content: fmt.Sprintf("Task content %d", j)}
	}

	wg.Wait()
	close(tasks)
	fmt.Println("All tasks completed")
}

package main

import "fmt"

type taskFunc func(Task)

type Task struct {
	name  string
	taskf taskFunc
	ch    chan int
	next  *Task
}

func (task Task) doTask() {
	for range task.ch {
		fmt.Print(task.name)
		if task.next == nil {
			sigChan <- 0
		} else {
			task.next.ch <- 0
		}
	}
}

var sigChan = make(chan int)

func main() {
	names := []string{"fish", "cat", "dog"}
	tasks := make([]*Task, len(names))

	for k, name := range names {
		tasks[k] = &Task{
			name: name,
			ch:   make(chan int),
			next: nil,
		}
	}

	for i, t := range tasks {
		if i != len(tasks)-1 {
			t.next = tasks[i+1]
		}
	}

	for _, task := range tasks {
		go func(t *Task) {
			t.doTask()
		}(task)
	}

	tasks[0].ch <- 0
	var counter int
	for range sigChan {
		counter++
		fmt.Println(counter)
		tasks[0].ch <- 0
		if counter == 100 {
			close(sigChan)
		}
	}
}

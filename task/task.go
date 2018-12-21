package task

import (
	"log"
	"sync"
)

// Task is basic task
type Task struct {
	wg *sync.WaitGroup
}

func (t *Task) initTask() {
	t.wg = &sync.WaitGroup{}
}

// Item is a task item holder that support context
type Item struct {
	item    interface{}
	context interface{}
}

type fn func()

func (t *Task) addItem(item Item, callback fn) {
	t.wg.Add(1)
	callback()
}

func (t *Task) onItemAsync(callback fn) {
	go func() {
		t.onItemSync(callback)
	}()
}

func (t *Task) onItemSync(callback fn) {
	defer t.wg.Done()
	callback()
}

func (t *Task) waitAllDone(callback fn) {
	t.wg.Add(1)
	callback()
	t.wg.Wait()
}

func (t *Task) Wait() {
	t.wg.Wait()
}

func (t *Task) Add(w Work) {
	log.Println("task added")
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		w.Run()
		log.Println("task finished")
	}()
}

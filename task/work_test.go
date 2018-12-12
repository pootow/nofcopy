package task

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestWaitWorkDone(t *testing.T) {
	st := NewSimpleTask()
	func(ws Scheduler) {
		ws.Wait()
	}(st)
}

type DummyWork struct {
	takes time.Duration
}

func (d *DummyWork) Run() {
	fmt.Println("work start....")
	time.Sleep(d.takes)
	fmt.Println("work done!!")
}

func TestAddWorkAndWaitAllDone(t *testing.T) {
	st := NewSimpleTask()
	func(xws WorkScheduler) {
		begin := time.Now()
		xws.Add(&DummyWork{takes: time.Millisecond * 10})
		xws.Wait()
		end := time.Now()
		duration := end.Sub(begin)
		if duration.Seconds() < 0.01 || duration.Seconds() > 0.02 {
			t.Error("work wait may not happend. duration(sec): ", duration.Seconds())
		}
	}(st)
}

func TestAddMultiWorkShouldWaitOnLongest(t *testing.T) {
	st := NewSimpleTask()
	func(xws WorkScheduler) {
		begin := time.Now()
		xws.Add(&DummyWork{takes: time.Millisecond * 10})
		xws.Add(&DummyWork{takes: time.Millisecond * 100})
		xws.Add(&DummyWork{takes: time.Millisecond * 50})
		xws.Wait()
		end := time.Now()
		duration := end.Sub(begin)
		if duration.Seconds() < 0.10 || duration.Seconds() > 0.11 {
			t.Error("work wait may not happend. duration(sec): ", duration.Seconds())
		}
	}(st)
}

type Echo string

func (*Echo) Run() {
	fmt.Println("vvvvvvvvvvv")
	log.Println("xxxxxxxxxxx")
}

func TestManyKindsOfWorkShouldWork(t *testing.T) {
	st := NewSimpleTask()
	func(xws WorkScheduler) {
		xws.Add(&DummyWork{takes: time.Millisecond * 10})
		echo := Echo("{}")
		xws.Add(&echo)
		xws.Wait()
	}(st)
}

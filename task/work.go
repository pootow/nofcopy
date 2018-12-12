package task

type Scheduler interface {
	Wait()
}

type WorkScheduler interface {
	Scheduler
	Add(Work)
}

type Work interface {
	Run()
}

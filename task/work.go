package task

type WorkScheduler interface {
	Wait()
}

type XXXWorkScheduler interface {
	WorkScheduler
	Add(Work)
}

type Work interface {
	Run()
}

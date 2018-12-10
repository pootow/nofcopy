package task

type SimpleTask struct {
	Task
}

func NewSimpleTask() *SimpleTask {
	t := new(SimpleTask)
	t.initTask()
	return t
}

package task

import (
	"testing"
)

func TestThatTaskCanWaitOnTaskItem(t *testing.T) {
	task := NewSimpleTask()
	task.onItemAsync(func() {})
	task.waitAllDone(func() {})
}

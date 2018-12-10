package task

import (
	"os"
)

// StatTask is a Task that handling a path.
type StatTask struct {
	DirWalkTask
	Total int64
}

// NewStatTask creates new StatTask
func NewStatTask() *StatTask {
	t := new(StatTask)
	t.initDirWalkTask(&statWalker{StatTask: t})
	return t
}

type statWalker struct {
	DefaultWalker
	*StatTask
}

// Stat total size of a path
func (t *StatTask) Stat(root string) {
	t.Walk(Item{item: root})
}

func (w *statWalker) onFile(file Item) {
	w.Total += file.context.(os.FileInfo).Size()
}

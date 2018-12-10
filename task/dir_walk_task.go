package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// DirWalkTask is a task that walk a dir tree.
type DirWalkTask struct {
	Task
	paths  chan Item
	dirs   chan Item
	files  chan Item
	walker Walker
}

func (t *DirWalkTask) initDirWalkTask(walker Walker) {
	t.Task.initTask()
	t.paths = make(chan Item)
	t.dirs = make(chan Item)
	t.files = make(chan Item)
	t.walker = walker
}

// Walker that do ths. when walk a path
type Walker interface {
	onPath(Item)
	onDir(Item)
	onFile(Item)
}

// DefaultWalker is noop walker
type DefaultWalker struct{}

func (t *DefaultWalker) onPath(path Item) {
}

func (t *DefaultWalker) onDir(dir Item) {
}

func (t *DefaultWalker) onFile(file Item) {
}

// Walk all of the files of a path
func (t *DirWalkTask) Walk(root Item) {
	fmt.Println("checksum: ", root)

	go func() {
		for {
			select {
			case path := <-t.paths:
				t.onItemAsync(func() {
					t.forPath(path)
				})
			case dir := <-t.dirs:
				t.onItemAsync(func() {
					t.forDir(dir)
				})
			case file := <-t.files:
				t.onItemAsync(func() {
					t.walker.onFile(file)
				})
			}
		}
	}()
	t.waitAllDone(func() {
		t.paths <- root
	})
}

func (t *DirWalkTask) forPath(path Item) {
	info, err := os.Lstat(path.item.(string))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
	fileMode := info.Mode()
	if fileMode&os.ModeSymlink != 0 {
		fmt.Fprintf(os.Stderr, "skip symlink: %s\n", path)
		return
	}
	path = Item{item: path.item, context: info}
	t.walker.onPath(path)
	if info.IsDir() {
		t.addItem(path, func() {
			t.dirs <- path
		})
	} else {
		t.addItem(path, func() {
			t.files <- path
		})
	}
}

func (t *DirWalkTask) forDir(dir Item) {
	t.walker.onDir(dir)
	dirPath := dir.item.(string)
	subs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
	for _, sub := range subs {
		subItem := Item{item: path.Join(dirPath, sub.Name())}
		t.addItem(subItem, func() {
			t.paths <- subItem
		})
	}
}

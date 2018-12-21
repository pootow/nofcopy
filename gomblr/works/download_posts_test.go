package works

import (
	"github.com/pootow/nofcopy/task"
	"testing"
)

func TestNewDownloadPosts(t *testing.T) {
	st := task.NewSimpleTask()
	posts := NewDownloadPosts(1, 1, st)
	st.Add(posts)
	st.Wait()
}

func TestGetSomePostFile(t *testing.T) {
	paths, err := getSomePostFile(10)
	if err != nil {
		t.Log(err)
	}
	for _, path := range paths {
		t.Log(path)
	}
}

func TestLoadPost(t *testing.T) {
	postInterface, err := loadPost("/Users/zhangyu/gomblr/posts/174531181312.json")
	if err != nil {
		t.Log(err)
	}
	t.Log(postInterface)
}

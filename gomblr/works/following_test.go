package works

import (
	. "github.com/pootow/nofcopy/gomblr/client"
	"github.com/pootow/nofcopy/task"
	"testing"
)

func TestFollowingWork(t *testing.T) {
	st := task.NewSimpleTask()
	client := NewGomblrClient()

	st.Add(&followingWork{client})

	st.Wait()

}

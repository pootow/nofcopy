package works

import (
	. "github.com/pootow/nofcopy/gomblr/client"
	"github.com/pootow/nofcopy/task"
	"testing"
)

func TestFavWork(t *testing.T) {
	st := task.NewSimpleTask()
	client := NewGomblrClient()

	st.Add(&likes{
		client,
	})

	st.Wait()

}

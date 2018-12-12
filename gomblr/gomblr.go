package gomblr

import (
	. "github.com/pootow/nofcopy/gomblr/client"
	"github.com/pootow/nofcopy/gomblr/works"
	"github.com/pootow/nofcopy/task"
	. "github.com/tumblr/tumblr.go"
)

func GetBlogPosts(blogs []string) {
	st := task.NewSimpleTask()
	client := NewGomblrClient()

	for _, blog := range blogs {
		st.Add(&works.BlogPosts{
			client,
			blog,
		})
	}
	st.Wait()
}

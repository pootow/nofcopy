package gomblr

import (
	. "github.com/pootow/nofcopy/gomblr/client"
	"github.com/pootow/nofcopy/gomblr/works"
	"github.com/pootow/nofcopy/task"
	"log"
	"strconv"
)

func GetBlogPosts(blogs []string) {
	st := task.NewSimpleTask()
	client := NewGomblrClient()

	for _, blog := range blogs {
		st.Add(works.NewBlogPosts(client, blog))
	}
	st.Wait()
}

func DownloadPosts(concurrency string) {
	st := task.NewSimpleTask()

	con64, err := strconv.ParseInt(concurrency, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	con := int(con64)
	st.Add(works.NewDownloadPosts(con, st))

	st.Wait()
}

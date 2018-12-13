package gomblr

import (
	"log"
	"strconv"

	. "github.com/pootow/nofcopy/gomblr/client"
	"github.com/pootow/nofcopy/gomblr/works"
	"github.com/pootow/nofcopy/task"
)

func GetBlogPosts(blogs []string) {
	st := task.NewSimpleTask()
	client := NewGomblrClient()

	for _, blog := range blogs {
		st.Add(works.NewBlogPosts(client, blog))
	}
	st.Wait()
}

func DownloadPosts(concurrency string, count string) {
	st := task.NewSimpleTask()

	con64, err := strconv.ParseInt(concurrency, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	con := int(con64)
	countInt, err := strconv.Atoi(count)

	for i := 0; i < con; i++ {
		st.Add(works.NewDownloadPosts(countInt, i, st))
	}

	st.Wait()
}

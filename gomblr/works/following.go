package works

import (
	"github.com/pootow/nofcopy/gomblr/client"
	"github.com/tumblr/tumblr.go"
	"strconv"
)

type followingWork struct {
	client *client.GomblrClient
}

func (f *followingWork) Run() {
	offset := uint(0)
	for {

		followingList, _ := tumblr.GetFollowing(f.client, offset, 20)

		if len(followingList.Blogs) == 0 {
			break
		}
		offset64, _ := strconv.ParseUint(followingList.Links.Next.QueryParams["offset"], 10, 0)
		println("offset: -------------------------->>", offset64)
		for _, blog := range followingList.Blogs {
			println(blog.Name)
		}
		offset = uint(offset64)
	}

}

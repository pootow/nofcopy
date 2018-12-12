package works

import (
	"github.com/pootow/nofcopy/gomblr/client"
	. "github.com/tumblr/tumblr.go"
	"net/url"
	"strconv"
)

type likes struct {
	client *client.GomblrClient
}

func (l *likes) Run() {
	before := ""
	for {
		println("==================================>>>>>>>>>>>>>my likes before: ", before)
		likes := l.getMyLikes(before, 20)
		if len(likes.Posts) == 0 {
			break
		}
		posts, _ := likes.Full()
		client.Store(posts)
		before = likes.Links.Next.QueryParams["before"]
	}

}

func (l *likes) getMyLikes(before string, limit int) *Likes {
	params := url.Values{}
	if before != "" {
		params.Set("before", before)
	}
	params.Set("limit", strconv.Itoa(limit))
	likes, _ := GetLikes(l.client, params)
	return likes
}

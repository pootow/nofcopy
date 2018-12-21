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

func NewLikes(client *client.GomblrClient) *likes {
	return &likes{client: client}
}

func (l *likes) Run() {
	before := ""
	for {
		println("==================================>>>>>>>>>>>>>my likes before: ", before)
		likes, err := l.getMyLikes(before, 20)
		if err != nil {
			println("error when get likes: ", err)
			continue
		}
		if len(likes.Posts) == 0 {
			break
		}
		posts, err := likes.Full()
		if err != nil {
			println("error when get full post props: ", err)
			continue
		}
		client.Store(posts)
		before = likes.Links.Next.QueryParams["before"]
	}

}

func (l *likes) getMyLikes(before string, limit int) (*Likes, error) {
	params := url.Values{}
	if before != "" {
		params.Set("before", before)
	}
	params.Set("limit", strconv.Itoa(limit))
	likes, err := GetLikes(l.client, params)
	return likes, err
}

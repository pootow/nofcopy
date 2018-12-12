package works

import (
	"github.com/pootow/nofcopy/gomblr/client"
	"github.com/tumblr/tumblr.go"
	"net/url"
	"strconv"
)

type BlogPosts struct {
	client   *client.GomblrClient
	blogName string
}

func (b *BlogPosts) Run() {
	offset := ""
	for {
		posts := b.getPosts(b.blogName, offset, 20)
		if len(posts.Posts) == 0 {
			break
		}
		postsFace, _ := posts.All()
		client.Store(postsFace)
		offset = posts.Links.Next.QueryParams["offset"]
	}

}

func (b *BlogPosts) getPosts(name string, offset string, limit int) *tumblr.Posts {
	params := url.Values{}
	if offset != "" {
		params.Set("offset", offset)
	}
	params.Set("limit", strconv.Itoa(limit))
	posts, _ := tumblr.GetPosts(b.client, name, params)
	return posts
}

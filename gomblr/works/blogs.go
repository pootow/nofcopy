package works

import (
	"github.com/pootow/nofcopy/gomblr/client"
)

type blogs struct {
	client    *client.GomblrClient
	blogNames []string
}

func NewBlogsWork(client *client.GomblrClient, blogNames []string) *blogs {
	return &blogs{client: client, blogNames: blogNames}
}

func (b *blogs) Run() {
	for _, blog := range b.blogNames {
		work := &BlogPosts{
			client:   b.client,
			blogName: blog,
		}
		work.Run()
	}
}

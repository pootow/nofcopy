package works

import (
	"github.com/pootow/nofcopy/gomblr/common"
	"log"
)

type DownloadPosts struct {
	index int
}

func NewDownloadPosts(index int) *DownloadPosts {
	return &DownloadPosts{index: index}
}

func (d *DownloadPosts) Run() {
	posts, err := common.LockPosts()
	if err != nil {
		log.Println(err)
	}
	for _, post := range posts {
		common.DownloadPost(post)
	}
}

package works

import (
	"github.com/pootow/nofcopy/gomblr/common"
	"github.com/pootow/nofcopy/gomblr/extractors"
	"log"
	"os"
	"path/filepath"
)

type downloadPost struct {
	post     extractors.PostResourceExtractor
	taskPath string
}

func (d *downloadPost) Run() {
	resourceUrls := d.post.GetResources()
	for _, resourceUrl := range resourceUrls {
		err := common.Download(resourceUrl)
		if err != nil {
			log.Println("error when downloading url: ", resourceUrl)
			return
		}
	}
	err := d.finish()
	if err != nil {
		log.Println("error when move task to finish: ", err)
	}
}

func (d *downloadPost) finish() error {
	finishDirPath, err := common.GetAndMakeInAppWorkingDir("finish")
	if err != nil {
		return err
	}
	_, fileName := filepath.Split(d.taskPath)
	finishPath := filepath.Join(finishDirPath, fileName)
	err = os.Rename(d.taskPath, finishPath)
	return err
}

func newDownloadPost(post extractors.PostResourceExtractor, taskPath string) *downloadPost {
	return &downloadPost{post: post, taskPath: taskPath}
}

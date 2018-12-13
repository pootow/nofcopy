package works

import (
	"encoding/json"
	"github.com/pootow/nofcopy/gomblr/common"
	"github.com/pootow/nofcopy/gomblr/extractors"
	"github.com/pootow/nofcopy/task"
	"github.com/tumblr/tumblr.go"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

type DownloadPosts struct {
	scheduler task.WorkScheduler
	Count     int
}

func NewDownloadPosts(count int, scheduler task.WorkScheduler) *DownloadPosts {
	//TODO should I move all locked post back to posts???
	return &DownloadPosts{scheduler: scheduler, Count: count}
}

func (d *DownloadPosts) Run() {
	for i := 0; i < d.Count; i++ {
		posts, err := d.lockPosts()
		if err != nil {
			log.Println("downloading error when lock post: ", err)
			// TODO wait for a random delay
			i--
			continue
		}
		for _, post := range posts {
			//d.scheduler.Add(post)
			post.Run()
		}
	}
}

func (d *DownloadPosts) lockPosts() ([]*downloadPost, error) {
	// TODO load locked posts first
	downloadPosts := make([]*downloadPost, d.Count)
	postPaths, err := getSomePostFile(d.Count)
	if err != nil {
		return nil, err
	}
	for i, postPath := range postPaths {
		newPath, err := d.lockPostPath(postPath)
		if err != nil {
			return nil, err // or print err and continue
		}
		postResourceExtractor, err := loadPost(newPath)
		if err != nil {
			return nil, err
		}
		downloadPosts[i] = newDownloadPost(postResourceExtractor, newPath)
	}
	return downloadPosts, nil
}

func loadPost(postPath string) (extractors.PostResourceExtractor, error) {
	postBytes, err := ioutil.ReadFile(postPath)
	if err != nil {
		return nil, err
	}

	miniPost := tumblr.MiniPost{}
	err = json.Unmarshal(postBytes, &miniPost)
	if err != nil {
		return nil, err
	}

	postInterface, err := tumblr.MakePostFromType(miniPost.Type)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(postBytes, &postInterface)
	if err != nil {
		return nil, err
	}

	return extractors.GetExtractor(postInterface), nil
}

func (d *DownloadPosts) lockPostPath(postPath string) (string, error) {
	lockerDir, err := common.GetAndMakeInAppWorkingDir("locker")
	if err != nil {
		return "", err
	}
	_, file := filepath.Split(postPath)
	newPath := path.Join(lockerDir, file)
	err = os.Rename(postPath, newPath)
	return newPath, err
}

func getSomePostFile(count int) ([]string, error) {
	var postFiles []string

	dirPath, err := common.GetAndMakeInAppWorkingDir("posts")
	if err != nil {
		return nil, err
	}

	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	infos, err := dir.Readdir(count)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		postFiles = append(postFiles, path.Join(dirPath, info.Name()))
	}
	return postFiles, nil
}

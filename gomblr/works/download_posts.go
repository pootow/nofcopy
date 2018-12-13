package works

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/pootow/nofcopy/gomblr/common"
	"github.com/pootow/nofcopy/gomblr/extractors"
	"github.com/pootow/nofcopy/task"
	tumblr "github.com/tumblr/tumblr.go"
)

type DownloadPosts struct {
	scheduler task.WorkScheduler
	Count     int
	index     int
	batchSize int
}

func NewDownloadPosts(count int, index int, scheduler task.WorkScheduler) *DownloadPosts {
	//TODO should I move all locked post back to posts???
	d := &DownloadPosts{scheduler: scheduler, Count: count, index: index}
	d.batchSize = 1
	return d
}

func (d *DownloadPosts) Run() {
	for i := 0; i < d.Count; i += d.batchSize {
		log.Println("Download posts batch, worker: ", d.index, " size: ", d.batchSize)
		waitRand := rand.Intn(10000000) % 2000
		time.Sleep(time.Millisecond * time.Duration(waitRand))
		posts, err := d.lockPosts()
		if err != nil {
			log.Println("downloading error when lock post: ", err)
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
	downloadPosts := make([]*downloadPost, d.batchSize)
	postPaths, err := getSomePostFile(d.batchSize)
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

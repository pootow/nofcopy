package common

import (
	"github.com/tumblr/tumblr.go"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func Download(targetUrl string) error {

	targetURL, err := url.Parse(targetUrl)
	if err != nil {
		return err
	}

	domain := targetURL.Hostname()
	fileDir, err := GetAndMakeInAppWorkingDir("videos", domain)
	if err != nil {
		return err
	}

	fileName := targetURL.Path
	filePath := path.Join(fileDir, fileName)

	head, err := http.Head(targetUrl)
	if err != nil {
		return err
	}

	remoteLength := head.ContentLength

	if downloaded, err := targetFileDownloaded(filePath, remoteLength); err != nil {
		return err
	} else {
		if downloaded {
			log.Println(filePath, " downloaded.")
			return nil
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(targetUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	written, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	println(written, " bytes written.")
	return nil
}

func targetFileDownloaded(filePath string, length int64) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	diskLength := info.Size()

	if diskLength == length {
		return true, nil
	}
	return false, nil
}

func DownloadPost(postInterface tumblr.PostInterface) {
	panic("implement this")
}

func LockPosts() ([]tumblr.PostInterface, error) {
	count := 3
	postInterfaces := make([]tumblr.PostInterface, count)
	postPaths, err := getSomePostFile(count)
	if err != nil {
		return nil, err
	}
	for i, postPath := range postPaths {
		newPath, err := lockPostPath(postPath)
		if err != nil {
			return nil, err // or print err and continue
		}
		postInterface, err := loadPost(newPath)
		if err != nil {
			return nil, err
		}
		postInterfaces[i] = postInterface
	}
	return nil, nil
}

func loadPost(postPath string) (tumblr.PostInterface, error) {
	panic("not implement")
}

func lockPostPath(postPath string) (string, error) {
	panic("not implement")
}

func getSomePostFile(count int) ([]string, error) {
	panic("implement this")
}

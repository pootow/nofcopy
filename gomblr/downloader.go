package gomblr

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
)

func Download(targetUrl string) error {

	targetURL, err := url.Parse(targetUrl)
	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	userHomePath := currentUser.HomeDir

	basePath := "gomblr/videos/"

	domain := targetURL.Hostname()

	fileName := targetURL.Path

	fileDir := path.Join(userHomePath, basePath, domain)
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return err
	}

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

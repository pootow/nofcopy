package common

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func Download(targetUrl string) error {

	println("start downloading: ", targetUrl)

	targetURL, err := url.Parse(targetUrl)
	if err != nil {
		return err
	}

	domain := targetURL.Hostname()
	fileDir, err := GetAndMakeInAppWorkingDir("resources", domain)
	if err != nil {
		return err
	}

	fileName := targetURL.Path
	_, fileName = filepath.Split(fileName)
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
			log.Println(filePath, " already downloaded.")
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
		return false, nil
	}
	diskLength := info.Size()

	if diskLength == length {
		return true, nil
	}
	return false, nil
}

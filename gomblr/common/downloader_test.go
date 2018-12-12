package common

import (
	"net/url"
	"os"
	"os/user"
	"path"
	"testing"
)

func TestDownloading(t *testing.T) {
	targetUrl := "https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4"

	if err := Download(targetUrl); err != nil {
		t.Fatal(err)
	}

}

func TestURLParsing(t *testing.T) {
	targetUrl := "https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4"
	targetURL, _ := url.Parse(targetUrl)
	println(targetURL.Host)
	println(targetURL.Hostname())

	println(targetURL.Path)

	appWorkingDir, err := GetAppWorkingDir()
	if err != nil {
		t.Fatal(err)
	}

	basePath := path.Join(appWorkingDir, "videos")

	domain := targetURL.Hostname()

	fileName := targetURL.Path

	fileDir := path.Join(basePath, domain)

	filePath := path.Join(fileDir, fileName)
	println(filePath)
}

func TestMakePath(t *testing.T) {
	current, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}

	testPath := path.Join(current.HomeDir, "a/b/c/notadir")
	err = os.MkdirAll(testPath, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

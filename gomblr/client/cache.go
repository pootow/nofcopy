package client

import (
	"bufio"
	"github.com/tumblr/tumblr.go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type cache struct {
	dir string
}

func newCache(dir string) Cache {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println("error when making sure cache dir, noopCache returned: ", err)
		return newNoopCache()
	}
	return &cache{dir: dir}
}

type Cache interface {
	Get(endpoint string, params *url.Values) (*tumblr.Response, error)
	Set(endpoint string, params *url.Values, body []byte) error
	GetCacheKey(endpoint string, params *url.Values) string
}

func (c *cache) Get(endpoint string, params *url.Values) (*tumblr.Response, error) {
	cacheFilePath := path.Join(c.dir, c.GetCacheKey(endpoint, params))

	cacheFile, err := os.Open(cacheFilePath)
	if err != nil {
		return nil, err
	}
	defer cacheFile.Close()

	reader := bufio.NewReader(cacheFile)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	response := tumblr.NewResponse(body, http.Header{})

	return response, nil
}

func (c *cache) Set(endpoint string, params *url.Values, body []byte) error {
	cacheFilePath := path.Join(c.dir, c.GetCacheKey(endpoint, params))

	file, err := os.Create(cacheFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(body); err != nil {
		return err
	}

	return nil
}

func (c *cache) GetCacheKey(endpoint string, params *url.Values) string {
	paramsString := ""
	if params != nil {
		paramsString = params.Encode()
	}
	return strings.Replace(endpoint, "/", "[", -1) + "$" + paramsString
}

type noopCache struct{}

func newNoopCache() *noopCache {
	return &noopCache{}
}

func (noopCache) Get(endpoint string, params *url.Values) (*tumblr.Response, error) {
	return nil, nil
}

func (noopCache) Set(endpoint string, params *url.Values, body []byte) error {
	return nil
}

func (noopCache) GetCacheKey(endpoint string, params *url.Values) string {
	return ""
}

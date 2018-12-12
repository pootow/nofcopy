package client

import (
	"github.com/pootow/nofcopy/gomblr/common"
	"github.com/tumblr/tumblr.go"
	client "github.com/tumblr/tumblrclient.go"
	"log"
	"net/url"
	"path"
)

type GomblrClient struct {
	*client.Client
	cb    callback
	cache Cache
}

type callback func(*tumblr.Response)

// Issue GET request to Tumblr API
func (g *GomblrClient) Get(endpoint string) (tumblr.Response, error) {
	var response *tumblr.Response
	var err error

	if response, err = g.cache.Get(endpoint, nil); !(err == nil && response != nil) {
		*response, err = g.Client.Get(endpoint)
		if err != nil {
			return tumblr.Response{}, err
		}

		err := g.cache.Set(endpoint, nil, response.GetBody())
		if err != nil {
			log.Println("cache error just log: ", err)
		}
	}

	if err == nil && g.cb != nil {
		g.cb(response)
	}

	return *response, err
}

// Issue GET request to Tumblr API with param values
func (g *GomblrClient) GetWithParams(endpoint string, params url.Values) (tumblr.Response, error) {
	var response *tumblr.Response
	var err error

	if response, err = g.cache.Get(endpoint, &params); err != nil || response == nil {
		freshResponse, freshErr := g.Client.GetWithParams(endpoint, params)
		if freshErr != nil {
			return tumblr.Response{}, freshErr
		}
		response = &freshResponse
		err = freshErr

		cacheErr := g.cache.Set(endpoint, &params, response.GetBody())
		if cacheErr != nil {
			log.Println("cache error just log: ", cacheErr)
		}
	}

	if err == nil && g.cb != nil {
		g.cb(response)
	}

	return *response, err
}

func NewGomblrClient() *GomblrClient {
	defaultClient := client.NewClientWithToken(
		//"8zrB1oPfwIKPbdzmXYkwfxlLLpJwcaa9vGrjVWrNKCcJhpLxPw",
		//"eDGjnvk3wUobkHa50cp4v5dXHbkGgLKtDg1QtBz4gXdyBaSliu",
		//"pFlbnVLpOZ5eqdwsUxsdw2525ODTjN3Lk9XuSqptlyWAXrF6sY",
		//"ON4ALkLIS49DGYVkIIwFKxXPmRkeNKk5qDc8H1EOvAFHT19ivw",
		"J0HHJ4vqE805MyOoVERgIxDCoLwzaMkLojhRsCNfltIiEuYR8Q",
		"ChafKsWr8Kr4pfRUtRNgcJYXkZD8cqiLkgoc01Sq0ndt3nBH2M",
		"v7hn1N4JFlYBHfW48isB8mPBPAuogxTxM6T3PTp1p4c5q72oIj",
		"tUhJKaFBO4LJn1Rci4XdVuOi7ddbT3qaswPLBxBjb3M0xmBuli",
	)

	var cache Cache = nil

	appWorkingDir, err := common.GetAppWorkingDir()
	if err != nil {
		log.Println("error get cache dir, disable cache and log it: ", err)
		cache = newNoopCache()
	} else {
		cacheDir := path.Join(appWorkingDir, "cache")
		cache = newCache(cacheDir)
	}

	gclient := &GomblrClient{
		Client: defaultClient,
		cache:  cache,
	}
	return gclient
}

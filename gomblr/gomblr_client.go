package gomblr

import (
	"net/url"

	tumblr "github.com/tumblr/tumblr.go"
	client "github.com/tumblr/tumblrclient.go"
)

type gomblrClient struct {
	*client.Client
	cb    callback
	cache Cache
}

type cache struct {
	dir string
}

type Cache interface {
	Get(url string) (tumblr.Response, error)
}

func (*cache) Get(url string) (tumblr.Response, error) {
	return tumblr.Response{}, nil
}

type callback func(*tumblr.Response)

// Issue GET request to Tumblr API
func (g *gomblrClient) Get(endpoint string) (tumblr.Response, error) {
	var response tumblr.Response
	var err error
	if response, err = g.cache.Get(endpoint); err != nil {
		response, err = g.Client.Get(endpoint)
	}
	if err == nil && g.cb != nil {
		g.cb(&response)
	}
	return response, err
}

// Issue GET request to Tumblr API with param values
func (g *gomblrClient) GetWithParams(endpoint string, params url.Values) (tumblr.Response, error) {
	response, err := g.Client.GetWithParams(endpoint, params)
	if err == nil && g.cb != nil {
		g.cb(&response)
	}
	return response, err
}

func newGomblrClient() *gomblrClient {
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
	client := &gomblrClient{
		Client: defaultClient,
		cache:  new(cache),
	}
	return client
}

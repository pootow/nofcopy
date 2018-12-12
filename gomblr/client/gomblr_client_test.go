package client

import (
	"github.com/tumblr/tumblr.go"
	"net/url"
	"testing"
)

func TestNewGomblrClient(t *testing.T) {
	client := NewGomblrClient()
	dashboard, _ := tumblr.GetDashboard(client, make(url.Values))
	t.Log(dashboard.Posts)
}

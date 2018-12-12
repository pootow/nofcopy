package gomblr

import (
	"github.com/tumblr/tumblr.go"
	"regexp"
)

type PostResourceExtractor interface {
	GetResources() []string
}

const HTTP_RESOURCE string = `http(s)?://[^\\\"'\<\?\&]+(\.jpg|\.png|\.gif|\.webp|\.mp4|\.amr|\.wav|\.3gp)`

type gTextPost struct {
	tumblr.TextPost
}

func (g *gTextPost) GetResources() []string {
	regex := regexp.MustCompile(HTTP_RESOURCE)

	allString := regex.FindAllString(g.Body, -1)
	set := make(map[string]struct{})
	for _, String := range allString {
		set[String] = struct{}{}
	}
	resources := make([]string, len(set))
	i := 0
	for e := range set {
		resources[i] = e
		i++
	}
	return resources
}

type gVideoPost struct {
	tumblr.VideoPost
}

func (g *gVideoPost) GetResources() []string {
	return []string{g.VideoUrl}
}

type gPhotoPost struct {
	tumblr.PhotoPost
}

func (g *gPhotoPost) GetResources() []string {
	photos := make([]string, len(g.Photos))
	for i, photo := range g.Photos {
		photos[i] = photo.OriginalSize.Url
	}
	return photos
}

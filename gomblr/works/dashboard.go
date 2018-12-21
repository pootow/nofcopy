package works

import (
	"errors"
	"github.com/pootow/nofcopy/gomblr/client"
	"github.com/tumblr/tumblr.go"
	"net/url"
	"strconv"
)

type dashboard struct {
	client *client.GomblrClient
}

func NewDashboard(client *client.GomblrClient) *dashboard {
	return &dashboard{client: client}
}

func (d *dashboard) Run() {
	earliestPostId, err := d.getEarliestPostId()
	if err != nil {
		println("can not get earliest post: ", err)
		return
	}
	firstTime := true
	var dashboard *tumblr.Dashboard
	for {
		if firstTime {
			dashboard, err = d.getMyDashboardSince(earliestPostId, 20)
			firstTime = false
		} else {
			dashboard, err = dashboard.NextBySinceId()
		}
		if err != nil {
			println("error when get dashboard: ", err)
			continue
		}
		if len(dashboard.Posts) == 0 {
			break
		}
		client.Store(dashboard.Posts)
	}
}

func (d *dashboard) getMyDashboard(limit int) (*tumblr.Dashboard, error) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	dashboard, err := tumblr.GetDashboard(d.client, params)
	return dashboard, err
}

func (d *dashboard) getEarliestPostId() (uint64, error) {
	params := url.Values{}
	params.Set("limit", "1")
	params.Set("offset", "9999999999")
	dashboard, err := tumblr.GetDashboard(d.client, params)
	if len(dashboard.Posts) < 1 {
		return 0, errors.New("server returned 0 posts")
	}
	return dashboard.Posts[0].GetSelf().Id, err
}

func (d *dashboard) getMyDashboardSince(postId uint64, limit int) (*tumblr.Dashboard, error) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	params.Set("since_id", strconv.Itoa(int(postId)))
	dashboard, err := tumblr.GetDashboard(d.client, params)
	return dashboard, err
}

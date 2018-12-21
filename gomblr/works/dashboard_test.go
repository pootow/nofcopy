package works

import (
	"github.com/pootow/nofcopy/gomblr/client"
	"github.com/tumblr/tumblr.go"
	"net/url"
	"testing"
)

func TestDashboardWork(t *testing.T) {
	dashboard := NewDashboard(client.NewGomblrClient())

	count := 10

	var dash *tumblr.Dashboard
	var err error
	for i := 0; i < count; i++ {
		if i == 0 {
			dash, err = dashboard.getMyDashboard(1)

		} else {
			dash, err = dash.NextByOffset()

		}
		if err != nil {
			t.Fatal(err)
		}
		for _, post := range dash.Posts {
			println("============================================")
			println(post.GetSelf().Caption)
			println("============================================")
		}
	}
}

func TestDashboardTimeline(t *testing.T) {
	client := client.NewGomblrClient()

	dash, _ := client.GetDashboard()
	for _, post := range dash.Posts {
		println(post.GetSelf().Id)
	}
	println("========================================")

	params := url.Values{}
	params.Set("limit", "10")
	params.Set("offset", "10")
	offsetDash, _ := client.GetDashboardWithParams(params)
	for _, post := range offsetDash.Posts {
		println(post.GetSelf().Id)
	}
	println("========================================")

}

func TestGetEarliestPostId(t *testing.T) {
	client := client.NewGomblrClient()

	dash := NewDashboard(client)
	postId, _ := dash.getEarliestPostId()
	println(postId)
}

func TestDashboard_Run(t *testing.T) {
	client := client.NewGomblrClient()

	dash := NewDashboard(client)
	dash.Run()
}

package gomblr

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"

	"github.com/pootow/nofcopy/task"

	tumblr "github.com/tumblr/tumblr.go"
)

func TestConnection(t *testing.T) {
	client := newGomblrClient()
	dashboard, _ := tumblr.GetDashboard(client, make(url.Values))
	t.Log(dashboard.Posts)
}

func TestFav(t *testing.T) {
	client := newGomblrClient()
	client.cb = func(r *tumblr.Response) {
		t.Log(string(r.GetBody()))
		t.Log("xxxxxxxxxxx")
		t.Log("xxxxxxxxxxx")
		t.Log("xxxxxxxxxxx")
		t.Log("xxxxxxxxxxx")
		t.Log("xxxxxxxxxxx")
	}

	params := url.Values{}
	params.Set("limit", "1")
	params.Set("before", "1543418485")
	fav, _ := tumblr.GetLikes(client, params)
	posts, _ := fav.Full()
	favBytes, _ := json.Marshal(posts)

	t.Log(string(favBytes))
}

func TestJson(t *testing.T) {
	bytes, _ := json.Marshal(struct {
		A int
		B string
	}{
		A: 3,
		B: "中文",
	})
	t.Log(string(bytes))

	mapp := make(map[string]int)
	mapp["key"] = 1
	t.Log(mapp)
	out, _ := json.Marshal(mapp)
	t.Log(string(out))
}

func TestFavWork(t *testing.T) {
	st := task.NewSimpleTask()
	client := newGomblrClient()

	st.Add(&likes{
		client,
	})

	//blogs := []string{
	//
	//}
	//
	//for _, blog := range blogs {
	//	st.Add(&blogPosts{
	//		client,
	//		blog,
	//	})
	//}

	//st.Add(&following{client})

	st.Wait()

}

func TestOsEnv(t *testing.T) {
	println(os.Getenv("HOME"))
}

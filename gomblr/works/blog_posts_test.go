package works

import (
	. "github.com/pootow/nofcopy/gomblr/client"
	"github.com/pootow/nofcopy/task"
	"github.com/pootow/nofcopy/utils"
	"strings"
	"testing"
)

func TestBlogPostsWork(t *testing.T) {
	st := task.NewSimpleTask()
	client := NewGomblrClient()

	blogsText := `
whole-video
sexy---heels
dog-bitcher
hupimao1010
blackbear7483
sm-vedio
sexygirl031
tugegegege
ws18s
yanhouzhongteng
pgb10
angela910617
rusuanjiang1
rusuanjiang123
1500t
nidemugou
veryunlikelysublimecollection
melhuu
yaweicat
mrpaoly
`

	blogs := strings.Split(strings.Trim(blogsText, "\n"), "\n")

	chunks := utils.StringSliceChunk(blogs, len(blogs)/2+1)

	for _, chunk := range chunks {
		st.Add(NewBlogsWork(client, chunk))
	}

	st.Wait()

}

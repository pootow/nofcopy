package gomblr

import (
	"encoding/json"
	"github.com/pootow/nofcopy/task"
	. "github.com/tumblr/tumblr.go"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
)

type likes struct {
	client *gomblrClient
}

func (l *likes) Run() {
	before := ""
	for {
		println("==================================>>>>>>>>>>>>>my likes before: ", before)
		likes := l.getMyLikes(before, 20)
		if len(likes.Posts) == 0 {
			break
		}
		posts, _ := likes.Full()
		store(posts)
		before = likes.Links.Next.QueryParams["before"]
	}

}

type following struct {
	client *gomblrClient
}

func (f *following) Run() {
	offset := uint(3120)
	for {

		followingList, _ := GetFollowing(f.client, offset, 20)

		if len(followingList.Blogs) == 0 {
			break
		}
		offset64, _ := strconv.ParseUint(followingList.Links.Next.QueryParams["offset"], 10, 0)
		println("offset: -------------------------->>", offset64)
		for _, blog := range followingList.Blogs {
			println(blog.Name)
		}
		offset = uint(offset64)
	}

}

type blogPosts struct {
	client   *gomblrClient
	blogName string
}

func (b *blogPosts) Run() {
	offset := ""
	for {
		posts := b.getPosts(b.blogName, offset, 20)
		if len(posts.Posts) == 0 {
			break
		}
		postsFace, _ := posts.All()
		store(postsFace)
		offset = posts.Links.Next.QueryParams["offset"]
	}

}

func (b *blogPosts) getPosts(name string, offset string, limit int) *Posts {
	params := url.Values{}
	if offset != "" {
		params.Set("offset", offset)
	}
	params.Set("limit", strconv.Itoa(limit))
	posts, _ := GetPosts(b.client, name, params)
	return posts
}

func store(posts []PostInterface) {
	for _, post := range posts {
		miniPost := post.GetSelf()
		println("------------------------------------")
		println("-=type=-: ", miniPost.Type)
		postBytes, _ := json.Marshal(post)
		postJson := string(postBytes)
		savePost(miniPost.Id, postJson)
		switch post.(type) {
		case *TextPost, *AnswerPost, *ChatPost, *LinkPost, *QuotePost:
			println(miniPost.Summary)
		case *PhotoPost:
			photo := post.(*PhotoPost)
			println("photo perm link: ", photo.ImagePermalink)
			println("photos: ")
			for _, img := range photo.Photos {
				println("[", img.Caption, "]: ", img.OriginalSize.Url)
			}
		case *AudioPost:
			audio := post.(*AudioPost)
			println("audio url: ", audio.AudioUrl)
			println("audio source url: ", audio.AudioSourceUrl)
			println("caption: ", audio.Caption)
			println("summary: ", audio.Summary)
		case *VideoPost:
			video := post.(*VideoPost)
			println("perm link: ", video.PermalinkUrl)
			println("video url: ", video.VideoUrl)
			println("video: ")
			for key, val := range video.Video {
				println(key, val.VideoId)
			}
		default:
			panic(post)
		}
		println("------------------------------------")
	}
}

func savePost(id uint64, jsonString string) {

	basePath := path.Join(os.Getenv("HOME"), "/gomblr/posts/")
	filePath := path.Join(basePath, strconv.FormatUint(id, 10)+".json")
	if stat, err := os.Stat(filePath); !os.IsNotExist(err) {
		log.Println("error check file exists: ", filePath, " > ", stat.Name())
		return
	}
	if file, err := os.Create(filePath); err != nil {
		log.Println("error when create post file: ", filePath, " err: ", err)
	} else {
		defer file.Close()
		if _, err := file.Write([]byte(jsonString)); err != nil {
			log.Println("error when save post to disk: ", filePath, " err: ", err)
		}
	}
}

func (l *likes) getMyLikes(before string, limit int) *Likes {
	params := url.Values{}
	if before != "" {
		params.Set("before", before)
	}
	params.Set("limit", strconv.Itoa(limit))
	likes, _ := GetLikes(l.client, params)
	return likes
}

func GetBlogPosts(blogs []string) {
	st := task.NewSimpleTask()
	client := newGomblrClient()

	for _, blog := range blogs {
		st.Add(&blogPosts{
			client,
			blog,
		})
	}
	st.Wait()
}

func DownloadPostResource() {

}

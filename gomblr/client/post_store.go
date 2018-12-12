package client

import (
	"encoding/json"
	"fmt"
	"github.com/pootow/nofcopy/gomblr/common"
	. "github.com/tumblr/tumblr.go"
	"log"
	"os"
	"path"
	"strconv"
)

func Store(posts []PostInterface) {
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

	appWorkingDir, err := common.GetAppWorkingDir()
	if err != nil {
		panic(fmt.Sprint("error get app dir: ", err))
	}
	basePath := path.Join(appWorkingDir, "posts")
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

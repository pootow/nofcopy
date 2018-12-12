package client

import (
	"encoding/json"
	"fmt"
	"github.com/pootow/nofcopy/gomblr/common"
	"github.com/pootow/nofcopy/gomblr/extractors"
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
		savePost(miniPost.Id, postBytes)

		println(miniPost.Summary)
		println("-----------")

		println("resources:")

		extractor := extractors.GetExtractor(post)
		if extractor == nil {
			println("no resource in this type of post: ", miniPost.Type)
			continue
		}

		resources := extractor.GetResources()
		for i, resource := range resources {
			println("[ ", strconv.Itoa(i), " ]\t", resource)
		}

		println("------------------------------------")
	}
}

func savePost(id uint64, postBytes []byte) {

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
		if _, err := file.Write(postBytes); err != nil {
			log.Println("error when save post to disk: ", filePath, " err: ", err)
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"github.com/pootow/nofcopy/gomblr"

	"github.com/pootow/nofcopy/task"
)

var (
	mode = flag.String("mode", "du", "du(disk usage) or cs(checksum)")
	sync = flag.Bool("sync", false, "sync will walk dir file one by one, best with non-ssd")
)

func main() {
	flag.Parse()

	fmt.Println("mode: ", *mode)

	switch *mode {
	case "du":
		if len(flag.Args()) > 0 {
			t := task.NewStatTask()
			t.Stat(flag.Arg(0))
			fmt.Printf("total: %d bytes\n", t.Total)
		} else {
			fmt.Println("please specify a path.")
		}
	case "cs":
		if len(flag.Args()) > 0 {
			t := task.NewChecksumTask()
			t.Checksum(flag.Arg(0), *sync)
		} else {
			fmt.Println("please specify a path.")
		}
	case "gf":
		gomblr.GetFavPosts()
	case "gg":
		if len(flag.Args()) > 0 {
			gomblr.GetBlogPosts(flag.Args())
		}
	case "gd":
		if len(flag.Args()) > 0 {
			gomblr.DownloadPosts(flag.Arg(0), flag.Arg(1))
		}
	case "gdash":
		gomblr.GetAllDashbordPosts()
	}

}

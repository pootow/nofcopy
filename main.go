package main

import (
	"flag"
	"fmt"
	"github.com/pootow/nofcopy/gomblr"

	"github.com/pootow/nofcopy/task"
)

var (
	mode = flag.String("mode", "du", "du(disk usage) or cs(checksum)")
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
			t.Checksum(flag.Arg(0))
		} else {
			fmt.Println("please specify a path.")
		}
	case "gg":
		if len(flag.Args()) > 0 {
			gomblr.GetBlogPosts(flag.Args())
		}
	case "gd":
		if len(flag.Args()) > 0 {
			gomblr.DownloadPosts(flag.Arg(0))
		}
	}

}

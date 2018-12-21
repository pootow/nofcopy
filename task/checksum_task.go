package task

import (
	"fmt"
	"os"
	"time"

	"github.com/pootow/nofcopy/utils"
)

// ChecksumTask is a Task that handling a path to do checksum of all its files.
type ChecksumTask struct {
	DirWalkTask
}

// NewChecksumTask creates new ChecksumTask
func NewChecksumTask() *ChecksumTask {
	t := new(ChecksumTask)
	t.initDirWalkTask(&walker{})
	return t
}

// Checksum shows sha1 checksum of all files for a path
func (t *ChecksumTask) Checksum(root string) {
	t.Walk(Item{item: root})
}

type walker struct {
	DefaultWalker
}

func (t *walker) onFile(file Item) {
	path := file.item.(string)
	hash := utils.GetSHA1Hash(path)
	fmt.Println(file, ": ", hash)

	hashFile, err := os.Create(path + hash + ".sha1.sum")
	if err != nil {
		fmt.Println("error when writing hash file: ", err)
		return
	}
	defer hashFile.Close()
	hashFile.WriteString(fmt.Sprintln("sha1 ", hash, " ", time.Now().UnixNano()))

}

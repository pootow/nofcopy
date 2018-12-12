package common

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

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

func TestOsEnv(t *testing.T) {
	println(os.Getenv("HOME"))
}

func TestTrim(t *testing.T) {
	println(strings.Trim("\n$$\n3\n4&&$$5$\n", "\n$"))
}

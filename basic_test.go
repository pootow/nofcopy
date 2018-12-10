package main

import (
	"testing"

	"github.com/pootow/nofcopy/utils"
)

func TestTrue(t *testing.T) {
}

func TestFileSHA1Hash(t *testing.T) {
	const path = "test_fixture/ee01f5f6ad1f881e1626242ab182e5d93f53de58.md"
	const hash = "ee01f5f6ad1f881e1626242ab182e5d93f53de58"
	actualHash := utils.GetSHA1Hash(path)
	if actualHash != hash {
		t.Errorf("got a error hash.\n actual: %s\n expected: %s\n", actualHash, hash)
	}
}

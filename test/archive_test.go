package test

import (
	"FastTdl/core"
	"testing"
)

// go test -v -timeout 0 -run TestArchiveAllFiles
func TestArchiveAllFiles(t *testing.T) {
	core.ArchiveAllFiles("/Users/zen/gitea/FastTdl/discussions")
}

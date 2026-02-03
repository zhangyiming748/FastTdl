package FastTdl

import "testing"

// go test -v -timeout 0 -run TestArchiveAllFiles
func TestArchiveAllFiles(t *testing.T) {
	ArchiveAllFiles("/Users/zen/gitea/FastTdl/discussions")
}
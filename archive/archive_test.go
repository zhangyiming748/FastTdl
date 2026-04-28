package archive

import "testing"
// go test -v -timeout 0 -run TestArchiveMusic
func TestArchiveMusic(t *testing.T) {
	root := "F:\\tdl\\ts"
	Videos(root,false)
}

package archive

import (
	"testing"
)

// go test -timeout 30h -v -run TestGetAllFiles
func TestGetAllFiles(t *testing.T) {
	files, _ := GetAllFiles("/Users/zen/github/FastTdl")
	for _, v := range files {
		ConvertH265(v)
	}
}

package archive

import (
	"testing"
)

// go test -timeout 30h -v -run TestGetAllFiles
func TestGetAllFiles(t *testing.T) {
	files, _ := GetAllVideoFiles("/Users/zen/github/FastTdl")
	for _, v := range files {
		ConvertH265(v)
	}
}

// go test -timeout 30h -v -run TestGetAllImages
func TestGetAllImages(t *testing.T) {
	files,_:=GetAllImageFiles("/Users/zen/Downloads/陈好")
	for _,v:=range files {
		ConvertAVIF(v)
	}
}
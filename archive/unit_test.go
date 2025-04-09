package archive

import (
	l "github.com/zhangyiming748/FastTdl/util"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

// go test -timeout 30h -v -run TestGetAllFiles
func TestGetAllFiles(t *testing.T) {
	l.SetLog("h265.log")
	files, _ := GetAllVideoFiles("D:\\pikpak\\Russia Funny show 俄羅斯电视台整蛊節目精选集第一季")
	for _, v := range files {
		ConvertH265(v)
	}
}

// go test -timeout 30h -v -run TestGetAllImages
func TestGetAllImages(t *testing.T) {
	files, _ := GetAllImageFiles("/Users/zen/Downloads/media")
	for _, v := range files {
		ConvertAVIF(v)
	}
}

package archive

import (
	l "github.com/zhangyiming748/FastTdl/util"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

// go test -timeout 30h -v -run TestArchiveAllVideoFiles
func TestArchiveAllVideoFiles(t *testing.T) {
	l.SetLog("h265.log")
	files, _ := GetAllVideoFiles("D:\\pikpak\\Russia Funny show 俄羅斯电视台整蛊節目精选集第一季")
	for _, v := range files {
		ConvertH265(v)
	}
}

// go test -timeout 30h -v -run TestArchiveAllImageFiles
func TestArchiveAllImageFiles(t *testing.T) {
	files, _ := GetAllImageFiles("/Users/zen/Downloads/media")
	for _, v := range files {
		ConvertAVIF(v)
	}
}

// go test -timeout 30h -v -run TestArchiveAllAudioBookFiles
func TestArchiveAllAudioBookFiles(t *testing.T) {
	files, _ := GetAllAudioFiles("/Users/zen/Downloads/media")
	for _, v := range files {
		ConvertAudio(v,AudioBookType)
	}
}

// go test -timeout 30h -v -run TestArchiveAllRapMusicFiles
func TestArchiveAllRapMusicFiles(t *testing.T) {
	files, _ := GetAllAudioFiles("/Volumes/Fanxiang/有声读物3/激情骚麦")
	for _, v := range files {
		ConvertAudio(v,RapMusicType)
	}
}
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
	files, _ := GetAllVideoFiles("/Volumes/Fanxiang/整理")
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
	files, _ := GetAllAudioFiles("/Volumes/Fanxiang/有声读物3/自慰催眠")
	for _, v := range files {
		ConvertAudio(v,AudioBookType)
	}
}

// go test -timeout 30h -v -run TestArchiveAllRapMusicFiles
func TestArchiveAllRapMusicFiles(t *testing.T) {
	files, _ := GetAllAudioFiles("/Volumes/Fanxiang/有声读物3/增大电平处理后")
	for _, v := range files {
		ConvertAudio(v,RapMusicType)
	}
}
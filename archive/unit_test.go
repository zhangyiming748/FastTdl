package archive

import (
	l "github.com/zhangyiming748/FastTdl/util"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	defer l.Alarm("单独的hevc归档完成")
	l.SetLog("h265.log")
}

// go test -timeout 30h -v -run TestArchiveAllVideoFiles
func TestArchiveAllVideoFiles(t *testing.T) {
	root := "/Users/zen/Downloads/media"
	dirs, e := GetFinalSubDirs(root)
	if e != nil {
		t.Error(e)
		return
	}
	for _, dir := range dirs {
		files, _ := GetAllVideoFiles(dir)
		for _, file := range files {
			ConvertH265(file)
		}
	}

}

// go test -timeout 30h -v -run TestArchiveAllImageFiles
func TestArchiveAllImageFiles(t *testing.T) {
	root := "/Users/zen/Downloads/10350524"
	dirs, e := GetFinalSubDirs(root)
	if e != nil {
		t.Error(e)
		return
	}
	for _, dir := range dirs {
		files, _ := GetAllImageFiles(dir)
		for _, file := range files {
			ConvertAVIF(file)
		}
	}
}

// go test -timeout 30h -v -run TestArchiveAllAudioBookFiles
func TestArchiveAllAudioBookFiles(t *testing.T) {
	root := "/Volumes/Fanxiang/有声读物3"
	dirs, e := GetFinalSubDirs(root)
	if e != nil {
		t.Error(e)
		return
	}
	for _, dir := range dirs {
		files, _ := GetAllAudioFiles(dir)
		for _, file := range files {
			ConvertAudio(file, AudioBookType)
		}
	}
}

// go test -timeout 30h -v -run TestArchiveAllRapMusicFiles
func TestArchiveAllRapMusicFiles(t *testing.T) {
	root := "/Volumes/Fanxiang/有声读物3"
	dirs, e := GetFinalSubDirs(root)
	if e != nil {
		t.Error(e)
		return
	}
	for _, dir := range dirs {
		files, _ := GetAllAudioFiles(dir)
		for _, file := range files {
			ConvertAudio(file, RapMusicType)
		}
	}
}

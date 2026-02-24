package archive

import (
	"log"
	"testing"

	"FastTdl/sqlite"
	"FastTdl/util"

	"github.com/zhangyiming748/finder"
)

func init() {
	util.SetLog("h265.log")
	sqlite.SetSqlite()
}

// go test -timeout 30h -v -run TestArchiveAllFiles
func TestArchiveAllFiles(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic: %v", err)
		}
	}()
	//roots := []string{"N:\\archive"}
	roots := []string{
		"/videos",
	}
	for _, root := range roots {
		Videos(root)
		Images(root)
	}
}
func TestArchiveAllMkvs(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic: %v", err)
		}
	}()
	//roots := []string{"N:\\archive"}
	roots := []string{
		"Q:\\dst",
	}
	for _, root := range roots {
		Movies(root)
	}
}

// go test -timeout 30h -v -run TestArchiveAllAudioBookFiles
func TestArchiveAllAudioBookFiles(t *testing.T) {
	root := "/Volumes/Fanxiang/有声读物3"
	dirs := finder.FindAllFolders(root)
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
	dirs := finder.FindAllFolders(root)
	for _, dir := range dirs {
		files, _ := GetAllAudioFiles(dir)
		for _, file := range files {
			ConvertAudio(file, RapMusicType)
		}
	}
}

func TestVideo(t *testing.T) {
	fp := "O:\\pikpak\\双子母性本能"
	videos := finder.FindAllVideos(fp)
	for _, video := range videos {
		log.Println(video)

	}
}

package archive
import (
	"testing"
)
// go test -v -run TestGetAllFiles
func TestGetAllFiles(t *testing.T) {
	files,_:=GetAllFiles("/Users/zen/github/FastTdl")
	t.Log(files)
}

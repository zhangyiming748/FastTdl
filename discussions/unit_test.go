package discussions

import (
	"testing"
)

func TestParse(t *testing.T) {
	u := "https://t.me/soqkkqossmsn/48334"
	Discussions(u, "")
}

// go test -timeout 30h -v -run TestDownloads
func TestDownloads(t *testing.T) {
	DownloadAllDiscussions("http://127.0.0.1:8889")
}

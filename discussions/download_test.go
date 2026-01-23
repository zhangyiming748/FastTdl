package discussions

import (
	"github.com/zhangyiming748/FastTdl/constant"
	"github.com/zhangyiming748/FastTdl/util"
	"testing"
)

// go test -v  -timeout 3h -run TestDownloadDiscussions
func TestDownloadDiscussions(t *testing.T) {
	p := constant.Parameter{
		Proxy:      constant.DEFAULT_PROXY,
		MainFolder: "./media",
	}
	uris := util.ReadByLine("discuss.txt")
	for _, uri := range uris {
		Discussion(uri, p)
	}
}

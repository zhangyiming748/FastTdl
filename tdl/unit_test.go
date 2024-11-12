package tdl

import "testing"

func TestSplit(t *testing.T) {
	u := "https://t.me/ciyuanb/43059 9"
	list := parseUrlWithOffset(u)
	t.Log(list)
}
func TestSplit2(t *testing.T) {
	u := "https://t.me/ciyuanb/43059#2b 9"
	list := parseUrlWithTagAndOffset(u)
	t.Log(list)
}

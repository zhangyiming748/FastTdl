package mysql

import "testing"

// go test -v -run TestSetMysql
func TestSetMysql(t *testing.T) {
	SetMysql()
	if !UseMysql() {
		t.Error("SetMysql failed")
	}
}

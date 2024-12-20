package util

import (
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"testing"
)

func TestFind(t *testing.T) {
	SetLevelDB()
	get, err := GetLevelDB().Get([]byte("key"), nil)
	if errors.Is(err, leveldb.ErrNotFound) {
		log.Printf("get = %v, err = %v\n", get, err)
	}
	log.Printf("get = %v, err = %v\n", get, err)
	err = GetLevelDB().Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		return
	}
}

package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func changeOwnership(path string, uid, gid int) error {
	// 更改文件或文件夹的属主和属组
	err := os.Chown(path, uid, gid)
	if err != nil {
		return fmt.Errorf("failed to change ownership of %s: %v", path, err)
	}
	return nil
}

func walkDir(dir string, uid, gid int) error {
	// 遍历目录及其子目录
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 更改当前文件或文件夹的属主和属组
		if err := changeOwnership(path, uid, gid); err != nil {
			return err
		}

		return nil
	})
}

func Chown(dir string, uid, gid int) {
	if err := walkDir(dir, uid, gid); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Ownership changed successfully.")
	}
}

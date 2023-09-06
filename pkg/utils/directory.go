package utils

import (
	"errors"
	"fmt"
	"os"
)

// PathExists 確認資料夾是否存在
func PathExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		if file.IsDir() {
			// 路徑存在，且為資料夾
			return true, nil
		}
		// 路徑存在，且為檔案
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MkdirIfNotExist 創建資料夾，如果資料夾不存在
func MkdirIfNotExist(path string) {
	if ok, _ := PathExists(path); !ok {
		// 沒有的話創建一個
		fmt.Printf("create %v directory\n", path)
		_ = os.Mkdir(path, os.ModePerm)
	}
}

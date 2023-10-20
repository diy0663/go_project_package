package file

import "os"

// 判断指定文件是否存在
func Exists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// 写数据到指定文件
func Put(data []byte, targetFile string) error {
	err := os.WriteFile(targetFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

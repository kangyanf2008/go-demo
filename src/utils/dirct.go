package utils

import (
	"fmt"
	"os"
)

// IfNotExistMk 目录不存在则创建
func IfNotExistMk(dirPath string) error {

	// 判断目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，则创建目录
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录失败: %v\n", err)
			return err
		}
		fmt.Printf("目录 %s 创建成功\n", dirPath)
	} else if err != nil {
		fmt.Printf("获取目录状态失败: %v\n", err)
		return err
	} else {
		return nil
	}
	return nil
}

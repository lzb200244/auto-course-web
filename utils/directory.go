package utils

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	检查路径是否存在的工具函数
*/

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

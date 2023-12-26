package utils

import "os"

func CreateDir(dirName string) (err error) {
	return os.Mkdir(dirName, os.ModePerm)
}

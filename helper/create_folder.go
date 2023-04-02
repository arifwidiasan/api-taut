package helper

import (
	"os"
)

func CreateFolder(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

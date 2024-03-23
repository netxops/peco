package util

import (
	"os"
)

func IsNamedPipe(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return fileInfo.Mode()&os.ModeNamedPipe != 0
}

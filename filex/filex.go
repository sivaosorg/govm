package filex

import "os"

func IsFileExisted(filename string) bool {
	i, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !i.IsDir()
}

func IsFilesExisted(filename ...string) bool {
	for _, v := range filename {
		if !IsFileExisted(v) {
			return false
		}
	}
	return true
}

func IsDirectoryExisted(directory string) bool {
	i, err := os.Stat(directory)

	if os.IsNotExist(err) {
		return false
	}
	return i.IsDir()
}

func IsDirectoriesExisted(directories ...string) bool {
	for _, v := range directories {
		if !IsDirectoryExisted(v) {
			return false
		}
	}
	return true
}

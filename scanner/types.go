package scanner

import (
	"os"
	"path/filepath"
)

type File struct {
	os.FileInfo
	AbsolutePath string
}

func (File) BuildFromFileInfo(fileInfo os.FileInfo, dir string) File {
	file := File{
		FileInfo: fileInfo,
	}

	if dir != "" {
		file.AbsolutePath, _ = filepath.Abs(dir + "/" + fileInfo.Name())
	}

	return file
}

type Job struct {
	Id     int
	Work   string
	Result string
}

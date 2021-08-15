package api

import (
	"fmt"
	"github.com/dghubble/sling"
)

var s *sling.Sling

func init() {
	s = sling.New().Base("http://localhost:8080")
}

type FileRequest struct {
	AbsolutePath string `json:"absolute_path,omitempty"`
	Size         uint16 `json:"size,omitempty"`
	IsDir        bool   `json:"is_dir,omitempty"`
}



func CreateFile(abs string, size uint16, isDir bool) {
	file := FileRequest{
		AbsolutePath: abs,
		Size:         size,
		IsDir:        isDir,
	}

	receive, err := s.Post("/file/").BodyJSON(file).Receive(nil, nil)
	if err != nil {
		return 
	}
}

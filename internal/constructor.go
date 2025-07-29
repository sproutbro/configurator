package internal

import "github.com/sproutbro/filex/internal/parser"

type FileManager struct {
	*parser.Manager
}

func New() *FileManager {

	return &FileManager{
		Manager: parser.New(),
	}
}

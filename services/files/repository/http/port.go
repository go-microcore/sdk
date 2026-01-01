package adapter

import (
	"context"
)

type Interface interface {
	// Dirs
	CreateDir(ctx context.Context, authToken string, data CreateDirData) error
	DeleteDir(ctx context.Context, authToken string, data DeleteDirData) error
	RenameDir(ctx context.Context, authToken string, data RenameDirData) error
	// Files
	CreateFile(ctx context.Context, authToken string, data CreateFileData) error
	GetFiles(ctx context.Context, authToken string, data GetFilesData) ([]FileResult, error)
	DeleteFile(ctx context.Context, authToken string, data DeleteFileData) error
	RenameFile(ctx context.Context, authToken string, data RenameFileData) error
}

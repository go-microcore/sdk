package adapter

import (
	"context"
)

type Interface interface {
	// Dirs
	CreateDir(ctx context.Context, authToken string, path string) error
	RenameDir(ctx context.Context, authToken string, data RenameDirData) error
	DeleteDir(ctx context.Context, authToken string, path string) error
	// Files
	GetFile(ctx context.Context, authToken string, path string) ([]byte, error)
	StreamFile(ctx context.Context, token string) ([]byte, error)
	DownloadFile(ctx context.Context, authToken string, path string) (*DownloadFileResult, error)
	ListFiles(ctx context.Context, authToken string, path string) ([]FileResult, error)
	CreateFile(ctx context.Context, authToken string, data CreateFileData) error
	RenameFile(ctx context.Context, authToken string, data RenameFileData) error
	DeleteFile(ctx context.Context, authToken string, path string) error
}

package adapter

import (
	"mime/multipart"
)

// Data

type CreateDirData struct {
	Path string `json:"path"`
}

type DeleteDirData struct {
	Path string `json:"path"`
}

type RenameDirData struct {
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
}

type CreateFileData struct {
	Path string
	File *multipart.FileHeader
}

type GetFilesData struct {
	Path string `json:"path"`
}

type DeleteFileData struct {
	Path string `json:"path"`
}

type RenameFileData struct {
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
}

// Results

type FileResult struct {
	Name     string  `json:"name"`
	IsDir    bool    `json:"is_dir"`
	Size     *int64  `json:"size"`
	MimeType *string `json:"mime_type"`
}

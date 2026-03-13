package adapter

import "io"

// Data

type RenameDirData struct {
	OldPath string `json:"oldPath"`
	NewPath string `json:"newPath"`
}

type CreateFileData struct {
	Path string
	File io.Reader
	Name string
}

type RenameFileData struct {
	OldPath string `json:"oldPath"`
	NewPath string `json:"newPath"`
}

// Results

type FileResult struct {
	Name     string  `json:"name"`
	IsDir    bool    `json:"isDir"`
	Size     *int64  `json:"size"`
	MimeType *string `json:"mimeType"`
}

type DownloadFileResult struct {
	Token string `json:"token"`
}

package adapter

import "io"

// Data

type RenameDirData struct {
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
}

type CreateFileData struct {
	Path string
	File io.Reader
	Name string
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

type DownloadFileResult struct {
	Token string `json:"token"`
}

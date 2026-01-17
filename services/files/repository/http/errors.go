package adapter

import "go.microcore.dev/framework/errors"

var (
	// Dirs
	ErrDirInvalidPath    = errors.New(errors.ErrBadRequest, "invalid_path")
	ErrDirExist          = errors.New(errors.ErrBadRequest, "dir_exist")
	ErrDirNotFound       = errors.New(errors.ErrBadRequest, "dir_not_found")
	ErrDirInvalidOldPath = errors.New(errors.ErrBadRequest, "invalid_old_path")
	ErrDirInvalidNewPath = errors.New(errors.ErrBadRequest, "invalid_new_path")
	ErrDirOldNotFound    = errors.New(errors.ErrBadRequest, "old_dir_not_found")
	ErrDirNewExist       = errors.New(errors.ErrBadRequest, "new_dir_exist")
	// Files
	ErrFileExist        = errors.New(errors.ErrBadRequest, "file_exist")
	ErrFileNotFound     = errors.New(errors.ErrBadRequest, "file_not_found")
	ErrFileOldNotFound  = errors.New(errors.ErrBadRequest, "old_file_not_found")
	ErrFileNewExist     = errors.New(errors.ErrBadRequest, "new_file_exist")
	ErrFileInvalidToken = errors.New(errors.ErrBadRequest, "invalid_token")
)

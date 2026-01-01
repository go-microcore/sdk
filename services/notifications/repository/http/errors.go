package adapter

import "go.microcore.dev/framework/errors"

var (
	ErrFolderInvalidParent = errors.New(errors.ErrBadRequest, "invalid_parent")
	ErrFolderInvalidName   = errors.New(errors.ErrBadRequest, "invalid_name")
	ErrFolderExist         = errors.New(errors.ErrBadRequest, "folder_exist")
	ErrFolderNotFound      = errors.New(errors.ErrBadRequest, "folder_not_found")

	ErrEmailInvalidName      = errors.New(errors.ErrBadRequest, "invalid_name")
	ErrEmailInvalidFolderId  = errors.New(errors.ErrBadRequest, "invalid_folder_id")
	ErrEmailInvalidFromEmail = errors.New(errors.ErrBadRequest, "invalid_from_email")
	ErrEmailInvalidFromName  = errors.New(errors.ErrBadRequest, "invalid_from_name")
	ErrEmailInvalidSubject   = errors.New(errors.ErrBadRequest, "invalid_subject")
	ErrEmailInvalidToEmail   = errors.New(errors.ErrBadRequest, "invalid_to_email")
	ErrEmailInvalidHtml      = errors.New(errors.ErrBadRequest, "invalid_html")
	ErrEmailInvalidText      = errors.New(errors.ErrBadRequest, "invalid_text")
	ErrEmailNotFound         = errors.New(errors.ErrBadRequest, "email_not_found")
	ErrEmailExist            = errors.New(errors.ErrBadRequest, "email_exist")
)

package adapter

import (
	"context"
)

type Interface interface {
	// Emails
	SendCustomEmail(ctx context.Context, authToken string, data SendCustomEmailData) (*SendCustomEmailResult, error)
	SendEmail(ctx context.Context, authToken string, data SendEmailData) (*SendEmailResult, error)
	FilterEmails(ctx context.Context, authToken string, data FilterEmailsData) ([]FilterEmailsResult, error)
	FilterEmailLogs(ctx context.Context, authToken string, data FilterEmailLogsData) ([]FilterEmailLogsResult, error)
	UpdateEmail(ctx context.Context, authToken string, id uint, data UpdateEmailData) error
	DeleteEmail(ctx context.Context, authToken string, id uint) error
	CreateEmail(ctx context.Context, authToken string, data CreateEmailData) (*CreateEmailResult, error)
	// Folders
	FilterFolders(ctx context.Context, authToken string, data FilterEmailFoldersData) ([]FilterEmailFoldersResult, error)
	UpdateFolder(ctx context.Context, authToken string, id uint, data UpdateEmailFolderData) error
	DeleteFolder(ctx context.Context, authToken string, id uint) error
	CreateFolder(ctx context.Context, authToken string, data CreateEmailFolderData) (*CreateEmailFolderResult, error)
}

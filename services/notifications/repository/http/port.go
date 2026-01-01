package adapter

import (
	"context"
)

type Interface interface {
	AdminEmailCreate(ctx context.Context, authToken string, data CreateEmailData) (*CreateEmailResult, error)
	AdminEmailFilter(ctx context.Context, authToken string, data FilterEmailsData) ([]FilterEmailsResult, error)
	AdminEmailUpdate(ctx context.Context, authToken string, id uint, data UpdateEmailData) error
	AdminEmailDelete(ctx context.Context, authToken string, id uint) error
	AdminEmailSend(ctx context.Context, authToken string, data SendEmailData) (*SendEmailResult, error)
	AdminEmailSendCustom(ctx context.Context, authToken string, data SendCustomEmailData) (*SendCustomEmailResult, error)
	AdminEmailLogFilter(ctx context.Context, authToken string, data FilterEmailLogsData) ([]FilterEmailLogsResult, error)
	AdminEmailFolderCreate(ctx context.Context, authToken string, data CreateEmailFolderData) (*CreateEmailFolderResult, error)
	AdminEmailFolderFilter(ctx context.Context, authToken string, data FilterEmailFoldersData) ([]FilterEmailFoldersResult, error)
	AdminEmailFolderUpdate(ctx context.Context, authToken string, id uint, data UpdateEmailFolderData) error
	AdminEmailFolderDelete(ctx context.Context, authToken string, id uint) error
}

package adapter

import (
	"encoding/json"
	"time"
)

// Data

type SendCustomEmailData struct {
	Name      string `json:"name"`
	FromEmail string `json:"fromEmail"`
	FromName  string `json:"fromName"`
	Subject   string `json:"subject"`
	ToEmail   string `json:"toEmail"`
	Html      string `json:"html"`
	Text      string `json:"text"`
}

type SendEmailData struct {
	Name    string           `json:"name"`
	ToEmail string           `json:"toEmail"`
	Vars    *json.RawMessage `json:"vars"`
}

type FilterEmailsData struct {
	ID         *[]uint   `json:"id,omitempty"`
	Name       *[]string `json:"name,omitempty"`
	FolderID   *[]*uint  `json:"folderId,omitempty"`
	FromEmail  *[]string `json:"fromEmail,omitempty"`
	FromName   *[]string `json:"fromName,omitempty"`
	Subject    *[]string `json:"subject,omitempty"`
	SystemFlag *bool     `json:"systemFlag,omitempty"`
}

type FilterEmailLogsData struct {
	ID        *[]uint   `json:"id,omitempty"`
	Name      *[]string `json:"name,omitempty"`
	FromEmail *[]string `json:"fromEmail,omitempty"`
	FromName  *[]string `json:"fromName,omitempty"`
	ToEmail   *[]string `json:"toEmail,omitempty"`
	Status    *[]string `json:"status,omitempty"`
	MessageID *[]string `json:"messageId,omitempty"`
}

type UpdateEmailData struct {
	Name        *string `json:"name,omitempty"`
	FolderID    *uint   `json:"folderId,omitempty"`
	FromEmail   *string `json:"fromEmail,omitempty"`
	FromName    *string `json:"fromName,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Html        *string `json:"html,omitempty"`
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateEmailData struct {
	Name        string `json:"name"`
	FolderID    *uint  `json:"folderId"`
	FromEmail   string `json:"fromEmail"`
	FromName    string `json:"fromName"`
	Subject     string `json:"subject"`
	Html        string `json:"html"`
	Text        string `json:"text"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"systemFlag"`
}

type FilterEmailFoldersData struct {
	ID         *[]uint   `json:"id,omitempty"`
	ParentID   *[]*uint  `json:"parentId,omitempty"`
	Name       *[]string `json:"name,omitempty"`
	SystemFlag *bool     `json:"systemFlag,omitempty"`
}

type UpdateEmailFolderData struct {
	ParentID    *uint   `json:"parentId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateEmailFolderData struct {
	ParentID    *uint  `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"systemFlag"`
}

// Results

type SendCustomEmailResult struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	FromEmail string    `json:"fromEmail"`
	FromName  string    `json:"fromName"`
	Subject   string    `json:"subject"`
	ToEmail   string    `json:"toEmail"`
	Html      string    `json:"html"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	MessageID *string   `json:"messageId"`
	Errors    *string   `json:"errors"`
	Created   time.Time `json:"created"`
}

type SendEmailResult struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	FromEmail string    `json:"fromEmail"`
	FromName  string    `json:"fromName"`
	Subject   string    `json:"subject"`
	ToEmail   string    `json:"toEmail"`
	Html      string    `json:"html"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	MessageID *string   `json:"messageId"`
	Errors    *string   `json:"errors"`
	Created   time.Time `json:"created"`
}

type FilterEmailsResult struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	FolderID    *uint     `json:"folderId"`
	FromEmail   string    `json:"fromEmail"`
	FromName    string    `json:"fromName"`
	Subject     string    `json:"subject"`
	Html        string    `json:"html"`
	Text        string    `json:"text"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"systemFlag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type FilterEmailLogsResult struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	FromEmail string    `json:"fromEmail"`
	FromName  string    `json:"fromName"`
	Subject   string    `json:"subject"`
	ToEmail   string    `json:"toEmail"`
	Html      string    `json:"html"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	MessageID *string   `json:"messageId"`
	Errors    *string   `json:"errors"`
	Created   time.Time `json:"created"`
}

type CreateEmailResult struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	FolderID    *uint     `json:"folderId"`
	FromEmail   string    `json:"fromEmail"`
	FromName    string    `json:"fromName"`
	Subject     string    `json:"subject"`
	Html        string    `json:"html"`
	Text        string    `json:"text"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"systemFlag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type FilterEmailFoldersResult struct {
	ID          uint      `json:"id"`
	ParentID    *uint     `json:"parentId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"systemFlag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type CreateEmailFolderResult struct {
	ID          uint      `json:"id"`
	ParentID    *uint     `json:"parentId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"systemFlag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

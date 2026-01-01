package adapter

import (
	"encoding/json"
	"time"

	"go.microcore.dev/sdk/types"
)

// Data

type CreateEmailFolderData struct {
	ParentId    *uint  `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"system_flag"`
}

type FilterEmailFoldersData struct {
	Id         *[]uint   `json:"id"`
	ParentId   *[]*uint  `json:"parent_id"`
	Name       *[]string `json:"name"`
	SystemFlag *bool     `json:"system_flag"`
}

type UpdateEmailFolderData struct {
	ParentId    types.Nullable[uint]   `json:"parent_id,omitempty"`
	Name        types.Nullable[string] `json:"name,omitempty"`
	Description types.Nullable[string] `json:"description,omitempty"`
}

type CreateEmailData struct {
	Name        string `json:"name"`
	FolderId    *uint  `json:"folder_id"`
	FromEmail   string `json:"from_email"`
	FromName    string `json:"from_name"`
	Subject     string `json:"subject"`
	Html        string `json:"html"`
	Text        string `json:"text"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"system_flag"`
}

type FilterEmailsData struct {
	Id         *[]uint   `json:"id"`
	Name       *[]string `json:"name"`
	FolderId   *[]*uint  `json:"folder_id"`
	SystemFlag *bool     `json:"system_flag"`
}

type UpdateEmailData struct {
	Name        types.Nullable[string] `json:"name,omitempty"`
	FolderId    types.Nullable[uint]   `json:"folder_id,omitempty"`
	FromEmail   types.Nullable[string] `json:"from_email,omitempty"`
	FromName    types.Nullable[string] `json:"from_name,omitempty"`
	Subject     types.Nullable[string] `json:"subject,omitempty"`
	Html        types.Nullable[string] `json:"html,omitempty"`
	Text        types.Nullable[string] `json:"text,omitempty"`
	Description types.Nullable[string] `json:"description,omitempty"`
}

type SendCustomEmailData struct {
	Name      string `json:"name"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	Subject   string `json:"subject"`
	ToEmail   string `json:"to_email"`
	Html      string `json:"html"`
	Text      string `json:"text"`
}

type SendEmailData struct {
	Name    string           `json:"name"`
	ToEmail string           `json:"to_email"`
	Vars    *json.RawMessage `json:"vars"`
}

type FilterEmailLogsData struct {
	Id        *[]uint   `json:"id"`
	Name      *[]string `json:"name"`
	FromEmail *[]string `json:"from_email"`
	FromName  *[]string `json:"from_name"`
	ToEmail   *[]string `json:"to_email"`
	Status    *[]string `json:"status"`
	MessageId *[]string `json:"message_id"`
}

// Results

type CreateEmailFolderResult struct {
	Id          uint      `json:"id"`
	ParentId    *uint     `json:"parent_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type FilterEmailFoldersResult struct {
	Id          uint      `json:"id"`
	ParentId    *uint     `json:"parent_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type CreateEmailResult struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	FolderId    *uint     `json:"folder_id"`
	FromEmail   string    `json:"from_email"`
	FromName    string    `json:"from_name"`
	Subject     string    `json:"subject"`
	Html        string    `json:"html"`
	Text        string    `json:"text"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type FilterEmailsResult struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	FolderId    *uint     `json:"folder_id"`
	FromEmail   string    `json:"from_email"`
	FromName    string    `json:"from_name"`
	Subject     string    `json:"subject"`
	Html        string    `json:"html"`
	Text        string    `json:"text"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type SendCustomEmailResult struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	FromEmail string    `json:"from_email"`
	FromName  string    `json:"from_name"`
	Subject   string    `json:"subject"`
	ToEmail   string    `json:"to_email"`
	Html      string    `json:"html"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	MessageId *string   `json:"message_id"`
	Errors    *string   `json:"errors"`
	Created   time.Time `json:"created"`
}

type SendEmailResult struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	FromEmail string    `json:"from_email"`
	FromName  string    `json:"from_name"`
	Subject   string    `json:"subject"`
	ToEmail   string    `json:"to_email"`
	Html      string    `json:"html"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	MessageId *string   `json:"message_id"`
	Errors    *string   `json:"errors"`
	Created   time.Time `json:"created"`
}

type FilterEmailLogsResult struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	FromEmail string    `json:"from_email"`
	FromName  string    `json:"from_name"`
	Subject   string    `json:"subject"`
	ToEmail   string    `json:"to_email"`
	Html      string    `json:"html"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	MessageId *string   `json:"message_id"`
	Errors    *string   `json:"errors"`
	Created   time.Time `json:"created"`
}

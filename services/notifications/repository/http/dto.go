package adapter

import (
	"encoding/json"
	"time"
)

// Data

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

type FilterEmailsData struct {
	Id         *[]uint   `json:"id,omitempty"`
	Name       *[]string `json:"name,omitempty"`
	FolderId   *[]*uint  `json:"folder_id,omitempty"`
	FromEmail  *[]string `json:"from_email,omitempty"`
	FromName   *[]string `json:"from_name,omitempty"`
	Subject    *[]string `json:"subject,omitempty"`
	SystemFlag *bool     `json:"system_flag,omitempty"`
}

type FilterEmailLogsData struct {
	Id        *[]uint   `json:"id,omitempty"`
	Name      *[]string `json:"name,omitempty"`
	FromEmail *[]string `json:"from_email,omitempty"`
	FromName  *[]string `json:"from_name,omitempty"`
	ToEmail   *[]string `json:"to_email,omitempty"`
	Status    *[]string `json:"status,omitempty"`
	MessageId *[]string `json:"message_id,omitempty"`
}

type UpdateEmailData struct {
	Name        *string `json:"name,omitempty"`
	FolderId    *uint   `json:"folder_id,omitempty"`
	FromEmail   *string `json:"from_email,omitempty"`
	FromName    *string `json:"from_name,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Html        *string `json:"html,omitempty"`
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
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

type FilterEmailFoldersData struct {
	Id         *[]uint   `json:"id,omitempty"`
	ParentId   *[]*uint  `json:"parent_id,omitempty"`
	Name       *[]string `json:"name,omitempty"`
	SystemFlag *bool     `json:"system_flag,omitempty"`
}

type UpdateEmailFolderData struct {
	ParentId    *uint   `json:"parent_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateEmailFolderData struct {
	ParentId    *uint  `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SystemFlag  bool   `json:"system_flag"`
}

// Results

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

type FilterEmailFoldersResult struct {
	Id          uint      `json:"id"`
	ParentId    *uint     `json:"parent_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

type CreateEmailFolderResult struct {
	Id          uint      `json:"id"`
	ParentId    *uint     `json:"parent_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemFlag  bool      `json:"system_flag"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

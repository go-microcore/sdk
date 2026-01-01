package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
)

type Config struct {
	HttpClientManager    client.Manager
	FilesServiceEndpoint string
}

func New(config *Config) Interface {
	return &adapter{
		config.HttpClientManager,
		config.FilesServiceEndpoint,
	}
}

type adapter struct {
	httpClientManager    client.Manager
	filesServiceEndpoint string
}

// Dirs

func (a *adapter) CreateDir(ctx context.Context, authToken string, data CreateDirData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/dir/",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_path": ErrDirInvalidPath,
		"bad_request:dir_exist":    ErrDirExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteDir(ctx context.Context, authToken string, data DeleteDirData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/dir/",
		client.WithRequestMethod(http.MethodDelete),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_path":  ErrDirInvalidPath,
		"bad_request:dir_not_found": ErrDirNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) RenameDir(ctx context.Context, authToken string, data RenameDirData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/dir/",
		client.WithRequestMethod(http.MethodPatch),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_old_path":  ErrDirInvalidOldPath,
		"bad_request:invalid_new_path":  ErrDirInvalidNewPath,
		"bad_request:old_dir_not_found": ErrDirOldNotFound,
		"bad_request:new_dir_exist":     ErrDirNewExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Files

func (a *adapter) CreateFile(ctx context.Context, authToken string, data CreateFileData) error {
	// Create buffer and multipart writer
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Encode meta json
	meta, err := json.Marshal(struct {
		Path string `json:"path"`
	}{
		Path: data.Path,
	})
	if err != nil {
		return fmt.Errorf("error parsing request meta: %v", err)
	}

	// Add field meta as string
	if err := writer.WriteField("meta", string(meta)); err != nil {
		return fmt.Errorf("write meta field: %v", err)
	}

	// Open file
	file, err := data.File.Open()
	if err != nil {
		return fmt.Errorf("open file: %v", err)
	}
	defer file.Close()

	// Create form file
	part, err := writer.CreateFormFile("file", data.File.Filename)
	if err != nil {
		return fmt.Errorf("create form file: %v", err)
	}

	// Copy file
	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("copy file: %v", err)
	}

	// Close writer
	if err := writer.Close(); err != nil {
		return fmt.Errorf("close writer: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body.Bytes()),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
			client.NewRequestHeader("Content-Type", writer.FormDataContentType()),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:dir_not_found": ErrDirNotFound,
		"bad_request:file_exist":    ErrFileExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) GetFiles(ctx context.Context, authToken string, data GetFilesData) ([]FileResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/list",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FileResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_path": ErrDirInvalidPath,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteFile(ctx context.Context, authToken string, data DeleteFileData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/",
		client.WithRequestMethod(http.MethodDelete),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_path":   ErrDirInvalidPath,
		"bad_request:file_not_found": ErrFileNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) RenameFile(ctx context.Context, authToken string, data RenameFileData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.filesServiceEndpoint+"/files/admin/",
		client.WithRequestMethod(http.MethodPatch),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_old_path":   ErrDirInvalidOldPath,
		"bad_request:invalid_new_path":   ErrDirInvalidNewPath,
		"bad_request:old_file_not_found": ErrFileOldNotFound,
		"bad_request:new_file_exist":     ErrFileNewExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

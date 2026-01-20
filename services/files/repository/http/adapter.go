package adapter

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

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

func (a *adapter) CreateDir(ctx context.Context, authToken string, path string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/dir/")
	url.WriteString(base64.RawURLEncoding.EncodeToString([]byte(path)))

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
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

func (a *adapter) RenameDir(ctx context.Context, authToken string, data RenameDirData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/dir/")

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
	if res.StatusCode() == 204 {
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

func (a *adapter) DeleteDir(ctx context.Context, authToken string, path string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/dir/")
	url.WriteString(path)

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodDelete),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
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

// Files

func (a *adapter) GetFile(ctx context.Context, authToken string, path string) ([]byte, error) {
	download, err := a.DownloadFile(ctx, authToken, path)
	if err != nil {
		return nil, err
	}

	stream, err := a.StreamFile(ctx, download.Token)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (a *adapter) StreamFile(ctx context.Context, token string) ([]byte, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/download/stream/")
	url.WriteString(token)

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodGet),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		return res.Body(), nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_token": ErrFileInvalidToken,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DownloadFile(ctx context.Context, authToken string, path string) (*DownloadFileResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/download/")
	url.WriteString(base64.RawURLEncoding.EncodeToString([]byte(path)))

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodGet),
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
		var response DownloadFileResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
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
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) ListFiles(ctx context.Context, authToken string, path string) ([]FileResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/list/")
	url.WriteString(base64.RawURLEncoding.EncodeToString([]byte(path)))

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodGet),
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

func (a *adapter) CreateFile(ctx context.Context, authToken string, data CreateFileData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/")
	url.WriteString(base64.RawURLEncoding.EncodeToString([]byte(data.Path)))

	// Multipart body
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Create form file
	part, err := writer.CreateFormFile("file", data.Name)
	if err != nil {
		return fmt.Errorf("create form file: %w", err)
	}

	// Copy file content
	if _, err := io.Copy(part, data.File); err != nil {
		return fmt.Errorf("copy file: %w", err)
	}

	// Close multipart writer
	if err := writer.Close(); err != nil {
		return fmt.Errorf("close writer: %w", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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

func (a *adapter) RenameFile(ctx context.Context, authToken string, data RenameFileData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/")

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
	if res.StatusCode() == 204 {
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

func (a *adapter) DeleteFile(ctx context.Context, authToken string, path string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.filesServiceEndpoint)
	url.WriteString("/files/")
	url.WriteString(base64.RawURLEncoding.EncodeToString([]byte(path)))

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodDelete),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.filesServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
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

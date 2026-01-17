package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
)

type Config struct {
	HttpClientManager            client.Manager
	NotificationsServiceEndpoint string
}

func New(config *Config) Interface {
	return &adapter{
		config.HttpClientManager,
		config.NotificationsServiceEndpoint,
	}
}

type adapter struct {
	httpClientManager            client.Manager
	notificationsServiceEndpoint string
}

// Emails

func (a *adapter) SendCustomEmail(ctx context.Context, authToken string, data SendCustomEmailData) (*SendCustomEmailResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/send/custom")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response SendCustomEmailResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":       ErrEmailInvalidName,
		"bad_request:invalid_from_email": ErrEmailInvalidFromEmail,
		"bad_request:invalid_from_name":  ErrEmailInvalidFromName,
		"bad_request:invalid_subject":    ErrEmailInvalidSubject,
		"bad_request:invalid_to_email":   ErrEmailInvalidToEmail,
		"bad_request:invalid_html":       ErrEmailInvalidHtml,
		"bad_request:invalid_text":       ErrEmailInvalidText,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) SendEmail(ctx context.Context, authToken string, data SendEmailData) (*SendEmailResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/send/")

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response SendEmailResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":     ErrEmailInvalidName,
		"bad_request:invalid_to_email": ErrEmailInvalidToEmail,
		"bad_request:email_not_found":  ErrEmailNotFound,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) FilterEmails(ctx context.Context, authToken string, data FilterEmailsData) ([]FilterEmailsResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/filter")

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FilterEmailsResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) FilterEmailLogs(ctx context.Context, authToken string, data FilterEmailLogsData) ([]FilterEmailLogsResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/log")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FilterEmailLogsResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) UpdateEmail(ctx context.Context, authToken string, id uint, data UpdateEmailData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/")
	url.WriteString(strconv.FormatUint(uint64(id), 10))

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
		return fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":       ErrEmailInvalidName,
		"bad_request:invalid_folder_id":  ErrEmailInvalidFolderId,
		"bad_request:invalid_from_email": ErrEmailInvalidFromEmail,
		"bad_request:invalid_from_name":  ErrEmailInvalidFromName,
		"bad_request:invalid_subject":    ErrEmailInvalidSubject,
		"bad_request:invalid_html":       ErrEmailInvalidHtml,
		"bad_request:invalid_text":       ErrEmailInvalidText,
		"bad_request:email_not_found":    ErrEmailNotFound,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) DeleteEmail(ctx context.Context, authToken string, id uint) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/")
	url.WriteString(strconv.FormatUint(uint64(id), 10))

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
		return fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:email_not_found": ErrEmailNotFound,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) CreateEmail(ctx context.Context, authToken string, data CreateEmailData) (*CreateEmailResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/emails/")

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response CreateEmailResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":       ErrEmailInvalidName,
		"bad_request:invalid_folder_id":  ErrEmailInvalidFolderId,
		"bad_request:invalid_from_email": ErrEmailInvalidFromEmail,
		"bad_request:invalid_from_name":  ErrEmailInvalidFromName,
		"bad_request:invalid_subject":    ErrEmailInvalidSubject,
		"bad_request:invalid_html":       ErrEmailInvalidHtml,
		"bad_request:invalid_text":       ErrEmailInvalidText,
		"bad_request:email_exist":        ErrEmailExist,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

// Folders

func (a *adapter) FilterFolders(ctx context.Context, authToken string, data FilterEmailFoldersData) ([]FilterEmailFoldersResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/folders/filter")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FilterEmailFoldersResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) UpdateFolder(ctx context.Context, authToken string, id uint, data UpdateEmailFolderData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/folders/")
	url.WriteString(strconv.FormatUint(uint64(id), 10))

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
		return fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":     ErrFolderInvalidName,
		"bad_request:folder_not_found": ErrFolderNotFound,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) DeleteFolder(ctx context.Context, authToken string, id uint) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/folders/")
	url.WriteString(strconv.FormatUint(uint64(id), 10))

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
		return fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:folder_not_found": ErrFolderNotFound,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) CreateFolder(ctx context.Context, authToken string, data CreateEmailFolderData) (*CreateEmailFolderResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.notificationsServiceEndpoint)
	url.WriteString("/notifications/folders/")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.notificationsServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response CreateEmailFolderResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_parent": ErrFolderInvalidParent,
		"bad_request:invalid_name":   ErrFolderInvalidName,
		"bad_request:folder_exist":   ErrFolderExist,
	}

	// Parse errors
	if err, ok := errMap[string(res.Body())]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

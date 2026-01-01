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

func (a *adapter) AdminEmailCreate(ctx context.Context, authToken string, data CreateEmailData) (*CreateEmailResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/",
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

	// Response message
	errMessage := string(res.Body())

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
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailFilter(ctx context.Context, authToken string, data FilterEmailsData) ([]FilterEmailsResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/filter",
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

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), string(res.Body()))
}

func (a *adapter) AdminEmailUpdate(ctx context.Context, authToken string, id uint, data UpdateEmailData) error {
	// Build path
	var path strings.Builder
	path.WriteString("/notifications/admin/email/")
	path.WriteString(strconv.FormatUint(uint64(id), 10))

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+path.String(),
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

	// Response message
	errMessage := string(res.Body())

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
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailDelete(ctx context.Context, authToken string, id uint) error {
	// Build path
	var path strings.Builder
	path.WriteString("/notifications/admin/email/")
	path.WriteString(strconv.FormatUint(uint64(id), 10))

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+path.String(),
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

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:email_not_found": ErrEmailNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailSend(ctx context.Context, authToken string, data SendEmailData) (*SendEmailResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/send/",
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

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":     ErrEmailInvalidName,
		"bad_request:invalid_to_email": ErrEmailInvalidToEmail,
		"bad_request:email_not_found":  ErrEmailNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailSendCustom(ctx context.Context, authToken string, data SendCustomEmailData) (*SendCustomEmailResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/send/custom",
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

	// Response message
	errMessage := string(res.Body())

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
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailLogFilter(ctx context.Context, authToken string, data FilterEmailLogsData) ([]FilterEmailLogsResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/log/filter",
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

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), string(res.Body()))
}

func (a *adapter) AdminEmailFolderCreate(ctx context.Context, authToken string, data CreateEmailFolderData) (*CreateEmailFolderResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/folder/",
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

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_parent": ErrFolderInvalidParent,
		"bad_request:invalid_name":   ErrFolderInvalidName,
		"bad_request:folder_exist":   ErrFolderExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailFolderFilter(ctx context.Context, authToken string, data FilterEmailFoldersData) ([]FilterEmailFoldersResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+"/notifications/admin/email/folder/filter",
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

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), string(res.Body()))
}

func (a *adapter) AdminEmailFolderUpdate(ctx context.Context, authToken string, id uint, data UpdateEmailFolderData) error {
	// Build path
	var path strings.Builder
	path.WriteString("/admin/notifications/emails/folders/")
	path.WriteString(strconv.FormatUint(uint64(id), 10))

	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+path.String(),
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

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":     ErrFolderInvalidName,
		"bad_request:folder_not_found": ErrFolderNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) AdminEmailFolderDelete(ctx context.Context, authToken string, id uint) error {
	// Build path
	var path strings.Builder
	path.WriteString("/notifications/admin/email/folder/")
	path.WriteString(strconv.FormatUint(uint64(id), 10))

	// Send service request
	res, err := a.httpClientManager.Request(
		a.notificationsServiceEndpoint+path.String(),
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

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:folder_not_found": ErrFolderNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// WorkspaceAPIService WorkspaceAPI service
type WorkspaceAPIService service

type ApiCreateWorkspaceRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	workspace *Workspace
}

func (r ApiCreateWorkspaceRequest) XTenant(xTenant string) ApiCreateWorkspaceRequest {
	r.xTenant = &xTenant
	return r
}

// The workspace to be created
func (r ApiCreateWorkspaceRequest) Workspace(workspace Workspace) ApiCreateWorkspaceRequest {
	r.workspace = &workspace
	return r
}

func (r ApiCreateWorkspaceRequest) Execute() (*Workspace, *http.Response, error) {
	return r.ApiService.CreateWorkspaceExecute(r)
}

/*
CreateWorkspace Create workspace

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWorkspaceRequest
*/
func (a *WorkspaceAPIService) CreateWorkspace(ctx context.Context) ApiCreateWorkspaceRequest {
	return ApiCreateWorkspaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Workspace
func (a *WorkspaceAPIService) CreateWorkspaceExecute(r ApiCreateWorkspaceRequest) (*Workspace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Workspace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.CreateWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}
	if r.workspace == nil {
		return localVarReturnValue, nil, reportError("workspace is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.workspace+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.workspace+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	// body params
	localVarPostBody = r.workspace
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteWorkspaceByIdRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	workspaceId string
}

func (r ApiDeleteWorkspaceByIdRequest) XTenant(xTenant string) ApiDeleteWorkspaceByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiDeleteWorkspaceByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWorkspaceByIdExecute(r)
}

/*
DeleteWorkspaceById Delete a workspace by an id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId Id of the workspace
 @return ApiDeleteWorkspaceByIdRequest
*/
func (a *WorkspaceAPIService) DeleteWorkspaceById(ctx context.Context, workspaceId string) ApiDeleteWorkspaceByIdRequest {
	return ApiDeleteWorkspaceByIdRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) DeleteWorkspaceByIdExecute(r ApiDeleteWorkspaceByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.DeleteWorkspaceById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.error+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFindWorkspaceByIdRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	workspaceId string
}

func (r ApiFindWorkspaceByIdRequest) XTenant(xTenant string) ApiFindWorkspaceByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindWorkspaceByIdRequest) Execute() (*Workspace, *http.Response, error) {
	return r.ApiService.FindWorkspaceByIdExecute(r)
}

/*
FindWorkspaceById Find workspace by Id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId Id of the workspace
 @return ApiFindWorkspaceByIdRequest
*/
func (a *WorkspaceAPIService) FindWorkspaceById(ctx context.Context, workspaceId string) ApiFindWorkspaceByIdRequest {
	return ApiFindWorkspaceByIdRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return Workspace
func (a *WorkspaceAPIService) FindWorkspaceByIdExecute(r ApiFindWorkspaceByIdRequest) (*Workspace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Workspace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.FindWorkspaceById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.workspace+json", "application/vnd.graal.systems.v1.error+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFindWorkspacesRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	page *int32
	size *int32
}

func (r ApiFindWorkspacesRequest) XTenant(xTenant string) ApiFindWorkspacesRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindWorkspacesRequest) Page(page int32) ApiFindWorkspacesRequest {
	r.page = &page
	return r
}

func (r ApiFindWorkspacesRequest) Size(size int32) ApiFindWorkspacesRequest {
	r.size = &size
	return r
}

func (r ApiFindWorkspacesRequest) Execute() (*WorkspacePage, *http.Response, error) {
	return r.ApiService.FindWorkspacesExecute(r)
}

/*
FindWorkspaces Retrieve all workspaces

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindWorkspacesRequest
*/
func (a *WorkspaceAPIService) FindWorkspaces(ctx context.Context) ApiFindWorkspacesRequest {
	return ApiFindWorkspacesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WorkspacePage
func (a *WorkspaceAPIService) FindWorkspacesExecute(r ApiFindWorkspacesRequest) (*WorkspacePage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkspacePage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.FindWorkspaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "")
	} else {
		var defaultValue int32 = 200
		r.size = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.workspace+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetLogsForWorkspaceRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	workspaceId string
	cursor *string
}

func (r ApiGetLogsForWorkspaceRequest) XTenant(xTenant string) ApiGetLogsForWorkspaceRequest {
	r.xTenant = &xTenant
	return r
}

// The cursor
func (r ApiGetLogsForWorkspaceRequest) Cursor(cursor string) ApiGetLogsForWorkspaceRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetLogsForWorkspaceRequest) Execute() ([]LogEntry, *http.Response, error) {
	return r.ApiService.GetLogsForWorkspaceExecute(r)
}

/*
GetLogsForWorkspace Get logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId Id of the workspace
 @return ApiGetLogsForWorkspaceRequest
*/
func (a *WorkspaceAPIService) GetLogsForWorkspace(ctx context.Context, workspaceId string) ApiGetLogsForWorkspaceRequest {
	return ApiGetLogsForWorkspaceRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return []LogEntry
func (a *WorkspaceAPIService) GetLogsForWorkspaceExecute(r ApiGetLogsForWorkspaceRequest) ([]LogEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LogEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.GetLogsForWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspaceId}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.log+json", "application/vnd.graal.systems.v1.error+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateWorkspaceRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	workspaceId string
	patch *[]Patch
}

func (r ApiUpdateWorkspaceRequest) XTenant(xTenant string) ApiUpdateWorkspaceRequest {
	r.xTenant = &xTenant
	return r
}

// The patch
func (r ApiUpdateWorkspaceRequest) Patch(patch []Patch) ApiUpdateWorkspaceRequest {
	r.patch = &patch
	return r
}

func (r ApiUpdateWorkspaceRequest) Execute() (*Workspace, *http.Response, error) {
	return r.ApiService.UpdateWorkspaceExecute(r)
}

/*
UpdateWorkspace Update a workspace

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId Id of the workspace
 @return ApiUpdateWorkspaceRequest
*/
func (a *WorkspaceAPIService) UpdateWorkspace(ctx context.Context, workspaceId string) ApiUpdateWorkspaceRequest {
	return ApiUpdateWorkspaceRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return Workspace
func (a *WorkspaceAPIService) UpdateWorkspaceExecute(r ApiUpdateWorkspaceRequest) (*Workspace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Workspace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.UpdateWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}
	if r.patch == nil {
		return localVarReturnValue, nil, reportError("patch is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json;charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.workspace+json", "application/vnd.graal.systems.v1.error+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	// body params
	localVarPostBody = r.patch
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUploadFileInWorkspaceRequest struct {
	ctx context.Context
	ApiService *WorkspaceAPIService
	xTenant *string
	workspaceId string
	path *string
	filename []*os.File
}

func (r ApiUploadFileInWorkspaceRequest) XTenant(xTenant string) ApiUploadFileInWorkspaceRequest {
	r.xTenant = &xTenant
	return r
}

// path
func (r ApiUploadFileInWorkspaceRequest) Path(path string) ApiUploadFileInWorkspaceRequest {
	r.path = &path
	return r
}

func (r ApiUploadFileInWorkspaceRequest) Filename(filename []*os.File) ApiUploadFileInWorkspaceRequest {
	r.filename = filename
	return r
}

func (r ApiUploadFileInWorkspaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadFileInWorkspaceExecute(r)
}

/*
UploadFileInWorkspace Upload files in a workspace

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspaceId Id of the bucket
 @return ApiUploadFileInWorkspaceRequest
*/
func (a *WorkspaceAPIService) UploadFileInWorkspace(ctx context.Context, workspaceId string) ApiUploadFileInWorkspaceRequest {
	return ApiUploadFileInWorkspaceRequest{
		ApiService: a,
		ctx: ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) UploadFileInWorkspaceExecute(r ApiUploadFileInWorkspaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.UploadFileInWorkspace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspaces/{workspaceId}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}

	if r.path != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	var filenameLocalVarFormFileName string
	var filenameLocalVarFileName     string
	var filenameLocalVarFileBytes    []byte

	filenameLocalVarFormFileName = "filename"
	filenameLocalVarFile := r.filename

	if filenameLocalVarFile != nil {
		// loop through the array to prepare multiple files upload
		for _, filenameLocalVarFileValue := range filenameLocalVarFile {
			fbs, _ := io.ReadAll(filenameLocalVarFileValue)

			filenameLocalVarFileBytes = fbs
			filenameLocalVarFileName = filenameLocalVarFileValue.Name()
			filenameLocalVarFileValue.Close()
			formFiles = append(formFiles, formFile{fileBytes: filenameLocalVarFileBytes, fileName: filenameLocalVarFileName, formFileName: filenameLocalVarFormFileName})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

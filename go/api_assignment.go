/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AssignmentApiService AssignmentApi service
type AssignmentApiService service

type ApiCreateRoleAssignmentRequest struct {
	ctx _context.Context
	ApiService *AssignmentApiService
	xTenant *string
	roleAssignment *RoleAssignment
}

func (r ApiCreateRoleAssignmentRequest) XTenant(xTenant string) ApiCreateRoleAssignmentRequest {
	r.xTenant = &xTenant
	return r
}
func (r ApiCreateRoleAssignmentRequest) RoleAssignment(roleAssignment RoleAssignment) ApiCreateRoleAssignmentRequest {
	r.roleAssignment = &roleAssignment
	return r
}

func (r ApiCreateRoleAssignmentRequest) Execute() (RoleAssignment, *_nethttp.Response, error) {
	return r.ApiService.CreateRoleAssignmentExecute(r)
}

/*
 * CreateRoleAssignment Create a assignment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateRoleAssignmentRequest
 */
func (a *AssignmentApiService) CreateRoleAssignment(ctx _context.Context) ApiCreateRoleAssignmentRequest {
	return ApiCreateRoleAssignmentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RoleAssignment
 */
func (a *AssignmentApiService) CreateRoleAssignmentExecute(r ApiCreateRoleAssignmentRequest) (RoleAssignment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RoleAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentApiService.CreateRoleAssignment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assignments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}
	if r.roleAssignment == nil {
		return localVarReturnValue, nil, reportError("roleAssignment is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.assignment+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.assignment+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(*r.xTenant, "")
	// body params
	localVarPostBody = r.roleAssignment
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteRoleAssignmentByIdRequest struct {
	ctx _context.Context
	ApiService *AssignmentApiService
	xTenant *string
	assignmentId string
}

func (r ApiDeleteRoleAssignmentByIdRequest) XTenant(xTenant string) ApiDeleteRoleAssignmentByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiDeleteRoleAssignmentByIdRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteRoleAssignmentByIdExecute(r)
}

/*
 * DeleteRoleAssignmentById Delete an assignment by an id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param assignmentId Id of the assignment
 * @return ApiDeleteRoleAssignmentByIdRequest
 */
func (a *AssignmentApiService) DeleteRoleAssignmentById(ctx _context.Context, assignmentId string) ApiDeleteRoleAssignmentByIdRequest {
	return ApiDeleteRoleAssignmentByIdRequest{
		ApiService: a,
		ctx: ctx,
		assignmentId: assignmentId,
	}
}

/*
 * Execute executes the request
 */
func (a *AssignmentApiService) DeleteRoleAssignmentByIdExecute(r ApiDeleteRoleAssignmentByIdRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentApiService.DeleteRoleAssignmentById")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assignments/{assignmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"assignmentId"+"}", _neturl.PathEscape(parameterToString(r.assignmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	localVarHeaderParams["X-Tenant"] = parameterToString(*r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFindRoleAssignmentByIdRequest struct {
	ctx _context.Context
	ApiService *AssignmentApiService
	xTenant *string
	assignmentId string
}

func (r ApiFindRoleAssignmentByIdRequest) XTenant(xTenant string) ApiFindRoleAssignmentByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindRoleAssignmentByIdRequest) Execute() (RoleAssignment, *_nethttp.Response, error) {
	return r.ApiService.FindRoleAssignmentByIdExecute(r)
}

/*
 * FindRoleAssignmentById Find assignment by Id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param assignmentId Id of the assignment
 * @return ApiFindRoleAssignmentByIdRequest
 */
func (a *AssignmentApiService) FindRoleAssignmentById(ctx _context.Context, assignmentId string) ApiFindRoleAssignmentByIdRequest {
	return ApiFindRoleAssignmentByIdRequest{
		ApiService: a,
		ctx: ctx,
		assignmentId: assignmentId,
	}
}

/*
 * Execute executes the request
 * @return RoleAssignment
 */
func (a *AssignmentApiService) FindRoleAssignmentByIdExecute(r ApiFindRoleAssignmentByIdRequest) (RoleAssignment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RoleAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentApiService.FindRoleAssignmentById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assignments/{assignmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"assignmentId"+"}", _neturl.PathEscape(parameterToString(r.assignmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.assignment+json", "application/vnd.graal.systems.v1.error+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(*r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFindRoleAssignmentsRequest struct {
	ctx _context.Context
	ApiService *AssignmentApiService
	xTenant *string
	resourceType *string
	resourceId *string
}

func (r ApiFindRoleAssignmentsRequest) XTenant(xTenant string) ApiFindRoleAssignmentsRequest {
	r.xTenant = &xTenant
	return r
}
func (r ApiFindRoleAssignmentsRequest) ResourceType(resourceType string) ApiFindRoleAssignmentsRequest {
	r.resourceType = &resourceType
	return r
}
func (r ApiFindRoleAssignmentsRequest) ResourceId(resourceId string) ApiFindRoleAssignmentsRequest {
	r.resourceId = &resourceId
	return r
}

func (r ApiFindRoleAssignmentsRequest) Execute() ([]RoleAndPrincipalAndAssignment, *_nethttp.Response, error) {
	return r.ApiService.FindRoleAssignmentsExecute(r)
}

/*
 * FindRoleAssignments Retrieve all assignments for a resource
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFindRoleAssignmentsRequest
 */
func (a *AssignmentApiService) FindRoleAssignments(ctx _context.Context) ApiFindRoleAssignmentsRequest {
	return ApiFindRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []RoleAndPrincipalAndAssignment
 */
func (a *AssignmentApiService) FindRoleAssignmentsExecute(r ApiFindRoleAssignmentsRequest) ([]RoleAndPrincipalAndAssignment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RoleAndPrincipalAndAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentApiService.FindRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assignments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}

	if r.resourceType != nil {
		localVarQueryParams.Add("resource_type", parameterToString(*r.resourceType, ""))
	}
	if r.resourceId != nil {
		localVarQueryParams.Add("resource_id", parameterToString(*r.resourceId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.roleprincipalassignment+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(*r.xTenant, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
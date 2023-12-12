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


// EndpointAPIService EndpointAPI service
type EndpointAPIService service

type ApiCallProxyInferRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpointId string
	body *map[string]interface{}
}

func (r ApiCallProxyInferRequest) XTenant(xTenant string) ApiCallProxyInferRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiCallProxyInferRequest) Body(body map[string]interface{}) ApiCallProxyInferRequest {
	r.body = &body
	return r
}

func (r ApiCallProxyInferRequest) Execute() (*http.Response, error) {
	return r.ApiService.CallProxyInferExecute(r)
}

/*
CallProxyInfer Call proxy infer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @param endpointId Id of the endpoint
 @return ApiCallProxyInferRequest
*/
func (a *EndpointAPIService) CallProxyInfer(ctx context.Context, projectId string, endpointId string) ApiCallProxyInferRequest {
	return ApiCallProxyInferRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		endpointId: endpointId,
	}
}

// Execute executes the request
func (a *EndpointAPIService) CallProxyInferExecute(r ApiCallProxyInferRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.CallProxyInfer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints/{endpointId}/proxy/infer"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type ApiCallProxyMetadataRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpointId string
}

func (r ApiCallProxyMetadataRequest) XTenant(xTenant string) ApiCallProxyMetadataRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiCallProxyMetadataRequest) Execute() (*http.Response, error) {
	return r.ApiService.CallProxyMetadataExecute(r)
}

/*
CallProxyMetadata Call proxy metadata

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @param endpointId Id of the endpoint
 @return ApiCallProxyMetadataRequest
*/
func (a *EndpointAPIService) CallProxyMetadata(ctx context.Context, projectId string, endpointId string) ApiCallProxyMetadataRequest {
	return ApiCallProxyMetadataRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		endpointId: endpointId,
	}
}

// Execute executes the request
func (a *EndpointAPIService) CallProxyMetadataExecute(r ApiCallProxyMetadataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.CallProxyMetadata")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints/{endpointId}/proxy/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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

type ApiCreateEndpointForProjectRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpoint *Endpoint
}

func (r ApiCreateEndpointForProjectRequest) XTenant(xTenant string) ApiCreateEndpointForProjectRequest {
	r.xTenant = &xTenant
	return r
}

// The endpoint to be created
func (r ApiCreateEndpointForProjectRequest) Endpoint(endpoint Endpoint) ApiCreateEndpointForProjectRequest {
	r.endpoint = &endpoint
	return r
}

func (r ApiCreateEndpointForProjectRequest) Execute() (*Endpoint, *http.Response, error) {
	return r.ApiService.CreateEndpointForProjectExecute(r)
}

/*
CreateEndpointForProject Create endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @return ApiCreateEndpointForProjectRequest
*/
func (a *EndpointAPIService) CreateEndpointForProject(ctx context.Context, projectId string) ApiCreateEndpointForProjectRequest {
	return ApiCreateEndpointForProjectRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Endpoint
func (a *EndpointAPIService) CreateEndpointForProjectExecute(r ApiCreateEndpointForProjectRequest) (*Endpoint, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Endpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.CreateEndpointForProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}
	if r.endpoint == nil {
		return localVarReturnValue, nil, reportError("endpoint is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.endpoint+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.endpoint+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	// body params
	localVarPostBody = r.endpoint
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

type ApiDeleteEndpointByIdAndProjectIdRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpointId string
}

func (r ApiDeleteEndpointByIdAndProjectIdRequest) XTenant(xTenant string) ApiDeleteEndpointByIdAndProjectIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiDeleteEndpointByIdAndProjectIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEndpointByIdAndProjectIdExecute(r)
}

/*
DeleteEndpointByIdAndProjectId Delete an endpoint by an id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @param endpointId Id of the endpoint
 @return ApiDeleteEndpointByIdAndProjectIdRequest
*/
func (a *EndpointAPIService) DeleteEndpointByIdAndProjectId(ctx context.Context, projectId string, endpointId string) ApiDeleteEndpointByIdAndProjectIdRequest {
	return ApiDeleteEndpointByIdAndProjectIdRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		endpointId: endpointId,
	}
}

// Execute executes the request
func (a *EndpointAPIService) DeleteEndpointByIdAndProjectIdExecute(r ApiDeleteEndpointByIdAndProjectIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.DeleteEndpointByIdAndProjectId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints/{endpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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

type ApiFindEndpointByIdAndProjectIdRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpointId string
}

func (r ApiFindEndpointByIdAndProjectIdRequest) XTenant(xTenant string) ApiFindEndpointByIdAndProjectIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindEndpointByIdAndProjectIdRequest) Execute() (*Endpoint, *http.Response, error) {
	return r.ApiService.FindEndpointByIdAndProjectIdExecute(r)
}

/*
FindEndpointByIdAndProjectId Find endpoint by Id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @param endpointId Id of the endpoint
 @return ApiFindEndpointByIdAndProjectIdRequest
*/
func (a *EndpointAPIService) FindEndpointByIdAndProjectId(ctx context.Context, projectId string, endpointId string) ApiFindEndpointByIdAndProjectIdRequest {
	return ApiFindEndpointByIdAndProjectIdRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		endpointId: endpointId,
	}
}

// Execute executes the request
//  @return Endpoint
func (a *EndpointAPIService) FindEndpointByIdAndProjectIdExecute(r ApiFindEndpointByIdAndProjectIdRequest) (*Endpoint, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Endpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.FindEndpointByIdAndProjectId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints/{endpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.endpoint+json", "application/vnd.graal.systems.v1.error+json"}

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

type ApiFindEndpointsByProjectIdRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	page *int32
	size *int32
}

func (r ApiFindEndpointsByProjectIdRequest) XTenant(xTenant string) ApiFindEndpointsByProjectIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindEndpointsByProjectIdRequest) Page(page int32) ApiFindEndpointsByProjectIdRequest {
	r.page = &page
	return r
}

func (r ApiFindEndpointsByProjectIdRequest) Size(size int32) ApiFindEndpointsByProjectIdRequest {
	r.size = &size
	return r
}

func (r ApiFindEndpointsByProjectIdRequest) Execute() (*EndpointPage, *http.Response, error) {
	return r.ApiService.FindEndpointsByProjectIdExecute(r)
}

/*
FindEndpointsByProjectId Retrieve all endpoints

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @return ApiFindEndpointsByProjectIdRequest
*/
func (a *EndpointAPIService) FindEndpointsByProjectId(ctx context.Context, projectId string) ApiFindEndpointsByProjectIdRequest {
	return ApiFindEndpointsByProjectIdRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return EndpointPage
func (a *EndpointAPIService) FindEndpointsByProjectIdExecute(r ApiFindEndpointsByProjectIdRequest) (*EndpointPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndpointPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.FindEndpointsByProjectId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.endpoint+json"}

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

type ApiGetLogsForEndpointRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpointId string
	cursor *string
}

func (r ApiGetLogsForEndpointRequest) XTenant(xTenant string) ApiGetLogsForEndpointRequest {
	r.xTenant = &xTenant
	return r
}

// The cursor
func (r ApiGetLogsForEndpointRequest) Cursor(cursor string) ApiGetLogsForEndpointRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetLogsForEndpointRequest) Execute() ([]LogEntry, *http.Response, error) {
	return r.ApiService.GetLogsForEndpointExecute(r)
}

/*
GetLogsForEndpoint Get logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @param endpointId Id of the endpoint
 @return ApiGetLogsForEndpointRequest
*/
func (a *EndpointAPIService) GetLogsForEndpoint(ctx context.Context, projectId string, endpointId string) ApiGetLogsForEndpointRequest {
	return ApiGetLogsForEndpointRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		endpointId: endpointId,
	}
}

// Execute executes the request
//  @return []LogEntry
func (a *EndpointAPIService) GetLogsForEndpointExecute(r ApiGetLogsForEndpointRequest) ([]LogEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LogEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.GetLogsForEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints/{endpointId}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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

type ApiUpdateEndpointByIdAndProjectIdRequest struct {
	ctx context.Context
	ApiService *EndpointAPIService
	xTenant *string
	projectId string
	endpointId string
	patch *[]Patch
}

func (r ApiUpdateEndpointByIdAndProjectIdRequest) XTenant(xTenant string) ApiUpdateEndpointByIdAndProjectIdRequest {
	r.xTenant = &xTenant
	return r
}

// The patch
func (r ApiUpdateEndpointByIdAndProjectIdRequest) Patch(patch []Patch) ApiUpdateEndpointByIdAndProjectIdRequest {
	r.patch = &patch
	return r
}

func (r ApiUpdateEndpointByIdAndProjectIdRequest) Execute() (*Endpoint, *http.Response, error) {
	return r.ApiService.UpdateEndpointByIdAndProjectIdExecute(r)
}

/*
UpdateEndpointByIdAndProjectId Update an endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of the project
 @param endpointId Id of the endpoint
 @return ApiUpdateEndpointByIdAndProjectIdRequest
*/
func (a *EndpointAPIService) UpdateEndpointByIdAndProjectId(ctx context.Context, projectId string, endpointId string) ApiUpdateEndpointByIdAndProjectIdRequest {
	return ApiUpdateEndpointByIdAndProjectIdRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		endpointId: endpointId,
	}
}

// Execute executes the request
//  @return Endpoint
func (a *EndpointAPIService) UpdateEndpointByIdAndProjectIdExecute(r ApiUpdateEndpointByIdAndProjectIdRequest) (*Endpoint, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Endpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointAPIService.UpdateEndpointByIdAndProjectId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{projectId}/endpoints/{endpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.endpoint+json", "application/vnd.graal.systems.v1.error+json"}

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

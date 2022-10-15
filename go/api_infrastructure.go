/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems/sdk

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

// InfrastructureApiService InfrastructureApi service
type InfrastructureApiService service

type ApiDeleteInfrastructureByIdRequest struct {
	ctx _context.Context
	ApiService *InfrastructureApiService
	xTenant *string
	infrastructureId string
}

func (r ApiDeleteInfrastructureByIdRequest) XTenant(xTenant string) ApiDeleteInfrastructureByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiDeleteInfrastructureByIdRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteInfrastructureByIdExecute(r)
}

/*
 * DeleteInfrastructureById Delete a infrastructure by an id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param infrastructureId Id of the infrastructure
 * @return ApiDeleteInfrastructureByIdRequest
 */
func (a *InfrastructureApiService) DeleteInfrastructureById(ctx _context.Context, infrastructureId string) ApiDeleteInfrastructureByIdRequest {
	return ApiDeleteInfrastructureByIdRequest{
		ApiService: a,
		ctx: ctx,
		infrastructureId: infrastructureId,
	}
}

/*
 * Execute executes the request
 */
func (a *InfrastructureApiService) DeleteInfrastructureByIdExecute(r ApiDeleteInfrastructureByIdRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.DeleteInfrastructureById")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infrastructures/{infrastructureId}"
	localVarPath = strings.Replace(localVarPath, "{"+"infrastructureId"+"}", _neturl.PathEscape(parameterToString(r.infrastructureId, "")), -1)

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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFindCapacityByInfrastructureByIdRequest struct {
	ctx _context.Context
	ApiService *InfrastructureApiService
	xTenant *string
	infrastructureId string
}

func (r ApiFindCapacityByInfrastructureByIdRequest) XTenant(xTenant string) ApiFindCapacityByInfrastructureByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindCapacityByInfrastructureByIdRequest) Execute() (TargetInfo, *_nethttp.Response, error) {
	return r.ApiService.FindCapacityByInfrastructureByIdExecute(r)
}

/*
 * FindCapacityByInfrastructureById Retrieve the capacity of an infrastructure
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param infrastructureId Id of the infrastructure
 * @return ApiFindCapacityByInfrastructureByIdRequest
 */
func (a *InfrastructureApiService) FindCapacityByInfrastructureById(ctx _context.Context, infrastructureId string) ApiFindCapacityByInfrastructureByIdRequest {
	return ApiFindCapacityByInfrastructureByIdRequest{
		ApiService: a,
		ctx: ctx,
		infrastructureId: infrastructureId,
	}
}

/*
 * Execute executes the request
 * @return TargetInfo
 */
func (a *InfrastructureApiService) FindCapacityByInfrastructureByIdExecute(r ApiFindCapacityByInfrastructureByIdRequest) (TargetInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.FindCapacityByInfrastructureById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infrastructures/{infrastructureId}/capacity"
	localVarPath = strings.Replace(localVarPath, "{"+"infrastructureId"+"}", _neturl.PathEscape(parameterToString(r.infrastructureId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.capacity+json"}

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

type ApiFindInfrastructureByInfrastructureIdRequest struct {
	ctx _context.Context
	ApiService *InfrastructureApiService
	xTenant *string
	infrastructureId string
}

func (r ApiFindInfrastructureByInfrastructureIdRequest) XTenant(xTenant string) ApiFindInfrastructureByInfrastructureIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindInfrastructureByInfrastructureIdRequest) Execute() (Infrastructure, *_nethttp.Response, error) {
	return r.ApiService.FindInfrastructureByInfrastructureIdExecute(r)
}

/*
 * FindInfrastructureByInfrastructureId Find infrastructure by Id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param infrastructureId Id of the infrastructure
 * @return ApiFindInfrastructureByInfrastructureIdRequest
 */
func (a *InfrastructureApiService) FindInfrastructureByInfrastructureId(ctx _context.Context, infrastructureId string) ApiFindInfrastructureByInfrastructureIdRequest {
	return ApiFindInfrastructureByInfrastructureIdRequest{
		ApiService: a,
		ctx: ctx,
		infrastructureId: infrastructureId,
	}
}

/*
 * Execute executes the request
 * @return Infrastructure
 */
func (a *InfrastructureApiService) FindInfrastructureByInfrastructureIdExecute(r ApiFindInfrastructureByInfrastructureIdRequest) (Infrastructure, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Infrastructure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.FindInfrastructureByInfrastructureId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infrastructures/{infrastructureId}"
	localVarPath = strings.Replace(localVarPath, "{"+"infrastructureId"+"}", _neturl.PathEscape(parameterToString(r.infrastructureId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.infrastructure+json", "application/vnd.graal.systems.v1.error+json"}

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

type ApiFindInfrastructuresRequest struct {
	ctx _context.Context
	ApiService *InfrastructureApiService
	xTenant *string
}

func (r ApiFindInfrastructuresRequest) XTenant(xTenant string) ApiFindInfrastructuresRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindInfrastructuresRequest) Execute() ([]Infrastructure, *_nethttp.Response, error) {
	return r.ApiService.FindInfrastructuresExecute(r)
}

/*
 * FindInfrastructures Retrieve all infrastructure
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFindInfrastructuresRequest
 */
func (a *InfrastructureApiService) FindInfrastructures(ctx _context.Context) ApiFindInfrastructuresRequest {
	return ApiFindInfrastructuresRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Infrastructure
 */
func (a *InfrastructureApiService) FindInfrastructuresExecute(r ApiFindInfrastructuresRequest) ([]Infrastructure, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Infrastructure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.FindInfrastructures")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infrastructures"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.infrastructure+json"}

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

type ApiFindPublicIpsByInfrastructureIdRequest struct {
	ctx _context.Context
	ApiService *InfrastructureApiService
	xTenant *string
	infrastructureId string
}

func (r ApiFindPublicIpsByInfrastructureIdRequest) XTenant(xTenant string) ApiFindPublicIpsByInfrastructureIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindPublicIpsByInfrastructureIdRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.FindPublicIpsByInfrastructureIdExecute(r)
}

/*
 * FindPublicIpsByInfrastructureId Find public ips by infrastructure id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param infrastructureId Id of the infrastructure
 * @return ApiFindPublicIpsByInfrastructureIdRequest
 */
func (a *InfrastructureApiService) FindPublicIpsByInfrastructureId(ctx _context.Context, infrastructureId string) ApiFindPublicIpsByInfrastructureIdRequest {
	return ApiFindPublicIpsByInfrastructureIdRequest{
		ApiService: a,
		ctx: ctx,
		infrastructureId: infrastructureId,
	}
}

/*
 * Execute executes the request
 * @return []string
 */
func (a *InfrastructureApiService) FindPublicIpsByInfrastructureIdExecute(r ApiFindPublicIpsByInfrastructureIdRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.FindPublicIpsByInfrastructureId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infrastructures/{infrastructureId}/public-ips"
	localVarPath = strings.Replace(localVarPath, "{"+"infrastructureId"+"}", _neturl.PathEscape(parameterToString(r.infrastructureId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.graal.systems.v1.error+json"}

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

type ApiUpdateInfrastructureRequest struct {
	ctx _context.Context
	ApiService *InfrastructureApiService
	xTenant *string
	infrastructureId string
	patch *[]Patch
}

func (r ApiUpdateInfrastructureRequest) XTenant(xTenant string) ApiUpdateInfrastructureRequest {
	r.xTenant = &xTenant
	return r
}
func (r ApiUpdateInfrastructureRequest) Patch(patch []Patch) ApiUpdateInfrastructureRequest {
	r.patch = &patch
	return r
}

func (r ApiUpdateInfrastructureRequest) Execute() (Infrastructure, *_nethttp.Response, error) {
	return r.ApiService.UpdateInfrastructureExecute(r)
}

/*
 * UpdateInfrastructure Update an infrastructure
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param infrastructureId Id of the infrastructure
 * @return ApiUpdateInfrastructureRequest
 */
func (a *InfrastructureApiService) UpdateInfrastructure(ctx _context.Context, infrastructureId string) ApiUpdateInfrastructureRequest {
	return ApiUpdateInfrastructureRequest{
		ApiService: a,
		ctx: ctx,
		infrastructureId: infrastructureId,
	}
}

/*
 * Execute executes the request
 * @return Infrastructure
 */
func (a *InfrastructureApiService) UpdateInfrastructureExecute(r ApiUpdateInfrastructureRequest) (Infrastructure, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Infrastructure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InfrastructureApiService.UpdateInfrastructure")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infrastructures/{infrastructureId}"
	localVarPath = strings.Replace(localVarPath, "{"+"infrastructureId"+"}", _neturl.PathEscape(parameterToString(r.infrastructureId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.infrastructure+json", "application/vnd.graal.systems.v1.error+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(*r.xTenant, "")
	// body params
	localVarPostBody = r.patch
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

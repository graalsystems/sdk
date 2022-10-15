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
)

// Linger please
var (
	_ _context.Context
)

// HistoryApiService HistoryApi service
type HistoryApiService service

type ApiAddResourceHistoryForUserRequest struct {
	ctx _context.Context
	ApiService *HistoryApiService
	xTenant *string
	type_ *string
	id *string
}

func (r ApiAddResourceHistoryForUserRequest) XTenant(xTenant string) ApiAddResourceHistoryForUserRequest {
	r.xTenant = &xTenant
	return r
}
func (r ApiAddResourceHistoryForUserRequest) Type_(type_ string) ApiAddResourceHistoryForUserRequest {
	r.type_ = &type_
	return r
}
func (r ApiAddResourceHistoryForUserRequest) Id(id string) ApiAddResourceHistoryForUserRequest {
	r.id = &id
	return r
}

func (r ApiAddResourceHistoryForUserRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AddResourceHistoryForUserExecute(r)
}

/*
 * AddResourceHistoryForUser Add history for user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAddResourceHistoryForUserRequest
 */
func (a *HistoryApiService) AddResourceHistoryForUser(ctx _context.Context) ApiAddResourceHistoryForUserRequest {
	return ApiAddResourceHistoryForUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *HistoryApiService) AddResourceHistoryForUserExecute(r ApiAddResourceHistoryForUserRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.AddResourceHistoryForUser")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}
	if r.type_ == nil {
		return nil, reportError("type_ is required and must be specified")
	}
	if r.id == nil {
		return nil, reportError("id is required and must be specified")
	}

	localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAddSearchHistoryForUserRequest struct {
	ctx _context.Context
	ApiService *HistoryApiService
	xTenant *string
	searchHistory *SearchHistory
}

func (r ApiAddSearchHistoryForUserRequest) XTenant(xTenant string) ApiAddSearchHistoryForUserRequest {
	r.xTenant = &xTenant
	return r
}
func (r ApiAddSearchHistoryForUserRequest) SearchHistory(searchHistory SearchHistory) ApiAddSearchHistoryForUserRequest {
	r.searchHistory = &searchHistory
	return r
}

func (r ApiAddSearchHistoryForUserRequest) Execute() ([]SearchHistory, *_nethttp.Response, error) {
	return r.ApiService.AddSearchHistoryForUserExecute(r)
}

/*
 * AddSearchHistoryForUser Add search history for user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAddSearchHistoryForUserRequest
 */
func (a *HistoryApiService) AddSearchHistoryForUser(ctx _context.Context) ApiAddSearchHistoryForUserRequest {
	return ApiAddSearchHistoryForUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []SearchHistory
 */
func (a *HistoryApiService) AddSearchHistoryForUserExecute(r ApiAddSearchHistoryForUserRequest) ([]SearchHistory, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SearchHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.AddSearchHistoryForUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}
	if r.searchHistory == nil {
		return localVarReturnValue, nil, reportError("searchHistory is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.searchhistory+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.searchhistory+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(*r.xTenant, "")
	// body params
	localVarPostBody = r.searchHistory
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

type ApiGetResourceHistoryRequest struct {
	ctx _context.Context
	ApiService *HistoryApiService
	xTenant *string
}

func (r ApiGetResourceHistoryRequest) XTenant(xTenant string) ApiGetResourceHistoryRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiGetResourceHistoryRequest) Execute() ([]Result, *_nethttp.Response, error) {
	return r.ApiService.GetResourceHistoryExecute(r)
}

/*
 * GetResourceHistory Get resource history for user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetResourceHistoryRequest
 */
func (a *HistoryApiService) GetResourceHistory(ctx _context.Context) ApiGetResourceHistoryRequest {
	return ApiGetResourceHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Result
 */
func (a *HistoryApiService) GetResourceHistoryExecute(r ApiGetResourceHistoryRequest) ([]Result, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Result
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.GetResourceHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.result+json", "application/vnd.graal.systems.v1.error+json"}

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

type ApiGetSearchHistoryRequest struct {
	ctx _context.Context
	ApiService *HistoryApiService
	xTenant *string
}

func (r ApiGetSearchHistoryRequest) XTenant(xTenant string) ApiGetSearchHistoryRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiGetSearchHistoryRequest) Execute() ([]SearchHistory, *_nethttp.Response, error) {
	return r.ApiService.GetSearchHistoryExecute(r)
}

/*
 * GetSearchHistory Search history for user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetSearchHistoryRequest
 */
func (a *HistoryApiService) GetSearchHistory(ctx _context.Context) ApiGetSearchHistoryRequest {
	return ApiGetSearchHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []SearchHistory
 */
func (a *HistoryApiService) GetSearchHistoryExecute(r ApiGetSearchHistoryRequest) ([]SearchHistory, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SearchHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryApiService.GetSearchHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/history"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.searchhistory+json", "application/vnd.graal.systems.v1.error+json"}

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
/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


// CodeAPIService CodeAPI service
type CodeAPIService service

type ApiDownloadCodeRequest struct {
	ctx context.Context
	ApiService *CodeAPIService
	xTenant *string
	source *string
	destination *string
	filename []*os.File
}

func (r ApiDownloadCodeRequest) XTenant(xTenant string) ApiDownloadCodeRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiDownloadCodeRequest) Source(source string) ApiDownloadCodeRequest {
	r.source = &source
	return r
}

func (r ApiDownloadCodeRequest) Destination(destination string) ApiDownloadCodeRequest {
	r.destination = &destination
	return r
}

func (r ApiDownloadCodeRequest) Filename(filename []*os.File) ApiDownloadCodeRequest {
	r.filename = filename
	return r
}

func (r ApiDownloadCodeRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.DownloadCodeExecute(r)
}

/*
DownloadCode Download code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDownloadCodeRequest
*/
func (a *CodeAPIService) DownloadCode(ctx context.Context) ApiDownloadCodeRequest {
	return ApiDownloadCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *CodeAPIService) DownloadCodeExecute(r ApiDownloadCodeRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodeAPIService.DownloadCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/code?download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}

	if r.source != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "source", r.source, "")
	}
	if r.destination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "destination", r.destination, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

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

type ApiVerifyCodeRequest struct {
	ctx context.Context
	ApiService *CodeAPIService
	xTenant *string
	source *string
	destination *string
	filename []*os.File
}

func (r ApiVerifyCodeRequest) XTenant(xTenant string) ApiVerifyCodeRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiVerifyCodeRequest) Source(source string) ApiVerifyCodeRequest {
	r.source = &source
	return r
}

func (r ApiVerifyCodeRequest) Destination(destination string) ApiVerifyCodeRequest {
	r.destination = &destination
	return r
}

func (r ApiVerifyCodeRequest) Filename(filename []*os.File) ApiVerifyCodeRequest {
	r.filename = filename
	return r
}

func (r ApiVerifyCodeRequest) Execute() (*http.Response, error) {
	return r.ApiService.VerifyCodeExecute(r)
}

/*
VerifyCode Verify code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyCodeRequest
*/
func (a *CodeAPIService) VerifyCode(ctx context.Context) ApiVerifyCodeRequest {
	return ApiVerifyCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CodeAPIService) VerifyCodeExecute(r ApiVerifyCodeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodeAPIService.VerifyCode")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/code?verify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}

	if r.source != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "source", r.source, "")
	}
	if r.destination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "destination", r.destination, "")
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

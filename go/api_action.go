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
)


// ActionAPIService ActionAPI service
type ActionAPIService service

type ApiExecuteActionRequest struct {
	ctx context.Context
	ApiService *ActionAPIService
	xTenant *string
	action *Action
}

func (r ApiExecuteActionRequest) XTenant(xTenant string) ApiExecuteActionRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiExecuteActionRequest) Action(action Action) ApiExecuteActionRequest {
	r.action = &action
	return r
}

func (r ApiExecuteActionRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExecuteActionExecute(r)
}

/*
ExecuteAction Execute an action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExecuteActionRequest
*/
func (a *ActionAPIService) ExecuteAction(ctx context.Context) ApiExecuteActionRequest {
	return ApiExecuteActionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ActionAPIService) ExecuteActionExecute(r ApiExecuteActionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionAPIService.ExecuteAction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}
	if r.action == nil {
		return nil, reportError("action is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.action+json"}

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
	// body params
	localVarPostBody = r.action
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

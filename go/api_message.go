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


// MessageAPIService MessageAPI service
type MessageAPIService service

type ApiSendMessageRequest struct {
	ctx context.Context
	ApiService *MessageAPIService
	xTenant *string
	message *Message
}

func (r ApiSendMessageRequest) XTenant(xTenant string) ApiSendMessageRequest {
	r.xTenant = &xTenant
	return r
}

// The message to be created
func (r ApiSendMessageRequest) Message(message Message) ApiSendMessageRequest {
	r.message = &message
	return r
}

func (r ApiSendMessageRequest) Execute() (*http.Response, error) {
	return r.ApiService.SendMessageExecute(r)
}

/*
SendMessage Send an message

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendMessageRequest
*/
func (a *MessageAPIService) SendMessage(ctx context.Context) ApiSendMessageRequest {
	return ApiSendMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MessageAPIService) SendMessageExecute(r ApiSendMessageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageAPIService.SendMessage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return nil, reportError("xTenant is required and must be specified")
	}
	if r.message == nil {
		return nil, reportError("message is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.message+json"}

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
	localVarPostBody = r.message
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

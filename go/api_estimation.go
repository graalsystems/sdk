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
)


// EstimationAPIService EstimationAPI service
type EstimationAPIService service

type ApiCreateEstimationRequest struct {
	ctx context.Context
	ApiService *EstimationAPIService
	xTenant *string
	estimation *Estimation
}

func (r ApiCreateEstimationRequest) XTenant(xTenant string) ApiCreateEstimationRequest {
	r.xTenant = &xTenant
	return r
}

// The estimation to be calculated
func (r ApiCreateEstimationRequest) Estimation(estimation Estimation) ApiCreateEstimationRequest {
	r.estimation = &estimation
	return r
}

func (r ApiCreateEstimationRequest) Execute() (*Estimation, *http.Response, error) {
	return r.ApiService.CreateEstimationExecute(r)
}

/*
CreateEstimation Create an estimation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateEstimationRequest
*/
func (a *EstimationAPIService) CreateEstimation(ctx context.Context) ApiCreateEstimationRequest {
	return ApiCreateEstimationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Estimation
func (a *EstimationAPIService) CreateEstimationExecute(r ApiCreateEstimationRequest) (*Estimation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Estimation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EstimationAPIService.CreateEstimation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/estimations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}
	if r.estimation == nil {
		return localVarReturnValue, nil, reportError("estimation is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.graal.systems.v1.estimation+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.estimation+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Tenant", r.xTenant, "")
	// body params
	localVarPostBody = r.estimation
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
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
	"time"
)


// CostAPIService CostAPI service
type CostAPIService service

type ApiGetCostsRequest struct {
	ctx context.Context
	ApiService *CostAPIService
	xTenant *string
	from *time.Time
	to *time.Time
	for_ *string
	params *map[string]string
}

func (r ApiGetCostsRequest) XTenant(xTenant string) ApiGetCostsRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiGetCostsRequest) From(from time.Time) ApiGetCostsRequest {
	r.from = &from
	return r
}

func (r ApiGetCostsRequest) To(to time.Time) ApiGetCostsRequest {
	r.to = &to
	return r
}

func (r ApiGetCostsRequest) For_(for_ string) ApiGetCostsRequest {
	r.for_ = &for_
	return r
}

func (r ApiGetCostsRequest) Params(params map[string]string) ApiGetCostsRequest {
	r.params = &params
	return r
}

func (r ApiGetCostsRequest) Execute() (*CostStats, *http.Response, error) {
	return r.ApiService.GetCostsExecute(r)
}

/*
GetCosts Find the costs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCostsRequest
*/
func (a *CostAPIService) GetCosts(ctx context.Context) ApiGetCostsRequest {
	return ApiGetCostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CostStats
func (a *CostAPIService) GetCostsExecute(r ApiGetCostsRequest) (*CostStats, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CostStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CostAPIService.GetCosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/costs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xTenant == nil {
		return localVarReturnValue, nil, reportError("xTenant is required and must be specified")
	}

	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.for_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "for", r.for_, "")
	}
	if r.params != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "params", r.params, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.costs+json", "application/vnd.graal.systems.v1.error+json"}

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

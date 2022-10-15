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

// FirewallApiService FirewallApi service
type FirewallApiService service

type ApiCheckAccessByIpRequest struct {
	ctx _context.Context
	ApiService *FirewallApiService
	xTenant *string
	ip string
}

func (r ApiCheckAccessByIpRequest) XTenant(xTenant string) ApiCheckAccessByIpRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiCheckAccessByIpRequest) Execute() (FirewallCheckStatus, *_nethttp.Response, error) {
	return r.ApiService.CheckAccessByIpExecute(r)
}

/*
 * CheckAccessByIp check access for firewall
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ip Input IP
 * @return ApiCheckAccessByIpRequest
 */
func (a *FirewallApiService) CheckAccessByIp(ctx _context.Context, ip string) ApiCheckAccessByIpRequest {
	return ApiCheckAccessByIpRequest{
		ApiService: a,
		ctx: ctx,
		ip: ip,
	}
}

/*
 * Execute executes the request
 * @return FirewallCheckStatus
 */
func (a *FirewallApiService) CheckAccessByIpExecute(r ApiCheckAccessByIpRequest) (FirewallCheckStatus, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FirewallCheckStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewallApiService.CheckAccessByIp")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/firewall/rules/check/{ip}"
	localVarPath = strings.Replace(localVarPath, "{"+"ip"+"}", _neturl.PathEscape(parameterToString(r.ip, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.firewall.rule+json", "application/vnd.graal.systems.v1.error+json"}

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

type ApiDeleteFirewallRuleByIdRequest struct {
	ctx _context.Context
	ApiService *FirewallApiService
	xTenant *string
	ruleId string
}

func (r ApiDeleteFirewallRuleByIdRequest) XTenant(xTenant string) ApiDeleteFirewallRuleByIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiDeleteFirewallRuleByIdRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteFirewallRuleByIdExecute(r)
}

/*
 * DeleteFirewallRuleById Delete a rule by its id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ruleId Id of the FirewallRule
 * @return ApiDeleteFirewallRuleByIdRequest
 */
func (a *FirewallApiService) DeleteFirewallRuleById(ctx _context.Context, ruleId string) ApiDeleteFirewallRuleByIdRequest {
	return ApiDeleteFirewallRuleByIdRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

/*
 * Execute executes the request
 */
func (a *FirewallApiService) DeleteFirewallRuleByIdExecute(r ApiDeleteFirewallRuleByIdRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewallApiService.DeleteFirewallRuleById")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/firewall/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", _neturl.PathEscape(parameterToString(r.ruleId, "")), -1)

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

type ApiFindFirewallRuleByFirewallRuleIdRequest struct {
	ctx _context.Context
	ApiService *FirewallApiService
	xTenant *string
	ruleId string
}

func (r ApiFindFirewallRuleByFirewallRuleIdRequest) XTenant(xTenant string) ApiFindFirewallRuleByFirewallRuleIdRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindFirewallRuleByFirewallRuleIdRequest) Execute() (FirewallRule, *_nethttp.Response, error) {
	return r.ApiService.FindFirewallRuleByFirewallRuleIdExecute(r)
}

/*
 * FindFirewallRuleByFirewallRuleId Find rule by Id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ruleId Id of the FirewallRule
 * @return ApiFindFirewallRuleByFirewallRuleIdRequest
 */
func (a *FirewallApiService) FindFirewallRuleByFirewallRuleId(ctx _context.Context, ruleId string) ApiFindFirewallRuleByFirewallRuleIdRequest {
	return ApiFindFirewallRuleByFirewallRuleIdRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

/*
 * Execute executes the request
 * @return FirewallRule
 */
func (a *FirewallApiService) FindFirewallRuleByFirewallRuleIdExecute(r ApiFindFirewallRuleByFirewallRuleIdRequest) (FirewallRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewallApiService.FindFirewallRuleByFirewallRuleId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/firewall/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", _neturl.PathEscape(parameterToString(r.ruleId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.firewall.rule+json", "application/vnd.graal.systems.v1.error+json"}

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

type ApiFindFirewallRulesRequest struct {
	ctx _context.Context
	ApiService *FirewallApiService
	xTenant *string
}

func (r ApiFindFirewallRulesRequest) XTenant(xTenant string) ApiFindFirewallRulesRequest {
	r.xTenant = &xTenant
	return r
}

func (r ApiFindFirewallRulesRequest) Execute() ([]FirewallRule, *_nethttp.Response, error) {
	return r.ApiService.FindFirewallRulesExecute(r)
}

/*
 * FindFirewallRules Retrieve all firewall rules
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFindFirewallRulesRequest
 */
func (a *FirewallApiService) FindFirewallRules(ctx _context.Context) ApiFindFirewallRulesRequest {
	return ApiFindFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []FirewallRule
 */
func (a *FirewallApiService) FindFirewallRulesExecute(r ApiFindFirewallRulesRequest) ([]FirewallRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []FirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewallApiService.FindFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/firewall/rules"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.firewall.rule+json"}

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

type ApiUpdateFirewallRuleRequest struct {
	ctx _context.Context
	ApiService *FirewallApiService
	xTenant *string
	ruleId string
	patch *[]Patch
}

func (r ApiUpdateFirewallRuleRequest) XTenant(xTenant string) ApiUpdateFirewallRuleRequest {
	r.xTenant = &xTenant
	return r
}
func (r ApiUpdateFirewallRuleRequest) Patch(patch []Patch) ApiUpdateFirewallRuleRequest {
	r.patch = &patch
	return r
}

func (r ApiUpdateFirewallRuleRequest) Execute() (FirewallRule, *_nethttp.Response, error) {
	return r.ApiService.UpdateFirewallRuleExecute(r)
}

/*
 * UpdateFirewallRule Update a rule
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ruleId Id of the FirewallRule
 * @return ApiUpdateFirewallRuleRequest
 */
func (a *FirewallApiService) UpdateFirewallRule(ctx _context.Context, ruleId string) ApiUpdateFirewallRuleRequest {
	return ApiUpdateFirewallRuleRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

/*
 * Execute executes the request
 * @return FirewallRule
 */
func (a *FirewallApiService) UpdateFirewallRuleExecute(r ApiUpdateFirewallRuleRequest) (FirewallRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewallApiService.UpdateFirewallRule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/firewall/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", _neturl.PathEscape(parameterToString(r.ruleId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.graal.systems.v1.firewall.rule+json", "application/vnd.graal.systems.v1.error+json"}

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

"""
    Tenant API

    Tenant API  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Contact: abc@layer.fr
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from openapi_client.api_client import ApiClient, Endpoint as _Endpoint
from openapi_client.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from openapi_client.model.activity import Activity
from openapi_client.model.http_validation_error import HTTPValidationError


class ActivityApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.get_timelines_v1_timelines_get_endpoint = _Endpoint(
            settings={
                'response_type': ([Activity],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/activities',
                'operation_id': 'get_timelines_v1_timelines_get',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'resource_id',
                    'resource_type',
                ],
                'required': [
                    'x_tenant',
                    'resource_id',
                    'resource_type',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'x_tenant':
                        (str,),
                    'resource_id':
                        (str,),
                    'resource_type':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'resource_id': 'resource_id',
                    'resource_type': 'resource_type',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'resource_id': 'query',
                    'resource_type': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )

    def get_timelines_v1_timelines_get(
        self,
        x_tenant,
        resource_id,
        resource_type,
        **kwargs
    ):
        """Get timelines.  # noqa: E501

        Gets all timelines associated with a given tenant and resource. The tenant is identified by its UUID, while the resource is represented by its UUID and resource type.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_timelines_v1_timelines_get(x_tenant, resource_id, resource_type, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            resource_id (str):
            resource_type (str):

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            async_req (bool): execute request asynchronously

        Returns:
            [Activity]
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['x_tenant'] = \
            x_tenant
        kwargs['resource_id'] = \
            resource_id
        kwargs['resource_type'] = \
            resource_type
        return self.get_timelines_v1_timelines_get_endpoint.call_with_http_info(**kwargs)

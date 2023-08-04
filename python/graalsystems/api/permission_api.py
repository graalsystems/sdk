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


class PermissionApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.find_permissions_endpoint = _Endpoint(
            settings={
                'response_type': ([str],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/permissions',
                'operation_id': 'find_permissions',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'resource_type',
                    'resource_id',
                    'principal_type',
                    'principal_id',
                ],
                'required': [
                    'x_tenant',
                    'resource_type',
                    'resource_id',
                    'principal_type',
                    'principal_id',
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
                    'resource_type':
                        (str,),
                    'resource_id':
                        (str,),
                    'principal_type':
                        (str,),
                    'principal_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'resource_type': 'resource_type',
                    'resource_id': 'resource_id',
                    'principal_type': 'principal_type',
                    'principal_id': 'principal_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'resource_type': 'query',
                    'resource_id': 'query',
                    'principal_type': 'query',
                    'principal_id': 'query',
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
        self.has_permission_endpoint = _Endpoint(
            settings={
                'response_type': (bool,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/permissions?metadata',
                'operation_id': 'has_permission',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'resource_type',
                    'resource_id',
                    'permission',
                ],
                'required': [
                    'x_tenant',
                    'resource_type',
                    'resource_id',
                    'permission',
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
                    'resource_type':
                        (str,),
                    'resource_id':
                        (str,),
                    'permission':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'resource_type': 'resource_type',
                    'resource_id': 'resource_id',
                    'permission': 'permission',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'resource_type': 'query',
                    'resource_id': 'query',
                    'permission': 'query',
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

    def find_permissions(
        self,
        x_tenant,
        resource_type,
        resource_id,
        principal_type,
        principal_id,
        **kwargs
    ):
        """Retrieve permissions for a specific principal  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.find_permissions(x_tenant, resource_type, resource_id, principal_type, principal_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            resource_type (str):
            resource_id (str):
            principal_type (str):
            principal_id (str):

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
            [str]
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
        kwargs['resource_type'] = \
            resource_type
        kwargs['resource_id'] = \
            resource_id
        kwargs['principal_type'] = \
            principal_type
        kwargs['principal_id'] = \
            principal_id
        return self.find_permissions_endpoint.call_with_http_info(**kwargs)

    def has_permission(
        self,
        x_tenant,
        resource_type,
        resource_id,
        permission,
        **kwargs
    ):
        """Retrieve permissions for a specific principal  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.has_permission(x_tenant, resource_type, resource_id, permission, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            resource_type (str):
            resource_id (str):
            permission (str):

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
            bool
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
        kwargs['resource_type'] = \
            resource_type
        kwargs['resource_id'] = \
            resource_id
        kwargs['permission'] = \
            permission
        return self.has_permission_endpoint.call_with_http_info(**kwargs)


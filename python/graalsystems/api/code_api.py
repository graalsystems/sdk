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


class CodeApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.download_code_endpoint = _Endpoint(
            settings={
                'response_type': (file_type,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/code?download',
                'operation_id': 'download_code',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'source',
                    'destination',
                    'filename',
                ],
                'required': [
                    'x_tenant',
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
                    'source':
                        (str,),
                    'destination':
                        (str,),
                    'filename':
                        ([file_type],),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'source': 'source',
                    'destination': 'destination',
                    'filename': 'filename',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'source': 'query',
                    'destination': 'query',
                    'filename': 'form',
                },
                'collection_format_map': {
                    'filename': 'csv',
                }
            },
            headers_map={
                'accept': [
                    '*/*'
                ],
                'content_type': [
                    'multipart/form-data'
                ]
            },
            api_client=api_client
        )
        self.verify_code_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/code?verify',
                'operation_id': 'verify_code',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'source',
                    'destination',
                    'filename',
                ],
                'required': [
                    'x_tenant',
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
                    'source':
                        (str,),
                    'destination':
                        (str,),
                    'filename':
                        ([file_type],),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'source': 'source',
                    'destination': 'destination',
                    'filename': 'filename',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'source': 'query',
                    'destination': 'query',
                    'filename': 'form',
                },
                'collection_format_map': {
                    'filename': 'csv',
                }
            },
            headers_map={
                'accept': [],
                'content_type': [
                    'multipart/form-data'
                ]
            },
            api_client=api_client
        )

    def download_code(
        self,
        x_tenant,
        **kwargs
    ):
        """Download code  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.download_code(x_tenant, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):

        Keyword Args:
            source (str): [optional]
            destination (str): [optional]
            filename ([file_type]): [optional]
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
            file_type
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
        return self.download_code_endpoint.call_with_http_info(**kwargs)

    def verify_code(
        self,
        x_tenant,
        **kwargs
    ):
        """Verify code  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.verify_code(x_tenant, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):

        Keyword Args:
            source (str): [optional]
            destination (str): [optional]
            filename ([file_type]): [optional]
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
            None
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
        return self.verify_code_endpoint.call_with_http_info(**kwargs)


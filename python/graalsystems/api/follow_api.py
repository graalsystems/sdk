"""
    Tenant API

    Tenant API  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Contact: abc@layer.fr
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from graalsystems.api_client import ApiClient, Endpoint as _Endpoint
from graalsystems.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from graalsystems.model.follow import Follow
from graalsystems.model.http_validation_error import HTTPValidationError
from graalsystems.model.user1 import User1


class FollowApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.create_project_followers_projects_project_id_followers_post_endpoint = _Endpoint(
            settings={
                'response_type': (Follow,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/projects/{project_id}/followers',
                'operation_id': 'create_project_followers_projects_project_id_followers_post',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'project_id',
                ],
                'required': [
                    'x_tenant',
                    'project_id',
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
                    'project_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'project_id': 'project_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'project_id': 'path',
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
        self.create_user_followers_users_user_id_followers_post_endpoint = _Endpoint(
            settings={
                'response_type': (Follow,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/users/{user_id}/followers',
                'operation_id': 'create_user_followers_users_user_id_followers_post',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'user_id',
                ],
                'required': [
                    'x_tenant',
                    'user_id',
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
                    'user_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'user_id': 'user_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'user_id': 'path',
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
        self.delete_project_followers_projects_project_id_followers_delete_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/following/projects/{project_id}',
                'operation_id': 'delete_project_followers_projects_project_id_followers_delete',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'project_id',
                ],
                'required': [
                    'x_tenant',
                    'project_id',
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
                    'project_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'project_id': 'project_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'project_id': 'path',
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
        self.delete_user_followers_users_user_id_followers_delete_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/following/users/{user_id}',
                'operation_id': 'delete_user_followers_users_user_id_followers_delete',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'user_id',
                ],
                'required': [
                    'x_tenant',
                    'user_id',
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
                    'user_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'user_id': 'user_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'user_id': 'path',
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
        self.get_project_followers_projects_project_id_followers_get_endpoint = _Endpoint(
            settings={
                'response_type': ([User1],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/projects/{project_id}/followers',
                'operation_id': 'get_project_followers_projects_project_id_followers_get',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'project_id',
                ],
                'required': [
                    'x_tenant',
                    'project_id',
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
                    'project_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'project_id': 'project_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'project_id': 'path',
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
        self.get_user_followers_users_user_id_followers_get_endpoint = _Endpoint(
            settings={
                'response_type': ([User1],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/users/{user_id}/followers',
                'operation_id': 'get_user_followers_users_user_id_followers_get',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'user_id',
                ],
                'required': [
                    'x_tenant',
                    'user_id',
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
                    'user_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'user_id': 'user_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'user_id': 'path',
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
        self.get_user_following_users_user_id_following_get_endpoint = _Endpoint(
            settings={
                'response_type': ([bool, date, datetime, dict, float, int, list, str, none_type],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/users/{user_id}/following',
                'operation_id': 'get_user_following_users_user_id_following_get',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'user_id',
                ],
                'required': [
                    'x_tenant',
                    'user_id',
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
                    'user_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'x-tenant',
                    'user_id': 'user_id',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'user_id': 'path',
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

    def create_project_followers_projects_project_id_followers_post(
        self,
        x_tenant,
        project_id,
        **kwargs
    ):
        """Create Project Followers  # noqa: E501

        Add an activity in the repository.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_project_followers_projects_project_id_followers_post(x_tenant, project_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            project_id (str):

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
            Follow
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
        kwargs['project_id'] = \
            project_id
        return self.create_project_followers_projects_project_id_followers_post_endpoint.call_with_http_info(**kwargs)

    def create_user_followers_users_user_id_followers_post(
        self,
        x_tenant,
        user_id,
        **kwargs
    ):
        """Create User Followers  # noqa: E501

        Add an activity in the repository.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_user_followers_users_user_id_followers_post(x_tenant, user_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            user_id (str):

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
            Follow
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
        kwargs['user_id'] = \
            user_id
        return self.create_user_followers_users_user_id_followers_post_endpoint.call_with_http_info(**kwargs)

    def delete_project_followers_projects_project_id_followers_delete(
        self,
        x_tenant,
        project_id,
        **kwargs
    ):
        """Delete Following User  # noqa: E501

        Add an activity in the repository.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.delete_project_followers_projects_project_id_followers_delete(x_tenant, project_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            project_id (str):

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
        kwargs['project_id'] = \
            project_id
        return self.delete_project_followers_projects_project_id_followers_delete_endpoint.call_with_http_info(**kwargs)

    def delete_user_followers_users_user_id_followers_delete(
        self,
        x_tenant,
        user_id,
        **kwargs
    ):
        """Delete Following User  # noqa: E501

        Add an activity in the repository.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.delete_user_followers_users_user_id_followers_delete(x_tenant, user_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            user_id (str):

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
        kwargs['user_id'] = \
            user_id
        return self.delete_user_followers_users_user_id_followers_delete_endpoint.call_with_http_info(**kwargs)

    def get_project_followers_projects_project_id_followers_get(
        self,
        x_tenant,
        project_id,
        **kwargs
    ):
        """Get Project Followers  # noqa: E501

        Route to GET user following.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_project_followers_projects_project_id_followers_get(x_tenant, project_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            project_id (str):

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
            [User1]
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
        kwargs['project_id'] = \
            project_id
        return self.get_project_followers_projects_project_id_followers_get_endpoint.call_with_http_info(**kwargs)

    def get_user_followers_users_user_id_followers_get(
        self,
        x_tenant,
        user_id,
        **kwargs
    ):
        """Get User Followers  # noqa: E501

        Route to GET user followers.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_user_followers_users_user_id_followers_get(x_tenant, user_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            user_id (str):

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
            [User1]
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
        kwargs['user_id'] = \
            user_id
        return self.get_user_followers_users_user_id_followers_get_endpoint.call_with_http_info(**kwargs)

    def get_user_following_users_user_id_following_get(
        self,
        x_tenant,
        user_id,
        **kwargs
    ):
        """Get User Following  # noqa: E501

        Route to GET user following.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_user_following_users_user_id_following_get(x_tenant, user_id, async_req=True)
        >>> result = thread.get()

        Args:
            x_tenant (str):
            user_id (str):

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
            [bool, date, datetime, dict, float, int, list, str, none_type]
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
        kwargs['user_id'] = \
            user_id
        return self.get_user_following_users_user_id_following_get_endpoint.call_with_http_info(**kwargs)


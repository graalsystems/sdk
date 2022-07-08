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
from graalsystems.model.cost_stats import CostStats
from graalsystems.model.error import Error
from graalsystems.model.pageable import Pageable
from graalsystems.model.patch import Patch
from graalsystems.model.run_stats import RunStats
from graalsystems.model.workflow import Workflow
from graalsystems.model.workflow_page import WorkflowPage
from graalsystems.model.workflow_run import WorkflowRun


class WorkflowApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

        def __delete_workflow_by_id(
            self,
            x_tenant,
            workflow_id,
            **kwargs
        ):
            """Delete a workflow by its id  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.delete_workflow_by_id(x_tenant, workflow_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                workflow_id (str): Id of the workflow

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            kwargs['workflow_id'] = \
                workflow_id
            return self.call_with_http_info(**kwargs)

        self.delete_workflow_by_id = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows/{workflowId}',
                'operation_id': 'delete_workflow_by_id',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'workflow_id',
                ],
                'required': [
                    'x_tenant',
                    'workflow_id',
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
                    'workflow_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'workflow_id': 'workflowId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'workflow_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__delete_workflow_by_id
        )

        def __find_run_by_workflow_id_and_run_id(
            self,
            x_tenant,
            workflow_id,
            run_id,
            **kwargs
        ):
            """Find the run by a workflowId and a runId  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.find_run_by_workflow_id_and_run_id(x_tenant, workflow_id, run_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                workflow_id (str): Id of the Workflow
                run_id (str): Id of the Run

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                WorkflowRun
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            kwargs['workflow_id'] = \
                workflow_id
            kwargs['run_id'] = \
                run_id
            return self.call_with_http_info(**kwargs)

        self.find_run_by_workflow_id_and_run_id = _Endpoint(
            settings={
                'response_type': (WorkflowRun,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows/{workflowId}/runs/{runId}',
                'operation_id': 'find_run_by_workflow_id_and_run_id',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'workflow_id',
                    'run_id',
                ],
                'required': [
                    'x_tenant',
                    'workflow_id',
                    'run_id',
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
                    'workflow_id':
                        (str,),
                    'run_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'workflow_id': 'workflowId',
                    'run_id': 'runId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'workflow_id': 'path',
                    'run_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.workflowrun+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__find_run_by_workflow_id_and_run_id
        )

        def __find_workflow_by_id(
            self,
            x_tenant,
            workflow_id,
            **kwargs
        ):
            """Find workflow by Id  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.find_workflow_by_id(x_tenant, workflow_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                workflow_id (str): Id of the workflow

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Workflow
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            kwargs['workflow_id'] = \
                workflow_id
            return self.call_with_http_info(**kwargs)

        self.find_workflow_by_id = _Endpoint(
            settings={
                'response_type': (Workflow,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows/{workflowId}',
                'operation_id': 'find_workflow_by_id',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'workflow_id',
                ],
                'required': [
                    'x_tenant',
                    'workflow_id',
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
                    'workflow_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'workflow_id': 'workflowId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'workflow_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.workflow+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__find_workflow_by_id
        )

        def __find_workflows(
            self,
            x_tenant,
            **kwargs
        ):
            """Retrieve all workflows  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.find_workflows(x_tenant, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):

            Keyword Args:
                pageable (Pageable): [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                WorkflowPage
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            return self.call_with_http_info(**kwargs)

        self.find_workflows = _Endpoint(
            settings={
                'response_type': (WorkflowPage,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows',
                'operation_id': 'find_workflows',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'pageable',
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
                    'pageable':
                        (Pageable,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'pageable': 'pageable',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'pageable': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.workflow+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__find_workflows
        )

        def __get_costs_by_workflow_id(
            self,
            x_tenant,
            workflow_id,
            **kwargs
        ):
            """Find the costs for a workflow  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_costs_by_workflow_id(x_tenant, workflow_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                workflow_id (str): Id of the Workflow

            Keyword Args:
                _from (datetime): [optional]
                to (datetime): [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                CostStats
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            kwargs['workflow_id'] = \
                workflow_id
            return self.call_with_http_info(**kwargs)

        self.get_costs_by_workflow_id = _Endpoint(
            settings={
                'response_type': (CostStats,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows/{workflowId}/costs',
                'operation_id': 'get_costs_by_workflow_id',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'workflow_id',
                    '_from',
                    'to',
                ],
                'required': [
                    'x_tenant',
                    'workflow_id',
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
                    'workflow_id':
                        (str,),
                    '_from':
                        (datetime,),
                    'to':
                        (datetime,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'workflow_id': 'workflowId',
                    '_from': 'from',
                    'to': 'to',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'workflow_id': 'path',
                    '_from': 'query',
                    'to': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.costs+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_costs_by_workflow_id
        )

        def __get_stats_by_workflow_id(
            self,
            x_tenant,
            workflow_id,
            **kwargs
        ):
            """Find the stats for a workflow  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_stats_by_workflow_id(x_tenant, workflow_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                workflow_id (str): Id of the Workflow

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                RunStats
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            kwargs['workflow_id'] = \
                workflow_id
            return self.call_with_http_info(**kwargs)

        self.get_stats_by_workflow_id = _Endpoint(
            settings={
                'response_type': (RunStats,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows/{workflowId}/stats',
                'operation_id': 'get_stats_by_workflow_id',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'workflow_id',
                ],
                'required': [
                    'x_tenant',
                    'workflow_id',
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
                    'workflow_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'workflow_id': 'workflowId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'workflow_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.stats+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_stats_by_workflow_id
        )

        def __update_workflow(
            self,
            x_tenant,
            workflow_id,
            patch,
            **kwargs
        ):
            """Update a workflow  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.update_workflow(x_tenant, workflow_id, patch, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                workflow_id (str): Id of the workflow
                patch ([Patch]): The patch

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Workflow
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['x_tenant'] = \
                x_tenant
            kwargs['workflow_id'] = \
                workflow_id
            kwargs['patch'] = \
                patch
            return self.call_with_http_info(**kwargs)

        self.update_workflow = _Endpoint(
            settings={
                'response_type': (Workflow,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/workflows/{workflowId}',
                'operation_id': 'update_workflow',
                'http_method': 'PATCH',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'workflow_id',
                    'patch',
                ],
                'required': [
                    'x_tenant',
                    'workflow_id',
                    'patch',
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
                    'workflow_id':
                        (str,),
                    'patch':
                        ([Patch],),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'workflow_id': 'workflowId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'workflow_id': 'path',
                    'patch': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.workflow+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [
                    'application/json-patch+json;charset=UTF-8'
                ]
            },
            api_client=api_client,
            callable=__update_workflow
        )

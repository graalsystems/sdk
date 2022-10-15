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
from graalsystems.model.error import Error
from graalsystems.model.job import Job
from graalsystems.model.metric import Metric
from graalsystems.model.patch import Patch
from graalsystems.model.run_stats import RunStats


class JobApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

        def __delete_job_by_id(
            self,
            x_tenant,
            job_id,
            **kwargs
        ):
            """Delete a job by its id  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.delete_job_by_id(x_tenant, job_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                job_id (str): Id of the Job

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
            kwargs['job_id'] = \
                job_id
            return self.call_with_http_info(**kwargs)

        self.delete_job_by_id = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/jobs/{jobId}',
                'operation_id': 'delete_job_by_id',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'job_id',
                ],
                'required': [
                    'x_tenant',
                    'job_id',
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
                    'job_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'job_id': 'jobId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'job_id': 'path',
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
            callable=__delete_job_by_id
        )

        def __find_job_by_job_id(
            self,
            x_tenant,
            job_id,
            **kwargs
        ):
            """Find job by Id  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.find_job_by_job_id(x_tenant, job_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                job_id (str): Id of the Job

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
                Job
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
            kwargs['job_id'] = \
                job_id
            return self.call_with_http_info(**kwargs)

        self.find_job_by_job_id = _Endpoint(
            settings={
                'response_type': (Job,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/jobs/{jobId}',
                'operation_id': 'find_job_by_job_id',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'job_id',
                ],
                'required': [
                    'x_tenant',
                    'job_id',
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
                    'job_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'job_id': 'jobId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'job_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.job+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__find_job_by_job_id
        )

        def __find_jobs(
            self,
            x_tenant,
            **kwargs
        ):
            """Retrieve all jobs  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.find_jobs(x_tenant, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):

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
                [Job]
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

        self.find_jobs = _Endpoint(
            settings={
                'response_type': ([Job],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/jobs',
                'operation_id': 'find_jobs',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
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
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                },
                'location_map': {
                    'x_tenant': 'header',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.job+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__find_jobs
        )

        def __get_metrics_for_job(
            self,
            x_tenant,
            job_id,
            metric,
            **kwargs
        ):
            """Find the metrics for a job  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_metrics_for_job(x_tenant, job_id, metric, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                job_id (str): Id of the Job
                metric (str): The metric name

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
                [Metric]
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
            kwargs['job_id'] = \
                job_id
            kwargs['metric'] = \
                metric
            return self.call_with_http_info(**kwargs)

        self.get_metrics_for_job = _Endpoint(
            settings={
                'response_type': ([Metric],),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/jobs/{jobId}/metrics',
                'operation_id': 'get_metrics_for_job',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'job_id',
                    'metric',
                    '_from',
                    'to',
                ],
                'required': [
                    'x_tenant',
                    'job_id',
                    'metric',
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
                    'job_id':
                        (str,),
                    'metric':
                        (str,),
                    '_from':
                        (datetime,),
                    'to':
                        (datetime,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'job_id': 'jobId',
                    'metric': 'metric',
                    '_from': 'from',
                    'to': 'to',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'job_id': 'path',
                    'metric': 'query',
                    '_from': 'query',
                    'to': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_metrics_for_job
        )

        def __get_stats_by_job_id(
            self,
            x_tenant,
            job_id,
            **kwargs
        ):
            """Find the stats for a job  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_stats_by_job_id(x_tenant, job_id, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                job_id (str): Id of the Job

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
            kwargs['job_id'] = \
                job_id
            return self.call_with_http_info(**kwargs)

        self.get_stats_by_job_id = _Endpoint(
            settings={
                'response_type': (RunStats,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/jobs/{jobId}/stats',
                'operation_id': 'get_stats_by_job_id',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'job_id',
                ],
                'required': [
                    'x_tenant',
                    'job_id',
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
                    'job_id':
                        (str,),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'job_id': 'jobId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'job_id': 'path',
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
            callable=__get_stats_by_job_id
        )

        def __update_job(
            self,
            x_tenant,
            job_id,
            patch,
            **kwargs
        ):
            """Update a job  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.update_job(x_tenant, job_id, patch, async_req=True)
            >>> result = thread.get()

            Args:
                x_tenant (str):
                job_id (str): Id of the Job
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
                Job
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
            kwargs['job_id'] = \
                job_id
            kwargs['patch'] = \
                patch
            return self.call_with_http_info(**kwargs)

        self.update_job = _Endpoint(
            settings={
                'response_type': (Job,),
                'auth': [
                    'internal'
                ],
                'endpoint_path': '/jobs/{jobId}',
                'operation_id': 'update_job',
                'http_method': 'PATCH',
                'servers': None,
            },
            params_map={
                'all': [
                    'x_tenant',
                    'job_id',
                    'patch',
                ],
                'required': [
                    'x_tenant',
                    'job_id',
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
                    'job_id':
                        (str,),
                    'patch':
                        ([Patch],),
                },
                'attribute_map': {
                    'x_tenant': 'X-Tenant',
                    'job_id': 'jobId',
                },
                'location_map': {
                    'x_tenant': 'header',
                    'job_id': 'path',
                    'patch': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/vnd.graal.systems.v1.job+json',
                    'application/vnd.graal.systems.v1.error+json'
                ],
                'content_type': [
                    'application/json-patch+json;charset=UTF-8'
                ]
            },
            api_client=api_client,
            callable=__update_job
        )
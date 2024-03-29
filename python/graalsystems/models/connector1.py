# coding: utf-8

"""
    Tenant API

    Tenant API

    The version of the OpenAPI document: 0.0.1
    Contact: abc@layer.fr
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
from inspect import getfullargspec
import json
import pprint
import re  # noqa: F401

from typing import Any, List, Optional
from pydantic import BaseModel, Field, StrictStr, ValidationError, field_validator
from graalsystems.models.azure_connector import AzureConnector
from graalsystems.models.local_connector import LocalConnector
from graalsystems.models.s3_connector import S3Connector
from typing import Union, Any, List, TYPE_CHECKING, Optional, Dict
from typing_extensions import Literal
from pydantic import StrictStr, Field
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

CONNECTOR1_ONE_OF_SCHEMAS = ["AzureConnector", "LocalConnector", "S3Connector"]

class Connector1(BaseModel):
    """
    Type of connector (local, S3, Azure).
    """
    # data type: S3Connector
    oneof_schema_1_validator: Optional[S3Connector] = None
    # data type: AzureConnector
    oneof_schema_2_validator: Optional[AzureConnector] = None
    # data type: LocalConnector
    oneof_schema_3_validator: Optional[LocalConnector] = None
    actual_instance: Optional[Union[AzureConnector, LocalConnector, S3Connector]] = None
    one_of_schemas: List[str] = Literal["AzureConnector", "LocalConnector", "S3Connector"]

    model_config = {
        "validate_assignment": True
    }


    discriminator_value_class_map: Dict[str, str] = {
    }

    def __init__(self, *args, **kwargs) -> None:
        if args:
            if len(args) > 1:
                raise ValueError("If a position argument is used, only 1 is allowed to set `actual_instance`")
            if kwargs:
                raise ValueError("If a position argument is used, keyword arguments cannot be used.")
            super().__init__(actual_instance=args[0])
        else:
            super().__init__(**kwargs)

    @field_validator('actual_instance')
    def actual_instance_must_validate_oneof(cls, v):
        instance = Connector1.model_construct()
        error_messages = []
        match = 0
        # validate data type: S3Connector
        if not isinstance(v, S3Connector):
            error_messages.append(f"Error! Input type `{type(v)}` is not `S3Connector`")
        else:
            match += 1
        # validate data type: AzureConnector
        if not isinstance(v, AzureConnector):
            error_messages.append(f"Error! Input type `{type(v)}` is not `AzureConnector`")
        else:
            match += 1
        # validate data type: LocalConnector
        if not isinstance(v, LocalConnector):
            error_messages.append(f"Error! Input type `{type(v)}` is not `LocalConnector`")
        else:
            match += 1
        if match > 1:
            # more than 1 match
            raise ValueError("Multiple matches found when setting `actual_instance` in Connector1 with oneOf schemas: AzureConnector, LocalConnector, S3Connector. Details: " + ", ".join(error_messages))
        elif match == 0:
            # no match
            raise ValueError("No match found when setting `actual_instance` in Connector1 with oneOf schemas: AzureConnector, LocalConnector, S3Connector. Details: " + ", ".join(error_messages))
        else:
            return v

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        return cls.from_json(json.dumps(obj))

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Returns the object represented by the json string"""
        instance = cls.model_construct()
        error_messages = []
        match = 0

        # deserialize data into S3Connector
        try:
            instance.actual_instance = S3Connector.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into AzureConnector
        try:
            instance.actual_instance = AzureConnector.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into LocalConnector
        try:
            instance.actual_instance = LocalConnector.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))

        if match > 1:
            # more than 1 match
            raise ValueError("Multiple matches found when deserializing the JSON string into Connector1 with oneOf schemas: AzureConnector, LocalConnector, S3Connector. Details: " + ", ".join(error_messages))
        elif match == 0:
            # no match
            raise ValueError("No match found when deserializing the JSON string into Connector1 with oneOf schemas: AzureConnector, LocalConnector, S3Connector. Details: " + ", ".join(error_messages))
        else:
            return instance

    def to_json(self) -> str:
        """Returns the JSON representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_json()
        else:
            return json.dumps(self.actual_instance)

    def to_dict(self) -> dict:
        """Returns the dict representation of the actual instance"""
        if self.actual_instance is None:
            return None

        to_dict = getattr(self.actual_instance, "to_dict", None)
        if callable(to_dict):
            return self.actual_instance.to_dict()
        else:
            # primitive type
            return self.actual_instance

    def to_str(self) -> str:
        """Returns the string representation of the actual instance"""
        return pprint.pformat(self.model_dump())



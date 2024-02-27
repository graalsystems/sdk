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
import pprint
import re  # noqa: F401
import json
import importlib


from typing import Optional, Union
from pydantic import StrictFloat, StrictInt, StrictStr
from pydantic import Field
from graalsystems.models.data_warehouse_options import DataWarehouseOptions
from typing import Dict, Any
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class DB2Options(DataWarehouseOptions):
    """
    DB2Options
    """
    host: Optional[StrictStr] = None
    port: Optional[Union[StrictFloat, StrictInt]] = None
    database: Optional[StrictStr] = None
    var_schema: Optional[StrictStr] = Field(default=None, alias="schema")
    type: Optional[StrictStr] = 'db2'
    __properties: ClassVar[List[str]] = ["type", "host", "port", "database", "schema"]

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of DB2Options from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
            },
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        """Create an instance of DB2Options from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "type": obj.get("type") if obj.get("type") is not None else 'db2',
            "host": obj.get("host"),
            "port": obj.get("port"),
            "database": obj.get("database"),
            "schema": obj.get("schema")
        })
        return _obj



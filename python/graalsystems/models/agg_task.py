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


from typing import List, Optional
from pydantic import BaseModel, StrictStr, field_validator
from pydantic import Field
from graalsystems.models.agg_params import AggParams
from typing import Dict, Any
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class AggTask(BaseModel):
    """
    Defines a task in a DAG.  Attributes ---------- id : str     Identifier of a task. params : Params     Parameters of a task.  Methods ------- accept(visitor)     Visit a task using a specified visitor. to_node()     Returns the information about the task (id and parameters). to_edge()     Gets all the dependencies of the task.  # noqa: E501
    """
    id: StrictStr = Field(description="Identifier of the task.")
    params: AggParams
    dependency: List[StrictStr] = Field(description="List of all dependencies of the task.")
    type: Optional[StrictStr] = Field(default='aggregation', description="Type of the aggregation task.")
    __properties: ClassVar[List[str]] = ["id", "params", "dependency", "type"]

    @field_validator('type')
    def type_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('aggregation', 'distinct', 'drop', 'filter', 'join', 'order', 'read', 'rename', 'train_model', 'write', 'predict_model', 'print'):
            raise ValueError("must be one of enum values ('aggregation', 'distinct', 'drop', 'filter', 'join', 'order', 'read', 'rename', 'train_model', 'write', 'predict_model', 'print')")
        return value

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
        """Create an instance of AggTask from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of params
        if self.params:
            _dict['params'] = self.params.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        """Create an instance of AggTask from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "id": obj.get("id"),
            "params": AggParams.from_dict(obj.get("params")) if obj.get("params") is not None else None,
            "dependency": obj.get("dependency"),
            "type": obj.get("type") if obj.get("type") is not None else 'aggregation'
        })
        return _obj



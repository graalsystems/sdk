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

from datetime import datetime
from typing import Any, Dict, Optional, Union
from pydantic import BaseModel, StrictBool, StrictStr
from graalsystems.models.details1 import Details1
from typing import Dict, Any
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class PaymentMethod(BaseModel):
    """
    PaymentMethod
    """
    id: Optional[StrictStr] = None
    details: Optional[Details1] = None
    provider: Optional[StrictStr] = None
    status: Optional[StrictStr] = None
    remote_payment_method_id: Optional[StrictStr] = None
    favorite: Optional[StrictBool] = None
    metadata: Optional[Dict[str, Union[str, Any]]] = None
    created: Optional[datetime] = None
    updated: Optional[datetime] = None
    __properties: ClassVar[List[str]] = ["id", "details", "provider", "status", "remote_payment_method_id", "favorite", "metadata", "created", "updated"]

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
        """Create an instance of PaymentMethod from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of details
        if self.details:
            _dict['details'] = self.details.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        """Create an instance of PaymentMethod from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "id": obj.get("id"),
            "details": Details1.from_dict(obj.get("details")) if obj.get("details") is not None else None,
            "provider": obj.get("provider"),
            "status": obj.get("status"),
            "remote_payment_method_id": obj.get("remote_payment_method_id"),
            "favorite": obj.get("favorite"),
            "metadata": obj.get("metadata"),
            "created": obj.get("created"),
            "updated": obj.get("updated")
        })
        return _obj


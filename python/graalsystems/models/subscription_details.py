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


from typing import Optional
from pydantic import StrictStr, field_validator
from graalsystems.models.company import Company
from graalsystems.models.contact import Contact
from graalsystems.models.details import Details
from graalsystems.models.payment_method import PaymentMethod
from typing import Dict, Any
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class SubscriptionDetails(Details):
    """
    SubscriptionDetails
    """
    type: Optional[StrictStr] = 'subscription'
    support_plan: Optional[StrictStr] = None
    account: Optional[StrictStr] = None
    company: Optional[Company] = None
    contact: Optional[Contact] = None
    payment_method: Optional[PaymentMethod] = None
    __properties: ClassVar[List[str]] = ["type", "support_plan", "account", "company", "contact", "payment_method"]

    @field_validator('support_plan')
    def support_plan_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('basic', 'developer', 'business', 'enterprise'):
            raise ValueError("must be one of enum values ('basic', 'developer', 'business', 'enterprise')")
        return value

    @field_validator('account')
    def account_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('personal', 'enterprise', 'gouv'):
            raise ValueError("must be one of enum values ('personal', 'enterprise', 'gouv')")
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
        """Create an instance of SubscriptionDetails from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of company
        if self.company:
            _dict['company'] = self.company.to_dict()
        # override the default output from pydantic by calling `to_dict()` of contact
        if self.contact:
            _dict['contact'] = self.contact.to_dict()
        # override the default output from pydantic by calling `to_dict()` of payment_method
        if self.payment_method:
            _dict['payment_method'] = self.payment_method.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        """Create an instance of SubscriptionDetails from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "type": obj.get("type") if obj.get("type") is not None else 'subscription',
            "support_plan": obj.get("support_plan"),
            "account": obj.get("account"),
            "company": Company.from_dict(obj.get("company")) if obj.get("company") is not None else None,
            "contact": Contact.from_dict(obj.get("contact")) if obj.get("contact") is not None else None,
            "payment_method": PaymentMethod.from_dict(obj.get("payment_method")) if obj.get("payment_method") is not None else None
        })
        return _obj



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

from typing import Optional
from pydantic import BaseModel, Field, StrictStr, ValidationError, field_validator
from graalsystems.models.and_filter import AndFilter
from graalsystems.models.between_filter import BetweenFilter
from graalsystems.models.equal_filter import EqualFilter
from graalsystems.models.greater_or_equal_than_filter import GreaterOrEqualThanFilter
from graalsystems.models.greater_than_filter import GreaterThanFilter
from graalsystems.models.in_filter import InFilter
from graalsystems.models.less_or_equal_than_filter import LessOrEqualThanFilter
from graalsystems.models.less_than_filter import LessThanFilter
from graalsystems.models.like_filter import LikeFilter
from graalsystems.models.not_equal_filter import NotEqualFilter
from graalsystems.models.not_filter import NotFilter
from graalsystems.models.not_null_filter import NotNullFilter
from graalsystems.models.or_filter import OrFilter
from typing import Union, Any, List, TYPE_CHECKING, Optional, Dict
from typing_extensions import Literal
from pydantic import StrictStr, Field
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

FILTERTASKPARAMETERS_ANY_OF_SCHEMAS = ["AndFilter", "BetweenFilter", "EqualFilter", "GreaterOrEqualThanFilter", "GreaterThanFilter", "InFilter", "LessOrEqualThanFilter", "LessThanFilter", "LikeFilter", "NotEqualFilter", "NotFilter", "NotNullFilter", "OrFilter"]

class FilterTaskParameters(BaseModel):
    """
    Parameters of the filter task.
    """

    # data type: GreaterOrEqualThanFilter
    anyof_schema_1_validator: Optional[GreaterOrEqualThanFilter] = None
    # data type: LessOrEqualThanFilter
    anyof_schema_2_validator: Optional[LessOrEqualThanFilter] = None
    # data type: GreaterThanFilter
    anyof_schema_3_validator: Optional[GreaterThanFilter] = None
    # data type: NotEqualFilter
    anyof_schema_4_validator: Optional[NotEqualFilter] = None
    # data type: LessThanFilter
    anyof_schema_5_validator: Optional[LessThanFilter] = None
    # data type: NotNullFilter
    anyof_schema_6_validator: Optional[NotNullFilter] = None
    # data type: BetweenFilter
    anyof_schema_7_validator: Optional[BetweenFilter] = None
    # data type: EqualFilter
    anyof_schema_8_validator: Optional[EqualFilter] = None
    # data type: LikeFilter
    anyof_schema_9_validator: Optional[LikeFilter] = None
    # data type: NotFilter
    anyof_schema_10_validator: Optional[NotFilter] = None
    # data type: AndFilter
    anyof_schema_11_validator: Optional[AndFilter] = None
    # data type: InFilter
    anyof_schema_12_validator: Optional[InFilter] = None
    # data type: OrFilter
    anyof_schema_13_validator: Optional[OrFilter] = None
    if TYPE_CHECKING:
        actual_instance: Optional[Union[AndFilter, BetweenFilter, EqualFilter, GreaterOrEqualThanFilter, GreaterThanFilter, InFilter, LessOrEqualThanFilter, LessThanFilter, LikeFilter, NotEqualFilter, NotFilter, NotNullFilter, OrFilter]] = None
    else:
        actual_instance: Any = None
    any_of_schemas: List[str] = Literal[FILTERTASKPARAMETERS_ANY_OF_SCHEMAS]

    model_config = {
        "validate_assignment": True
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
    def actual_instance_must_validate_anyof(cls, v):
        instance = FilterTaskParameters.model_construct()
        error_messages = []
        # validate data type: GreaterOrEqualThanFilter
        if not isinstance(v, GreaterOrEqualThanFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `GreaterOrEqualThanFilter`")
        else:
            return v

        # validate data type: LessOrEqualThanFilter
        if not isinstance(v, LessOrEqualThanFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `LessOrEqualThanFilter`")
        else:
            return v

        # validate data type: GreaterThanFilter
        if not isinstance(v, GreaterThanFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `GreaterThanFilter`")
        else:
            return v

        # validate data type: NotEqualFilter
        if not isinstance(v, NotEqualFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `NotEqualFilter`")
        else:
            return v

        # validate data type: LessThanFilter
        if not isinstance(v, LessThanFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `LessThanFilter`")
        else:
            return v

        # validate data type: NotNullFilter
        if not isinstance(v, NotNullFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `NotNullFilter`")
        else:
            return v

        # validate data type: BetweenFilter
        if not isinstance(v, BetweenFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `BetweenFilter`")
        else:
            return v

        # validate data type: EqualFilter
        if not isinstance(v, EqualFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `EqualFilter`")
        else:
            return v

        # validate data type: LikeFilter
        if not isinstance(v, LikeFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `LikeFilter`")
        else:
            return v

        # validate data type: NotFilter
        if not isinstance(v, NotFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `NotFilter`")
        else:
            return v

        # validate data type: AndFilter
        if not isinstance(v, AndFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `AndFilter`")
        else:
            return v

        # validate data type: InFilter
        if not isinstance(v, InFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `InFilter`")
        else:
            return v

        # validate data type: OrFilter
        if not isinstance(v, OrFilter):
            error_messages.append(f"Error! Input type `{type(v)}` is not `OrFilter`")
        else:
            return v

        if error_messages:
            # no match
            raise ValueError("No match found when setting the actual_instance in FilterTaskParameters with anyOf schemas: AndFilter, BetweenFilter, EqualFilter, GreaterOrEqualThanFilter, GreaterThanFilter, InFilter, LessOrEqualThanFilter, LessThanFilter, LikeFilter, NotEqualFilter, NotFilter, NotNullFilter, OrFilter. Details: " + ", ".join(error_messages))
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
        # anyof_schema_1_validator: Optional[GreaterOrEqualThanFilter] = None
        try:
            instance.actual_instance = GreaterOrEqualThanFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_2_validator: Optional[LessOrEqualThanFilter] = None
        try:
            instance.actual_instance = LessOrEqualThanFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_3_validator: Optional[GreaterThanFilter] = None
        try:
            instance.actual_instance = GreaterThanFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_4_validator: Optional[NotEqualFilter] = None
        try:
            instance.actual_instance = NotEqualFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_5_validator: Optional[LessThanFilter] = None
        try:
            instance.actual_instance = LessThanFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_6_validator: Optional[NotNullFilter] = None
        try:
            instance.actual_instance = NotNullFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_7_validator: Optional[BetweenFilter] = None
        try:
            instance.actual_instance = BetweenFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_8_validator: Optional[EqualFilter] = None
        try:
            instance.actual_instance = EqualFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_9_validator: Optional[LikeFilter] = None
        try:
            instance.actual_instance = LikeFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_10_validator: Optional[NotFilter] = None
        try:
            instance.actual_instance = NotFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_11_validator: Optional[AndFilter] = None
        try:
            instance.actual_instance = AndFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_12_validator: Optional[InFilter] = None
        try:
            instance.actual_instance = InFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))
        # anyof_schema_13_validator: Optional[OrFilter] = None
        try:
            instance.actual_instance = OrFilter.from_json(json_str)
            return instance
        except (ValidationError, ValueError) as e:
             error_messages.append(str(e))

        if error_messages:
            # no match
            raise ValueError("No match found when deserializing the JSON string into FilterTaskParameters with anyOf schemas: AndFilter, BetweenFilter, EqualFilter, GreaterOrEqualThanFilter, GreaterThanFilter, InFilter, LessOrEqualThanFilter, LessThanFilter, LikeFilter, NotEqualFilter, NotFilter, NotNullFilter, OrFilter. Details: " + ", ".join(error_messages))
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
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_dict()
        else:
            return json.dumps(self.actual_instance)

    def to_str(self) -> str:
        """Returns the string representation of the actual instance"""
        return pprint.pformat(self.model_dump())



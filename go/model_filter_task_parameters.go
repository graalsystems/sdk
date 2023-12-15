/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// FilterTaskParameters Parameters of the filter task.
type FilterTaskParameters struct {
	AndFilter *AndFilter
	BetweenFilter *BetweenFilter
	EqualFilter *EqualFilter
	GreaterOrEqualThanFilter *GreaterOrEqualThanFilter
	GreaterThanFilter *GreaterThanFilter
	InFilter *InFilter
	LessOrEqualThanFilter *LessOrEqualThanFilter
	LessThanFilter *LessThanFilter
	LikeFilter *LikeFilter
	NotEqualFilter *NotEqualFilter
	NotFilter *NotFilter
	NotNullFilter *NotNullFilter
	OrFilter *OrFilter
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FilterTaskParameters) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AndFilter
	err = json.Unmarshal(data, &dst.AndFilter);
	if err == nil {
		jsonAndFilter, _ := json.Marshal(dst.AndFilter)
		if string(jsonAndFilter) == "{}" { // empty struct
			dst.AndFilter = nil
		} else {
			return nil // data stored in dst.AndFilter, return on the first match
		}
	} else {
		dst.AndFilter = nil
	}

	// try to unmarshal JSON data into BetweenFilter
	err = json.Unmarshal(data, &dst.BetweenFilter);
	if err == nil {
		jsonBetweenFilter, _ := json.Marshal(dst.BetweenFilter)
		if string(jsonBetweenFilter) == "{}" { // empty struct
			dst.BetweenFilter = nil
		} else {
			return nil // data stored in dst.BetweenFilter, return on the first match
		}
	} else {
		dst.BetweenFilter = nil
	}

	// try to unmarshal JSON data into EqualFilter
	err = json.Unmarshal(data, &dst.EqualFilter);
	if err == nil {
		jsonEqualFilter, _ := json.Marshal(dst.EqualFilter)
		if string(jsonEqualFilter) == "{}" { // empty struct
			dst.EqualFilter = nil
		} else {
			return nil // data stored in dst.EqualFilter, return on the first match
		}
	} else {
		dst.EqualFilter = nil
	}

	// try to unmarshal JSON data into GreaterOrEqualThanFilter
	err = json.Unmarshal(data, &dst.GreaterOrEqualThanFilter);
	if err == nil {
		jsonGreaterOrEqualThanFilter, _ := json.Marshal(dst.GreaterOrEqualThanFilter)
		if string(jsonGreaterOrEqualThanFilter) == "{}" { // empty struct
			dst.GreaterOrEqualThanFilter = nil
		} else {
			return nil // data stored in dst.GreaterOrEqualThanFilter, return on the first match
		}
	} else {
		dst.GreaterOrEqualThanFilter = nil
	}

	// try to unmarshal JSON data into GreaterThanFilter
	err = json.Unmarshal(data, &dst.GreaterThanFilter);
	if err == nil {
		jsonGreaterThanFilter, _ := json.Marshal(dst.GreaterThanFilter)
		if string(jsonGreaterThanFilter) == "{}" { // empty struct
			dst.GreaterThanFilter = nil
		} else {
			return nil // data stored in dst.GreaterThanFilter, return on the first match
		}
	} else {
		dst.GreaterThanFilter = nil
	}

	// try to unmarshal JSON data into InFilter
	err = json.Unmarshal(data, &dst.InFilter);
	if err == nil {
		jsonInFilter, _ := json.Marshal(dst.InFilter)
		if string(jsonInFilter) == "{}" { // empty struct
			dst.InFilter = nil
		} else {
			return nil // data stored in dst.InFilter, return on the first match
		}
	} else {
		dst.InFilter = nil
	}

	// try to unmarshal JSON data into LessOrEqualThanFilter
	err = json.Unmarshal(data, &dst.LessOrEqualThanFilter);
	if err == nil {
		jsonLessOrEqualThanFilter, _ := json.Marshal(dst.LessOrEqualThanFilter)
		if string(jsonLessOrEqualThanFilter) == "{}" { // empty struct
			dst.LessOrEqualThanFilter = nil
		} else {
			return nil // data stored in dst.LessOrEqualThanFilter, return on the first match
		}
	} else {
		dst.LessOrEqualThanFilter = nil
	}

	// try to unmarshal JSON data into LessThanFilter
	err = json.Unmarshal(data, &dst.LessThanFilter);
	if err == nil {
		jsonLessThanFilter, _ := json.Marshal(dst.LessThanFilter)
		if string(jsonLessThanFilter) == "{}" { // empty struct
			dst.LessThanFilter = nil
		} else {
			return nil // data stored in dst.LessThanFilter, return on the first match
		}
	} else {
		dst.LessThanFilter = nil
	}

	// try to unmarshal JSON data into LikeFilter
	err = json.Unmarshal(data, &dst.LikeFilter);
	if err == nil {
		jsonLikeFilter, _ := json.Marshal(dst.LikeFilter)
		if string(jsonLikeFilter) == "{}" { // empty struct
			dst.LikeFilter = nil
		} else {
			return nil // data stored in dst.LikeFilter, return on the first match
		}
	} else {
		dst.LikeFilter = nil
	}

	// try to unmarshal JSON data into NotEqualFilter
	err = json.Unmarshal(data, &dst.NotEqualFilter);
	if err == nil {
		jsonNotEqualFilter, _ := json.Marshal(dst.NotEqualFilter)
		if string(jsonNotEqualFilter) == "{}" { // empty struct
			dst.NotEqualFilter = nil
		} else {
			return nil // data stored in dst.NotEqualFilter, return on the first match
		}
	} else {
		dst.NotEqualFilter = nil
	}

	// try to unmarshal JSON data into NotFilter
	err = json.Unmarshal(data, &dst.NotFilter);
	if err == nil {
		jsonNotFilter, _ := json.Marshal(dst.NotFilter)
		if string(jsonNotFilter) == "{}" { // empty struct
			dst.NotFilter = nil
		} else {
			return nil // data stored in dst.NotFilter, return on the first match
		}
	} else {
		dst.NotFilter = nil
	}

	// try to unmarshal JSON data into NotNullFilter
	err = json.Unmarshal(data, &dst.NotNullFilter);
	if err == nil {
		jsonNotNullFilter, _ := json.Marshal(dst.NotNullFilter)
		if string(jsonNotNullFilter) == "{}" { // empty struct
			dst.NotNullFilter = nil
		} else {
			return nil // data stored in dst.NotNullFilter, return on the first match
		}
	} else {
		dst.NotNullFilter = nil
	}

	// try to unmarshal JSON data into OrFilter
	err = json.Unmarshal(data, &dst.OrFilter);
	if err == nil {
		jsonOrFilter, _ := json.Marshal(dst.OrFilter)
		if string(jsonOrFilter) == "{}" { // empty struct
			dst.OrFilter = nil
		} else {
			return nil // data stored in dst.OrFilter, return on the first match
		}
	} else {
		dst.OrFilter = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(FilterTaskParameters)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FilterTaskParameters) MarshalJSON() ([]byte, error) {
	if src.AndFilter != nil {
		return json.Marshal(&src.AndFilter)
	}

	if src.BetweenFilter != nil {
		return json.Marshal(&src.BetweenFilter)
	}

	if src.EqualFilter != nil {
		return json.Marshal(&src.EqualFilter)
	}

	if src.GreaterOrEqualThanFilter != nil {
		return json.Marshal(&src.GreaterOrEqualThanFilter)
	}

	if src.GreaterThanFilter != nil {
		return json.Marshal(&src.GreaterThanFilter)
	}

	if src.InFilter != nil {
		return json.Marshal(&src.InFilter)
	}

	if src.LessOrEqualThanFilter != nil {
		return json.Marshal(&src.LessOrEqualThanFilter)
	}

	if src.LessThanFilter != nil {
		return json.Marshal(&src.LessThanFilter)
	}

	if src.LikeFilter != nil {
		return json.Marshal(&src.LikeFilter)
	}

	if src.NotEqualFilter != nil {
		return json.Marshal(&src.NotEqualFilter)
	}

	if src.NotFilter != nil {
		return json.Marshal(&src.NotFilter)
	}

	if src.NotNullFilter != nil {
		return json.Marshal(&src.NotNullFilter)
	}

	if src.OrFilter != nil {
		return json.Marshal(&src.OrFilter)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFilterTaskParameters struct {
	value *FilterTaskParameters
	isSet bool
}

func (v NullableFilterTaskParameters) Get() *FilterTaskParameters {
	return v.value
}

func (v *NullableFilterTaskParameters) Set(val *FilterTaskParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterTaskParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterTaskParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterTaskParameters(val *FilterTaskParameters) *NullableFilterTaskParameters {
	return &NullableFilterTaskParameters{value: val, isSet: true}
}

func (v NullableFilterTaskParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterTaskParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



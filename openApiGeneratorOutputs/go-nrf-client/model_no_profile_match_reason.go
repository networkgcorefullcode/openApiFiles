/*
NRF NFManagement Service

NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// NoProfileMatchReason No Profile Match Reason
type NoProfileMatchReason struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NoProfileMatchReason) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NoProfileMatchReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NoProfileMatchReason) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNoProfileMatchReason struct {
	value *NoProfileMatchReason
	isSet bool
}

func (v NullableNoProfileMatchReason) Get() *NoProfileMatchReason {
	return v.value
}

func (v *NullableNoProfileMatchReason) Set(val *NoProfileMatchReason) {
	v.value = val
	v.isSet = true
}

func (v NullableNoProfileMatchReason) IsSet() bool {
	return v.isSet
}

func (v *NullableNoProfileMatchReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoProfileMatchReason(val *NoProfileMatchReason) *NullableNoProfileMatchReason {
	return &NullableNoProfileMatchReason{value: val, isSet: true}
}

func (v NullableNoProfileMatchReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoProfileMatchReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
NRF NFManagement Service

NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PfdData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PfdData{}

// PfdData List of Application IDs and/or AF IDs managed by a given NEF Instance
type PfdData struct {
	AppIds []string `json:"appIds,omitempty"`
	AfIds []string `json:"afIds,omitempty"`
}

// NewPfdData instantiates a new PfdData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPfdData() *PfdData {
	this := PfdData{}
	return &this
}

// NewPfdDataWithDefaults instantiates a new PfdData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPfdDataWithDefaults() *PfdData {
	this := PfdData{}
	return &this
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *PfdData) GetAppIds() []string {
	if o == nil || IsNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdData) GetAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *PfdData) HasAppIds() bool {
	if o != nil && !IsNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *PfdData) SetAppIds(v []string) {
	o.AppIds = v
}

// GetAfIds returns the AfIds field value if set, zero value otherwise.
func (o *PfdData) GetAfIds() []string {
	if o == nil || IsNil(o.AfIds) {
		var ret []string
		return ret
	}
	return o.AfIds
}

// GetAfIdsOk returns a tuple with the AfIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdData) GetAfIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AfIds) {
		return nil, false
	}
	return o.AfIds, true
}

// HasAfIds returns a boolean if a field has been set.
func (o *PfdData) HasAfIds() bool {
	if o != nil && !IsNil(o.AfIds) {
		return true
	}

	return false
}

// SetAfIds gets a reference to the given []string and assigns it to the AfIds field.
func (o *PfdData) SetAfIds(v []string) {
	o.AfIds = v
}

func (o PfdData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PfdData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !IsNil(o.AfIds) {
		toSerialize["afIds"] = o.AfIds
	}
	return toSerialize, nil
}

type NullablePfdData struct {
	value *PfdData
	isSet bool
}

func (v NullablePfdData) Get() *PfdData {
	return v.value
}

func (v *NullablePfdData) Set(val *PfdData) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdData) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdData(val *PfdData) *NullablePfdData {
	return &NullablePfdData{value: val, isSet: true}
}

func (v NullablePfdData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



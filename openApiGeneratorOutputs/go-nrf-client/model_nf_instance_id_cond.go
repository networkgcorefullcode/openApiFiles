/*
NRF NFManagement Service

NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NfInstanceIdCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NfInstanceIdCond{}

// NfInstanceIdCond Subscription to a given NF Instance Id
type NfInstanceIdCond struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`
}

type _NfInstanceIdCond NfInstanceIdCond

// NewNfInstanceIdCond instantiates a new NfInstanceIdCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfInstanceIdCond(nfInstanceId string) *NfInstanceIdCond {
	this := NfInstanceIdCond{}
	this.NfInstanceId = nfInstanceId
	return &this
}

// NewNfInstanceIdCondWithDefaults instantiates a new NfInstanceIdCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfInstanceIdCondWithDefaults() *NfInstanceIdCond {
	this := NfInstanceIdCond{}
	return &this
}

// GetNfInstanceId returns the NfInstanceId field value
func (o *NfInstanceIdCond) GetNfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value
// and a boolean to check if the value has been set.
func (o *NfInstanceIdCond) GetNfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfInstanceId, true
}

// SetNfInstanceId sets field value
func (o *NfInstanceIdCond) SetNfInstanceId(v string) {
	o.NfInstanceId = v
}

func (o NfInstanceIdCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NfInstanceIdCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfInstanceId"] = o.NfInstanceId
	return toSerialize, nil
}

func (o *NfInstanceIdCond) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nfInstanceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNfInstanceIdCond := _NfInstanceIdCond{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNfInstanceIdCond)

	if err != nil {
		return err
	}

	*o = NfInstanceIdCond(varNfInstanceIdCond)

	return err
}

type NullableNfInstanceIdCond struct {
	value *NfInstanceIdCond
	isSet bool
}

func (v NullableNfInstanceIdCond) Get() *NfInstanceIdCond {
	return v.value
}

func (v *NullableNfInstanceIdCond) Set(val *NfInstanceIdCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfInstanceIdCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfInstanceIdCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfInstanceIdCond(val *NfInstanceIdCond) *NullableNfInstanceIdCond {
	return &NullableNfInstanceIdCond{value: val, isSet: true}
}

func (v NullableNfInstanceIdCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfInstanceIdCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



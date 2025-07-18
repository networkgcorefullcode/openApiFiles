# DnnUpfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** |  | 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 

## Methods

### NewDnnUpfInfoItem

`func NewDnnUpfInfoItem(dnn string, ) *DnnUpfInfoItem`

NewDnnUpfInfoItem instantiates a new DnnUpfInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnUpfInfoItemWithDefaults

`func NewDnnUpfInfoItemWithDefaults() *DnnUpfInfoItem`

NewDnnUpfInfoItemWithDefaults instantiates a new DnnUpfInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *DnnUpfInfoItem) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnnUpfInfoItem) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnnUpfInfoItem) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetDnaiList

`func (o *DnnUpfInfoItem) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *DnnUpfInfoItem) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *DnnUpfInfoItem) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *DnnUpfInfoItem) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetPduSessionTypes

`func (o *DnnUpfInfoItem) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *DnnUpfInfoItem) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *DnnUpfInfoItem) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *DnnUpfInfoItem) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



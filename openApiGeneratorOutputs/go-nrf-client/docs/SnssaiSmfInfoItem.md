# SnssaiSmfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssai** | [**ExtSnssai**](ExtSnssai.md) |  | 
**DnnSmfInfoList** | Pointer to [**[]DnnSmfInfoItem**](DnnSmfInfoItem.md) |  | [optional] 
**DnnSmfInfoListId** | Pointer to **int32** |  | [optional] 

## Methods

### NewSnssaiSmfInfoItem

`func NewSnssaiSmfInfoItem(sNssai ExtSnssai, ) *SnssaiSmfInfoItem`

NewSnssaiSmfInfoItem instantiates a new SnssaiSmfInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnssaiSmfInfoItemWithDefaults

`func NewSnssaiSmfInfoItemWithDefaults() *SnssaiSmfInfoItem`

NewSnssaiSmfInfoItemWithDefaults instantiates a new SnssaiSmfInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssai

`func (o *SnssaiSmfInfoItem) GetSNssai() ExtSnssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SnssaiSmfInfoItem) GetSNssaiOk() (*ExtSnssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SnssaiSmfInfoItem) SetSNssai(v ExtSnssai)`

SetSNssai sets SNssai field to given value.


### GetDnnSmfInfoList

`func (o *SnssaiSmfInfoItem) GetDnnSmfInfoList() []DnnSmfInfoItem`

GetDnnSmfInfoList returns the DnnSmfInfoList field if non-nil, zero value otherwise.

### GetDnnSmfInfoListOk

`func (o *SnssaiSmfInfoItem) GetDnnSmfInfoListOk() (*[]DnnSmfInfoItem, bool)`

GetDnnSmfInfoListOk returns a tuple with the DnnSmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSmfInfoList

`func (o *SnssaiSmfInfoItem) SetDnnSmfInfoList(v []DnnSmfInfoItem)`

SetDnnSmfInfoList sets DnnSmfInfoList field to given value.

### HasDnnSmfInfoList

`func (o *SnssaiSmfInfoItem) HasDnnSmfInfoList() bool`

HasDnnSmfInfoList returns a boolean if a field has been set.

### GetDnnSmfInfoListId

`func (o *SnssaiSmfInfoItem) GetDnnSmfInfoListId() int32`

GetDnnSmfInfoListId returns the DnnSmfInfoListId field if non-nil, zero value otherwise.

### GetDnnSmfInfoListIdOk

`func (o *SnssaiSmfInfoItem) GetDnnSmfInfoListIdOk() (*int32, bool)`

GetDnnSmfInfoListIdOk returns a tuple with the DnnSmfInfoListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSmfInfoListId

`func (o *SnssaiSmfInfoItem) SetDnnSmfInfoListId(v int32)`

SetDnnSmfInfoListId sets DnnSmfInfoListId field to given value.

### HasDnnSmfInfoListId

`func (o *SnssaiSmfInfoItem) HasDnnSmfInfoListId() bool`

HasDnnSmfInfoListId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



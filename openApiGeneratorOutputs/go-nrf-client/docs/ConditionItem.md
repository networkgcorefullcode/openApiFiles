# ConditionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**ServiceFeature** | Pointer to **int32** |  | [optional] 
**VsServiceFeature** | Pointer to **int32** |  | [optional] 
**SupiRangeList** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ImpuRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ImpiRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**PeiList** | Pointer to **[]string** |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConditionItem

`func NewConditionItem() *ConditionItem`

NewConditionItem instantiates a new ConditionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionItemWithDefaults

`func NewConditionItemWithDefaults() *ConditionItem`

NewConditionItemWithDefaults instantiates a new ConditionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerNfTypes

`func (o *ConditionItem) GetConsumerNfTypes() []NFType`

GetConsumerNfTypes returns the ConsumerNfTypes field if non-nil, zero value otherwise.

### GetConsumerNfTypesOk

`func (o *ConditionItem) GetConsumerNfTypesOk() (*[]NFType, bool)`

GetConsumerNfTypesOk returns a tuple with the ConsumerNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerNfTypes

`func (o *ConditionItem) SetConsumerNfTypes(v []NFType)`

SetConsumerNfTypes sets ConsumerNfTypes field to given value.

### HasConsumerNfTypes

`func (o *ConditionItem) HasConsumerNfTypes() bool`

HasConsumerNfTypes returns a boolean if a field has been set.

### GetServiceFeature

`func (o *ConditionItem) GetServiceFeature() int32`

GetServiceFeature returns the ServiceFeature field if non-nil, zero value otherwise.

### GetServiceFeatureOk

`func (o *ConditionItem) GetServiceFeatureOk() (*int32, bool)`

GetServiceFeatureOk returns a tuple with the ServiceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFeature

`func (o *ConditionItem) SetServiceFeature(v int32)`

SetServiceFeature sets ServiceFeature field to given value.

### HasServiceFeature

`func (o *ConditionItem) HasServiceFeature() bool`

HasServiceFeature returns a boolean if a field has been set.

### GetVsServiceFeature

`func (o *ConditionItem) GetVsServiceFeature() int32`

GetVsServiceFeature returns the VsServiceFeature field if non-nil, zero value otherwise.

### GetVsServiceFeatureOk

`func (o *ConditionItem) GetVsServiceFeatureOk() (*int32, bool)`

GetVsServiceFeatureOk returns a tuple with the VsServiceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsServiceFeature

`func (o *ConditionItem) SetVsServiceFeature(v int32)`

SetVsServiceFeature sets VsServiceFeature field to given value.

### HasVsServiceFeature

`func (o *ConditionItem) HasVsServiceFeature() bool`

HasVsServiceFeature returns a boolean if a field has been set.

### GetSupiRangeList

`func (o *ConditionItem) GetSupiRangeList() []SupiRange`

GetSupiRangeList returns the SupiRangeList field if non-nil, zero value otherwise.

### GetSupiRangeListOk

`func (o *ConditionItem) GetSupiRangeListOk() (*[]SupiRange, bool)`

GetSupiRangeListOk returns a tuple with the SupiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRangeList

`func (o *ConditionItem) SetSupiRangeList(v []SupiRange)`

SetSupiRangeList sets SupiRangeList field to given value.

### HasSupiRangeList

`func (o *ConditionItem) HasSupiRangeList() bool`

HasSupiRangeList returns a boolean if a field has been set.

### GetGpsiRangeList

`func (o *ConditionItem) GetGpsiRangeList() []IdentityRange`

GetGpsiRangeList returns the GpsiRangeList field if non-nil, zero value otherwise.

### GetGpsiRangeListOk

`func (o *ConditionItem) GetGpsiRangeListOk() (*[]IdentityRange, bool)`

GetGpsiRangeListOk returns a tuple with the GpsiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRangeList

`func (o *ConditionItem) SetGpsiRangeList(v []IdentityRange)`

SetGpsiRangeList sets GpsiRangeList field to given value.

### HasGpsiRangeList

`func (o *ConditionItem) HasGpsiRangeList() bool`

HasGpsiRangeList returns a boolean if a field has been set.

### GetImpuRangeList

`func (o *ConditionItem) GetImpuRangeList() []IdentityRange`

GetImpuRangeList returns the ImpuRangeList field if non-nil, zero value otherwise.

### GetImpuRangeListOk

`func (o *ConditionItem) GetImpuRangeListOk() (*[]IdentityRange, bool)`

GetImpuRangeListOk returns a tuple with the ImpuRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpuRangeList

`func (o *ConditionItem) SetImpuRangeList(v []IdentityRange)`

SetImpuRangeList sets ImpuRangeList field to given value.

### HasImpuRangeList

`func (o *ConditionItem) HasImpuRangeList() bool`

HasImpuRangeList returns a boolean if a field has been set.

### GetImpiRangeList

`func (o *ConditionItem) GetImpiRangeList() []IdentityRange`

GetImpiRangeList returns the ImpiRangeList field if non-nil, zero value otherwise.

### GetImpiRangeListOk

`func (o *ConditionItem) GetImpiRangeListOk() (*[]IdentityRange, bool)`

GetImpiRangeListOk returns a tuple with the ImpiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpiRangeList

`func (o *ConditionItem) SetImpiRangeList(v []IdentityRange)`

SetImpiRangeList sets ImpiRangeList field to given value.

### HasImpiRangeList

`func (o *ConditionItem) HasImpiRangeList() bool`

HasImpiRangeList returns a boolean if a field has been set.

### GetPeiList

`func (o *ConditionItem) GetPeiList() []string`

GetPeiList returns the PeiList field if non-nil, zero value otherwise.

### GetPeiListOk

`func (o *ConditionItem) GetPeiListOk() (*[]string, bool)`

GetPeiListOk returns a tuple with the PeiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeiList

`func (o *ConditionItem) SetPeiList(v []string)`

SetPeiList sets PeiList field to given value.

### HasPeiList

`func (o *ConditionItem) HasPeiList() bool`

HasPeiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *ConditionItem) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *ConditionItem) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *ConditionItem) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *ConditionItem) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetDnnList

`func (o *ConditionItem) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *ConditionItem) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *ConditionItem) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *ConditionItem) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



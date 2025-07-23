# SelectionConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerNfTypes** | Pointer to [**[]NfType**](NfType.md) |  | [optional] 
**ServiceFeature** | Pointer to **int32** |  | [optional] 
**VsServiceFeature** | Pointer to **int32** |  | [optional] 
**SupiRangeList** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ImpuRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ImpiRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**PeiList** | Pointer to **[]string** |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 
**And** | Pointer to [**[]SelectionConditions**](SelectionConditions.md) |  | [optional] 
**Or** | Pointer to [**[]SelectionConditions**](SelectionConditions.md) |  | [optional] 

## Methods

### NewSelectionConditions

`func NewSelectionConditions() *SelectionConditions`

NewSelectionConditions instantiates a new SelectionConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectionConditionsWithDefaults

`func NewSelectionConditionsWithDefaults() *SelectionConditions`

NewSelectionConditionsWithDefaults instantiates a new SelectionConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerNfTypes

`func (o *SelectionConditions) GetConsumerNfTypes() []NfType`

GetConsumerNfTypes returns the ConsumerNfTypes field if non-nil, zero value otherwise.

### GetConsumerNfTypesOk

`func (o *SelectionConditions) GetConsumerNfTypesOk() (*[]NfType, bool)`

GetConsumerNfTypesOk returns a tuple with the ConsumerNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerNfTypes

`func (o *SelectionConditions) SetConsumerNfTypes(v []NfType)`

SetConsumerNfTypes sets ConsumerNfTypes field to given value.

### HasConsumerNfTypes

`func (o *SelectionConditions) HasConsumerNfTypes() bool`

HasConsumerNfTypes returns a boolean if a field has been set.

### GetServiceFeature

`func (o *SelectionConditions) GetServiceFeature() int32`

GetServiceFeature returns the ServiceFeature field if non-nil, zero value otherwise.

### GetServiceFeatureOk

`func (o *SelectionConditions) GetServiceFeatureOk() (*int32, bool)`

GetServiceFeatureOk returns a tuple with the ServiceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFeature

`func (o *SelectionConditions) SetServiceFeature(v int32)`

SetServiceFeature sets ServiceFeature field to given value.

### HasServiceFeature

`func (o *SelectionConditions) HasServiceFeature() bool`

HasServiceFeature returns a boolean if a field has been set.

### GetVsServiceFeature

`func (o *SelectionConditions) GetVsServiceFeature() int32`

GetVsServiceFeature returns the VsServiceFeature field if non-nil, zero value otherwise.

### GetVsServiceFeatureOk

`func (o *SelectionConditions) GetVsServiceFeatureOk() (*int32, bool)`

GetVsServiceFeatureOk returns a tuple with the VsServiceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsServiceFeature

`func (o *SelectionConditions) SetVsServiceFeature(v int32)`

SetVsServiceFeature sets VsServiceFeature field to given value.

### HasVsServiceFeature

`func (o *SelectionConditions) HasVsServiceFeature() bool`

HasVsServiceFeature returns a boolean if a field has been set.

### GetSupiRangeList

`func (o *SelectionConditions) GetSupiRangeList() []SupiRange`

GetSupiRangeList returns the SupiRangeList field if non-nil, zero value otherwise.

### GetSupiRangeListOk

`func (o *SelectionConditions) GetSupiRangeListOk() (*[]SupiRange, bool)`

GetSupiRangeListOk returns a tuple with the SupiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRangeList

`func (o *SelectionConditions) SetSupiRangeList(v []SupiRange)`

SetSupiRangeList sets SupiRangeList field to given value.

### HasSupiRangeList

`func (o *SelectionConditions) HasSupiRangeList() bool`

HasSupiRangeList returns a boolean if a field has been set.

### GetGpsiRangeList

`func (o *SelectionConditions) GetGpsiRangeList() []IdentityRange`

GetGpsiRangeList returns the GpsiRangeList field if non-nil, zero value otherwise.

### GetGpsiRangeListOk

`func (o *SelectionConditions) GetGpsiRangeListOk() (*[]IdentityRange, bool)`

GetGpsiRangeListOk returns a tuple with the GpsiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRangeList

`func (o *SelectionConditions) SetGpsiRangeList(v []IdentityRange)`

SetGpsiRangeList sets GpsiRangeList field to given value.

### HasGpsiRangeList

`func (o *SelectionConditions) HasGpsiRangeList() bool`

HasGpsiRangeList returns a boolean if a field has been set.

### GetImpuRangeList

`func (o *SelectionConditions) GetImpuRangeList() []IdentityRange`

GetImpuRangeList returns the ImpuRangeList field if non-nil, zero value otherwise.

### GetImpuRangeListOk

`func (o *SelectionConditions) GetImpuRangeListOk() (*[]IdentityRange, bool)`

GetImpuRangeListOk returns a tuple with the ImpuRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpuRangeList

`func (o *SelectionConditions) SetImpuRangeList(v []IdentityRange)`

SetImpuRangeList sets ImpuRangeList field to given value.

### HasImpuRangeList

`func (o *SelectionConditions) HasImpuRangeList() bool`

HasImpuRangeList returns a boolean if a field has been set.

### GetImpiRangeList

`func (o *SelectionConditions) GetImpiRangeList() []IdentityRange`

GetImpiRangeList returns the ImpiRangeList field if non-nil, zero value otherwise.

### GetImpiRangeListOk

`func (o *SelectionConditions) GetImpiRangeListOk() (*[]IdentityRange, bool)`

GetImpiRangeListOk returns a tuple with the ImpiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpiRangeList

`func (o *SelectionConditions) SetImpiRangeList(v []IdentityRange)`

SetImpiRangeList sets ImpiRangeList field to given value.

### HasImpiRangeList

`func (o *SelectionConditions) HasImpiRangeList() bool`

HasImpiRangeList returns a boolean if a field has been set.

### GetPeiList

`func (o *SelectionConditions) GetPeiList() []string`

GetPeiList returns the PeiList field if non-nil, zero value otherwise.

### GetPeiListOk

`func (o *SelectionConditions) GetPeiListOk() (*[]string, bool)`

GetPeiListOk returns a tuple with the PeiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeiList

`func (o *SelectionConditions) SetPeiList(v []string)`

SetPeiList sets PeiList field to given value.

### HasPeiList

`func (o *SelectionConditions) HasPeiList() bool`

HasPeiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *SelectionConditions) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *SelectionConditions) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *SelectionConditions) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *SelectionConditions) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetDnnList

`func (o *SelectionConditions) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *SelectionConditions) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *SelectionConditions) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *SelectionConditions) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetAnd

`func (o *SelectionConditions) GetAnd() []SelectionConditions`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SelectionConditions) GetAndOk() (*[]SelectionConditions, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SelectionConditions) SetAnd(v []SelectionConditions)`

SetAnd sets And field to given value.

### HasAnd

`func (o *SelectionConditions) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *SelectionConditions) GetOr() []SelectionConditions`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SelectionConditions) GetOrOk() (*[]SelectionConditions, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SelectionConditions) SetOr(v []SelectionConditions)`

SetOr sets Or field to given value.

### HasOr

`func (o *SelectionConditions) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NrfInfoServedPcfInfoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**RxDiamHost** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**RxDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**V2xSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**ProseSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**ProseCapability** | Pointer to [**ProSeCapability**](ProSeCapability.md) |  | [optional] 
**V2xCapability** | Pointer to [**V2xCapability**](V2xCapability.md) |  | [optional] 
**A2xSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**A2xCapability** | Pointer to [**A2xCapability**](A2xCapability.md) |  | [optional] 
**RangingSlPosSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**UrspEpsSupport** | Pointer to **bool** | URSP delivery in EPS is supported by the PCF | [optional] [default to false]

## Methods

### NewNrfInfoServedPcfInfoValue

`func NewNrfInfoServedPcfInfoValue() *NrfInfoServedPcfInfoValue`

NewNrfInfoServedPcfInfoValue instantiates a new NrfInfoServedPcfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedPcfInfoValueWithDefaults

`func NewNrfInfoServedPcfInfoValueWithDefaults() *NrfInfoServedPcfInfoValue`

NewNrfInfoServedPcfInfoValueWithDefaults instantiates a new NrfInfoServedPcfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *NrfInfoServedPcfInfoValue) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NrfInfoServedPcfInfoValue) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NrfInfoServedPcfInfoValue) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NrfInfoServedPcfInfoValue) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDnnList

`func (o *NrfInfoServedPcfInfoValue) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *NrfInfoServedPcfInfoValue) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *NrfInfoServedPcfInfoValue) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *NrfInfoServedPcfInfoValue) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetSupiRanges

`func (o *NrfInfoServedPcfInfoValue) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *NrfInfoServedPcfInfoValue) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *NrfInfoServedPcfInfoValue) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *NrfInfoServedPcfInfoValue) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NrfInfoServedPcfInfoValue) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NrfInfoServedPcfInfoValue) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NrfInfoServedPcfInfoValue) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NrfInfoServedPcfInfoValue) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetRxDiamHost

`func (o *NrfInfoServedPcfInfoValue) GetRxDiamHost() string`

GetRxDiamHost returns the RxDiamHost field if non-nil, zero value otherwise.

### GetRxDiamHostOk

`func (o *NrfInfoServedPcfInfoValue) GetRxDiamHostOk() (*string, bool)`

GetRxDiamHostOk returns a tuple with the RxDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamHost

`func (o *NrfInfoServedPcfInfoValue) SetRxDiamHost(v string)`

SetRxDiamHost sets RxDiamHost field to given value.

### HasRxDiamHost

`func (o *NrfInfoServedPcfInfoValue) HasRxDiamHost() bool`

HasRxDiamHost returns a boolean if a field has been set.

### GetRxDiamRealm

`func (o *NrfInfoServedPcfInfoValue) GetRxDiamRealm() string`

GetRxDiamRealm returns the RxDiamRealm field if non-nil, zero value otherwise.

### GetRxDiamRealmOk

`func (o *NrfInfoServedPcfInfoValue) GetRxDiamRealmOk() (*string, bool)`

GetRxDiamRealmOk returns a tuple with the RxDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamRealm

`func (o *NrfInfoServedPcfInfoValue) SetRxDiamRealm(v string)`

SetRxDiamRealm sets RxDiamRealm field to given value.

### HasRxDiamRealm

`func (o *NrfInfoServedPcfInfoValue) HasRxDiamRealm() bool`

HasRxDiamRealm returns a boolean if a field has been set.

### GetV2xSupportInd

`func (o *NrfInfoServedPcfInfoValue) GetV2xSupportInd() bool`

GetV2xSupportInd returns the V2xSupportInd field if non-nil, zero value otherwise.

### GetV2xSupportIndOk

`func (o *NrfInfoServedPcfInfoValue) GetV2xSupportIndOk() (*bool, bool)`

GetV2xSupportIndOk returns a tuple with the V2xSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xSupportInd

`func (o *NrfInfoServedPcfInfoValue) SetV2xSupportInd(v bool)`

SetV2xSupportInd sets V2xSupportInd field to given value.

### HasV2xSupportInd

`func (o *NrfInfoServedPcfInfoValue) HasV2xSupportInd() bool`

HasV2xSupportInd returns a boolean if a field has been set.

### GetProseSupportInd

`func (o *NrfInfoServedPcfInfoValue) GetProseSupportInd() bool`

GetProseSupportInd returns the ProseSupportInd field if non-nil, zero value otherwise.

### GetProseSupportIndOk

`func (o *NrfInfoServedPcfInfoValue) GetProseSupportIndOk() (*bool, bool)`

GetProseSupportIndOk returns a tuple with the ProseSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseSupportInd

`func (o *NrfInfoServedPcfInfoValue) SetProseSupportInd(v bool)`

SetProseSupportInd sets ProseSupportInd field to given value.

### HasProseSupportInd

`func (o *NrfInfoServedPcfInfoValue) HasProseSupportInd() bool`

HasProseSupportInd returns a boolean if a field has been set.

### GetProseCapability

`func (o *NrfInfoServedPcfInfoValue) GetProseCapability() ProSeCapability`

GetProseCapability returns the ProseCapability field if non-nil, zero value otherwise.

### GetProseCapabilityOk

`func (o *NrfInfoServedPcfInfoValue) GetProseCapabilityOk() (*ProSeCapability, bool)`

GetProseCapabilityOk returns a tuple with the ProseCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseCapability

`func (o *NrfInfoServedPcfInfoValue) SetProseCapability(v ProSeCapability)`

SetProseCapability sets ProseCapability field to given value.

### HasProseCapability

`func (o *NrfInfoServedPcfInfoValue) HasProseCapability() bool`

HasProseCapability returns a boolean if a field has been set.

### GetV2xCapability

`func (o *NrfInfoServedPcfInfoValue) GetV2xCapability() V2xCapability`

GetV2xCapability returns the V2xCapability field if non-nil, zero value otherwise.

### GetV2xCapabilityOk

`func (o *NrfInfoServedPcfInfoValue) GetV2xCapabilityOk() (*V2xCapability, bool)`

GetV2xCapabilityOk returns a tuple with the V2xCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xCapability

`func (o *NrfInfoServedPcfInfoValue) SetV2xCapability(v V2xCapability)`

SetV2xCapability sets V2xCapability field to given value.

### HasV2xCapability

`func (o *NrfInfoServedPcfInfoValue) HasV2xCapability() bool`

HasV2xCapability returns a boolean if a field has been set.

### GetA2xSupportInd

`func (o *NrfInfoServedPcfInfoValue) GetA2xSupportInd() bool`

GetA2xSupportInd returns the A2xSupportInd field if non-nil, zero value otherwise.

### GetA2xSupportIndOk

`func (o *NrfInfoServedPcfInfoValue) GetA2xSupportIndOk() (*bool, bool)`

GetA2xSupportIndOk returns a tuple with the A2xSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2xSupportInd

`func (o *NrfInfoServedPcfInfoValue) SetA2xSupportInd(v bool)`

SetA2xSupportInd sets A2xSupportInd field to given value.

### HasA2xSupportInd

`func (o *NrfInfoServedPcfInfoValue) HasA2xSupportInd() bool`

HasA2xSupportInd returns a boolean if a field has been set.

### GetA2xCapability

`func (o *NrfInfoServedPcfInfoValue) GetA2xCapability() A2xCapability`

GetA2xCapability returns the A2xCapability field if non-nil, zero value otherwise.

### GetA2xCapabilityOk

`func (o *NrfInfoServedPcfInfoValue) GetA2xCapabilityOk() (*A2xCapability, bool)`

GetA2xCapabilityOk returns a tuple with the A2xCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2xCapability

`func (o *NrfInfoServedPcfInfoValue) SetA2xCapability(v A2xCapability)`

SetA2xCapability sets A2xCapability field to given value.

### HasA2xCapability

`func (o *NrfInfoServedPcfInfoValue) HasA2xCapability() bool`

HasA2xCapability returns a boolean if a field has been set.

### GetRangingSlPosSupportInd

`func (o *NrfInfoServedPcfInfoValue) GetRangingSlPosSupportInd() bool`

GetRangingSlPosSupportInd returns the RangingSlPosSupportInd field if non-nil, zero value otherwise.

### GetRangingSlPosSupportIndOk

`func (o *NrfInfoServedPcfInfoValue) GetRangingSlPosSupportIndOk() (*bool, bool)`

GetRangingSlPosSupportIndOk returns a tuple with the RangingSlPosSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangingSlPosSupportInd

`func (o *NrfInfoServedPcfInfoValue) SetRangingSlPosSupportInd(v bool)`

SetRangingSlPosSupportInd sets RangingSlPosSupportInd field to given value.

### HasRangingSlPosSupportInd

`func (o *NrfInfoServedPcfInfoValue) HasRangingSlPosSupportInd() bool`

HasRangingSlPosSupportInd returns a boolean if a field has been set.

### GetUrspEpsSupport

`func (o *NrfInfoServedPcfInfoValue) GetUrspEpsSupport() bool`

GetUrspEpsSupport returns the UrspEpsSupport field if non-nil, zero value otherwise.

### GetUrspEpsSupportOk

`func (o *NrfInfoServedPcfInfoValue) GetUrspEpsSupportOk() (*bool, bool)`

GetUrspEpsSupportOk returns a tuple with the UrspEpsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrspEpsSupport

`func (o *NrfInfoServedPcfInfoValue) SetUrspEpsSupport(v bool)`

SetUrspEpsSupport sets UrspEpsSupport field to given value.

### HasUrspEpsSupport

`func (o *NrfInfoServedPcfInfoValue) HasUrspEpsSupport() bool`

HasUrspEpsSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



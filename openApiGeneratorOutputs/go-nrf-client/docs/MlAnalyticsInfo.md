# MlAnalyticsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MlAnalyticsIds** | Pointer to [**[]NwdafEvent**](NwdafEvent.md) |  | [optional] 
**SnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**TrackingAreaList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**MlModelInterInfo** | Pointer to [**MlModelInterInfo**](MlModelInterInfo.md) |  | [optional] 
**FlCapabilityType** | Pointer to [**FlCapabilityType**](FlCapabilityType.md) |  | [optional] 
**FlTimeInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**NfSetIdList** | Pointer to **[]string** |  | [optional] 
**VflCapabilityType** | Pointer to [**VflCapabilityType**](VflCapabilityType.md) |  | [optional] 

## Methods

### NewMlAnalyticsInfo

`func NewMlAnalyticsInfo() *MlAnalyticsInfo`

NewMlAnalyticsInfo instantiates a new MlAnalyticsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlAnalyticsInfoWithDefaults

`func NewMlAnalyticsInfoWithDefaults() *MlAnalyticsInfo`

NewMlAnalyticsInfoWithDefaults instantiates a new MlAnalyticsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMlAnalyticsIds

`func (o *MlAnalyticsInfo) GetMlAnalyticsIds() []NwdafEvent`

GetMlAnalyticsIds returns the MlAnalyticsIds field if non-nil, zero value otherwise.

### GetMlAnalyticsIdsOk

`func (o *MlAnalyticsInfo) GetMlAnalyticsIdsOk() (*[]NwdafEvent, bool)`

GetMlAnalyticsIdsOk returns a tuple with the MlAnalyticsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlAnalyticsIds

`func (o *MlAnalyticsInfo) SetMlAnalyticsIds(v []NwdafEvent)`

SetMlAnalyticsIds sets MlAnalyticsIds field to given value.

### HasMlAnalyticsIds

`func (o *MlAnalyticsInfo) HasMlAnalyticsIds() bool`

HasMlAnalyticsIds returns a boolean if a field has been set.

### GetSnssaiList

`func (o *MlAnalyticsInfo) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *MlAnalyticsInfo) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *MlAnalyticsInfo) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.

### HasSnssaiList

`func (o *MlAnalyticsInfo) HasSnssaiList() bool`

HasSnssaiList returns a boolean if a field has been set.

### GetTrackingAreaList

`func (o *MlAnalyticsInfo) GetTrackingAreaList() []Tai`

GetTrackingAreaList returns the TrackingAreaList field if non-nil, zero value otherwise.

### GetTrackingAreaListOk

`func (o *MlAnalyticsInfo) GetTrackingAreaListOk() (*[]Tai, bool)`

GetTrackingAreaListOk returns a tuple with the TrackingAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingAreaList

`func (o *MlAnalyticsInfo) SetTrackingAreaList(v []Tai)`

SetTrackingAreaList sets TrackingAreaList field to given value.

### HasTrackingAreaList

`func (o *MlAnalyticsInfo) HasTrackingAreaList() bool`

HasTrackingAreaList returns a boolean if a field has been set.

### GetMlModelInterInfo

`func (o *MlAnalyticsInfo) GetMlModelInterInfo() MlModelInterInfo`

GetMlModelInterInfo returns the MlModelInterInfo field if non-nil, zero value otherwise.

### GetMlModelInterInfoOk

`func (o *MlAnalyticsInfo) GetMlModelInterInfoOk() (*MlModelInterInfo, bool)`

GetMlModelInterInfoOk returns a tuple with the MlModelInterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlModelInterInfo

`func (o *MlAnalyticsInfo) SetMlModelInterInfo(v MlModelInterInfo)`

SetMlModelInterInfo sets MlModelInterInfo field to given value.

### HasMlModelInterInfo

`func (o *MlAnalyticsInfo) HasMlModelInterInfo() bool`

HasMlModelInterInfo returns a boolean if a field has been set.

### GetFlCapabilityType

`func (o *MlAnalyticsInfo) GetFlCapabilityType() FlCapabilityType`

GetFlCapabilityType returns the FlCapabilityType field if non-nil, zero value otherwise.

### GetFlCapabilityTypeOk

`func (o *MlAnalyticsInfo) GetFlCapabilityTypeOk() (*FlCapabilityType, bool)`

GetFlCapabilityTypeOk returns a tuple with the FlCapabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlCapabilityType

`func (o *MlAnalyticsInfo) SetFlCapabilityType(v FlCapabilityType)`

SetFlCapabilityType sets FlCapabilityType field to given value.

### HasFlCapabilityType

`func (o *MlAnalyticsInfo) HasFlCapabilityType() bool`

HasFlCapabilityType returns a boolean if a field has been set.

### GetFlTimeInterval

`func (o *MlAnalyticsInfo) GetFlTimeInterval() int32`

GetFlTimeInterval returns the FlTimeInterval field if non-nil, zero value otherwise.

### GetFlTimeIntervalOk

`func (o *MlAnalyticsInfo) GetFlTimeIntervalOk() (*int32, bool)`

GetFlTimeIntervalOk returns a tuple with the FlTimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlTimeInterval

`func (o *MlAnalyticsInfo) SetFlTimeInterval(v int32)`

SetFlTimeInterval sets FlTimeInterval field to given value.

### HasFlTimeInterval

`func (o *MlAnalyticsInfo) HasFlTimeInterval() bool`

HasFlTimeInterval returns a boolean if a field has been set.

### GetNfTypeList

`func (o *MlAnalyticsInfo) GetNfTypeList() []NFType`

GetNfTypeList returns the NfTypeList field if non-nil, zero value otherwise.

### GetNfTypeListOk

`func (o *MlAnalyticsInfo) GetNfTypeListOk() (*[]NFType, bool)`

GetNfTypeListOk returns a tuple with the NfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypeList

`func (o *MlAnalyticsInfo) SetNfTypeList(v []NFType)`

SetNfTypeList sets NfTypeList field to given value.

### HasNfTypeList

`func (o *MlAnalyticsInfo) HasNfTypeList() bool`

HasNfTypeList returns a boolean if a field has been set.

### GetNfSetIdList

`func (o *MlAnalyticsInfo) GetNfSetIdList() []string`

GetNfSetIdList returns the NfSetIdList field if non-nil, zero value otherwise.

### GetNfSetIdListOk

`func (o *MlAnalyticsInfo) GetNfSetIdListOk() (*[]string, bool)`

GetNfSetIdListOk returns a tuple with the NfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIdList

`func (o *MlAnalyticsInfo) SetNfSetIdList(v []string)`

SetNfSetIdList sets NfSetIdList field to given value.

### HasNfSetIdList

`func (o *MlAnalyticsInfo) HasNfSetIdList() bool`

HasNfSetIdList returns a boolean if a field has been set.

### GetVflCapabilityType

`func (o *MlAnalyticsInfo) GetVflCapabilityType() VflCapabilityType`

GetVflCapabilityType returns the VflCapabilityType field if non-nil, zero value otherwise.

### GetVflCapabilityTypeOk

`func (o *MlAnalyticsInfo) GetVflCapabilityTypeOk() (*VflCapabilityType, bool)`

GetVflCapabilityTypeOk returns a tuple with the VflCapabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVflCapabilityType

`func (o *MlAnalyticsInfo) SetVflCapabilityType(v VflCapabilityType)`

SetVflCapabilityType sets VflCapabilityType field to given value.

### HasVflCapabilityType

`func (o *MlAnalyticsInfo) HasVflCapabilityType() bool`

HasVflCapabilityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



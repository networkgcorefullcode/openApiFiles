# NrfInfoServedUpfInfoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiUpfInfoList** | [**[]SnssaiUpfInfoItem**](SnssaiUpfInfoItem.md) |  | 
**SmfServingArea** | Pointer to **[]string** |  | [optional] 
**InterfaceUpfInfoList** | Pointer to [**[]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) |  | [optional] 
**N6TunnelInfoList** | Pointer to [**map[string]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) | A map of InterfaceUpfInfoItems containing the N6 tunnelling information for establishing  a N6 tunnel between the V-UPF and the V-EASDF, where a valid JSON string serves as key.  | [optional] 
**IwkEpsInd** | Pointer to **bool** |  | [optional] [default to false]
**SxaInd** | Pointer to **bool** |  | [optional] 
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**AtsssCapability** | Pointer to [**AtsssCapability**](AtsssCapability.md) |  | [optional] 
**UeIpAddrInd** | Pointer to **bool** |  | [optional] [default to false]
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**WAgfInfo** | Pointer to [**NullableWAgfInfo**](WAgfInfo.md) |  | [optional] 
**TngfInfo** | Pointer to [**NullableTngfInfo**](TngfInfo.md) |  | [optional] 
**TwifInfo** | Pointer to [**NullableTwifInfo**](TwifInfo.md) |  | [optional] 
**PreferredEpdgInfoList** | Pointer to [**[]EpdgInfo**](EpdgInfo.md) |  | [optional] 
**PreferredWAgfInfoList** | Pointer to [**[]WAgfInfo**](WAgfInfo.md) |  | [optional] 
**PreferredTngfInfoList** | Pointer to [**[]TngfInfo**](TngfInfo.md) |  | [optional] 
**PreferredTwifInfoList** | Pointer to [**[]TwifInfo**](TwifInfo.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RedundantGtpu** | Pointer to **bool** |  | [optional] [default to false]
**Ipups** | Pointer to **bool** |  | [optional] [default to false]
**DataForwarding** | Pointer to **bool** |  | [optional] [default to false]
**SupportedPfcpFeatures** | Pointer to **string** |  | [optional] 
**UpfEvents** | Pointer to [**[]EventType**](EventType.md) |  | [optional] 
**OpConfigCaps** | Pointer to **[]string** |  | [optional] 
**PacketInspectionFunctionalities** | Pointer to [**[]UpfPacketInspectionFunctionality**](UpfPacketInspectionFunctionality.md) |  | [optional] 
**N6DelayMeastProtocs** | Pointer to [**[]N6DelayMeasurementProtocol**](N6DelayMeasurementProtocol.md) |  | [optional] 
**GeranUtranInd** | Pointer to **bool** |  | [optional] 
**Var2g3gLocationAreaList** | Pointer to [**[]Model2G3GLocationArea**](Model2G3GLocationArea.md) |  | [optional] 
**Var2g3gLocationAreaRangeList** | Pointer to [**[]Model2G3GLocationAreaRange**](Model2G3GLocationAreaRange.md) |  | [optional] 

## Methods

### NewNrfInfoServedUpfInfoValue

`func NewNrfInfoServedUpfInfoValue(sNssaiUpfInfoList []SnssaiUpfInfoItem, ) *NrfInfoServedUpfInfoValue`

NewNrfInfoServedUpfInfoValue instantiates a new NrfInfoServedUpfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedUpfInfoValueWithDefaults

`func NewNrfInfoServedUpfInfoValueWithDefaults() *NrfInfoServedUpfInfoValue`

NewNrfInfoServedUpfInfoValueWithDefaults instantiates a new NrfInfoServedUpfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) GetSNssaiUpfInfoList() []SnssaiUpfInfoItem`

GetSNssaiUpfInfoList returns the SNssaiUpfInfoList field if non-nil, zero value otherwise.

### GetSNssaiUpfInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetSNssaiUpfInfoListOk() (*[]SnssaiUpfInfoItem, bool)`

GetSNssaiUpfInfoListOk returns a tuple with the SNssaiUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) SetSNssaiUpfInfoList(v []SnssaiUpfInfoItem)`

SetSNssaiUpfInfoList sets SNssaiUpfInfoList field to given value.


### GetSmfServingArea

`func (o *NrfInfoServedUpfInfoValue) GetSmfServingArea() []string`

GetSmfServingArea returns the SmfServingArea field if non-nil, zero value otherwise.

### GetSmfServingAreaOk

`func (o *NrfInfoServedUpfInfoValue) GetSmfServingAreaOk() (*[]string, bool)`

GetSmfServingAreaOk returns a tuple with the SmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServingArea

`func (o *NrfInfoServedUpfInfoValue) SetSmfServingArea(v []string)`

SetSmfServingArea sets SmfServingArea field to given value.

### HasSmfServingArea

`func (o *NrfInfoServedUpfInfoValue) HasSmfServingArea() bool`

HasSmfServingArea returns a boolean if a field has been set.

### GetInterfaceUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem`

GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field if non-nil, zero value otherwise.

### GetInterfaceUpfInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetInterfaceUpfInfoListOk() (*[]InterfaceUpfInfoItem, bool)`

GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem)`

SetInterfaceUpfInfoList sets InterfaceUpfInfoList field to given value.

### HasInterfaceUpfInfoList

`func (o *NrfInfoServedUpfInfoValue) HasInterfaceUpfInfoList() bool`

HasInterfaceUpfInfoList returns a boolean if a field has been set.

### GetN6TunnelInfoList

`func (o *NrfInfoServedUpfInfoValue) GetN6TunnelInfoList() map[string]InterfaceUpfInfoItem`

GetN6TunnelInfoList returns the N6TunnelInfoList field if non-nil, zero value otherwise.

### GetN6TunnelInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetN6TunnelInfoListOk() (*map[string]InterfaceUpfInfoItem, bool)`

GetN6TunnelInfoListOk returns a tuple with the N6TunnelInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6TunnelInfoList

`func (o *NrfInfoServedUpfInfoValue) SetN6TunnelInfoList(v map[string]InterfaceUpfInfoItem)`

SetN6TunnelInfoList sets N6TunnelInfoList field to given value.

### HasN6TunnelInfoList

`func (o *NrfInfoServedUpfInfoValue) HasN6TunnelInfoList() bool`

HasN6TunnelInfoList returns a boolean if a field has been set.

### GetIwkEpsInd

`func (o *NrfInfoServedUpfInfoValue) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *NrfInfoServedUpfInfoValue) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *NrfInfoServedUpfInfoValue) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *NrfInfoServedUpfInfoValue) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetSxaInd

`func (o *NrfInfoServedUpfInfoValue) GetSxaInd() bool`

GetSxaInd returns the SxaInd field if non-nil, zero value otherwise.

### GetSxaIndOk

`func (o *NrfInfoServedUpfInfoValue) GetSxaIndOk() (*bool, bool)`

GetSxaIndOk returns a tuple with the SxaInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSxaInd

`func (o *NrfInfoServedUpfInfoValue) SetSxaInd(v bool)`

SetSxaInd sets SxaInd field to given value.

### HasSxaInd

`func (o *NrfInfoServedUpfInfoValue) HasSxaInd() bool`

HasSxaInd returns a boolean if a field has been set.

### GetPduSessionTypes

`func (o *NrfInfoServedUpfInfoValue) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *NrfInfoServedUpfInfoValue) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *NrfInfoServedUpfInfoValue) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *NrfInfoServedUpfInfoValue) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetAtsssCapability

`func (o *NrfInfoServedUpfInfoValue) GetAtsssCapability() AtsssCapability`

GetAtsssCapability returns the AtsssCapability field if non-nil, zero value otherwise.

### GetAtsssCapabilityOk

`func (o *NrfInfoServedUpfInfoValue) GetAtsssCapabilityOk() (*AtsssCapability, bool)`

GetAtsssCapabilityOk returns a tuple with the AtsssCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssCapability

`func (o *NrfInfoServedUpfInfoValue) SetAtsssCapability(v AtsssCapability)`

SetAtsssCapability sets AtsssCapability field to given value.

### HasAtsssCapability

`func (o *NrfInfoServedUpfInfoValue) HasAtsssCapability() bool`

HasAtsssCapability returns a boolean if a field has been set.

### GetUeIpAddrInd

`func (o *NrfInfoServedUpfInfoValue) GetUeIpAddrInd() bool`

GetUeIpAddrInd returns the UeIpAddrInd field if non-nil, zero value otherwise.

### GetUeIpAddrIndOk

`func (o *NrfInfoServedUpfInfoValue) GetUeIpAddrIndOk() (*bool, bool)`

GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddrInd

`func (o *NrfInfoServedUpfInfoValue) SetUeIpAddrInd(v bool)`

SetUeIpAddrInd sets UeIpAddrInd field to given value.

### HasUeIpAddrInd

`func (o *NrfInfoServedUpfInfoValue) HasUeIpAddrInd() bool`

HasUeIpAddrInd returns a boolean if a field has been set.

### GetTaiList

`func (o *NrfInfoServedUpfInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedUpfInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedUpfInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedUpfInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedUpfInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedUpfInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedUpfInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedUpfInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetWAgfInfo

`func (o *NrfInfoServedUpfInfoValue) GetWAgfInfo() WAgfInfo`

GetWAgfInfo returns the WAgfInfo field if non-nil, zero value otherwise.

### GetWAgfInfoOk

`func (o *NrfInfoServedUpfInfoValue) GetWAgfInfoOk() (*WAgfInfo, bool)`

GetWAgfInfoOk returns a tuple with the WAgfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAgfInfo

`func (o *NrfInfoServedUpfInfoValue) SetWAgfInfo(v WAgfInfo)`

SetWAgfInfo sets WAgfInfo field to given value.

### HasWAgfInfo

`func (o *NrfInfoServedUpfInfoValue) HasWAgfInfo() bool`

HasWAgfInfo returns a boolean if a field has been set.

### SetWAgfInfoNil

`func (o *NrfInfoServedUpfInfoValue) SetWAgfInfoNil(b bool)`

 SetWAgfInfoNil sets the value for WAgfInfo to be an explicit nil

### UnsetWAgfInfo
`func (o *NrfInfoServedUpfInfoValue) UnsetWAgfInfo()`

UnsetWAgfInfo ensures that no value is present for WAgfInfo, not even an explicit nil
### GetTngfInfo

`func (o *NrfInfoServedUpfInfoValue) GetTngfInfo() TngfInfo`

GetTngfInfo returns the TngfInfo field if non-nil, zero value otherwise.

### GetTngfInfoOk

`func (o *NrfInfoServedUpfInfoValue) GetTngfInfoOk() (*TngfInfo, bool)`

GetTngfInfoOk returns a tuple with the TngfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfInfo

`func (o *NrfInfoServedUpfInfoValue) SetTngfInfo(v TngfInfo)`

SetTngfInfo sets TngfInfo field to given value.

### HasTngfInfo

`func (o *NrfInfoServedUpfInfoValue) HasTngfInfo() bool`

HasTngfInfo returns a boolean if a field has been set.

### SetTngfInfoNil

`func (o *NrfInfoServedUpfInfoValue) SetTngfInfoNil(b bool)`

 SetTngfInfoNil sets the value for TngfInfo to be an explicit nil

### UnsetTngfInfo
`func (o *NrfInfoServedUpfInfoValue) UnsetTngfInfo()`

UnsetTngfInfo ensures that no value is present for TngfInfo, not even an explicit nil
### GetTwifInfo

`func (o *NrfInfoServedUpfInfoValue) GetTwifInfo() TwifInfo`

GetTwifInfo returns the TwifInfo field if non-nil, zero value otherwise.

### GetTwifInfoOk

`func (o *NrfInfoServedUpfInfoValue) GetTwifInfoOk() (*TwifInfo, bool)`

GetTwifInfoOk returns a tuple with the TwifInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwifInfo

`func (o *NrfInfoServedUpfInfoValue) SetTwifInfo(v TwifInfo)`

SetTwifInfo sets TwifInfo field to given value.

### HasTwifInfo

`func (o *NrfInfoServedUpfInfoValue) HasTwifInfo() bool`

HasTwifInfo returns a boolean if a field has been set.

### SetTwifInfoNil

`func (o *NrfInfoServedUpfInfoValue) SetTwifInfoNil(b bool)`

 SetTwifInfoNil sets the value for TwifInfo to be an explicit nil

### UnsetTwifInfo
`func (o *NrfInfoServedUpfInfoValue) UnsetTwifInfo()`

UnsetTwifInfo ensures that no value is present for TwifInfo, not even an explicit nil
### GetPreferredEpdgInfoList

`func (o *NrfInfoServedUpfInfoValue) GetPreferredEpdgInfoList() []EpdgInfo`

GetPreferredEpdgInfoList returns the PreferredEpdgInfoList field if non-nil, zero value otherwise.

### GetPreferredEpdgInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetPreferredEpdgInfoListOk() (*[]EpdgInfo, bool)`

GetPreferredEpdgInfoListOk returns a tuple with the PreferredEpdgInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEpdgInfoList

`func (o *NrfInfoServedUpfInfoValue) SetPreferredEpdgInfoList(v []EpdgInfo)`

SetPreferredEpdgInfoList sets PreferredEpdgInfoList field to given value.

### HasPreferredEpdgInfoList

`func (o *NrfInfoServedUpfInfoValue) HasPreferredEpdgInfoList() bool`

HasPreferredEpdgInfoList returns a boolean if a field has been set.

### GetPreferredWAgfInfoList

`func (o *NrfInfoServedUpfInfoValue) GetPreferredWAgfInfoList() []WAgfInfo`

GetPreferredWAgfInfoList returns the PreferredWAgfInfoList field if non-nil, zero value otherwise.

### GetPreferredWAgfInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetPreferredWAgfInfoListOk() (*[]WAgfInfo, bool)`

GetPreferredWAgfInfoListOk returns a tuple with the PreferredWAgfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredWAgfInfoList

`func (o *NrfInfoServedUpfInfoValue) SetPreferredWAgfInfoList(v []WAgfInfo)`

SetPreferredWAgfInfoList sets PreferredWAgfInfoList field to given value.

### HasPreferredWAgfInfoList

`func (o *NrfInfoServedUpfInfoValue) HasPreferredWAgfInfoList() bool`

HasPreferredWAgfInfoList returns a boolean if a field has been set.

### GetPreferredTngfInfoList

`func (o *NrfInfoServedUpfInfoValue) GetPreferredTngfInfoList() []TngfInfo`

GetPreferredTngfInfoList returns the PreferredTngfInfoList field if non-nil, zero value otherwise.

### GetPreferredTngfInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetPreferredTngfInfoListOk() (*[]TngfInfo, bool)`

GetPreferredTngfInfoListOk returns a tuple with the PreferredTngfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTngfInfoList

`func (o *NrfInfoServedUpfInfoValue) SetPreferredTngfInfoList(v []TngfInfo)`

SetPreferredTngfInfoList sets PreferredTngfInfoList field to given value.

### HasPreferredTngfInfoList

`func (o *NrfInfoServedUpfInfoValue) HasPreferredTngfInfoList() bool`

HasPreferredTngfInfoList returns a boolean if a field has been set.

### GetPreferredTwifInfoList

`func (o *NrfInfoServedUpfInfoValue) GetPreferredTwifInfoList() []TwifInfo`

GetPreferredTwifInfoList returns the PreferredTwifInfoList field if non-nil, zero value otherwise.

### GetPreferredTwifInfoListOk

`func (o *NrfInfoServedUpfInfoValue) GetPreferredTwifInfoListOk() (*[]TwifInfo, bool)`

GetPreferredTwifInfoListOk returns a tuple with the PreferredTwifInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTwifInfoList

`func (o *NrfInfoServedUpfInfoValue) SetPreferredTwifInfoList(v []TwifInfo)`

SetPreferredTwifInfoList sets PreferredTwifInfoList field to given value.

### HasPreferredTwifInfoList

`func (o *NrfInfoServedUpfInfoValue) HasPreferredTwifInfoList() bool`

HasPreferredTwifInfoList returns a boolean if a field has been set.

### GetPriority

`func (o *NrfInfoServedUpfInfoValue) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NrfInfoServedUpfInfoValue) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NrfInfoServedUpfInfoValue) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NrfInfoServedUpfInfoValue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRedundantGtpu

`func (o *NrfInfoServedUpfInfoValue) GetRedundantGtpu() bool`

GetRedundantGtpu returns the RedundantGtpu field if non-nil, zero value otherwise.

### GetRedundantGtpuOk

`func (o *NrfInfoServedUpfInfoValue) GetRedundantGtpuOk() (*bool, bool)`

GetRedundantGtpuOk returns a tuple with the RedundantGtpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantGtpu

`func (o *NrfInfoServedUpfInfoValue) SetRedundantGtpu(v bool)`

SetRedundantGtpu sets RedundantGtpu field to given value.

### HasRedundantGtpu

`func (o *NrfInfoServedUpfInfoValue) HasRedundantGtpu() bool`

HasRedundantGtpu returns a boolean if a field has been set.

### GetIpups

`func (o *NrfInfoServedUpfInfoValue) GetIpups() bool`

GetIpups returns the Ipups field if non-nil, zero value otherwise.

### GetIpupsOk

`func (o *NrfInfoServedUpfInfoValue) GetIpupsOk() (*bool, bool)`

GetIpupsOk returns a tuple with the Ipups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpups

`func (o *NrfInfoServedUpfInfoValue) SetIpups(v bool)`

SetIpups sets Ipups field to given value.

### HasIpups

`func (o *NrfInfoServedUpfInfoValue) HasIpups() bool`

HasIpups returns a boolean if a field has been set.

### GetDataForwarding

`func (o *NrfInfoServedUpfInfoValue) GetDataForwarding() bool`

GetDataForwarding returns the DataForwarding field if non-nil, zero value otherwise.

### GetDataForwardingOk

`func (o *NrfInfoServedUpfInfoValue) GetDataForwardingOk() (*bool, bool)`

GetDataForwardingOk returns a tuple with the DataForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwarding

`func (o *NrfInfoServedUpfInfoValue) SetDataForwarding(v bool)`

SetDataForwarding sets DataForwarding field to given value.

### HasDataForwarding

`func (o *NrfInfoServedUpfInfoValue) HasDataForwarding() bool`

HasDataForwarding returns a boolean if a field has been set.

### GetSupportedPfcpFeatures

`func (o *NrfInfoServedUpfInfoValue) GetSupportedPfcpFeatures() string`

GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field if non-nil, zero value otherwise.

### GetSupportedPfcpFeaturesOk

`func (o *NrfInfoServedUpfInfoValue) GetSupportedPfcpFeaturesOk() (*string, bool)`

GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPfcpFeatures

`func (o *NrfInfoServedUpfInfoValue) SetSupportedPfcpFeatures(v string)`

SetSupportedPfcpFeatures sets SupportedPfcpFeatures field to given value.

### HasSupportedPfcpFeatures

`func (o *NrfInfoServedUpfInfoValue) HasSupportedPfcpFeatures() bool`

HasSupportedPfcpFeatures returns a boolean if a field has been set.

### GetUpfEvents

`func (o *NrfInfoServedUpfInfoValue) GetUpfEvents() []EventType`

GetUpfEvents returns the UpfEvents field if non-nil, zero value otherwise.

### GetUpfEventsOk

`func (o *NrfInfoServedUpfInfoValue) GetUpfEventsOk() (*[]EventType, bool)`

GetUpfEventsOk returns a tuple with the UpfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfEvents

`func (o *NrfInfoServedUpfInfoValue) SetUpfEvents(v []EventType)`

SetUpfEvents sets UpfEvents field to given value.

### HasUpfEvents

`func (o *NrfInfoServedUpfInfoValue) HasUpfEvents() bool`

HasUpfEvents returns a boolean if a field has been set.

### GetOpConfigCaps

`func (o *NrfInfoServedUpfInfoValue) GetOpConfigCaps() []string`

GetOpConfigCaps returns the OpConfigCaps field if non-nil, zero value otherwise.

### GetOpConfigCapsOk

`func (o *NrfInfoServedUpfInfoValue) GetOpConfigCapsOk() (*[]string, bool)`

GetOpConfigCapsOk returns a tuple with the OpConfigCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpConfigCaps

`func (o *NrfInfoServedUpfInfoValue) SetOpConfigCaps(v []string)`

SetOpConfigCaps sets OpConfigCaps field to given value.

### HasOpConfigCaps

`func (o *NrfInfoServedUpfInfoValue) HasOpConfigCaps() bool`

HasOpConfigCaps returns a boolean if a field has been set.

### GetPacketInspectionFunctionalities

`func (o *NrfInfoServedUpfInfoValue) GetPacketInspectionFunctionalities() []UpfPacketInspectionFunctionality`

GetPacketInspectionFunctionalities returns the PacketInspectionFunctionalities field if non-nil, zero value otherwise.

### GetPacketInspectionFunctionalitiesOk

`func (o *NrfInfoServedUpfInfoValue) GetPacketInspectionFunctionalitiesOk() (*[]UpfPacketInspectionFunctionality, bool)`

GetPacketInspectionFunctionalitiesOk returns a tuple with the PacketInspectionFunctionalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketInspectionFunctionalities

`func (o *NrfInfoServedUpfInfoValue) SetPacketInspectionFunctionalities(v []UpfPacketInspectionFunctionality)`

SetPacketInspectionFunctionalities sets PacketInspectionFunctionalities field to given value.

### HasPacketInspectionFunctionalities

`func (o *NrfInfoServedUpfInfoValue) HasPacketInspectionFunctionalities() bool`

HasPacketInspectionFunctionalities returns a boolean if a field has been set.

### GetN6DelayMeastProtocs

`func (o *NrfInfoServedUpfInfoValue) GetN6DelayMeastProtocs() []N6DelayMeasurementProtocol`

GetN6DelayMeastProtocs returns the N6DelayMeastProtocs field if non-nil, zero value otherwise.

### GetN6DelayMeastProtocsOk

`func (o *NrfInfoServedUpfInfoValue) GetN6DelayMeastProtocsOk() (*[]N6DelayMeasurementProtocol, bool)`

GetN6DelayMeastProtocsOk returns a tuple with the N6DelayMeastProtocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6DelayMeastProtocs

`func (o *NrfInfoServedUpfInfoValue) SetN6DelayMeastProtocs(v []N6DelayMeasurementProtocol)`

SetN6DelayMeastProtocs sets N6DelayMeastProtocs field to given value.

### HasN6DelayMeastProtocs

`func (o *NrfInfoServedUpfInfoValue) HasN6DelayMeastProtocs() bool`

HasN6DelayMeastProtocs returns a boolean if a field has been set.

### GetGeranUtranInd

`func (o *NrfInfoServedUpfInfoValue) GetGeranUtranInd() bool`

GetGeranUtranInd returns the GeranUtranInd field if non-nil, zero value otherwise.

### GetGeranUtranIndOk

`func (o *NrfInfoServedUpfInfoValue) GetGeranUtranIndOk() (*bool, bool)`

GetGeranUtranIndOk returns a tuple with the GeranUtranInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeranUtranInd

`func (o *NrfInfoServedUpfInfoValue) SetGeranUtranInd(v bool)`

SetGeranUtranInd sets GeranUtranInd field to given value.

### HasGeranUtranInd

`func (o *NrfInfoServedUpfInfoValue) HasGeranUtranInd() bool`

HasGeranUtranInd returns a boolean if a field has been set.

### GetVar2g3gLocationAreaList

`func (o *NrfInfoServedUpfInfoValue) GetVar2g3gLocationAreaList() []Model2G3GLocationArea`

GetVar2g3gLocationAreaList returns the Var2g3gLocationAreaList field if non-nil, zero value otherwise.

### GetVar2g3gLocationAreaListOk

`func (o *NrfInfoServedUpfInfoValue) GetVar2g3gLocationAreaListOk() (*[]Model2G3GLocationArea, bool)`

GetVar2g3gLocationAreaListOk returns a tuple with the Var2g3gLocationAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2g3gLocationAreaList

`func (o *NrfInfoServedUpfInfoValue) SetVar2g3gLocationAreaList(v []Model2G3GLocationArea)`

SetVar2g3gLocationAreaList sets Var2g3gLocationAreaList field to given value.

### HasVar2g3gLocationAreaList

`func (o *NrfInfoServedUpfInfoValue) HasVar2g3gLocationAreaList() bool`

HasVar2g3gLocationAreaList returns a boolean if a field has been set.

### GetVar2g3gLocationAreaRangeList

`func (o *NrfInfoServedUpfInfoValue) GetVar2g3gLocationAreaRangeList() []Model2G3GLocationAreaRange`

GetVar2g3gLocationAreaRangeList returns the Var2g3gLocationAreaRangeList field if non-nil, zero value otherwise.

### GetVar2g3gLocationAreaRangeListOk

`func (o *NrfInfoServedUpfInfoValue) GetVar2g3gLocationAreaRangeListOk() (*[]Model2G3GLocationAreaRange, bool)`

GetVar2g3gLocationAreaRangeListOk returns a tuple with the Var2g3gLocationAreaRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2g3gLocationAreaRangeList

`func (o *NrfInfoServedUpfInfoValue) SetVar2g3gLocationAreaRangeList(v []Model2G3GLocationAreaRange)`

SetVar2g3gLocationAreaRangeList sets Var2g3gLocationAreaRangeList field to given value.

### HasVar2g3gLocationAreaRangeList

`func (o *NrfInfoServedUpfInfoValue) HasVar2g3gLocationAreaRangeList() bool`

HasVar2g3gLocationAreaRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



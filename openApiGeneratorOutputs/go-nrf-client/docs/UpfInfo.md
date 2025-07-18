# UpfInfo

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

### NewUpfInfo

`func NewUpfInfo(sNssaiUpfInfoList []SnssaiUpfInfoItem, ) *UpfInfo`

NewUpfInfo instantiates a new UpfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpfInfoWithDefaults

`func NewUpfInfoWithDefaults() *UpfInfo`

NewUpfInfoWithDefaults instantiates a new UpfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiUpfInfoList

`func (o *UpfInfo) GetSNssaiUpfInfoList() []SnssaiUpfInfoItem`

GetSNssaiUpfInfoList returns the SNssaiUpfInfoList field if non-nil, zero value otherwise.

### GetSNssaiUpfInfoListOk

`func (o *UpfInfo) GetSNssaiUpfInfoListOk() (*[]SnssaiUpfInfoItem, bool)`

GetSNssaiUpfInfoListOk returns a tuple with the SNssaiUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiUpfInfoList

`func (o *UpfInfo) SetSNssaiUpfInfoList(v []SnssaiUpfInfoItem)`

SetSNssaiUpfInfoList sets SNssaiUpfInfoList field to given value.


### GetSmfServingArea

`func (o *UpfInfo) GetSmfServingArea() []string`

GetSmfServingArea returns the SmfServingArea field if non-nil, zero value otherwise.

### GetSmfServingAreaOk

`func (o *UpfInfo) GetSmfServingAreaOk() (*[]string, bool)`

GetSmfServingAreaOk returns a tuple with the SmfServingArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServingArea

`func (o *UpfInfo) SetSmfServingArea(v []string)`

SetSmfServingArea sets SmfServingArea field to given value.

### HasSmfServingArea

`func (o *UpfInfo) HasSmfServingArea() bool`

HasSmfServingArea returns a boolean if a field has been set.

### GetInterfaceUpfInfoList

`func (o *UpfInfo) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem`

GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field if non-nil, zero value otherwise.

### GetInterfaceUpfInfoListOk

`func (o *UpfInfo) GetInterfaceUpfInfoListOk() (*[]InterfaceUpfInfoItem, bool)`

GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUpfInfoList

`func (o *UpfInfo) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem)`

SetInterfaceUpfInfoList sets InterfaceUpfInfoList field to given value.

### HasInterfaceUpfInfoList

`func (o *UpfInfo) HasInterfaceUpfInfoList() bool`

HasInterfaceUpfInfoList returns a boolean if a field has been set.

### GetN6TunnelInfoList

`func (o *UpfInfo) GetN6TunnelInfoList() map[string]InterfaceUpfInfoItem`

GetN6TunnelInfoList returns the N6TunnelInfoList field if non-nil, zero value otherwise.

### GetN6TunnelInfoListOk

`func (o *UpfInfo) GetN6TunnelInfoListOk() (*map[string]InterfaceUpfInfoItem, bool)`

GetN6TunnelInfoListOk returns a tuple with the N6TunnelInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6TunnelInfoList

`func (o *UpfInfo) SetN6TunnelInfoList(v map[string]InterfaceUpfInfoItem)`

SetN6TunnelInfoList sets N6TunnelInfoList field to given value.

### HasN6TunnelInfoList

`func (o *UpfInfo) HasN6TunnelInfoList() bool`

HasN6TunnelInfoList returns a boolean if a field has been set.

### GetIwkEpsInd

`func (o *UpfInfo) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *UpfInfo) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *UpfInfo) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *UpfInfo) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetSxaInd

`func (o *UpfInfo) GetSxaInd() bool`

GetSxaInd returns the SxaInd field if non-nil, zero value otherwise.

### GetSxaIndOk

`func (o *UpfInfo) GetSxaIndOk() (*bool, bool)`

GetSxaIndOk returns a tuple with the SxaInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSxaInd

`func (o *UpfInfo) SetSxaInd(v bool)`

SetSxaInd sets SxaInd field to given value.

### HasSxaInd

`func (o *UpfInfo) HasSxaInd() bool`

HasSxaInd returns a boolean if a field has been set.

### GetPduSessionTypes

`func (o *UpfInfo) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *UpfInfo) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *UpfInfo) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *UpfInfo) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetAtsssCapability

`func (o *UpfInfo) GetAtsssCapability() AtsssCapability`

GetAtsssCapability returns the AtsssCapability field if non-nil, zero value otherwise.

### GetAtsssCapabilityOk

`func (o *UpfInfo) GetAtsssCapabilityOk() (*AtsssCapability, bool)`

GetAtsssCapabilityOk returns a tuple with the AtsssCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssCapability

`func (o *UpfInfo) SetAtsssCapability(v AtsssCapability)`

SetAtsssCapability sets AtsssCapability field to given value.

### HasAtsssCapability

`func (o *UpfInfo) HasAtsssCapability() bool`

HasAtsssCapability returns a boolean if a field has been set.

### GetUeIpAddrInd

`func (o *UpfInfo) GetUeIpAddrInd() bool`

GetUeIpAddrInd returns the UeIpAddrInd field if non-nil, zero value otherwise.

### GetUeIpAddrIndOk

`func (o *UpfInfo) GetUeIpAddrIndOk() (*bool, bool)`

GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddrInd

`func (o *UpfInfo) SetUeIpAddrInd(v bool)`

SetUeIpAddrInd sets UeIpAddrInd field to given value.

### HasUeIpAddrInd

`func (o *UpfInfo) HasUeIpAddrInd() bool`

HasUeIpAddrInd returns a boolean if a field has been set.

### GetTaiList

`func (o *UpfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *UpfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *UpfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *UpfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *UpfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *UpfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *UpfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *UpfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetWAgfInfo

`func (o *UpfInfo) GetWAgfInfo() WAgfInfo`

GetWAgfInfo returns the WAgfInfo field if non-nil, zero value otherwise.

### GetWAgfInfoOk

`func (o *UpfInfo) GetWAgfInfoOk() (*WAgfInfo, bool)`

GetWAgfInfoOk returns a tuple with the WAgfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAgfInfo

`func (o *UpfInfo) SetWAgfInfo(v WAgfInfo)`

SetWAgfInfo sets WAgfInfo field to given value.

### HasWAgfInfo

`func (o *UpfInfo) HasWAgfInfo() bool`

HasWAgfInfo returns a boolean if a field has been set.

### SetWAgfInfoNil

`func (o *UpfInfo) SetWAgfInfoNil(b bool)`

 SetWAgfInfoNil sets the value for WAgfInfo to be an explicit nil

### UnsetWAgfInfo
`func (o *UpfInfo) UnsetWAgfInfo()`

UnsetWAgfInfo ensures that no value is present for WAgfInfo, not even an explicit nil
### GetTngfInfo

`func (o *UpfInfo) GetTngfInfo() TngfInfo`

GetTngfInfo returns the TngfInfo field if non-nil, zero value otherwise.

### GetTngfInfoOk

`func (o *UpfInfo) GetTngfInfoOk() (*TngfInfo, bool)`

GetTngfInfoOk returns a tuple with the TngfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfInfo

`func (o *UpfInfo) SetTngfInfo(v TngfInfo)`

SetTngfInfo sets TngfInfo field to given value.

### HasTngfInfo

`func (o *UpfInfo) HasTngfInfo() bool`

HasTngfInfo returns a boolean if a field has been set.

### SetTngfInfoNil

`func (o *UpfInfo) SetTngfInfoNil(b bool)`

 SetTngfInfoNil sets the value for TngfInfo to be an explicit nil

### UnsetTngfInfo
`func (o *UpfInfo) UnsetTngfInfo()`

UnsetTngfInfo ensures that no value is present for TngfInfo, not even an explicit nil
### GetTwifInfo

`func (o *UpfInfo) GetTwifInfo() TwifInfo`

GetTwifInfo returns the TwifInfo field if non-nil, zero value otherwise.

### GetTwifInfoOk

`func (o *UpfInfo) GetTwifInfoOk() (*TwifInfo, bool)`

GetTwifInfoOk returns a tuple with the TwifInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwifInfo

`func (o *UpfInfo) SetTwifInfo(v TwifInfo)`

SetTwifInfo sets TwifInfo field to given value.

### HasTwifInfo

`func (o *UpfInfo) HasTwifInfo() bool`

HasTwifInfo returns a boolean if a field has been set.

### SetTwifInfoNil

`func (o *UpfInfo) SetTwifInfoNil(b bool)`

 SetTwifInfoNil sets the value for TwifInfo to be an explicit nil

### UnsetTwifInfo
`func (o *UpfInfo) UnsetTwifInfo()`

UnsetTwifInfo ensures that no value is present for TwifInfo, not even an explicit nil
### GetPreferredEpdgInfoList

`func (o *UpfInfo) GetPreferredEpdgInfoList() []EpdgInfo`

GetPreferredEpdgInfoList returns the PreferredEpdgInfoList field if non-nil, zero value otherwise.

### GetPreferredEpdgInfoListOk

`func (o *UpfInfo) GetPreferredEpdgInfoListOk() (*[]EpdgInfo, bool)`

GetPreferredEpdgInfoListOk returns a tuple with the PreferredEpdgInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEpdgInfoList

`func (o *UpfInfo) SetPreferredEpdgInfoList(v []EpdgInfo)`

SetPreferredEpdgInfoList sets PreferredEpdgInfoList field to given value.

### HasPreferredEpdgInfoList

`func (o *UpfInfo) HasPreferredEpdgInfoList() bool`

HasPreferredEpdgInfoList returns a boolean if a field has been set.

### GetPreferredWAgfInfoList

`func (o *UpfInfo) GetPreferredWAgfInfoList() []WAgfInfo`

GetPreferredWAgfInfoList returns the PreferredWAgfInfoList field if non-nil, zero value otherwise.

### GetPreferredWAgfInfoListOk

`func (o *UpfInfo) GetPreferredWAgfInfoListOk() (*[]WAgfInfo, bool)`

GetPreferredWAgfInfoListOk returns a tuple with the PreferredWAgfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredWAgfInfoList

`func (o *UpfInfo) SetPreferredWAgfInfoList(v []WAgfInfo)`

SetPreferredWAgfInfoList sets PreferredWAgfInfoList field to given value.

### HasPreferredWAgfInfoList

`func (o *UpfInfo) HasPreferredWAgfInfoList() bool`

HasPreferredWAgfInfoList returns a boolean if a field has been set.

### GetPreferredTngfInfoList

`func (o *UpfInfo) GetPreferredTngfInfoList() []TngfInfo`

GetPreferredTngfInfoList returns the PreferredTngfInfoList field if non-nil, zero value otherwise.

### GetPreferredTngfInfoListOk

`func (o *UpfInfo) GetPreferredTngfInfoListOk() (*[]TngfInfo, bool)`

GetPreferredTngfInfoListOk returns a tuple with the PreferredTngfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTngfInfoList

`func (o *UpfInfo) SetPreferredTngfInfoList(v []TngfInfo)`

SetPreferredTngfInfoList sets PreferredTngfInfoList field to given value.

### HasPreferredTngfInfoList

`func (o *UpfInfo) HasPreferredTngfInfoList() bool`

HasPreferredTngfInfoList returns a boolean if a field has been set.

### GetPreferredTwifInfoList

`func (o *UpfInfo) GetPreferredTwifInfoList() []TwifInfo`

GetPreferredTwifInfoList returns the PreferredTwifInfoList field if non-nil, zero value otherwise.

### GetPreferredTwifInfoListOk

`func (o *UpfInfo) GetPreferredTwifInfoListOk() (*[]TwifInfo, bool)`

GetPreferredTwifInfoListOk returns a tuple with the PreferredTwifInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTwifInfoList

`func (o *UpfInfo) SetPreferredTwifInfoList(v []TwifInfo)`

SetPreferredTwifInfoList sets PreferredTwifInfoList field to given value.

### HasPreferredTwifInfoList

`func (o *UpfInfo) HasPreferredTwifInfoList() bool`

HasPreferredTwifInfoList returns a boolean if a field has been set.

### GetPriority

`func (o *UpfInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpfInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpfInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpfInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRedundantGtpu

`func (o *UpfInfo) GetRedundantGtpu() bool`

GetRedundantGtpu returns the RedundantGtpu field if non-nil, zero value otherwise.

### GetRedundantGtpuOk

`func (o *UpfInfo) GetRedundantGtpuOk() (*bool, bool)`

GetRedundantGtpuOk returns a tuple with the RedundantGtpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantGtpu

`func (o *UpfInfo) SetRedundantGtpu(v bool)`

SetRedundantGtpu sets RedundantGtpu field to given value.

### HasRedundantGtpu

`func (o *UpfInfo) HasRedundantGtpu() bool`

HasRedundantGtpu returns a boolean if a field has been set.

### GetIpups

`func (o *UpfInfo) GetIpups() bool`

GetIpups returns the Ipups field if non-nil, zero value otherwise.

### GetIpupsOk

`func (o *UpfInfo) GetIpupsOk() (*bool, bool)`

GetIpupsOk returns a tuple with the Ipups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpups

`func (o *UpfInfo) SetIpups(v bool)`

SetIpups sets Ipups field to given value.

### HasIpups

`func (o *UpfInfo) HasIpups() bool`

HasIpups returns a boolean if a field has been set.

### GetDataForwarding

`func (o *UpfInfo) GetDataForwarding() bool`

GetDataForwarding returns the DataForwarding field if non-nil, zero value otherwise.

### GetDataForwardingOk

`func (o *UpfInfo) GetDataForwardingOk() (*bool, bool)`

GetDataForwardingOk returns a tuple with the DataForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwarding

`func (o *UpfInfo) SetDataForwarding(v bool)`

SetDataForwarding sets DataForwarding field to given value.

### HasDataForwarding

`func (o *UpfInfo) HasDataForwarding() bool`

HasDataForwarding returns a boolean if a field has been set.

### GetSupportedPfcpFeatures

`func (o *UpfInfo) GetSupportedPfcpFeatures() string`

GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field if non-nil, zero value otherwise.

### GetSupportedPfcpFeaturesOk

`func (o *UpfInfo) GetSupportedPfcpFeaturesOk() (*string, bool)`

GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPfcpFeatures

`func (o *UpfInfo) SetSupportedPfcpFeatures(v string)`

SetSupportedPfcpFeatures sets SupportedPfcpFeatures field to given value.

### HasSupportedPfcpFeatures

`func (o *UpfInfo) HasSupportedPfcpFeatures() bool`

HasSupportedPfcpFeatures returns a boolean if a field has been set.

### GetUpfEvents

`func (o *UpfInfo) GetUpfEvents() []EventType`

GetUpfEvents returns the UpfEvents field if non-nil, zero value otherwise.

### GetUpfEventsOk

`func (o *UpfInfo) GetUpfEventsOk() (*[]EventType, bool)`

GetUpfEventsOk returns a tuple with the UpfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfEvents

`func (o *UpfInfo) SetUpfEvents(v []EventType)`

SetUpfEvents sets UpfEvents field to given value.

### HasUpfEvents

`func (o *UpfInfo) HasUpfEvents() bool`

HasUpfEvents returns a boolean if a field has been set.

### GetOpConfigCaps

`func (o *UpfInfo) GetOpConfigCaps() []string`

GetOpConfigCaps returns the OpConfigCaps field if non-nil, zero value otherwise.

### GetOpConfigCapsOk

`func (o *UpfInfo) GetOpConfigCapsOk() (*[]string, bool)`

GetOpConfigCapsOk returns a tuple with the OpConfigCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpConfigCaps

`func (o *UpfInfo) SetOpConfigCaps(v []string)`

SetOpConfigCaps sets OpConfigCaps field to given value.

### HasOpConfigCaps

`func (o *UpfInfo) HasOpConfigCaps() bool`

HasOpConfigCaps returns a boolean if a field has been set.

### GetPacketInspectionFunctionalities

`func (o *UpfInfo) GetPacketInspectionFunctionalities() []UpfPacketInspectionFunctionality`

GetPacketInspectionFunctionalities returns the PacketInspectionFunctionalities field if non-nil, zero value otherwise.

### GetPacketInspectionFunctionalitiesOk

`func (o *UpfInfo) GetPacketInspectionFunctionalitiesOk() (*[]UpfPacketInspectionFunctionality, bool)`

GetPacketInspectionFunctionalitiesOk returns a tuple with the PacketInspectionFunctionalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketInspectionFunctionalities

`func (o *UpfInfo) SetPacketInspectionFunctionalities(v []UpfPacketInspectionFunctionality)`

SetPacketInspectionFunctionalities sets PacketInspectionFunctionalities field to given value.

### HasPacketInspectionFunctionalities

`func (o *UpfInfo) HasPacketInspectionFunctionalities() bool`

HasPacketInspectionFunctionalities returns a boolean if a field has been set.

### GetN6DelayMeastProtocs

`func (o *UpfInfo) GetN6DelayMeastProtocs() []N6DelayMeasurementProtocol`

GetN6DelayMeastProtocs returns the N6DelayMeastProtocs field if non-nil, zero value otherwise.

### GetN6DelayMeastProtocsOk

`func (o *UpfInfo) GetN6DelayMeastProtocsOk() (*[]N6DelayMeasurementProtocol, bool)`

GetN6DelayMeastProtocsOk returns a tuple with the N6DelayMeastProtocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6DelayMeastProtocs

`func (o *UpfInfo) SetN6DelayMeastProtocs(v []N6DelayMeasurementProtocol)`

SetN6DelayMeastProtocs sets N6DelayMeastProtocs field to given value.

### HasN6DelayMeastProtocs

`func (o *UpfInfo) HasN6DelayMeastProtocs() bool`

HasN6DelayMeastProtocs returns a boolean if a field has been set.

### GetGeranUtranInd

`func (o *UpfInfo) GetGeranUtranInd() bool`

GetGeranUtranInd returns the GeranUtranInd field if non-nil, zero value otherwise.

### GetGeranUtranIndOk

`func (o *UpfInfo) GetGeranUtranIndOk() (*bool, bool)`

GetGeranUtranIndOk returns a tuple with the GeranUtranInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeranUtranInd

`func (o *UpfInfo) SetGeranUtranInd(v bool)`

SetGeranUtranInd sets GeranUtranInd field to given value.

### HasGeranUtranInd

`func (o *UpfInfo) HasGeranUtranInd() bool`

HasGeranUtranInd returns a boolean if a field has been set.

### GetVar2g3gLocationAreaList

`func (o *UpfInfo) GetVar2g3gLocationAreaList() []Model2G3GLocationArea`

GetVar2g3gLocationAreaList returns the Var2g3gLocationAreaList field if non-nil, zero value otherwise.

### GetVar2g3gLocationAreaListOk

`func (o *UpfInfo) GetVar2g3gLocationAreaListOk() (*[]Model2G3GLocationArea, bool)`

GetVar2g3gLocationAreaListOk returns a tuple with the Var2g3gLocationAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2g3gLocationAreaList

`func (o *UpfInfo) SetVar2g3gLocationAreaList(v []Model2G3GLocationArea)`

SetVar2g3gLocationAreaList sets Var2g3gLocationAreaList field to given value.

### HasVar2g3gLocationAreaList

`func (o *UpfInfo) HasVar2g3gLocationAreaList() bool`

HasVar2g3gLocationAreaList returns a boolean if a field has been set.

### GetVar2g3gLocationAreaRangeList

`func (o *UpfInfo) GetVar2g3gLocationAreaRangeList() []Model2G3GLocationAreaRange`

GetVar2g3gLocationAreaRangeList returns the Var2g3gLocationAreaRangeList field if non-nil, zero value otherwise.

### GetVar2g3gLocationAreaRangeListOk

`func (o *UpfInfo) GetVar2g3gLocationAreaRangeListOk() (*[]Model2G3GLocationAreaRange, bool)`

GetVar2g3gLocationAreaRangeListOk returns a tuple with the Var2g3gLocationAreaRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2g3gLocationAreaRangeList

`func (o *UpfInfo) SetVar2g3gLocationAreaRangeList(v []Model2G3GLocationAreaRange)`

SetVar2g3gLocationAreaRangeList sets Var2g3gLocationAreaRangeList field to given value.

### HasVar2g3gLocationAreaRangeList

`func (o *UpfInfo) HasVar2g3gLocationAreaRangeList() bool`

HasVar2g3gLocationAreaRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



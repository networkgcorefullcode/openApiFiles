# UpfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiUpfInfoList** | [**[]SnssaiUpfInfoItem**](SnssaiUpfInfoItem.md) |  | 
**SmfServingArea** | Pointer to **[]string** |  | [optional] 
**InterfaceUpfInfoList** | Pointer to [**[]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) |  | [optional] 
**IwkEpsInd** | Pointer to **bool** |  | [optional] [default to false]
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**AtsssCapability** | Pointer to [**AtsssCapability**](AtsssCapability.md) |  | [optional] 
**UeIpAddrInd** | Pointer to **bool** |  | [optional] [default to false]

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



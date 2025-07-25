# DnnUpfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**Ipv4AddressRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**Ipv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 
**NatedIpv4AddressRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**NatedIpv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 
**Ipv4IndexList** | Pointer to [**[]IpIndex**](IpIndex.md) |  | [optional] 
**Ipv6IndexList** | Pointer to [**[]IpIndex**](IpIndex.md) |  | [optional] 
**NetworkInstance** | Pointer to **string** | The N6 Network Instance associated with the S-NSSAI and DNN.  | [optional] 
**DnaiNwInstanceList** | Pointer to **map[string]string** | Map of network instance per DNAI for the DNN, where the key of the map is the DNAI. When present, the value of each entry of the map shall contain a N6 network instance that is configured for the DNAI indicated by the key.  | [optional] 
**InterfaceUpfInfoList** | Pointer to [**[]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) |  | [optional] 
**PrivateIpv4AddressRangesPerIpDomain** | Pointer to [**map[string][]Ipv4AddressRange**](array.md) | Map of private IPv4 Address Ranges Per Ip Domain, where the key of the map is the IP. Domain. When present, the value of each entry of the map shall contain a IPv4 private address ranges configured for that IP domain.  | [optional] 

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

### GetIpv4AddressRanges

`func (o *DnnUpfInfoItem) GetIpv4AddressRanges() []Ipv4AddressRange`

GetIpv4AddressRanges returns the Ipv4AddressRanges field if non-nil, zero value otherwise.

### GetIpv4AddressRangesOk

`func (o *DnnUpfInfoItem) GetIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool)`

GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressRanges

`func (o *DnnUpfInfoItem) SetIpv4AddressRanges(v []Ipv4AddressRange)`

SetIpv4AddressRanges sets Ipv4AddressRanges field to given value.

### HasIpv4AddressRanges

`func (o *DnnUpfInfoItem) HasIpv4AddressRanges() bool`

HasIpv4AddressRanges returns a boolean if a field has been set.

### GetIpv6PrefixRanges

`func (o *DnnUpfInfoItem) GetIpv6PrefixRanges() []Ipv6PrefixRange`

GetIpv6PrefixRanges returns the Ipv6PrefixRanges field if non-nil, zero value otherwise.

### GetIpv6PrefixRangesOk

`func (o *DnnUpfInfoItem) GetIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixRanges

`func (o *DnnUpfInfoItem) SetIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetIpv6PrefixRanges sets Ipv6PrefixRanges field to given value.

### HasIpv6PrefixRanges

`func (o *DnnUpfInfoItem) HasIpv6PrefixRanges() bool`

HasIpv6PrefixRanges returns a boolean if a field has been set.

### GetNatedIpv4AddressRanges

`func (o *DnnUpfInfoItem) GetNatedIpv4AddressRanges() []Ipv4AddressRange`

GetNatedIpv4AddressRanges returns the NatedIpv4AddressRanges field if non-nil, zero value otherwise.

### GetNatedIpv4AddressRangesOk

`func (o *DnnUpfInfoItem) GetNatedIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool)`

GetNatedIpv4AddressRangesOk returns a tuple with the NatedIpv4AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatedIpv4AddressRanges

`func (o *DnnUpfInfoItem) SetNatedIpv4AddressRanges(v []Ipv4AddressRange)`

SetNatedIpv4AddressRanges sets NatedIpv4AddressRanges field to given value.

### HasNatedIpv4AddressRanges

`func (o *DnnUpfInfoItem) HasNatedIpv4AddressRanges() bool`

HasNatedIpv4AddressRanges returns a boolean if a field has been set.

### GetNatedIpv6PrefixRanges

`func (o *DnnUpfInfoItem) GetNatedIpv6PrefixRanges() []Ipv6PrefixRange`

GetNatedIpv6PrefixRanges returns the NatedIpv6PrefixRanges field if non-nil, zero value otherwise.

### GetNatedIpv6PrefixRangesOk

`func (o *DnnUpfInfoItem) GetNatedIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetNatedIpv6PrefixRangesOk returns a tuple with the NatedIpv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatedIpv6PrefixRanges

`func (o *DnnUpfInfoItem) SetNatedIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetNatedIpv6PrefixRanges sets NatedIpv6PrefixRanges field to given value.

### HasNatedIpv6PrefixRanges

`func (o *DnnUpfInfoItem) HasNatedIpv6PrefixRanges() bool`

HasNatedIpv6PrefixRanges returns a boolean if a field has been set.

### GetIpv4IndexList

`func (o *DnnUpfInfoItem) GetIpv4IndexList() []IpIndex`

GetIpv4IndexList returns the Ipv4IndexList field if non-nil, zero value otherwise.

### GetIpv4IndexListOk

`func (o *DnnUpfInfoItem) GetIpv4IndexListOk() (*[]IpIndex, bool)`

GetIpv4IndexListOk returns a tuple with the Ipv4IndexList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4IndexList

`func (o *DnnUpfInfoItem) SetIpv4IndexList(v []IpIndex)`

SetIpv4IndexList sets Ipv4IndexList field to given value.

### HasIpv4IndexList

`func (o *DnnUpfInfoItem) HasIpv4IndexList() bool`

HasIpv4IndexList returns a boolean if a field has been set.

### GetIpv6IndexList

`func (o *DnnUpfInfoItem) GetIpv6IndexList() []IpIndex`

GetIpv6IndexList returns the Ipv6IndexList field if non-nil, zero value otherwise.

### GetIpv6IndexListOk

`func (o *DnnUpfInfoItem) GetIpv6IndexListOk() (*[]IpIndex, bool)`

GetIpv6IndexListOk returns a tuple with the Ipv6IndexList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6IndexList

`func (o *DnnUpfInfoItem) SetIpv6IndexList(v []IpIndex)`

SetIpv6IndexList sets Ipv6IndexList field to given value.

### HasIpv6IndexList

`func (o *DnnUpfInfoItem) HasIpv6IndexList() bool`

HasIpv6IndexList returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *DnnUpfInfoItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *DnnUpfInfoItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *DnnUpfInfoItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *DnnUpfInfoItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetDnaiNwInstanceList

`func (o *DnnUpfInfoItem) GetDnaiNwInstanceList() map[string]string`

GetDnaiNwInstanceList returns the DnaiNwInstanceList field if non-nil, zero value otherwise.

### GetDnaiNwInstanceListOk

`func (o *DnnUpfInfoItem) GetDnaiNwInstanceListOk() (*map[string]string, bool)`

GetDnaiNwInstanceListOk returns a tuple with the DnaiNwInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiNwInstanceList

`func (o *DnnUpfInfoItem) SetDnaiNwInstanceList(v map[string]string)`

SetDnaiNwInstanceList sets DnaiNwInstanceList field to given value.

### HasDnaiNwInstanceList

`func (o *DnnUpfInfoItem) HasDnaiNwInstanceList() bool`

HasDnaiNwInstanceList returns a boolean if a field has been set.

### GetInterfaceUpfInfoList

`func (o *DnnUpfInfoItem) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem`

GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field if non-nil, zero value otherwise.

### GetInterfaceUpfInfoListOk

`func (o *DnnUpfInfoItem) GetInterfaceUpfInfoListOk() (*[]InterfaceUpfInfoItem, bool)`

GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUpfInfoList

`func (o *DnnUpfInfoItem) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem)`

SetInterfaceUpfInfoList sets InterfaceUpfInfoList field to given value.

### HasInterfaceUpfInfoList

`func (o *DnnUpfInfoItem) HasInterfaceUpfInfoList() bool`

HasInterfaceUpfInfoList returns a boolean if a field has been set.

### GetPrivateIpv4AddressRangesPerIpDomain

`func (o *DnnUpfInfoItem) GetPrivateIpv4AddressRangesPerIpDomain() map[string][]Ipv4AddressRange`

GetPrivateIpv4AddressRangesPerIpDomain returns the PrivateIpv4AddressRangesPerIpDomain field if non-nil, zero value otherwise.

### GetPrivateIpv4AddressRangesPerIpDomainOk

`func (o *DnnUpfInfoItem) GetPrivateIpv4AddressRangesPerIpDomainOk() (*map[string][]Ipv4AddressRange, bool)`

GetPrivateIpv4AddressRangesPerIpDomainOk returns a tuple with the PrivateIpv4AddressRangesPerIpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4AddressRangesPerIpDomain

`func (o *DnnUpfInfoItem) SetPrivateIpv4AddressRangesPerIpDomain(v map[string][]Ipv4AddressRange)`

SetPrivateIpv4AddressRangesPerIpDomain sets PrivateIpv4AddressRangesPerIpDomain field to given value.

### HasPrivateIpv4AddressRangesPerIpDomain

`func (o *DnnUpfInfoItem) HasPrivateIpv4AddressRangesPerIpDomain() bool`

HasPrivateIpv4AddressRangesPerIpDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



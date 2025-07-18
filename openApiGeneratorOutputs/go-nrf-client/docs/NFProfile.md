# NFProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** |  | 
**NfInstanceName** | Pointer to **string** |  | [optional] 
**NfType** | [**NFType**](NFType.md) |  | 
**NfStatus** | [**NFStatus**](NFStatus.md) |  | 
**HeartBeatTimer** | Pointer to **int32** |  | [optional] 
**PlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**SNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**PerPlmnSnssaiList** | Pointer to [**[]PlmnSnssai**](PlmnSnssai.md) |  | [optional] 
**NsiList** | Pointer to **[]string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**InterPlmnFqdn** | Pointer to **string** |  | [optional] 
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AllowedPlmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AllowedNfDomains** | Pointer to **[]string** |  | [optional] 
**AllowedNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Load** | Pointer to **int32** |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**UdrInfo** | Pointer to [**UdrInfo**](UdrInfo.md) |  | [optional] 
**UdrInfoExt** | Pointer to [**[]UdrInfo**](UdrInfo.md) |  | [optional] 
**UdmInfo** | Pointer to [**UdmInfo**](UdmInfo.md) |  | [optional] 
**UdmInfoExt** | Pointer to [**[]UdmInfo**](UdmInfo.md) |  | [optional] 
**AusfInfo** | Pointer to [**AusfInfo**](AusfInfo.md) |  | [optional] 
**AusfInfoExt** | Pointer to [**[]AusfInfo**](AusfInfo.md) |  | [optional] 
**AmfInfo** | Pointer to [**AmfInfo**](AmfInfo.md) |  | [optional] 
**AmfInfoExt** | Pointer to [**[]AmfInfo**](AmfInfo.md) |  | [optional] 
**SmfInfo** | Pointer to [**SmfInfo**](SmfInfo.md) |  | [optional] 
**SmfInfoExt** | Pointer to [**[]SmfInfo**](SmfInfo.md) |  | [optional] 
**UpfInfo** | Pointer to [**UpfInfo**](UpfInfo.md) |  | [optional] 
**UpfInfoExt** | Pointer to [**[]UpfInfo**](UpfInfo.md) |  | [optional] 
**PcfInfo** | Pointer to [**PcfInfo**](PcfInfo.md) |  | [optional] 
**PcfInfoExt** | Pointer to [**[]PcfInfo**](PcfInfo.md) |  | [optional] 
**BsfInfo** | Pointer to [**BsfInfo**](BsfInfo.md) |  | [optional] 
**BsfInfoExt** | Pointer to [**[]BsfInfo**](BsfInfo.md) |  | [optional] 
**ChfInfo** | Pointer to [**ChfInfo**](ChfInfo.md) |  | [optional] 
**ChfInfoExt** | Pointer to [**[]ChfInfo**](ChfInfo.md) |  | [optional] 
**NrfInfo** | Pointer to [**NrfInfo**](NrfInfo.md) |  | [optional] 
**NwdafInfo** | Pointer to [**NwdafInfo**](NwdafInfo.md) |  | [optional] 
**CustomInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 
**NfServicePersistence** | Pointer to **bool** |  | [optional] [default to false]
**NfServices** | Pointer to [**[]NFService**](NFService.md) |  | [optional] 
**NfProfileChangesSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**NfProfileChangesInd** | Pointer to **bool** |  | [optional] [readonly] [default to false]
**DefaultNotificationSubscriptions** | Pointer to [**[]DefaultNotificationSubscription**](DefaultNotificationSubscription.md) |  | [optional] 

## Methods

### NewNFProfile

`func NewNFProfile(nfInstanceId string, nfType NFType, nfStatus NFStatus, ) *NFProfile`

NewNFProfile instantiates a new NFProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFProfileWithDefaults

`func NewNFProfileWithDefaults() *NFProfile`

NewNFProfileWithDefaults instantiates a new NFProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *NFProfile) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NFProfile) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NFProfile) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfInstanceName

`func (o *NFProfile) GetNfInstanceName() string`

GetNfInstanceName returns the NfInstanceName field if non-nil, zero value otherwise.

### GetNfInstanceNameOk

`func (o *NFProfile) GetNfInstanceNameOk() (*string, bool)`

GetNfInstanceNameOk returns a tuple with the NfInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceName

`func (o *NFProfile) SetNfInstanceName(v string)`

SetNfInstanceName sets NfInstanceName field to given value.

### HasNfInstanceName

`func (o *NFProfile) HasNfInstanceName() bool`

HasNfInstanceName returns a boolean if a field has been set.

### GetNfType

`func (o *NFProfile) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *NFProfile) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *NFProfile) SetNfType(v NFType)`

SetNfType sets NfType field to given value.


### GetNfStatus

`func (o *NFProfile) GetNfStatus() NFStatus`

GetNfStatus returns the NfStatus field if non-nil, zero value otherwise.

### GetNfStatusOk

`func (o *NFProfile) GetNfStatusOk() (*NFStatus, bool)`

GetNfStatusOk returns a tuple with the NfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStatus

`func (o *NFProfile) SetNfStatus(v NFStatus)`

SetNfStatus sets NfStatus field to given value.


### GetHeartBeatTimer

`func (o *NFProfile) GetHeartBeatTimer() int32`

GetHeartBeatTimer returns the HeartBeatTimer field if non-nil, zero value otherwise.

### GetHeartBeatTimerOk

`func (o *NFProfile) GetHeartBeatTimerOk() (*int32, bool)`

GetHeartBeatTimerOk returns a tuple with the HeartBeatTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTimer

`func (o *NFProfile) SetHeartBeatTimer(v int32)`

SetHeartBeatTimer sets HeartBeatTimer field to given value.

### HasHeartBeatTimer

`func (o *NFProfile) HasHeartBeatTimer() bool`

HasHeartBeatTimer returns a boolean if a field has been set.

### GetPlmnList

`func (o *NFProfile) GetPlmnList() []PlmnId`

GetPlmnList returns the PlmnList field if non-nil, zero value otherwise.

### GetPlmnListOk

`func (o *NFProfile) GetPlmnListOk() (*[]PlmnId, bool)`

GetPlmnListOk returns a tuple with the PlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnList

`func (o *NFProfile) SetPlmnList(v []PlmnId)`

SetPlmnList sets PlmnList field to given value.

### HasPlmnList

`func (o *NFProfile) HasPlmnList() bool`

HasPlmnList returns a boolean if a field has been set.

### GetSNssais

`func (o *NFProfile) GetSNssais() []Snssai`

GetSNssais returns the SNssais field if non-nil, zero value otherwise.

### GetSNssaisOk

`func (o *NFProfile) GetSNssaisOk() (*[]Snssai, bool)`

GetSNssaisOk returns a tuple with the SNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssais

`func (o *NFProfile) SetSNssais(v []Snssai)`

SetSNssais sets SNssais field to given value.

### HasSNssais

`func (o *NFProfile) HasSNssais() bool`

HasSNssais returns a boolean if a field has been set.

### GetPerPlmnSnssaiList

`func (o *NFProfile) GetPerPlmnSnssaiList() []PlmnSnssai`

GetPerPlmnSnssaiList returns the PerPlmnSnssaiList field if non-nil, zero value otherwise.

### GetPerPlmnSnssaiListOk

`func (o *NFProfile) GetPerPlmnSnssaiListOk() (*[]PlmnSnssai, bool)`

GetPerPlmnSnssaiListOk returns a tuple with the PerPlmnSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPlmnSnssaiList

`func (o *NFProfile) SetPerPlmnSnssaiList(v []PlmnSnssai)`

SetPerPlmnSnssaiList sets PerPlmnSnssaiList field to given value.

### HasPerPlmnSnssaiList

`func (o *NFProfile) HasPerPlmnSnssaiList() bool`

HasPerPlmnSnssaiList returns a boolean if a field has been set.

### GetNsiList

`func (o *NFProfile) GetNsiList() []string`

GetNsiList returns the NsiList field if non-nil, zero value otherwise.

### GetNsiListOk

`func (o *NFProfile) GetNsiListOk() (*[]string, bool)`

GetNsiListOk returns a tuple with the NsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiList

`func (o *NFProfile) SetNsiList(v []string)`

SetNsiList sets NsiList field to given value.

### HasNsiList

`func (o *NFProfile) HasNsiList() bool`

HasNsiList returns a boolean if a field has been set.

### GetFqdn

`func (o *NFProfile) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NFProfile) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NFProfile) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NFProfile) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NFProfile) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NFProfile) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NFProfile) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NFProfile) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *NFProfile) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *NFProfile) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *NFProfile) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *NFProfile) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *NFProfile) GetIpv6Addresses() []Ipv6Addr`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *NFProfile) GetIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *NFProfile) SetIpv6Addresses(v []Ipv6Addr)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *NFProfile) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetAllowedPlmns

`func (o *NFProfile) GetAllowedPlmns() []PlmnId`

GetAllowedPlmns returns the AllowedPlmns field if non-nil, zero value otherwise.

### GetAllowedPlmnsOk

`func (o *NFProfile) GetAllowedPlmnsOk() (*[]PlmnId, bool)`

GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmns

`func (o *NFProfile) SetAllowedPlmns(v []PlmnId)`

SetAllowedPlmns sets AllowedPlmns field to given value.

### HasAllowedPlmns

`func (o *NFProfile) HasAllowedPlmns() bool`

HasAllowedPlmns returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *NFProfile) GetAllowedNfTypes() []NFType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *NFProfile) GetAllowedNfTypesOk() (*[]NFType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *NFProfile) SetAllowedNfTypes(v []NFType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *NFProfile) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNfDomains

`func (o *NFProfile) GetAllowedNfDomains() []string`

GetAllowedNfDomains returns the AllowedNfDomains field if non-nil, zero value otherwise.

### GetAllowedNfDomainsOk

`func (o *NFProfile) GetAllowedNfDomainsOk() (*[]string, bool)`

GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfDomains

`func (o *NFProfile) SetAllowedNfDomains(v []string)`

SetAllowedNfDomains sets AllowedNfDomains field to given value.

### HasAllowedNfDomains

`func (o *NFProfile) HasAllowedNfDomains() bool`

HasAllowedNfDomains returns a boolean if a field has been set.

### GetAllowedNssais

`func (o *NFProfile) GetAllowedNssais() []Snssai`

GetAllowedNssais returns the AllowedNssais field if non-nil, zero value otherwise.

### GetAllowedNssaisOk

`func (o *NFProfile) GetAllowedNssaisOk() (*[]Snssai, bool)`

GetAllowedNssaisOk returns a tuple with the AllowedNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssais

`func (o *NFProfile) SetAllowedNssais(v []Snssai)`

SetAllowedNssais sets AllowedNssais field to given value.

### HasAllowedNssais

`func (o *NFProfile) HasAllowedNssais() bool`

HasAllowedNssais returns a boolean if a field has been set.

### GetPriority

`func (o *NFProfile) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NFProfile) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NFProfile) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NFProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCapacity

`func (o *NFProfile) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NFProfile) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NFProfile) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *NFProfile) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetLoad

`func (o *NFProfile) GetLoad() int32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *NFProfile) GetLoadOk() (*int32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *NFProfile) SetLoad(v int32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *NFProfile) HasLoad() bool`

HasLoad returns a boolean if a field has been set.

### GetLocality

`func (o *NFProfile) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *NFProfile) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *NFProfile) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *NFProfile) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetUdrInfo

`func (o *NFProfile) GetUdrInfo() UdrInfo`

GetUdrInfo returns the UdrInfo field if non-nil, zero value otherwise.

### GetUdrInfoOk

`func (o *NFProfile) GetUdrInfoOk() (*UdrInfo, bool)`

GetUdrInfoOk returns a tuple with the UdrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrInfo

`func (o *NFProfile) SetUdrInfo(v UdrInfo)`

SetUdrInfo sets UdrInfo field to given value.

### HasUdrInfo

`func (o *NFProfile) HasUdrInfo() bool`

HasUdrInfo returns a boolean if a field has been set.

### GetUdrInfoExt

`func (o *NFProfile) GetUdrInfoExt() []UdrInfo`

GetUdrInfoExt returns the UdrInfoExt field if non-nil, zero value otherwise.

### GetUdrInfoExtOk

`func (o *NFProfile) GetUdrInfoExtOk() (*[]UdrInfo, bool)`

GetUdrInfoExtOk returns a tuple with the UdrInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrInfoExt

`func (o *NFProfile) SetUdrInfoExt(v []UdrInfo)`

SetUdrInfoExt sets UdrInfoExt field to given value.

### HasUdrInfoExt

`func (o *NFProfile) HasUdrInfoExt() bool`

HasUdrInfoExt returns a boolean if a field has been set.

### GetUdmInfo

`func (o *NFProfile) GetUdmInfo() UdmInfo`

GetUdmInfo returns the UdmInfo field if non-nil, zero value otherwise.

### GetUdmInfoOk

`func (o *NFProfile) GetUdmInfoOk() (*UdmInfo, bool)`

GetUdmInfoOk returns a tuple with the UdmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmInfo

`func (o *NFProfile) SetUdmInfo(v UdmInfo)`

SetUdmInfo sets UdmInfo field to given value.

### HasUdmInfo

`func (o *NFProfile) HasUdmInfo() bool`

HasUdmInfo returns a boolean if a field has been set.

### GetUdmInfoExt

`func (o *NFProfile) GetUdmInfoExt() []UdmInfo`

GetUdmInfoExt returns the UdmInfoExt field if non-nil, zero value otherwise.

### GetUdmInfoExtOk

`func (o *NFProfile) GetUdmInfoExtOk() (*[]UdmInfo, bool)`

GetUdmInfoExtOk returns a tuple with the UdmInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmInfoExt

`func (o *NFProfile) SetUdmInfoExt(v []UdmInfo)`

SetUdmInfoExt sets UdmInfoExt field to given value.

### HasUdmInfoExt

`func (o *NFProfile) HasUdmInfoExt() bool`

HasUdmInfoExt returns a boolean if a field has been set.

### GetAusfInfo

`func (o *NFProfile) GetAusfInfo() AusfInfo`

GetAusfInfo returns the AusfInfo field if non-nil, zero value otherwise.

### GetAusfInfoOk

`func (o *NFProfile) GetAusfInfoOk() (*AusfInfo, bool)`

GetAusfInfoOk returns a tuple with the AusfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInfo

`func (o *NFProfile) SetAusfInfo(v AusfInfo)`

SetAusfInfo sets AusfInfo field to given value.

### HasAusfInfo

`func (o *NFProfile) HasAusfInfo() bool`

HasAusfInfo returns a boolean if a field has been set.

### GetAusfInfoExt

`func (o *NFProfile) GetAusfInfoExt() []AusfInfo`

GetAusfInfoExt returns the AusfInfoExt field if non-nil, zero value otherwise.

### GetAusfInfoExtOk

`func (o *NFProfile) GetAusfInfoExtOk() (*[]AusfInfo, bool)`

GetAusfInfoExtOk returns a tuple with the AusfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInfoExt

`func (o *NFProfile) SetAusfInfoExt(v []AusfInfo)`

SetAusfInfoExt sets AusfInfoExt field to given value.

### HasAusfInfoExt

`func (o *NFProfile) HasAusfInfoExt() bool`

HasAusfInfoExt returns a boolean if a field has been set.

### GetAmfInfo

`func (o *NFProfile) GetAmfInfo() AmfInfo`

GetAmfInfo returns the AmfInfo field if non-nil, zero value otherwise.

### GetAmfInfoOk

`func (o *NFProfile) GetAmfInfoOk() (*AmfInfo, bool)`

GetAmfInfoOk returns a tuple with the AmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInfo

`func (o *NFProfile) SetAmfInfo(v AmfInfo)`

SetAmfInfo sets AmfInfo field to given value.

### HasAmfInfo

`func (o *NFProfile) HasAmfInfo() bool`

HasAmfInfo returns a boolean if a field has been set.

### GetAmfInfoExt

`func (o *NFProfile) GetAmfInfoExt() []AmfInfo`

GetAmfInfoExt returns the AmfInfoExt field if non-nil, zero value otherwise.

### GetAmfInfoExtOk

`func (o *NFProfile) GetAmfInfoExtOk() (*[]AmfInfo, bool)`

GetAmfInfoExtOk returns a tuple with the AmfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInfoExt

`func (o *NFProfile) SetAmfInfoExt(v []AmfInfo)`

SetAmfInfoExt sets AmfInfoExt field to given value.

### HasAmfInfoExt

`func (o *NFProfile) HasAmfInfoExt() bool`

HasAmfInfoExt returns a boolean if a field has been set.

### GetSmfInfo

`func (o *NFProfile) GetSmfInfo() SmfInfo`

GetSmfInfo returns the SmfInfo field if non-nil, zero value otherwise.

### GetSmfInfoOk

`func (o *NFProfile) GetSmfInfoOk() (*SmfInfo, bool)`

GetSmfInfoOk returns a tuple with the SmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInfo

`func (o *NFProfile) SetSmfInfo(v SmfInfo)`

SetSmfInfo sets SmfInfo field to given value.

### HasSmfInfo

`func (o *NFProfile) HasSmfInfo() bool`

HasSmfInfo returns a boolean if a field has been set.

### GetSmfInfoExt

`func (o *NFProfile) GetSmfInfoExt() []SmfInfo`

GetSmfInfoExt returns the SmfInfoExt field if non-nil, zero value otherwise.

### GetSmfInfoExtOk

`func (o *NFProfile) GetSmfInfoExtOk() (*[]SmfInfo, bool)`

GetSmfInfoExtOk returns a tuple with the SmfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInfoExt

`func (o *NFProfile) SetSmfInfoExt(v []SmfInfo)`

SetSmfInfoExt sets SmfInfoExt field to given value.

### HasSmfInfoExt

`func (o *NFProfile) HasSmfInfoExt() bool`

HasSmfInfoExt returns a boolean if a field has been set.

### GetUpfInfo

`func (o *NFProfile) GetUpfInfo() UpfInfo`

GetUpfInfo returns the UpfInfo field if non-nil, zero value otherwise.

### GetUpfInfoOk

`func (o *NFProfile) GetUpfInfoOk() (*UpfInfo, bool)`

GetUpfInfoOk returns a tuple with the UpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfo

`func (o *NFProfile) SetUpfInfo(v UpfInfo)`

SetUpfInfo sets UpfInfo field to given value.

### HasUpfInfo

`func (o *NFProfile) HasUpfInfo() bool`

HasUpfInfo returns a boolean if a field has been set.

### GetUpfInfoExt

`func (o *NFProfile) GetUpfInfoExt() []UpfInfo`

GetUpfInfoExt returns the UpfInfoExt field if non-nil, zero value otherwise.

### GetUpfInfoExtOk

`func (o *NFProfile) GetUpfInfoExtOk() (*[]UpfInfo, bool)`

GetUpfInfoExtOk returns a tuple with the UpfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfoExt

`func (o *NFProfile) SetUpfInfoExt(v []UpfInfo)`

SetUpfInfoExt sets UpfInfoExt field to given value.

### HasUpfInfoExt

`func (o *NFProfile) HasUpfInfoExt() bool`

HasUpfInfoExt returns a boolean if a field has been set.

### GetPcfInfo

`func (o *NFProfile) GetPcfInfo() PcfInfo`

GetPcfInfo returns the PcfInfo field if non-nil, zero value otherwise.

### GetPcfInfoOk

`func (o *NFProfile) GetPcfInfoOk() (*PcfInfo, bool)`

GetPcfInfoOk returns a tuple with the PcfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfo

`func (o *NFProfile) SetPcfInfo(v PcfInfo)`

SetPcfInfo sets PcfInfo field to given value.

### HasPcfInfo

`func (o *NFProfile) HasPcfInfo() bool`

HasPcfInfo returns a boolean if a field has been set.

### GetPcfInfoExt

`func (o *NFProfile) GetPcfInfoExt() []PcfInfo`

GetPcfInfoExt returns the PcfInfoExt field if non-nil, zero value otherwise.

### GetPcfInfoExtOk

`func (o *NFProfile) GetPcfInfoExtOk() (*[]PcfInfo, bool)`

GetPcfInfoExtOk returns a tuple with the PcfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfoExt

`func (o *NFProfile) SetPcfInfoExt(v []PcfInfo)`

SetPcfInfoExt sets PcfInfoExt field to given value.

### HasPcfInfoExt

`func (o *NFProfile) HasPcfInfoExt() bool`

HasPcfInfoExt returns a boolean if a field has been set.

### GetBsfInfo

`func (o *NFProfile) GetBsfInfo() BsfInfo`

GetBsfInfo returns the BsfInfo field if non-nil, zero value otherwise.

### GetBsfInfoOk

`func (o *NFProfile) GetBsfInfoOk() (*BsfInfo, bool)`

GetBsfInfoOk returns a tuple with the BsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfo

`func (o *NFProfile) SetBsfInfo(v BsfInfo)`

SetBsfInfo sets BsfInfo field to given value.

### HasBsfInfo

`func (o *NFProfile) HasBsfInfo() bool`

HasBsfInfo returns a boolean if a field has been set.

### GetBsfInfoExt

`func (o *NFProfile) GetBsfInfoExt() []BsfInfo`

GetBsfInfoExt returns the BsfInfoExt field if non-nil, zero value otherwise.

### GetBsfInfoExtOk

`func (o *NFProfile) GetBsfInfoExtOk() (*[]BsfInfo, bool)`

GetBsfInfoExtOk returns a tuple with the BsfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfoExt

`func (o *NFProfile) SetBsfInfoExt(v []BsfInfo)`

SetBsfInfoExt sets BsfInfoExt field to given value.

### HasBsfInfoExt

`func (o *NFProfile) HasBsfInfoExt() bool`

HasBsfInfoExt returns a boolean if a field has been set.

### GetChfInfo

`func (o *NFProfile) GetChfInfo() ChfInfo`

GetChfInfo returns the ChfInfo field if non-nil, zero value otherwise.

### GetChfInfoOk

`func (o *NFProfile) GetChfInfoOk() (*ChfInfo, bool)`

GetChfInfoOk returns a tuple with the ChfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfo

`func (o *NFProfile) SetChfInfo(v ChfInfo)`

SetChfInfo sets ChfInfo field to given value.

### HasChfInfo

`func (o *NFProfile) HasChfInfo() bool`

HasChfInfo returns a boolean if a field has been set.

### GetChfInfoExt

`func (o *NFProfile) GetChfInfoExt() []ChfInfo`

GetChfInfoExt returns the ChfInfoExt field if non-nil, zero value otherwise.

### GetChfInfoExtOk

`func (o *NFProfile) GetChfInfoExtOk() (*[]ChfInfo, bool)`

GetChfInfoExtOk returns a tuple with the ChfInfoExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfoExt

`func (o *NFProfile) SetChfInfoExt(v []ChfInfo)`

SetChfInfoExt sets ChfInfoExt field to given value.

### HasChfInfoExt

`func (o *NFProfile) HasChfInfoExt() bool`

HasChfInfoExt returns a boolean if a field has been set.

### GetNrfInfo

`func (o *NFProfile) GetNrfInfo() NrfInfo`

GetNrfInfo returns the NrfInfo field if non-nil, zero value otherwise.

### GetNrfInfoOk

`func (o *NFProfile) GetNrfInfoOk() (*NrfInfo, bool)`

GetNrfInfoOk returns a tuple with the NrfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfInfo

`func (o *NFProfile) SetNrfInfo(v NrfInfo)`

SetNrfInfo sets NrfInfo field to given value.

### HasNrfInfo

`func (o *NFProfile) HasNrfInfo() bool`

HasNrfInfo returns a boolean if a field has been set.

### GetNwdafInfo

`func (o *NFProfile) GetNwdafInfo() NwdafInfo`

GetNwdafInfo returns the NwdafInfo field if non-nil, zero value otherwise.

### GetNwdafInfoOk

`func (o *NFProfile) GetNwdafInfoOk() (*NwdafInfo, bool)`

GetNwdafInfoOk returns a tuple with the NwdafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInfo

`func (o *NFProfile) SetNwdafInfo(v NwdafInfo)`

SetNwdafInfo sets NwdafInfo field to given value.

### HasNwdafInfo

`func (o *NFProfile) HasNwdafInfo() bool`

HasNwdafInfo returns a boolean if a field has been set.

### GetCustomInfo

`func (o *NFProfile) GetCustomInfo() map[string]interface{}`

GetCustomInfo returns the CustomInfo field if non-nil, zero value otherwise.

### GetCustomInfoOk

`func (o *NFProfile) GetCustomInfoOk() (*map[string]interface{}, bool)`

GetCustomInfoOk returns a tuple with the CustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInfo

`func (o *NFProfile) SetCustomInfo(v map[string]interface{})`

SetCustomInfo sets CustomInfo field to given value.

### HasCustomInfo

`func (o *NFProfile) HasCustomInfo() bool`

HasCustomInfo returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *NFProfile) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *NFProfile) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *NFProfile) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *NFProfile) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetNfServicePersistence

`func (o *NFProfile) GetNfServicePersistence() bool`

GetNfServicePersistence returns the NfServicePersistence field if non-nil, zero value otherwise.

### GetNfServicePersistenceOk

`func (o *NFProfile) GetNfServicePersistenceOk() (*bool, bool)`

GetNfServicePersistenceOk returns a tuple with the NfServicePersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServicePersistence

`func (o *NFProfile) SetNfServicePersistence(v bool)`

SetNfServicePersistence sets NfServicePersistence field to given value.

### HasNfServicePersistence

`func (o *NFProfile) HasNfServicePersistence() bool`

HasNfServicePersistence returns a boolean if a field has been set.

### GetNfServices

`func (o *NFProfile) GetNfServices() []NFService`

GetNfServices returns the NfServices field if non-nil, zero value otherwise.

### GetNfServicesOk

`func (o *NFProfile) GetNfServicesOk() (*[]NFService, bool)`

GetNfServicesOk returns a tuple with the NfServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServices

`func (o *NFProfile) SetNfServices(v []NFService)`

SetNfServices sets NfServices field to given value.

### HasNfServices

`func (o *NFProfile) HasNfServices() bool`

HasNfServices returns a boolean if a field has been set.

### GetNfProfileChangesSupportInd

`func (o *NFProfile) GetNfProfileChangesSupportInd() bool`

GetNfProfileChangesSupportInd returns the NfProfileChangesSupportInd field if non-nil, zero value otherwise.

### GetNfProfileChangesSupportIndOk

`func (o *NFProfile) GetNfProfileChangesSupportIndOk() (*bool, bool)`

GetNfProfileChangesSupportIndOk returns a tuple with the NfProfileChangesSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfileChangesSupportInd

`func (o *NFProfile) SetNfProfileChangesSupportInd(v bool)`

SetNfProfileChangesSupportInd sets NfProfileChangesSupportInd field to given value.

### HasNfProfileChangesSupportInd

`func (o *NFProfile) HasNfProfileChangesSupportInd() bool`

HasNfProfileChangesSupportInd returns a boolean if a field has been set.

### GetNfProfileChangesInd

`func (o *NFProfile) GetNfProfileChangesInd() bool`

GetNfProfileChangesInd returns the NfProfileChangesInd field if non-nil, zero value otherwise.

### GetNfProfileChangesIndOk

`func (o *NFProfile) GetNfProfileChangesIndOk() (*bool, bool)`

GetNfProfileChangesIndOk returns a tuple with the NfProfileChangesInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfileChangesInd

`func (o *NFProfile) SetNfProfileChangesInd(v bool)`

SetNfProfileChangesInd sets NfProfileChangesInd field to given value.

### HasNfProfileChangesInd

`func (o *NFProfile) HasNfProfileChangesInd() bool`

HasNfProfileChangesInd returns a boolean if a field has been set.

### GetDefaultNotificationSubscriptions

`func (o *NFProfile) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription`

GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field if non-nil, zero value otherwise.

### GetDefaultNotificationSubscriptionsOk

`func (o *NFProfile) GetDefaultNotificationSubscriptionsOk() (*[]DefaultNotificationSubscription, bool)`

GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNotificationSubscriptions

`func (o *NFProfile) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription)`

SetDefaultNotificationSubscriptions sets DefaultNotificationSubscriptions field to given value.

### HasDefaultNotificationSubscriptions

`func (o *NFProfile) HasDefaultNotificationSubscriptions() bool`

HasDefaultNotificationSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



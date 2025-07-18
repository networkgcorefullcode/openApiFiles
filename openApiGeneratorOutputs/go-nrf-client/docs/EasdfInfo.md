# EasdfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiEasdfInfoList** | Pointer to [**[]SnssaiEasdfInfoItem**](SnssaiEasdfInfoItem.md) |  | [optional] 
**EasdfN6IpAddressList** | Pointer to [**[]IpAddr**](IpAddr.md) |  | [optional] 
**UpfN6IpAddressList** | Pointer to [**[]IpAddr**](IpAddr.md) |  | [optional] 
**N6TunnelInfoList** | Pointer to [**map[string]InterfaceUpfInfoItem**](InterfaceUpfInfoItem.md) | A map of InterfaceUpfInfoItems containing the N6 tunnelling information for establishing  a N6 tunnel between the V-UPF and the V-EASDF, where a valid JSON string serves as key.  | [optional] 
**DnsSecurityProtocols** | Pointer to [**[]DnsSecurityProtocol**](DnsSecurityProtocol.md) |  | [optional] 

## Methods

### NewEasdfInfo

`func NewEasdfInfo() *EasdfInfo`

NewEasdfInfo instantiates a new EasdfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasdfInfoWithDefaults

`func NewEasdfInfoWithDefaults() *EasdfInfo`

NewEasdfInfoWithDefaults instantiates a new EasdfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiEasdfInfoList

`func (o *EasdfInfo) GetSNssaiEasdfInfoList() []SnssaiEasdfInfoItem`

GetSNssaiEasdfInfoList returns the SNssaiEasdfInfoList field if non-nil, zero value otherwise.

### GetSNssaiEasdfInfoListOk

`func (o *EasdfInfo) GetSNssaiEasdfInfoListOk() (*[]SnssaiEasdfInfoItem, bool)`

GetSNssaiEasdfInfoListOk returns a tuple with the SNssaiEasdfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiEasdfInfoList

`func (o *EasdfInfo) SetSNssaiEasdfInfoList(v []SnssaiEasdfInfoItem)`

SetSNssaiEasdfInfoList sets SNssaiEasdfInfoList field to given value.

### HasSNssaiEasdfInfoList

`func (o *EasdfInfo) HasSNssaiEasdfInfoList() bool`

HasSNssaiEasdfInfoList returns a boolean if a field has been set.

### GetEasdfN6IpAddressList

`func (o *EasdfInfo) GetEasdfN6IpAddressList() []IpAddr`

GetEasdfN6IpAddressList returns the EasdfN6IpAddressList field if non-nil, zero value otherwise.

### GetEasdfN6IpAddressListOk

`func (o *EasdfInfo) GetEasdfN6IpAddressListOk() (*[]IpAddr, bool)`

GetEasdfN6IpAddressListOk returns a tuple with the EasdfN6IpAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasdfN6IpAddressList

`func (o *EasdfInfo) SetEasdfN6IpAddressList(v []IpAddr)`

SetEasdfN6IpAddressList sets EasdfN6IpAddressList field to given value.

### HasEasdfN6IpAddressList

`func (o *EasdfInfo) HasEasdfN6IpAddressList() bool`

HasEasdfN6IpAddressList returns a boolean if a field has been set.

### GetUpfN6IpAddressList

`func (o *EasdfInfo) GetUpfN6IpAddressList() []IpAddr`

GetUpfN6IpAddressList returns the UpfN6IpAddressList field if non-nil, zero value otherwise.

### GetUpfN6IpAddressListOk

`func (o *EasdfInfo) GetUpfN6IpAddressListOk() (*[]IpAddr, bool)`

GetUpfN6IpAddressListOk returns a tuple with the UpfN6IpAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfN6IpAddressList

`func (o *EasdfInfo) SetUpfN6IpAddressList(v []IpAddr)`

SetUpfN6IpAddressList sets UpfN6IpAddressList field to given value.

### HasUpfN6IpAddressList

`func (o *EasdfInfo) HasUpfN6IpAddressList() bool`

HasUpfN6IpAddressList returns a boolean if a field has been set.

### GetN6TunnelInfoList

`func (o *EasdfInfo) GetN6TunnelInfoList() map[string]InterfaceUpfInfoItem`

GetN6TunnelInfoList returns the N6TunnelInfoList field if non-nil, zero value otherwise.

### GetN6TunnelInfoListOk

`func (o *EasdfInfo) GetN6TunnelInfoListOk() (*map[string]InterfaceUpfInfoItem, bool)`

GetN6TunnelInfoListOk returns a tuple with the N6TunnelInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6TunnelInfoList

`func (o *EasdfInfo) SetN6TunnelInfoList(v map[string]InterfaceUpfInfoItem)`

SetN6TunnelInfoList sets N6TunnelInfoList field to given value.

### HasN6TunnelInfoList

`func (o *EasdfInfo) HasN6TunnelInfoList() bool`

HasN6TunnelInfoList returns a boolean if a field has been set.

### GetDnsSecurityProtocols

`func (o *EasdfInfo) GetDnsSecurityProtocols() []DnsSecurityProtocol`

GetDnsSecurityProtocols returns the DnsSecurityProtocols field if non-nil, zero value otherwise.

### GetDnsSecurityProtocolsOk

`func (o *EasdfInfo) GetDnsSecurityProtocolsOk() (*[]DnsSecurityProtocol, bool)`

GetDnsSecurityProtocolsOk returns a tuple with the DnsSecurityProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecurityProtocols

`func (o *EasdfInfo) SetDnsSecurityProtocols(v []DnsSecurityProtocol)`

SetDnsSecurityProtocols sets DnsSecurityProtocols field to given value.

### HasDnsSecurityProtocols

`func (o *EasdfInfo) HasDnsSecurityProtocols() bool`

HasDnsSecurityProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NrfInfoServedSeppInfoListValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeppPrefix** | Pointer to **string** |  | [optional] 
**SeppPorts** | Pointer to **map[string]int32** | Port numbers for HTTP and HTTPS. The key of the map shall be \&quot;http\&quot; or \&quot;https\&quot;.  | [optional] 
**RemotePlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**RemoteSnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**N32Purposes** | Pointer to [**[]N32Purpose**](N32Purpose.md) | N32 purposes supported by the SEPP | [optional] 

## Methods

### NewNrfInfoServedSeppInfoListValue

`func NewNrfInfoServedSeppInfoListValue() *NrfInfoServedSeppInfoListValue`

NewNrfInfoServedSeppInfoListValue instantiates a new NrfInfoServedSeppInfoListValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedSeppInfoListValueWithDefaults

`func NewNrfInfoServedSeppInfoListValueWithDefaults() *NrfInfoServedSeppInfoListValue`

NewNrfInfoServedSeppInfoListValueWithDefaults instantiates a new NrfInfoServedSeppInfoListValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeppPrefix

`func (o *NrfInfoServedSeppInfoListValue) GetSeppPrefix() string`

GetSeppPrefix returns the SeppPrefix field if non-nil, zero value otherwise.

### GetSeppPrefixOk

`func (o *NrfInfoServedSeppInfoListValue) GetSeppPrefixOk() (*string, bool)`

GetSeppPrefixOk returns a tuple with the SeppPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppPrefix

`func (o *NrfInfoServedSeppInfoListValue) SetSeppPrefix(v string)`

SetSeppPrefix sets SeppPrefix field to given value.

### HasSeppPrefix

`func (o *NrfInfoServedSeppInfoListValue) HasSeppPrefix() bool`

HasSeppPrefix returns a boolean if a field has been set.

### GetSeppPorts

`func (o *NrfInfoServedSeppInfoListValue) GetSeppPorts() map[string]int32`

GetSeppPorts returns the SeppPorts field if non-nil, zero value otherwise.

### GetSeppPortsOk

`func (o *NrfInfoServedSeppInfoListValue) GetSeppPortsOk() (*map[string]int32, bool)`

GetSeppPortsOk returns a tuple with the SeppPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppPorts

`func (o *NrfInfoServedSeppInfoListValue) SetSeppPorts(v map[string]int32)`

SetSeppPorts sets SeppPorts field to given value.

### HasSeppPorts

`func (o *NrfInfoServedSeppInfoListValue) HasSeppPorts() bool`

HasSeppPorts returns a boolean if a field has been set.

### GetRemotePlmnList

`func (o *NrfInfoServedSeppInfoListValue) GetRemotePlmnList() []PlmnId`

GetRemotePlmnList returns the RemotePlmnList field if non-nil, zero value otherwise.

### GetRemotePlmnListOk

`func (o *NrfInfoServedSeppInfoListValue) GetRemotePlmnListOk() (*[]PlmnId, bool)`

GetRemotePlmnListOk returns a tuple with the RemotePlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePlmnList

`func (o *NrfInfoServedSeppInfoListValue) SetRemotePlmnList(v []PlmnId)`

SetRemotePlmnList sets RemotePlmnList field to given value.

### HasRemotePlmnList

`func (o *NrfInfoServedSeppInfoListValue) HasRemotePlmnList() bool`

HasRemotePlmnList returns a boolean if a field has been set.

### GetRemoteSnpnList

`func (o *NrfInfoServedSeppInfoListValue) GetRemoteSnpnList() []PlmnIdNid`

GetRemoteSnpnList returns the RemoteSnpnList field if non-nil, zero value otherwise.

### GetRemoteSnpnListOk

`func (o *NrfInfoServedSeppInfoListValue) GetRemoteSnpnListOk() (*[]PlmnIdNid, bool)`

GetRemoteSnpnListOk returns a tuple with the RemoteSnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSnpnList

`func (o *NrfInfoServedSeppInfoListValue) SetRemoteSnpnList(v []PlmnIdNid)`

SetRemoteSnpnList sets RemoteSnpnList field to given value.

### HasRemoteSnpnList

`func (o *NrfInfoServedSeppInfoListValue) HasRemoteSnpnList() bool`

HasRemoteSnpnList returns a boolean if a field has been set.

### GetN32Purposes

`func (o *NrfInfoServedSeppInfoListValue) GetN32Purposes() []N32Purpose`

GetN32Purposes returns the N32Purposes field if non-nil, zero value otherwise.

### GetN32PurposesOk

`func (o *NrfInfoServedSeppInfoListValue) GetN32PurposesOk() (*[]N32Purpose, bool)`

GetN32PurposesOk returns a tuple with the N32Purposes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN32Purposes

`func (o *NrfInfoServedSeppInfoListValue) SetN32Purposes(v []N32Purpose)`

SetN32Purposes sets N32Purposes field to given value.

### HasN32Purposes

`func (o *NrfInfoServedSeppInfoListValue) HasN32Purposes() bool`

HasN32Purposes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



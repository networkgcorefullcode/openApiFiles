# NrfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServedUdrInfo** | Pointer to [**map[string]UdrInfo**](UdrInfo.md) |  | [optional] 
**ServedUdmInfo** | Pointer to [**map[string]UdmInfo**](UdmInfo.md) |  | [optional] 
**ServedAusfInfo** | Pointer to [**map[string]AusfInfo**](AusfInfo.md) |  | [optional] 
**ServedAmfInfo** | Pointer to [**map[string]AmfInfo**](AmfInfo.md) |  | [optional] 
**ServedSmfInfo** | Pointer to [**map[string]SmfInfo**](SmfInfo.md) |  | [optional] 
**ServedUpfInfo** | Pointer to [**map[string]UpfInfo**](UpfInfo.md) |  | [optional] 
**ServedPcfInfo** | Pointer to [**map[string]PcfInfo**](PcfInfo.md) |  | [optional] 
**ServedBsfInfo** | Pointer to [**map[string]BsfInfo**](BsfInfo.md) |  | [optional] 
**ServedChfInfo** | Pointer to [**map[string]ChfInfo**](ChfInfo.md) |  | [optional] 
**ServedNwdafInfo** | Pointer to [**map[string]NwdafInfo**](NwdafInfo.md) |  | [optional] 

## Methods

### NewNrfInfo

`func NewNrfInfo() *NrfInfo`

NewNrfInfo instantiates a new NrfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoWithDefaults

`func NewNrfInfoWithDefaults() *NrfInfo`

NewNrfInfoWithDefaults instantiates a new NrfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServedUdrInfo

`func (o *NrfInfo) GetServedUdrInfo() map[string]UdrInfo`

GetServedUdrInfo returns the ServedUdrInfo field if non-nil, zero value otherwise.

### GetServedUdrInfoOk

`func (o *NrfInfo) GetServedUdrInfoOk() (*map[string]UdrInfo, bool)`

GetServedUdrInfoOk returns a tuple with the ServedUdrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedUdrInfo

`func (o *NrfInfo) SetServedUdrInfo(v map[string]UdrInfo)`

SetServedUdrInfo sets ServedUdrInfo field to given value.

### HasServedUdrInfo

`func (o *NrfInfo) HasServedUdrInfo() bool`

HasServedUdrInfo returns a boolean if a field has been set.

### GetServedUdmInfo

`func (o *NrfInfo) GetServedUdmInfo() map[string]UdmInfo`

GetServedUdmInfo returns the ServedUdmInfo field if non-nil, zero value otherwise.

### GetServedUdmInfoOk

`func (o *NrfInfo) GetServedUdmInfoOk() (*map[string]UdmInfo, bool)`

GetServedUdmInfoOk returns a tuple with the ServedUdmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedUdmInfo

`func (o *NrfInfo) SetServedUdmInfo(v map[string]UdmInfo)`

SetServedUdmInfo sets ServedUdmInfo field to given value.

### HasServedUdmInfo

`func (o *NrfInfo) HasServedUdmInfo() bool`

HasServedUdmInfo returns a boolean if a field has been set.

### GetServedAusfInfo

`func (o *NrfInfo) GetServedAusfInfo() map[string]AusfInfo`

GetServedAusfInfo returns the ServedAusfInfo field if non-nil, zero value otherwise.

### GetServedAusfInfoOk

`func (o *NrfInfo) GetServedAusfInfoOk() (*map[string]AusfInfo, bool)`

GetServedAusfInfoOk returns a tuple with the ServedAusfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedAusfInfo

`func (o *NrfInfo) SetServedAusfInfo(v map[string]AusfInfo)`

SetServedAusfInfo sets ServedAusfInfo field to given value.

### HasServedAusfInfo

`func (o *NrfInfo) HasServedAusfInfo() bool`

HasServedAusfInfo returns a boolean if a field has been set.

### GetServedAmfInfo

`func (o *NrfInfo) GetServedAmfInfo() map[string]AmfInfo`

GetServedAmfInfo returns the ServedAmfInfo field if non-nil, zero value otherwise.

### GetServedAmfInfoOk

`func (o *NrfInfo) GetServedAmfInfoOk() (*map[string]AmfInfo, bool)`

GetServedAmfInfoOk returns a tuple with the ServedAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedAmfInfo

`func (o *NrfInfo) SetServedAmfInfo(v map[string]AmfInfo)`

SetServedAmfInfo sets ServedAmfInfo field to given value.

### HasServedAmfInfo

`func (o *NrfInfo) HasServedAmfInfo() bool`

HasServedAmfInfo returns a boolean if a field has been set.

### GetServedSmfInfo

`func (o *NrfInfo) GetServedSmfInfo() map[string]SmfInfo`

GetServedSmfInfo returns the ServedSmfInfo field if non-nil, zero value otherwise.

### GetServedSmfInfoOk

`func (o *NrfInfo) GetServedSmfInfoOk() (*map[string]SmfInfo, bool)`

GetServedSmfInfoOk returns a tuple with the ServedSmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedSmfInfo

`func (o *NrfInfo) SetServedSmfInfo(v map[string]SmfInfo)`

SetServedSmfInfo sets ServedSmfInfo field to given value.

### HasServedSmfInfo

`func (o *NrfInfo) HasServedSmfInfo() bool`

HasServedSmfInfo returns a boolean if a field has been set.

### GetServedUpfInfo

`func (o *NrfInfo) GetServedUpfInfo() map[string]UpfInfo`

GetServedUpfInfo returns the ServedUpfInfo field if non-nil, zero value otherwise.

### GetServedUpfInfoOk

`func (o *NrfInfo) GetServedUpfInfoOk() (*map[string]UpfInfo, bool)`

GetServedUpfInfoOk returns a tuple with the ServedUpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedUpfInfo

`func (o *NrfInfo) SetServedUpfInfo(v map[string]UpfInfo)`

SetServedUpfInfo sets ServedUpfInfo field to given value.

### HasServedUpfInfo

`func (o *NrfInfo) HasServedUpfInfo() bool`

HasServedUpfInfo returns a boolean if a field has been set.

### GetServedPcfInfo

`func (o *NrfInfo) GetServedPcfInfo() map[string]PcfInfo`

GetServedPcfInfo returns the ServedPcfInfo field if non-nil, zero value otherwise.

### GetServedPcfInfoOk

`func (o *NrfInfo) GetServedPcfInfoOk() (*map[string]PcfInfo, bool)`

GetServedPcfInfoOk returns a tuple with the ServedPcfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedPcfInfo

`func (o *NrfInfo) SetServedPcfInfo(v map[string]PcfInfo)`

SetServedPcfInfo sets ServedPcfInfo field to given value.

### HasServedPcfInfo

`func (o *NrfInfo) HasServedPcfInfo() bool`

HasServedPcfInfo returns a boolean if a field has been set.

### GetServedBsfInfo

`func (o *NrfInfo) GetServedBsfInfo() map[string]BsfInfo`

GetServedBsfInfo returns the ServedBsfInfo field if non-nil, zero value otherwise.

### GetServedBsfInfoOk

`func (o *NrfInfo) GetServedBsfInfoOk() (*map[string]BsfInfo, bool)`

GetServedBsfInfoOk returns a tuple with the ServedBsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedBsfInfo

`func (o *NrfInfo) SetServedBsfInfo(v map[string]BsfInfo)`

SetServedBsfInfo sets ServedBsfInfo field to given value.

### HasServedBsfInfo

`func (o *NrfInfo) HasServedBsfInfo() bool`

HasServedBsfInfo returns a boolean if a field has been set.

### GetServedChfInfo

`func (o *NrfInfo) GetServedChfInfo() map[string]ChfInfo`

GetServedChfInfo returns the ServedChfInfo field if non-nil, zero value otherwise.

### GetServedChfInfoOk

`func (o *NrfInfo) GetServedChfInfoOk() (*map[string]ChfInfo, bool)`

GetServedChfInfoOk returns a tuple with the ServedChfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedChfInfo

`func (o *NrfInfo) SetServedChfInfo(v map[string]ChfInfo)`

SetServedChfInfo sets ServedChfInfo field to given value.

### HasServedChfInfo

`func (o *NrfInfo) HasServedChfInfo() bool`

HasServedChfInfo returns a boolean if a field has been set.

### GetServedNwdafInfo

`func (o *NrfInfo) GetServedNwdafInfo() map[string]NwdafInfo`

GetServedNwdafInfo returns the ServedNwdafInfo field if non-nil, zero value otherwise.

### GetServedNwdafInfoOk

`func (o *NrfInfo) GetServedNwdafInfoOk() (*map[string]NwdafInfo, bool)`

GetServedNwdafInfoOk returns a tuple with the ServedNwdafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedNwdafInfo

`func (o *NrfInfo) SetServedNwdafInfo(v map[string]NwdafInfo)`

SetServedNwdafInfo sets ServedNwdafInfo field to given value.

### HasServedNwdafInfo

`func (o *NrfInfo) HasServedNwdafInfo() bool`

HasServedNwdafInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



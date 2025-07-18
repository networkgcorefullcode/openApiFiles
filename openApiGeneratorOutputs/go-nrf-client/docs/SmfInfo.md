# SmfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiSmfInfoList** | [**[]SnssaiSmfInfoItem**](SnssaiSmfInfoItem.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**PgwFqdn** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewSmfInfo

`func NewSmfInfo(sNssaiSmfInfoList []SnssaiSmfInfoItem, ) *SmfInfo`

NewSmfInfo instantiates a new SmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfInfoWithDefaults

`func NewSmfInfoWithDefaults() *SmfInfo`

NewSmfInfoWithDefaults instantiates a new SmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiSmfInfoList

`func (o *SmfInfo) GetSNssaiSmfInfoList() []SnssaiSmfInfoItem`

GetSNssaiSmfInfoList returns the SNssaiSmfInfoList field if non-nil, zero value otherwise.

### GetSNssaiSmfInfoListOk

`func (o *SmfInfo) GetSNssaiSmfInfoListOk() (*[]SnssaiSmfInfoItem, bool)`

GetSNssaiSmfInfoListOk returns a tuple with the SNssaiSmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiSmfInfoList

`func (o *SmfInfo) SetSNssaiSmfInfoList(v []SnssaiSmfInfoItem)`

SetSNssaiSmfInfoList sets SNssaiSmfInfoList field to given value.


### GetTaiList

`func (o *SmfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *SmfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *SmfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *SmfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *SmfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *SmfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *SmfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *SmfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetPgwFqdn

`func (o *SmfInfo) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *SmfInfo) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *SmfInfo) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *SmfInfo) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetAccessType

`func (o *SmfInfo) GetAccessType() []AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SmfInfo) GetAccessTypeOk() (*[]AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SmfInfo) SetAccessType(v []AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SmfInfo) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



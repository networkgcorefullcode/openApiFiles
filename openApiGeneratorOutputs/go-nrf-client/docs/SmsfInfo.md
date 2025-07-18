# SmsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoamingUeInd** | Pointer to **bool** |  | [optional] 
**RemotePlmnRangeList** | Pointer to [**[]PlmnRange**](PlmnRange.md) |  | [optional] 

## Methods

### NewSmsfInfo

`func NewSmsfInfo() *SmsfInfo`

NewSmsfInfo instantiates a new SmsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsfInfoWithDefaults

`func NewSmsfInfoWithDefaults() *SmsfInfo`

NewSmsfInfoWithDefaults instantiates a new SmsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoamingUeInd

`func (o *SmsfInfo) GetRoamingUeInd() bool`

GetRoamingUeInd returns the RoamingUeInd field if non-nil, zero value otherwise.

### GetRoamingUeIndOk

`func (o *SmsfInfo) GetRoamingUeIndOk() (*bool, bool)`

GetRoamingUeIndOk returns a tuple with the RoamingUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingUeInd

`func (o *SmsfInfo) SetRoamingUeInd(v bool)`

SetRoamingUeInd sets RoamingUeInd field to given value.

### HasRoamingUeInd

`func (o *SmsfInfo) HasRoamingUeInd() bool`

HasRoamingUeInd returns a boolean if a field has been set.

### GetRemotePlmnRangeList

`func (o *SmsfInfo) GetRemotePlmnRangeList() []PlmnRange`

GetRemotePlmnRangeList returns the RemotePlmnRangeList field if non-nil, zero value otherwise.

### GetRemotePlmnRangeListOk

`func (o *SmsfInfo) GetRemotePlmnRangeListOk() (*[]PlmnRange, bool)`

GetRemotePlmnRangeListOk returns a tuple with the RemotePlmnRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePlmnRangeList

`func (o *SmsfInfo) SetRemotePlmnRangeList(v []PlmnRange)`

SetRemotePlmnRangeList sets RemotePlmnRangeList field to given value.

### HasRemotePlmnRangeList

`func (o *SmsfInfo) HasRemotePlmnRangeList() bool`

HasRemotePlmnRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SharedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedDataId** | **string** |  | 
**SharedProfileData** | Pointer to [**NullableNFProfile**](NFProfile.md) |  | [optional] 
**SharedServiceData** | Pointer to [**NFService**](NFService.md) |  | [optional] 
**AuthorizedWriteScope** | Pointer to [**SharedScope**](SharedScope.md) |  | [optional] 

## Methods

### NewSharedData

`func NewSharedData(sharedDataId string, ) *SharedData`

NewSharedData instantiates a new SharedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDataWithDefaults

`func NewSharedDataWithDefaults() *SharedData`

NewSharedDataWithDefaults instantiates a new SharedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedDataId

`func (o *SharedData) GetSharedDataId() string`

GetSharedDataId returns the SharedDataId field if non-nil, zero value otherwise.

### GetSharedDataIdOk

`func (o *SharedData) GetSharedDataIdOk() (*string, bool)`

GetSharedDataIdOk returns a tuple with the SharedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataId

`func (o *SharedData) SetSharedDataId(v string)`

SetSharedDataId sets SharedDataId field to given value.


### GetSharedProfileData

`func (o *SharedData) GetSharedProfileData() NFProfile`

GetSharedProfileData returns the SharedProfileData field if non-nil, zero value otherwise.

### GetSharedProfileDataOk

`func (o *SharedData) GetSharedProfileDataOk() (*NFProfile, bool)`

GetSharedProfileDataOk returns a tuple with the SharedProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedProfileData

`func (o *SharedData) SetSharedProfileData(v NFProfile)`

SetSharedProfileData sets SharedProfileData field to given value.

### HasSharedProfileData

`func (o *SharedData) HasSharedProfileData() bool`

HasSharedProfileData returns a boolean if a field has been set.

### SetSharedProfileDataNil

`func (o *SharedData) SetSharedProfileDataNil(b bool)`

 SetSharedProfileDataNil sets the value for SharedProfileData to be an explicit nil

### UnsetSharedProfileData
`func (o *SharedData) UnsetSharedProfileData()`

UnsetSharedProfileData ensures that no value is present for SharedProfileData, not even an explicit nil
### GetSharedServiceData

`func (o *SharedData) GetSharedServiceData() NFService`

GetSharedServiceData returns the SharedServiceData field if non-nil, zero value otherwise.

### GetSharedServiceDataOk

`func (o *SharedData) GetSharedServiceDataOk() (*NFService, bool)`

GetSharedServiceDataOk returns a tuple with the SharedServiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedServiceData

`func (o *SharedData) SetSharedServiceData(v NFService)`

SetSharedServiceData sets SharedServiceData field to given value.

### HasSharedServiceData

`func (o *SharedData) HasSharedServiceData() bool`

HasSharedServiceData returns a boolean if a field has been set.

### GetAuthorizedWriteScope

`func (o *SharedData) GetAuthorizedWriteScope() SharedScope`

GetAuthorizedWriteScope returns the AuthorizedWriteScope field if non-nil, zero value otherwise.

### GetAuthorizedWriteScopeOk

`func (o *SharedData) GetAuthorizedWriteScopeOk() (*SharedScope, bool)`

GetAuthorizedWriteScopeOk returns a tuple with the AuthorizedWriteScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedWriteScope

`func (o *SharedData) SetAuthorizedWriteScope(v SharedScope)`

SetAuthorizedWriteScope sets AuthorizedWriteScope field to given value.

### HasAuthorizedWriteScope

`func (o *SharedData) HasAuthorizedWriteScope() bool`

HasAuthorizedWriteScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



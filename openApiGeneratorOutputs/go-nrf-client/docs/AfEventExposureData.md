# AfEventExposureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfEvents** | [**[]AfEvent**](AfEvent.md) |  | 
**AfIds** | Pointer to **[]string** |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 

## Methods

### NewAfEventExposureData

`func NewAfEventExposureData(afEvents []AfEvent, ) *AfEventExposureData`

NewAfEventExposureData instantiates a new AfEventExposureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventExposureDataWithDefaults

`func NewAfEventExposureDataWithDefaults() *AfEventExposureData`

NewAfEventExposureDataWithDefaults instantiates a new AfEventExposureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfEvents

`func (o *AfEventExposureData) GetAfEvents() []AfEvent`

GetAfEvents returns the AfEvents field if non-nil, zero value otherwise.

### GetAfEventsOk

`func (o *AfEventExposureData) GetAfEventsOk() (*[]AfEvent, bool)`

GetAfEventsOk returns a tuple with the AfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEvents

`func (o *AfEventExposureData) SetAfEvents(v []AfEvent)`

SetAfEvents sets AfEvents field to given value.


### GetAfIds

`func (o *AfEventExposureData) GetAfIds() []string`

GetAfIds returns the AfIds field if non-nil, zero value otherwise.

### GetAfIdsOk

`func (o *AfEventExposureData) GetAfIdsOk() (*[]string, bool)`

GetAfIdsOk returns a tuple with the AfIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfIds

`func (o *AfEventExposureData) SetAfIds(v []string)`

SetAfIds sets AfIds field to given value.

### HasAfIds

`func (o *AfEventExposureData) HasAfIds() bool`

HasAfIds returns a boolean if a field has been set.

### GetAppIds

`func (o *AfEventExposureData) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AfEventExposureData) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AfEventExposureData) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AfEventExposureData) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetTaiList

`func (o *AfEventExposureData) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *AfEventExposureData) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *AfEventExposureData) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *AfEventExposureData) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *AfEventExposureData) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *AfEventExposureData) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *AfEventExposureData) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *AfEventExposureData) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



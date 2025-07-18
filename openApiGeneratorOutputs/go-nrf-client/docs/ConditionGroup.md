# ConditionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]SelectionConditions**](SelectionConditions.md) |  | [optional] 
**Or** | Pointer to [**[]SelectionConditions**](SelectionConditions.md) |  | [optional] 

## Methods

### NewConditionGroup

`func NewConditionGroup() *ConditionGroup`

NewConditionGroup instantiates a new ConditionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionGroupWithDefaults

`func NewConditionGroupWithDefaults() *ConditionGroup`

NewConditionGroupWithDefaults instantiates a new ConditionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ConditionGroup) GetAnd() []SelectionConditions`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ConditionGroup) GetAndOk() (*[]SelectionConditions, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ConditionGroup) SetAnd(v []SelectionConditions)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ConditionGroup) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *ConditionGroup) GetOr() []SelectionConditions`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *ConditionGroup) GetOrOk() (*[]SelectionConditions, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *ConditionGroup) SetOr(v []SelectionConditions)`

SetOr sets Or field to given value.

### HasOr

`func (o *ConditionGroup) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



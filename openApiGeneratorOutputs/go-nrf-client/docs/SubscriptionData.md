# SubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfStatusNotificationUri** | **string** |  | 
**ReqNfInstanceId** | Pointer to **string** |  | [optional] 
**SubscrCond** | Pointer to [**SubscriptionDataSubscrCond**](SubscriptionDataSubscrCond.md) |  | [optional] 
**SubscriptionId** | **string** |  | [readonly] 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 
**ReqNotifEvents** | Pointer to [**[]NotificationEventType**](NotificationEventType.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**NotifCondition** | Pointer to [**NotifCondition**](NotifCondition.md) |  | [optional] 
**ReqNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**ReqNfFqdn** | Pointer to **string** |  | [optional] 
**ReqSnssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewSubscriptionData

`func NewSubscriptionData(nfStatusNotificationUri string, subscriptionId string, ) *SubscriptionData`

NewSubscriptionData instantiates a new SubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataWithDefaults

`func NewSubscriptionDataWithDefaults() *SubscriptionData`

NewSubscriptionDataWithDefaults instantiates a new SubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfStatusNotificationUri

`func (o *SubscriptionData) GetNfStatusNotificationUri() string`

GetNfStatusNotificationUri returns the NfStatusNotificationUri field if non-nil, zero value otherwise.

### GetNfStatusNotificationUriOk

`func (o *SubscriptionData) GetNfStatusNotificationUriOk() (*string, bool)`

GetNfStatusNotificationUriOk returns a tuple with the NfStatusNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStatusNotificationUri

`func (o *SubscriptionData) SetNfStatusNotificationUri(v string)`

SetNfStatusNotificationUri sets NfStatusNotificationUri field to given value.


### GetReqNfInstanceId

`func (o *SubscriptionData) GetReqNfInstanceId() string`

GetReqNfInstanceId returns the ReqNfInstanceId field if non-nil, zero value otherwise.

### GetReqNfInstanceIdOk

`func (o *SubscriptionData) GetReqNfInstanceIdOk() (*string, bool)`

GetReqNfInstanceIdOk returns a tuple with the ReqNfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNfInstanceId

`func (o *SubscriptionData) SetReqNfInstanceId(v string)`

SetReqNfInstanceId sets ReqNfInstanceId field to given value.

### HasReqNfInstanceId

`func (o *SubscriptionData) HasReqNfInstanceId() bool`

HasReqNfInstanceId returns a boolean if a field has been set.

### GetSubscrCond

`func (o *SubscriptionData) GetSubscrCond() SubscriptionDataSubscrCond`

GetSubscrCond returns the SubscrCond field if non-nil, zero value otherwise.

### GetSubscrCondOk

`func (o *SubscriptionData) GetSubscrCondOk() (*SubscriptionDataSubscrCond, bool)`

GetSubscrCondOk returns a tuple with the SubscrCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscrCond

`func (o *SubscriptionData) SetSubscrCond(v SubscriptionDataSubscrCond)`

SetSubscrCond sets SubscrCond field to given value.

### HasSubscrCond

`func (o *SubscriptionData) HasSubscrCond() bool`

HasSubscrCond returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SubscriptionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetValidityTime

`func (o *SubscriptionData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *SubscriptionData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *SubscriptionData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *SubscriptionData) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetReqNotifEvents

`func (o *SubscriptionData) GetReqNotifEvents() []NotificationEventType`

GetReqNotifEvents returns the ReqNotifEvents field if non-nil, zero value otherwise.

### GetReqNotifEventsOk

`func (o *SubscriptionData) GetReqNotifEventsOk() (*[]NotificationEventType, bool)`

GetReqNotifEventsOk returns a tuple with the ReqNotifEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNotifEvents

`func (o *SubscriptionData) SetReqNotifEvents(v []NotificationEventType)`

SetReqNotifEvents sets ReqNotifEvents field to given value.

### HasReqNotifEvents

`func (o *SubscriptionData) HasReqNotifEvents() bool`

HasReqNotifEvents returns a boolean if a field has been set.

### GetPlmnId

`func (o *SubscriptionData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SubscriptionData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SubscriptionData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *SubscriptionData) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetNotifCondition

`func (o *SubscriptionData) GetNotifCondition() NotifCondition`

GetNotifCondition returns the NotifCondition field if non-nil, zero value otherwise.

### GetNotifConditionOk

`func (o *SubscriptionData) GetNotifConditionOk() (*NotifCondition, bool)`

GetNotifConditionOk returns a tuple with the NotifCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCondition

`func (o *SubscriptionData) SetNotifCondition(v NotifCondition)`

SetNotifCondition sets NotifCondition field to given value.

### HasNotifCondition

`func (o *SubscriptionData) HasNotifCondition() bool`

HasNotifCondition returns a boolean if a field has been set.

### GetReqNfType

`func (o *SubscriptionData) GetReqNfType() NFType`

GetReqNfType returns the ReqNfType field if non-nil, zero value otherwise.

### GetReqNfTypeOk

`func (o *SubscriptionData) GetReqNfTypeOk() (*NFType, bool)`

GetReqNfTypeOk returns a tuple with the ReqNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNfType

`func (o *SubscriptionData) SetReqNfType(v NFType)`

SetReqNfType sets ReqNfType field to given value.

### HasReqNfType

`func (o *SubscriptionData) HasReqNfType() bool`

HasReqNfType returns a boolean if a field has been set.

### GetReqNfFqdn

`func (o *SubscriptionData) GetReqNfFqdn() string`

GetReqNfFqdn returns the ReqNfFqdn field if non-nil, zero value otherwise.

### GetReqNfFqdnOk

`func (o *SubscriptionData) GetReqNfFqdnOk() (*string, bool)`

GetReqNfFqdnOk returns a tuple with the ReqNfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNfFqdn

`func (o *SubscriptionData) SetReqNfFqdn(v string)`

SetReqNfFqdn sets ReqNfFqdn field to given value.

### HasReqNfFqdn

`func (o *SubscriptionData) HasReqNfFqdn() bool`

HasReqNfFqdn returns a boolean if a field has been set.

### GetReqSnssais

`func (o *SubscriptionData) GetReqSnssais() []Snssai`

GetReqSnssais returns the ReqSnssais field if non-nil, zero value otherwise.

### GetReqSnssaisOk

`func (o *SubscriptionData) GetReqSnssaisOk() (*[]Snssai, bool)`

GetReqSnssaisOk returns a tuple with the ReqSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSnssais

`func (o *SubscriptionData) SetReqSnssais(v []Snssai)`

SetReqSnssais sets ReqSnssais field to given value.

### HasReqSnssais

`func (o *SubscriptionData) HasReqSnssais() bool`

HasReqSnssais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



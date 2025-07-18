# DefaultNotificationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**CallbackUri** | **string** |  | 
**N1MessageClass** | Pointer to [**N1MessageClass**](N1MessageClass.md) |  | [optional] 
**N2InformationClass** | Pointer to [**N2InformationClass**](N2InformationClass.md) |  | [optional] 

## Methods

### NewDefaultNotificationSubscription

`func NewDefaultNotificationSubscription(notificationType NotificationType, callbackUri string, ) *DefaultNotificationSubscription`

NewDefaultNotificationSubscription instantiates a new DefaultNotificationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultNotificationSubscriptionWithDefaults

`func NewDefaultNotificationSubscriptionWithDefaults() *DefaultNotificationSubscription`

NewDefaultNotificationSubscriptionWithDefaults instantiates a new DefaultNotificationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *DefaultNotificationSubscription) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *DefaultNotificationSubscription) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *DefaultNotificationSubscription) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetCallbackUri

`func (o *DefaultNotificationSubscription) GetCallbackUri() string`

GetCallbackUri returns the CallbackUri field if non-nil, zero value otherwise.

### GetCallbackUriOk

`func (o *DefaultNotificationSubscription) GetCallbackUriOk() (*string, bool)`

GetCallbackUriOk returns a tuple with the CallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUri

`func (o *DefaultNotificationSubscription) SetCallbackUri(v string)`

SetCallbackUri sets CallbackUri field to given value.


### GetN1MessageClass

`func (o *DefaultNotificationSubscription) GetN1MessageClass() N1MessageClass`

GetN1MessageClass returns the N1MessageClass field if non-nil, zero value otherwise.

### GetN1MessageClassOk

`func (o *DefaultNotificationSubscription) GetN1MessageClassOk() (*N1MessageClass, bool)`

GetN1MessageClassOk returns a tuple with the N1MessageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageClass

`func (o *DefaultNotificationSubscription) SetN1MessageClass(v N1MessageClass)`

SetN1MessageClass sets N1MessageClass field to given value.

### HasN1MessageClass

`func (o *DefaultNotificationSubscription) HasN1MessageClass() bool`

HasN1MessageClass returns a boolean if a field has been set.

### GetN2InformationClass

`func (o *DefaultNotificationSubscription) GetN2InformationClass() N2InformationClass`

GetN2InformationClass returns the N2InformationClass field if non-nil, zero value otherwise.

### GetN2InformationClassOk

`func (o *DefaultNotificationSubscription) GetN2InformationClassOk() (*N2InformationClass, bool)`

GetN2InformationClassOk returns a tuple with the N2InformationClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InformationClass

`func (o *DefaultNotificationSubscription) SetN2InformationClass(v N2InformationClass)`

SetN2InformationClass sets N2InformationClass field to given value.

### HasN2InformationClass

`func (o *DefaultNotificationSubscription) HasN2InformationClass() bool`

HasN2InformationClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



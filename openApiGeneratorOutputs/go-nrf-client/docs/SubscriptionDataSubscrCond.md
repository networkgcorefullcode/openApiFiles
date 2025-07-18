# SubscriptionDataSubscrCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** |  | 
**NfType** | **string** |  | 
**ServiceName** | [**ServiceName**](ServiceName.md) |  | 
**AmfSetId** | Pointer to **string** |  | [optional] 
**AmfRegionId** | Pointer to **string** |  | [optional] 
**GuamiList** | [**[]Guami**](Guami.md) |  | 
**SnssaiList** | [**[]Snssai**](Snssai.md) |  | 
**NsiList** | Pointer to **[]string** |  | [optional] 
**NfGroupId** | **string** |  | 

## Methods

### NewSubscriptionDataSubscrCond

`func NewSubscriptionDataSubscrCond(nfInstanceId string, nfType string, serviceName ServiceName, guamiList []Guami, snssaiList []Snssai, nfGroupId string, ) *SubscriptionDataSubscrCond`

NewSubscriptionDataSubscrCond instantiates a new SubscriptionDataSubscrCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataSubscrCondWithDefaults

`func NewSubscriptionDataSubscrCondWithDefaults() *SubscriptionDataSubscrCond`

NewSubscriptionDataSubscrCondWithDefaults instantiates a new SubscriptionDataSubscrCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *SubscriptionDataSubscrCond) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *SubscriptionDataSubscrCond) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *SubscriptionDataSubscrCond) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfType

`func (o *SubscriptionDataSubscrCond) GetNfType() string`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *SubscriptionDataSubscrCond) GetNfTypeOk() (*string, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *SubscriptionDataSubscrCond) SetNfType(v string)`

SetNfType sets NfType field to given value.


### GetServiceName

`func (o *SubscriptionDataSubscrCond) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SubscriptionDataSubscrCond) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SubscriptionDataSubscrCond) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.


### GetAmfSetId

`func (o *SubscriptionDataSubscrCond) GetAmfSetId() string`

GetAmfSetId returns the AmfSetId field if non-nil, zero value otherwise.

### GetAmfSetIdOk

`func (o *SubscriptionDataSubscrCond) GetAmfSetIdOk() (*string, bool)`

GetAmfSetIdOk returns a tuple with the AmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSetId

`func (o *SubscriptionDataSubscrCond) SetAmfSetId(v string)`

SetAmfSetId sets AmfSetId field to given value.

### HasAmfSetId

`func (o *SubscriptionDataSubscrCond) HasAmfSetId() bool`

HasAmfSetId returns a boolean if a field has been set.

### GetAmfRegionId

`func (o *SubscriptionDataSubscrCond) GetAmfRegionId() string`

GetAmfRegionId returns the AmfRegionId field if non-nil, zero value otherwise.

### GetAmfRegionIdOk

`func (o *SubscriptionDataSubscrCond) GetAmfRegionIdOk() (*string, bool)`

GetAmfRegionIdOk returns a tuple with the AmfRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfRegionId

`func (o *SubscriptionDataSubscrCond) SetAmfRegionId(v string)`

SetAmfRegionId sets AmfRegionId field to given value.

### HasAmfRegionId

`func (o *SubscriptionDataSubscrCond) HasAmfRegionId() bool`

HasAmfRegionId returns a boolean if a field has been set.

### GetGuamiList

`func (o *SubscriptionDataSubscrCond) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *SubscriptionDataSubscrCond) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *SubscriptionDataSubscrCond) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.


### GetSnssaiList

`func (o *SubscriptionDataSubscrCond) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *SubscriptionDataSubscrCond) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *SubscriptionDataSubscrCond) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.


### GetNsiList

`func (o *SubscriptionDataSubscrCond) GetNsiList() []string`

GetNsiList returns the NsiList field if non-nil, zero value otherwise.

### GetNsiListOk

`func (o *SubscriptionDataSubscrCond) GetNsiListOk() (*[]string, bool)`

GetNsiListOk returns a tuple with the NsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiList

`func (o *SubscriptionDataSubscrCond) SetNsiList(v []string)`

SetNsiList sets NsiList field to given value.

### HasNsiList

`func (o *SubscriptionDataSubscrCond) HasNsiList() bool`

HasNsiList returns a boolean if a field has been set.

### GetNfGroupId

`func (o *SubscriptionDataSubscrCond) GetNfGroupId() string`

GetNfGroupId returns the NfGroupId field if non-nil, zero value otherwise.

### GetNfGroupIdOk

`func (o *SubscriptionDataSubscrCond) GetNfGroupIdOk() (*string, bool)`

GetNfGroupIdOk returns a tuple with the NfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfGroupId

`func (o *SubscriptionDataSubscrCond) SetNfGroupId(v string)`

SetNfGroupId sets NfGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NFService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstanceId** | **string** |  | 
**ServiceName** | [**ServiceName**](ServiceName.md) |  | 
**Versions** | [**[]NFServiceVersion**](NFServiceVersion.md) |  | 
**Scheme** | [**UriScheme**](UriScheme.md) |  | 
**NfServiceStatus** | [**NFServiceStatus**](NFServiceStatus.md) |  | 
**Fqdn** | Pointer to **string** |  | [optional] 
**InterPlmnFqdn** | Pointer to **string** |  | [optional] 
**IpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) |  | [optional] 
**ApiPrefix** | Pointer to **string** |  | [optional] 
**DefaultNotificationSubscriptions** | Pointer to [**[]DefaultNotificationSubscription**](DefaultNotificationSubscription.md) |  | [optional] 
**AllowedPlmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AllowedNfDomains** | Pointer to **[]string** |  | [optional] 
**AllowedNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Load** | Pointer to **int32** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 
**ChfServiceInfo** | Pointer to [**ChfServiceInfo**](ChfServiceInfo.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewNFService

`func NewNFService(serviceInstanceId string, serviceName ServiceName, versions []NFServiceVersion, scheme UriScheme, nfServiceStatus NFServiceStatus, ) *NFService`

NewNFService instantiates a new NFService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFServiceWithDefaults

`func NewNFServiceWithDefaults() *NFService`

NewNFServiceWithDefaults instantiates a new NFService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstanceId

`func (o *NFService) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *NFService) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *NFService) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.


### GetServiceName

`func (o *NFService) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NFService) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NFService) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.


### GetVersions

`func (o *NFService) GetVersions() []NFServiceVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *NFService) GetVersionsOk() (*[]NFServiceVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *NFService) SetVersions(v []NFServiceVersion)`

SetVersions sets Versions field to given value.


### GetScheme

`func (o *NFService) GetScheme() UriScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *NFService) GetSchemeOk() (*UriScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *NFService) SetScheme(v UriScheme)`

SetScheme sets Scheme field to given value.


### GetNfServiceStatus

`func (o *NFService) GetNfServiceStatus() NFServiceStatus`

GetNfServiceStatus returns the NfServiceStatus field if non-nil, zero value otherwise.

### GetNfServiceStatusOk

`func (o *NFService) GetNfServiceStatusOk() (*NFServiceStatus, bool)`

GetNfServiceStatusOk returns a tuple with the NfServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceStatus

`func (o *NFService) SetNfServiceStatus(v NFServiceStatus)`

SetNfServiceStatus sets NfServiceStatus field to given value.


### GetFqdn

`func (o *NFService) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NFService) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NFService) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NFService) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NFService) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NFService) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NFService) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NFService) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetIpEndPoints

`func (o *NFService) GetIpEndPoints() []IpEndPoint`

GetIpEndPoints returns the IpEndPoints field if non-nil, zero value otherwise.

### GetIpEndPointsOk

`func (o *NFService) GetIpEndPointsOk() (*[]IpEndPoint, bool)`

GetIpEndPointsOk returns a tuple with the IpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEndPoints

`func (o *NFService) SetIpEndPoints(v []IpEndPoint)`

SetIpEndPoints sets IpEndPoints field to given value.

### HasIpEndPoints

`func (o *NFService) HasIpEndPoints() bool`

HasIpEndPoints returns a boolean if a field has been set.

### GetApiPrefix

`func (o *NFService) GetApiPrefix() string`

GetApiPrefix returns the ApiPrefix field if non-nil, zero value otherwise.

### GetApiPrefixOk

`func (o *NFService) GetApiPrefixOk() (*string, bool)`

GetApiPrefixOk returns a tuple with the ApiPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPrefix

`func (o *NFService) SetApiPrefix(v string)`

SetApiPrefix sets ApiPrefix field to given value.

### HasApiPrefix

`func (o *NFService) HasApiPrefix() bool`

HasApiPrefix returns a boolean if a field has been set.

### GetDefaultNotificationSubscriptions

`func (o *NFService) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription`

GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field if non-nil, zero value otherwise.

### GetDefaultNotificationSubscriptionsOk

`func (o *NFService) GetDefaultNotificationSubscriptionsOk() (*[]DefaultNotificationSubscription, bool)`

GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNotificationSubscriptions

`func (o *NFService) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription)`

SetDefaultNotificationSubscriptions sets DefaultNotificationSubscriptions field to given value.

### HasDefaultNotificationSubscriptions

`func (o *NFService) HasDefaultNotificationSubscriptions() bool`

HasDefaultNotificationSubscriptions returns a boolean if a field has been set.

### GetAllowedPlmns

`func (o *NFService) GetAllowedPlmns() []PlmnId`

GetAllowedPlmns returns the AllowedPlmns field if non-nil, zero value otherwise.

### GetAllowedPlmnsOk

`func (o *NFService) GetAllowedPlmnsOk() (*[]PlmnId, bool)`

GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmns

`func (o *NFService) SetAllowedPlmns(v []PlmnId)`

SetAllowedPlmns sets AllowedPlmns field to given value.

### HasAllowedPlmns

`func (o *NFService) HasAllowedPlmns() bool`

HasAllowedPlmns returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *NFService) GetAllowedNfTypes() []NFType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *NFService) GetAllowedNfTypesOk() (*[]NFType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *NFService) SetAllowedNfTypes(v []NFType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *NFService) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNfDomains

`func (o *NFService) GetAllowedNfDomains() []string`

GetAllowedNfDomains returns the AllowedNfDomains field if non-nil, zero value otherwise.

### GetAllowedNfDomainsOk

`func (o *NFService) GetAllowedNfDomainsOk() (*[]string, bool)`

GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfDomains

`func (o *NFService) SetAllowedNfDomains(v []string)`

SetAllowedNfDomains sets AllowedNfDomains field to given value.

### HasAllowedNfDomains

`func (o *NFService) HasAllowedNfDomains() bool`

HasAllowedNfDomains returns a boolean if a field has been set.

### GetAllowedNssais

`func (o *NFService) GetAllowedNssais() []Snssai`

GetAllowedNssais returns the AllowedNssais field if non-nil, zero value otherwise.

### GetAllowedNssaisOk

`func (o *NFService) GetAllowedNssaisOk() (*[]Snssai, bool)`

GetAllowedNssaisOk returns a tuple with the AllowedNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssais

`func (o *NFService) SetAllowedNssais(v []Snssai)`

SetAllowedNssais sets AllowedNssais field to given value.

### HasAllowedNssais

`func (o *NFService) HasAllowedNssais() bool`

HasAllowedNssais returns a boolean if a field has been set.

### GetPriority

`func (o *NFService) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NFService) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NFService) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NFService) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCapacity

`func (o *NFService) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NFService) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NFService) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *NFService) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetLoad

`func (o *NFService) GetLoad() int32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *NFService) GetLoadOk() (*int32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *NFService) SetLoad(v int32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *NFService) HasLoad() bool`

HasLoad returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *NFService) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *NFService) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *NFService) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *NFService) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetChfServiceInfo

`func (o *NFService) GetChfServiceInfo() ChfServiceInfo`

GetChfServiceInfo returns the ChfServiceInfo field if non-nil, zero value otherwise.

### GetChfServiceInfoOk

`func (o *NFService) GetChfServiceInfoOk() (*ChfServiceInfo, bool)`

GetChfServiceInfoOk returns a tuple with the ChfServiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfServiceInfo

`func (o *NFService) SetChfServiceInfo(v ChfServiceInfo)`

SetChfServiceInfo sets ChfServiceInfo field to given value.

### HasChfServiceInfo

`func (o *NFService) HasChfServiceInfo() bool`

HasChfServiceInfo returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NFService) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NFService) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NFService) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NFService) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



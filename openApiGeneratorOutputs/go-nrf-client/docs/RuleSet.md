# RuleSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **int32** |  | 
**Plmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**Snpns** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**NfTypes** | Pointer to [**[]NfType**](NfType.md) |  | [optional] 
**NfDomains** | Pointer to **[]string** |  | [optional] 
**Nssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**NfInstances** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Action** | [**RuleSetAction**](RuleSetAction.md) |  | 

## Methods

### NewRuleSet

`func NewRuleSet(priority int32, action RuleSetAction, ) *RuleSet`

NewRuleSet instantiates a new RuleSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetWithDefaults

`func NewRuleSetWithDefaults() *RuleSet`

NewRuleSetWithDefaults instantiates a new RuleSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *RuleSet) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RuleSet) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RuleSet) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPlmns

`func (o *RuleSet) GetPlmns() []PlmnId`

GetPlmns returns the Plmns field if non-nil, zero value otherwise.

### GetPlmnsOk

`func (o *RuleSet) GetPlmnsOk() (*[]PlmnId, bool)`

GetPlmnsOk returns a tuple with the Plmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmns

`func (o *RuleSet) SetPlmns(v []PlmnId)`

SetPlmns sets Plmns field to given value.

### HasPlmns

`func (o *RuleSet) HasPlmns() bool`

HasPlmns returns a boolean if a field has been set.

### GetSnpns

`func (o *RuleSet) GetSnpns() []PlmnIdNid`

GetSnpns returns the Snpns field if non-nil, zero value otherwise.

### GetSnpnsOk

`func (o *RuleSet) GetSnpnsOk() (*[]PlmnIdNid, bool)`

GetSnpnsOk returns a tuple with the Snpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpns

`func (o *RuleSet) SetSnpns(v []PlmnIdNid)`

SetSnpns sets Snpns field to given value.

### HasSnpns

`func (o *RuleSet) HasSnpns() bool`

HasSnpns returns a boolean if a field has been set.

### GetNfTypes

`func (o *RuleSet) GetNfTypes() []NfType`

GetNfTypes returns the NfTypes field if non-nil, zero value otherwise.

### GetNfTypesOk

`func (o *RuleSet) GetNfTypesOk() (*[]NfType, bool)`

GetNfTypesOk returns a tuple with the NfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypes

`func (o *RuleSet) SetNfTypes(v []NfType)`

SetNfTypes sets NfTypes field to given value.

### HasNfTypes

`func (o *RuleSet) HasNfTypes() bool`

HasNfTypes returns a boolean if a field has been set.

### GetNfDomains

`func (o *RuleSet) GetNfDomains() []string`

GetNfDomains returns the NfDomains field if non-nil, zero value otherwise.

### GetNfDomainsOk

`func (o *RuleSet) GetNfDomainsOk() (*[]string, bool)`

GetNfDomainsOk returns a tuple with the NfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfDomains

`func (o *RuleSet) SetNfDomains(v []string)`

SetNfDomains sets NfDomains field to given value.

### HasNfDomains

`func (o *RuleSet) HasNfDomains() bool`

HasNfDomains returns a boolean if a field has been set.

### GetNssais

`func (o *RuleSet) GetNssais() []ExtSnssai`

GetNssais returns the Nssais field if non-nil, zero value otherwise.

### GetNssaisOk

`func (o *RuleSet) GetNssaisOk() (*[]ExtSnssai, bool)`

GetNssaisOk returns a tuple with the Nssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssais

`func (o *RuleSet) SetNssais(v []ExtSnssai)`

SetNssais sets Nssais field to given value.

### HasNssais

`func (o *RuleSet) HasNssais() bool`

HasNssais returns a boolean if a field has been set.

### GetNfInstances

`func (o *RuleSet) GetNfInstances() []string`

GetNfInstances returns the NfInstances field if non-nil, zero value otherwise.

### GetNfInstancesOk

`func (o *RuleSet) GetNfInstancesOk() (*[]string, bool)`

GetNfInstancesOk returns a tuple with the NfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstances

`func (o *RuleSet) SetNfInstances(v []string)`

SetNfInstances sets NfInstances field to given value.

### HasNfInstances

`func (o *RuleSet) HasNfInstances() bool`

HasNfInstances returns a boolean if a field has been set.

### GetScopes

`func (o *RuleSet) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RuleSet) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RuleSet) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *RuleSet) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAction

`func (o *RuleSet) GetAction() RuleSetAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleSet) GetActionOk() (*RuleSetAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleSet) SetAction(v RuleSetAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



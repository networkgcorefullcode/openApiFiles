# NFProfileRegistrationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Instance** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Cause** | Pointer to **string** | A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available.  | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam**](InvalidParam.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AccessTokenError** | Pointer to [**AccessTokenErr**](AccessTokenErr.md) |  | [optional] 
**AccessTokenRequest** | Pointer to [**AccessTokenReq**](AccessTokenReq.md) |  | [optional] 
**NrfId** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**SupportedApiVersions** | Pointer to **[]string** |  | [optional] 
**NoProfileMatchInfo** | Pointer to [**NoProfileMatchInfo**](NoProfileMatchInfo.md) |  | [optional] 
**SharedDataIds** | **[]string** |  | 

## Methods

### NewNFProfileRegistrationError

`func NewNFProfileRegistrationError(sharedDataIds []string, ) *NFProfileRegistrationError`

NewNFProfileRegistrationError instantiates a new NFProfileRegistrationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFProfileRegistrationErrorWithDefaults

`func NewNFProfileRegistrationErrorWithDefaults() *NFProfileRegistrationError`

NewNFProfileRegistrationErrorWithDefaults instantiates a new NFProfileRegistrationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NFProfileRegistrationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NFProfileRegistrationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NFProfileRegistrationError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NFProfileRegistrationError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *NFProfileRegistrationError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NFProfileRegistrationError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NFProfileRegistrationError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NFProfileRegistrationError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *NFProfileRegistrationError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NFProfileRegistrationError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NFProfileRegistrationError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NFProfileRegistrationError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *NFProfileRegistrationError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *NFProfileRegistrationError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *NFProfileRegistrationError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *NFProfileRegistrationError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *NFProfileRegistrationError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NFProfileRegistrationError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NFProfileRegistrationError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *NFProfileRegistrationError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *NFProfileRegistrationError) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *NFProfileRegistrationError) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *NFProfileRegistrationError) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *NFProfileRegistrationError) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *NFProfileRegistrationError) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *NFProfileRegistrationError) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *NFProfileRegistrationError) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *NFProfileRegistrationError) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NFProfileRegistrationError) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NFProfileRegistrationError) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NFProfileRegistrationError) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NFProfileRegistrationError) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *NFProfileRegistrationError) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *NFProfileRegistrationError) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *NFProfileRegistrationError) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *NFProfileRegistrationError) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *NFProfileRegistrationError) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *NFProfileRegistrationError) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *NFProfileRegistrationError) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *NFProfileRegistrationError) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *NFProfileRegistrationError) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *NFProfileRegistrationError) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *NFProfileRegistrationError) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *NFProfileRegistrationError) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.

### GetSupportedApiVersions

`func (o *NFProfileRegistrationError) GetSupportedApiVersions() []string`

GetSupportedApiVersions returns the SupportedApiVersions field if non-nil, zero value otherwise.

### GetSupportedApiVersionsOk

`func (o *NFProfileRegistrationError) GetSupportedApiVersionsOk() (*[]string, bool)`

GetSupportedApiVersionsOk returns a tuple with the SupportedApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedApiVersions

`func (o *NFProfileRegistrationError) SetSupportedApiVersions(v []string)`

SetSupportedApiVersions sets SupportedApiVersions field to given value.

### HasSupportedApiVersions

`func (o *NFProfileRegistrationError) HasSupportedApiVersions() bool`

HasSupportedApiVersions returns a boolean if a field has been set.

### GetNoProfileMatchInfo

`func (o *NFProfileRegistrationError) GetNoProfileMatchInfo() NoProfileMatchInfo`

GetNoProfileMatchInfo returns the NoProfileMatchInfo field if non-nil, zero value otherwise.

### GetNoProfileMatchInfoOk

`func (o *NFProfileRegistrationError) GetNoProfileMatchInfoOk() (*NoProfileMatchInfo, bool)`

GetNoProfileMatchInfoOk returns a tuple with the NoProfileMatchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProfileMatchInfo

`func (o *NFProfileRegistrationError) SetNoProfileMatchInfo(v NoProfileMatchInfo)`

SetNoProfileMatchInfo sets NoProfileMatchInfo field to given value.

### HasNoProfileMatchInfo

`func (o *NFProfileRegistrationError) HasNoProfileMatchInfo() bool`

HasNoProfileMatchInfo returns a boolean if a field has been set.

### GetSharedDataIds

`func (o *NFProfileRegistrationError) GetSharedDataIds() []string`

GetSharedDataIds returns the SharedDataIds field if non-nil, zero value otherwise.

### GetSharedDataIdsOk

`func (o *NFProfileRegistrationError) GetSharedDataIdsOk() (*[]string, bool)`

GetSharedDataIdsOk returns a tuple with the SharedDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataIds

`func (o *NFProfileRegistrationError) SetSharedDataIds(v []string)`

SetSharedDataIds sets SharedDataIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



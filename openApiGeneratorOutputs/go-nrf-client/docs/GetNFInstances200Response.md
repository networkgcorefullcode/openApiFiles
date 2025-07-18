# GetNFInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksValueSchema**](LinksValueSchema.md) | List of the URI of NF instances. It has two members whose names are item and self. The item one contains an array of URIs. | [optional] 

## Methods

### NewGetNFInstances200Response

`func NewGetNFInstances200Response() *GetNFInstances200Response`

NewGetNFInstances200Response instantiates a new GetNFInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNFInstances200ResponseWithDefaults

`func NewGetNFInstances200ResponseWithDefaults() *GetNFInstances200Response`

NewGetNFInstances200ResponseWithDefaults instantiates a new GetNFInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetNFInstances200Response) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetNFInstances200Response) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetNFInstances200Response) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetNFInstances200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



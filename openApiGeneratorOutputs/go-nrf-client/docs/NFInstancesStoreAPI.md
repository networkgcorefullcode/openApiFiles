# \NFInstancesStoreAPI

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNFInstances**](NFInstancesStoreAPI.md#GetNFInstances) | **Get** /nf-instances | Retrieves a collection of NF Instances
[**OptionsNFInstances**](NFInstancesStoreAPI.md#OptionsNFInstances) | **Options** /nf-instances | Discover communication options supported by NRF for NF Instances



## GetNFInstances

> GetNFInstances200Response GetNFInstances(ctx).NfType(nfType).Limit(limit).Execute()

Retrieves a collection of NF Instances

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nfType := *openapiclient.NewNFType() // NFType | Type of NF (optional)
	limit := int32(56) // int32 | How many items to return at one time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFInstancesStoreAPI.GetNFInstances(context.Background()).NfType(nfType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFInstancesStoreAPI.GetNFInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNFInstances`: GetNFInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `NFInstancesStoreAPI.GetNFInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNFInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfType** | [**NFType**](NFType.md) | Type of NF | 
 **limit** | **int32** | How many items to return at one time | 

### Return type

[**GetNFInstances200Response**](GetNFInstances200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/3gppHal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsNFInstances

> OptionsNFInstances(ctx).Execute()

Discover communication options supported by NRF for NF Instances

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NFInstancesStoreAPI.OptionsNFInstances(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFInstancesStoreAPI.OptionsNFInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsNFInstancesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


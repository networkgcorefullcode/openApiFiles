# \SharedDataStoreAPI

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSharedDataCollection**](SharedDataStoreAPI.md#GetSharedDataCollection) | **Get** /shared-data | Retrieves a collection of Shared Data



## GetSharedDataCollection

> []SharedData GetSharedDataCollection(ctx).SharedDataIds(sharedDataIds).SupportedFeatures(supportedFeatures).Execute()

Retrieves a collection of Shared Data

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
	sharedDataIds := *openapiclient.NewSharedDataIdList([]string{"SharedDataIds_example"}) // SharedDataIdList | List of shared data IDs
	supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedDataStoreAPI.GetSharedDataCollection(context.Background()).SharedDataIds(sharedDataIds).SupportedFeatures(supportedFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedDataStoreAPI.GetSharedDataCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedDataCollection`: []SharedData
	fmt.Fprintf(os.Stdout, "Response from `SharedDataStoreAPI.GetSharedDataCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedDataCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedDataIds** | [**SharedDataIdList**](SharedDataIdList.md) | List of shared data IDs | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**[]SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


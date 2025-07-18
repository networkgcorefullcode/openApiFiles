# \SharedDataDocumentAPI

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSharedData**](SharedDataDocumentAPI.md#DeleteSharedData) | **Delete** /shared-data/{sharedDataId} | Delete Shared Data identified by a given sharedDataId
[**GetSharedData**](SharedDataDocumentAPI.md#GetSharedData) | **Get** /shared-data/{sharedDataId} | Read the shared data identified by a given NF sharedDataId
[**RegisterSharedData**](SharedDataDocumentAPI.md#RegisterSharedData) | **Put** /shared-data/{sharedDataId} | Register new Shared Data
[**UpdateSharedData**](SharedDataDocumentAPI.md#UpdateSharedData) | **Patch** /shared-data/{sharedDataId} | Update Shared Data



## DeleteSharedData

> DeleteSharedData(ctx, sharedDataId).Execute()

Delete Shared Data identified by a given sharedDataId

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
	sharedDataId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the Shared Data to deregister

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SharedDataDocumentAPI.DeleteSharedData(context.Background(), sharedDataId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedDataDocumentAPI.DeleteSharedData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDataId** | **string** | Unique ID of the Shared Data to deregister | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedData

> SharedData GetSharedData(ctx, sharedDataId).RequesterFeatures(requesterFeatures).Execute()

Read the shared data identified by a given NF sharedDataId

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
	sharedDataId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the Shared Data
	requesterFeatures := "requesterFeatures_example" // string | Features supported by the NF Service Consumer (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedDataDocumentAPI.GetSharedData(context.Background(), sharedDataId).RequesterFeatures(requesterFeatures).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedDataDocumentAPI.GetSharedData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedData`: SharedData
	fmt.Fprintf(os.Stdout, "Response from `SharedDataDocumentAPI.GetSharedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDataId** | **string** | Unique ID of the Shared Data | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requesterFeatures** | **string** | Features supported by the NF Service Consumer | 

### Return type

[**SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSharedData

> SharedData RegisterSharedData(ctx, sharedDataId).SharedData(sharedData).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Register new Shared Data

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
	sharedDataId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the Shared Data to register
	sharedData := *openapiclient.NewSharedData("SharedDataId_example") // SharedData | 
	contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 9110 (optional)
	acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 9110 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedDataDocumentAPI.RegisterSharedData(context.Background(), sharedDataId).SharedData(sharedData).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedDataDocumentAPI.RegisterSharedData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterSharedData`: SharedData
	fmt.Fprintf(os.Stdout, "Response from `SharedDataDocumentAPI.RegisterSharedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDataId** | **string** | Unique ID of the Shared Data to register | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedData** | [**SharedData**](SharedData.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 9110 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 9110 | 

### Return type

[**SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSharedData

> SharedData UpdateSharedData(ctx, sharedDataId).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).IfMatch(ifMatch).Execute()

Update Shared Data

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
	sharedDataId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of shared data to update
	patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
	contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 9110 (optional)
	acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 9110 (optional)
	ifMatch := "ifMatch_example" // string | Validator for conditional requests, as described in IETF RFC 9110, 13.1.1 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedDataDocumentAPI.UpdateSharedData(context.Background(), sharedDataId).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedDataDocumentAPI.UpdateSharedData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSharedData`: SharedData
	fmt.Fprintf(os.Stdout, "Response from `SharedDataDocumentAPI.UpdateSharedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDataId** | **string** | Unique ID of shared data to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 9110 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 9110 | 
 **ifMatch** | **string** | Validator for conditional requests, as described in IETF RFC 9110, 13.1.1 | 

### Return type

[**SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


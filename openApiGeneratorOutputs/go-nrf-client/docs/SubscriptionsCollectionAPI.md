# \SubscriptionsCollectionAPI

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionAPI.md#CreateSubscription) | **Post** /subscriptions | Create a new subscription



## CreateSubscription

> SubscriptionData CreateSubscription(ctx).SubscriptionData(subscriptionData).Execute()

Create a new subscription

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
	subscriptionData := *openapiclient.NewSubscriptionData("NfStatusNotificationUri_example", "SubscriptionId_example") // SubscriptionData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsCollectionAPI.CreateSubscription(context.Background()).SubscriptionData(subscriptionData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: SubscriptionData
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionData** | [**SubscriptionData**](SubscriptionData.md) |  | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


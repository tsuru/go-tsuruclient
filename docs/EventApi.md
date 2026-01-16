# \EventApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventCancel**](EventApi.md#EventCancel) | **Post** /1.1/events/{eventid}/cancel | 
[**EventInfo**](EventApi.md#EventInfo) | **Get** /1.1/events/{eventid} | 
[**EventList**](EventApi.md#EventList) | **Get** /1.1/events | 
[**WebhookCreate**](EventApi.md#WebhookCreate) | **Post** /1.6/events/webhooks | 
[**WebhookDelete**](EventApi.md#WebhookDelete) | **Delete** /1.6/events/webhooks/{name} | 
[**WebhookGet**](EventApi.md#WebhookGet) | **Get** /1.6/events/webhooks/{name} | 
[**WebhookList**](EventApi.md#WebhookList) | **Get** /1.6/events/webhooks | 
[**WebhookUpdate**](EventApi.md#WebhookUpdate) | **Put** /1.6/events/webhooks/{name} | 


# **EventCancel**
> EventCancel(ctx, eventid, eventCancelArgs)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventid** | **string**|  | 
  **eventCancelArgs** | [**EventCancelArgs**](EventCancelArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventInfo**
> Event EventInfo(ctx, eventid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventid** | **string**|  | 

### Return type

[**Event**](Event.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventList**
> []Event EventList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetType** | **optional.String**| Filter by target type | 
 **targetValue** | **optional.String**| Filter by target value | 
 **kindType** | **optional.String**| Filter by kind type | 
 **kindNames** | [**optional.Interface of []string**](string.md)| Filter by kind names | 
 **ownerType** | **optional.String**| Filter by owner type | 
 **ownerName** | **optional.String**| Filter by owner name | 
 **since** | **optional.Time**| Filter events since this date | 
 **until** | **optional.Time**| Filter events until this date | 
 **running** | **optional.Bool**| Filter by running status | 
 **errorOnly** | **optional.Bool**| Filter only events with errors | 
 **limit** | **optional.Int32**| Limit number of results | 
 **skip** | **optional.Int32**| Skip number of results | 
 **sort** | **optional.String**| Sort order | 

### Return type

[**[]Event**](Event.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhookCreate**
> WebhookCreate(ctx, webhook)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhook** | [**Webhook**](Webhook.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhookDelete**
> WebhookDelete(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Webhook name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhookGet**
> Webhook WebhookGet(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Webhook name. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhookList**
> []Webhook WebhookList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhookUpdate**
> WebhookUpdate(ctx, name, webhook)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Webhook name. | 
  **webhook** | [**Webhook**](Webhook.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


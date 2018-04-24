# \ServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceDelete**](ServiceApi.md#InstanceDelete) | **Delete** /1.0/services/{service}/instances/{instance} | 
[**InstanceGet**](ServiceApi.md#InstanceGet) | **Get** /1.0/services/{service}/instances/{instance} | 
[**InstanceUpdate**](ServiceApi.md#InstanceUpdate) | **Put** /1.0/services/{service}/instances/{instance} | 
[**InstancesList**](ServiceApi.md#InstancesList) | **Get** /1.0/services/instances | 


# **InstanceDelete**
> InstanceDelete(ctx, service, instance, unbindall)


Remove service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **unbindall** | **bool**| Remove current binds to this instance | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceGet**
> []ServiceInstanceInfo InstanceGet(ctx, service, instance)


Get service instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 

### Return type

[**[]ServiceInstanceInfo**](ServiceInstanceInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceUpdate**
> InstanceUpdate(ctx, service, instance, optional)


Update a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string**| Service name. | 
 **instance** | **string**| Instance name. | 
 **updateData** | [**ServiceInstanceUpdateData**](ServiceInstanceUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesList**
> []Service InstancesList(ctx, optional)


List service instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **string**| Filter instances by app name | 

### Return type

[**[]Service**](Service.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


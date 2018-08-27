# \ServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceDelete**](ServiceApi.md#InstanceDelete) | **Delete** /1.0/services/{service}/instances/{instance} | 
[**InstanceGet**](ServiceApi.md#InstanceGet) | **Get** /1.0/services/{service}/instances/{instance} | 
[**InstanceUpdate**](ServiceApi.md#InstanceUpdate) | **Put** /1.0/services/{service}/instances/{instance} | 
[**InstancesList**](ServiceApi.md#InstancesList) | **Get** /1.0/services/instances | 
[**ServiceBrokerCreate**](ServiceApi.md#ServiceBrokerCreate) | **Post** /1.7/brokers | 
[**ServiceBrokerDelete**](ServiceApi.md#ServiceBrokerDelete) | **Delete** /1.7/brokers/{name} | 
[**ServiceBrokerList**](ServiceApi.md#ServiceBrokerList) | **Get** /1.7/brokers | 
[**ServiceBrokerUpdate**](ServiceApi.md#ServiceBrokerUpdate) | **Put** /1.7/brokers/{name} | 
[**ServicesList**](ServiceApi.md#ServicesList) | **Get** /1.0/services | 


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
> ServiceInstanceInfo InstanceGet(ctx, service, instance)


Get service instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 

### Return type

[**ServiceInstanceInfo**](ServiceInstanceInfo.md)

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

# **ServiceBrokerCreate**
> ServiceBrokerCreate(ctx, broker)


Create service broker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **broker** | [**ServiceBroker**](ServiceBroker.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBrokerDelete**
> ServiceBrokerDelete(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Service Broker name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBrokerList**
> ServiceBrokerList ServiceBrokerList(ctx, )


List service brokers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceBrokerList**](ServiceBrokerList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBrokerUpdate**
> ServiceBrokerUpdate(ctx, name, broker)


Update service broker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Service Broker name. | 
  **broker** | [**ServiceBroker**](ServiceBroker.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesList**
> []Service ServicesList(ctx, )


List services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Service**](Service.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


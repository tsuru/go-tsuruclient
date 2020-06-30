# \ServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceCreate**](ServiceApi.md#InstanceCreate) | **Post** /1.0/services/{service}/instances | 
[**InstanceDelete**](ServiceApi.md#InstanceDelete) | **Delete** /1.0/services/{service}/instances/{instance} | 
[**InstanceGet**](ServiceApi.md#InstanceGet) | **Get** /1.0/services/{service}/instances/{instance} | 
[**InstanceUpdate**](ServiceApi.md#InstanceUpdate) | **Put** /1.0/services/{service}/instances/{instance} | 
[**InstancesList**](ServiceApi.md#InstancesList) | **Get** /1.0/services/instances | 
[**ServiceBrokerCreate**](ServiceApi.md#ServiceBrokerCreate) | **Post** /1.7/brokers | 
[**ServiceBrokerDelete**](ServiceApi.md#ServiceBrokerDelete) | **Delete** /1.7/brokers/{name} | 
[**ServiceBrokerList**](ServiceApi.md#ServiceBrokerList) | **Get** /1.7/brokers | 
[**ServiceBrokerUpdate**](ServiceApi.md#ServiceBrokerUpdate) | **Put** /1.7/brokers/{name} | 
[**ServiceInstanceBind**](ServiceApi.md#ServiceInstanceBind) | **Put** /1.0/services/{service}/instances/{instance}/{app} | 
[**ServiceInstanceUnbind**](ServiceApi.md#ServiceInstanceUnbind) | **Delete** /1.0/services/{service}/instances/{instance}/{app} | 
[**ServicesList**](ServiceApi.md#ServicesList) | **Get** /1.0/services | 


# **InstanceCreate**
> InstanceCreate(ctx, service, serviceInstance)


Create a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **serviceInstance** | [**ServiceInstance**](ServiceInstance.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceDelete**
> InstanceDelete(ctx, service, instance, unbindall)


Remove service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
 **optional** | ***InstanceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceInstanceUpdateData** | [**optional.Interface of ServiceInstanceUpdateData**](ServiceInstanceUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesList**
> []Service InstancesList(ctx, optional)


List service instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstancesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstancesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **optional.String**| Filter instances by app name | 

### Return type

[**[]Service**](Service.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBrokerCreate**
> ServiceBrokerCreate(ctx, serviceBroker)


Create service broker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceBroker** | [**ServiceBroker**](ServiceBroker.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBrokerDelete**
> ServiceBrokerDelete(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service Broker name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

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
> ServiceBrokerUpdate(ctx, name, serviceBroker)


Update service broker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service Broker name. | 
  **serviceBroker** | [**ServiceBroker**](ServiceBroker.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceBind**
> ServiceInstanceBind(ctx, service, instance, app, optional)


Bind the service instance to app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **app** | **string**| App name. | 
 **optional** | ***ServiceInstanceBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstanceBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **serviceInstanceBind** | [**optional.Interface of ServiceInstanceBind**](ServiceInstanceBind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceUnbind**
> ServiceInstanceUnbind(ctx, service, instance, app, optional)


Unbind the service instance from app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **app** | **string**| App name. | 
 **optional** | ***ServiceInstanceUnbindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstanceUnbindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **serviceInstanceUnbind** | [**optional.Interface of ServiceInstanceUnbind**](ServiceInstanceUnbind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

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


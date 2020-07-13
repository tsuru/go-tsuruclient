# \RouterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RouterCreate**](RouterApi.md#RouterCreate) | **Post** /1.8/routers | 
[**RouterDelete**](RouterApi.md#RouterDelete) | **Delete** /1.8/routers/{name} | 
[**RouterList**](RouterApi.md#RouterList) | **Get** /1.3/routers | 
[**RouterUpdate**](RouterApi.md#RouterUpdate) | **Put** /1.8/routers/{name} | 


# **RouterCreate**
> RouterCreate(ctx, dynamicRouter)


Adds a new dynamic router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dynamicRouter** | [**DynamicRouter**](DynamicRouter.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouterDelete**
> RouterDelete(ctx, name)


Deletes a dynamic router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Dynamic router name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouterList**
> []PlanRouter RouterList(ctx, )


List available routers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PlanRouter**](PlanRouter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouterUpdate**
> RouterUpdate(ctx, name, dynamicRouter)


Updates a dynamic router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Dynamic router name. | 
  **dynamicRouter** | [**DynamicRouter**](DynamicRouter.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


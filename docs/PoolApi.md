# \PoolApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoolCreate**](PoolApi.md#PoolCreate) | **Post** /1.0/pools | 
[**PoolDelete**](PoolApi.md#PoolDelete) | **Delete** /pools/{pool} | 
[**PoolList**](PoolApi.md#PoolList) | **Get** /1.0/pools | 
[**PoolUpdate**](PoolApi.md#PoolUpdate) | **Put** /pools/{pool} | 


# **PoolCreate**
> PoolCreate(ctx, poolCreateData)


Creates a pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolCreateData** | [**PoolCreateData**](PoolCreateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolDelete**
> PoolDelete(ctx, pool)


Deletes a pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pool** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolList**
> []Pool PoolList(ctx, )


List pools.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Pool**](Pool.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolUpdate**
> PoolUpdate(ctx, pool, poolUpdateData)


Updates a pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pool** | **string**|  | 
  **poolUpdateData** | [**PoolUpdateData**](PoolUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


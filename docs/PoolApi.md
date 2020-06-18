# \PoolApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConstraintList**](PoolApi.md#ConstraintList) | **Get** /1.3/constraints | 
[**ConstraintSet**](PoolApi.md#ConstraintSet) | **Put** /1.3/constraints | 
[**PoolCreate**](PoolApi.md#PoolCreate) | **Post** /1.0/pools | 
[**PoolDelete**](PoolApi.md#PoolDelete) | **Delete** /pools/{pool} | 
[**PoolGet**](PoolApi.md#PoolGet) | **Get** /pools/{pool} | 
[**PoolList**](PoolApi.md#PoolList) | **Get** /1.0/pools | 
[**PoolUpdate**](PoolApi.md#PoolUpdate) | **Put** /pools/{pool} | 


# **ConstraintList**
> []PoolConstraint ConstraintList(ctx, )


List pool constraints

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PoolConstraint**](PoolConstraint.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstraintSet**
> ConstraintSet(ctx, optional)


Update a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConstraintSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConstraintSetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolConstraintSet** | [**optional.Interface of PoolConstraintSet**](PoolConstraintSet.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolCreate**
> PoolCreate(ctx, poolCreateData)


Creates a pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolCreateData** | [**PoolCreateData**](PoolCreateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolDelete**
> PoolDelete(ctx, pool)


Deletes a pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pool** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolGet**
> Pool PoolGet(ctx, pool)


Get pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pool** | **string**|  | 

### Return type

[**Pool**](Pool.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

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
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pool** | **string**|  | 
  **poolUpdateData** | [**PoolUpdateData**](PoolUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \NodeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NodeAdd**](NodeApi.md#NodeAdd) | **Post** /1.2/node | 
[**NodeDelete**](NodeApi.md#NodeDelete) | **Delete** /1.2/node/{address} | 
[**NodeGet**](NodeApi.md#NodeGet) | **Get** /1.2/node/{address} | 
[**NodeList**](NodeApi.md#NodeList) | **Get** /1.2/node | 
[**NodeUpdate**](NodeApi.md#NodeUpdate) | **Put** /1.2/node | 


# **NodeAdd**
> NodeAdd(ctx, nodeAddData)


Add a node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeAddData** | [**NodeAddData**](NodeAddData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeDelete**
> NodeDelete(ctx, address, noRebalance, removeIaas)


Remove node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Node address. | 
  **noRebalance** | **bool**| Trigger node rebalance. | 
  **removeIaas** | **bool**| Remove machine from IaaS. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeGet**
> NodeGetResponse NodeGet(ctx, address)


Get node information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Node address. | 

### Return type

[**NodeGetResponse**](NodeGetResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeList**
> NodeListResponse NodeList(ctx, )


List nodes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeListResponse**](NodeListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeUpdate**
> NodeUpdate(ctx, nodeUpdateData)


Update node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeUpdateData** | [**NodeUpdateData**](NodeUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


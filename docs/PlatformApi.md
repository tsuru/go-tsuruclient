# \PlatformApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformAdd**](PlatformApi.md#PlatformAdd) | **Post** /1.0/platforms | 
[**PlatformDelete**](PlatformApi.md#PlatformDelete) | **Delete** /1.0/platforms/{platform} | 
[**PlatformInfo**](PlatformApi.md#PlatformInfo) | **Get** /1.6/platforms/{platform} | 
[**PlatformList**](PlatformApi.md#PlatformList) | **Get** /1.0/platforms | 
[**PlatformUpdate**](PlatformApi.md#PlatformUpdate) | **Put** /1.0/platforms/{platform} | 


# **PlatformAdd**
> PlatformAdd(ctx, name, dockerfileContent)


Add new platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **dockerfileContent** | ***os.File**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformDelete**
> PlatformDelete(ctx, platform)


Delete platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **platform** | **string**| Platform name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformInfo**
> PlatformInfo PlatformInfo(ctx, platform)


Platform info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **platform** | **string**| Platform info. | 

### Return type

[**PlatformInfo**](PlatformInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformList**
> []Platform PlatformList(ctx, )


List platforms.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Platform**](Platform.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformUpdate**
> PlatformUpdate(ctx, platform, dockerfileContent)


Update platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **platform** | **string**| Platform name. | 
  **dockerfileContent** | ***os.File**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


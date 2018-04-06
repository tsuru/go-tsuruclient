# \AppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppList**](AppApi.md#AppList) | **Get** /1.0/apps | 
[**Create**](AppApi.md#Create) | **Post** /1.0/apps | 
[**EnvSet**](AppApi.md#EnvSet) | **Post** /1.0/apps/{app}/env | 


# **AppList**
> []MiniApp AppList(ctx, optional)


List apps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locked** | **bool**| Filter applications by lock status. | 
 **name** | **string**| Filter applications by name. | 
 **owner** | **string**| Filter applications by owner. | 
 **platform** | **string**| Filter applications by platform. | 
 **pool** | **string**| Filter applications by pool. | 
 **status** | **string**| Filter applications by unit status. | 
 **tag** | [**[]string**](string.md)| Filter applications by tag. | 
 **teamOwner** | **string**| Filter applications by team owner. | 

### Return type

[**[]MiniApp**](MiniApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create**
> AppCreateResponse Create(ctx, app)


Create a new app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **app** | [**App**](App.md)|  | 

### Return type

[**AppCreateResponse**](AppCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvSet**
> EnvSetResponse EnvSet(ctx, app, envs)


Set new environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **app** | **string**| App name. | 
  **envs** | [**EnvSetData**](EnvSetData.md)| Environment variables. | 

### Return type

[**EnvSetResponse**](EnvSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


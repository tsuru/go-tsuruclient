# \AppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](AppApi.md#Create) | **Post** /apps | 
[**List**](AppApi.md#List) | **Get** /apps | 


# **Create**
> AppCreateResponse Create(ctx, app)


create a new app

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

# **List**
> List(ctx, optional)


list apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locked** | **bool**| Filter applications by lock status | 
 **name** | **string**| Filter applications by name | 
 **owner** | **string**| Filter applications by owner | 
 **platform** | **string**| Filter applications by platform | 
 **pool** | **string**| Filter applications by pool | 
 **status** | **string**| Filter applications by unit status. | 
 **tag** | [**[]string**](string.md)| Filter applications by tag. | 
 **teamOwner** | **string**| Filter applications by team owner | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


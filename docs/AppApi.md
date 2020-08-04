# \AppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCreate**](AppApi.md#AppCreate) | **Post** /1.0/apps | 
[**AppDelete**](AppApi.md#AppDelete) | **Delete** /1.0/apps/{app} | 
[**AppGet**](AppApi.md#AppGet) | **Get** /1.0/apps/{app} | 
[**AppList**](AppApi.md#AppList) | **Get** /1.0/apps | 
[**AppQuotaChange**](AppApi.md#AppQuotaChange) | **Put** /1.0/apps/{app}/quota | 
[**AppQuotaGet**](AppApi.md#AppQuotaGet) | **Get** /1.0/apps/{app}/quota | 
[**AppRestart**](AppApi.md#AppRestart) | **Post** /1.0/apps/{app}/restart | 
[**AppSetRoutable**](AppApi.md#AppSetRoutable) | **Post** /1.8/apps/{app}/routable | 
[**AppUpdate**](AppApi.md#AppUpdate) | **Put** /1.0/apps/{app} | 
[**EnvGet**](AppApi.md#EnvGet) | **Get** /1.0/apps/{app}/env | 
[**EnvSet**](AppApi.md#EnvSet) | **Post** /1.0/apps/{app}/env | 
[**EnvUnset**](AppApi.md#EnvUnset) | **Delete** /1.0/apps/{app}/env | 


# **AppCreate**
> AppCreateResponse AppCreate(ctx, inputApp)


Create a new app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputApp** | [**InputApp**](InputApp.md)|  | 

### Return type

[**AppCreateResponse**](AppCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppDelete**
> AppDelete(ctx, app)


Delete a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppGet**
> App AppGet(ctx, app)


Get info about a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

[**App**](App.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppList**
> []MiniApp AppList(ctx, optional)


List apps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locked** | **optional.Bool**| Filter applications by lock status. | 
 **name** | **optional.String**| Filter applications by name. | 
 **owner** | **optional.String**| Filter applications by owner. | 
 **platform** | **optional.String**| Filter applications by platform. | 
 **pool** | **optional.String**| Filter applications by pool. | 
 **status** | **optional.String**| Filter applications by unit status. | 
 **tag** | [**optional.Interface of []string**](string.md)| Filter applications by tag. | 
 **teamOwner** | **optional.String**| Filter applications by team owner. | 
 **simplified** | **optional.Bool**| Returns applications without units list. | 

### Return type

[**[]MiniApp**](MiniApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppQuotaChange**
> AppQuotaChange(ctx, app, limit)


Changes the maximum limit of units allowed for use.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **limit** | **float32**| Number of units allowed for use by the current app. Negative number indicates unlimited. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppQuotaGet**
> Quota AppQuotaGet(ctx, app)


Shows app usage info and its quota limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

[**Quota**](Quota.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRestart**
> AppRestart(ctx, app, optional)


Restart App.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
 **optional** | ***AppRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppRestartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **process** | **optional.String**| Process number to be restarted. If the process number will not informed, whole application will be restarted. | 
 **version** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppSetRoutable**
> AppSetRoutable(ctx, app, setRoutableArgs)


Sets a version as routable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **setRoutableArgs** | [**SetRoutableArgs**](SetRoutableArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppUpdate**
> AppUpdate(ctx, app, updateApp)


Update a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **updateApp** | [**UpdateApp**](UpdateApp.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvGet**
> []Env EnvGet(ctx, app, optional)


Get app environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
 **optional** | ***EnvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.String**| Environment variable name. | 

### Return type

[**[]Env**](Env.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvSet**
> []map[string]interface{} EnvSet(ctx, app, envSetData)


Set new environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **envSetData** | [**EnvSetData**](EnvSetData.md)| Environment variables. | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvUnset**
> EnvUnset(ctx, app, env, norestart)


Unset app environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **env** | [**[]string**](string.md)|  | 
  **norestart** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


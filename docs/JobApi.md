# \JobApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJob**](JobApi.md#CreateJob) | **Post** /1.13/jobs | 
[**DeleteJob**](JobApi.md#DeleteJob) | **Delete** /1.13/jobs/{name} | 
[**GetJob**](JobApi.md#GetJob) | **Get** /1.13/jobs/{name} | 
[**JobEnvGet**](JobApi.md#JobEnvGet) | **Get** /1.16/jobs/{name}/env | 
[**JobEnvSet**](JobApi.md#JobEnvSet) | **Post** /1.13/jobs/{name}/env | 
[**JobEnvUnset**](JobApi.md#JobEnvUnset) | **Delete** /1.13/jobs/{name}/env | 
[**JobLog**](JobApi.md#JobLog) | **Get** /1.13/jobs/{name}/log | 
[**ListJob**](JobApi.md#ListJob) | **Get** /1.13/jobs | 
[**TriggerJob**](JobApi.md#TriggerJob) | **Post** /1.13/jobs/{name}/trigger | 
[**UpdateJob**](JobApi.md#UpdateJob) | **Put** /1.13/jobs/{name} | 


# **CreateJob**
> CreateJob(ctx, inputJob)


Create a job that runs periodically

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputJob** | [**InputJob**](InputJob.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> DeleteJob(ctx, name)


Remove a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of job | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJob**
> JobInfo GetJob(ctx, name)


Get a job that runs periodically

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of job | 

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobEnvGet**
> []EnvVar JobEnvGet(ctx, name, optional)


Get job environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Job name. | 
 **optional** | ***JobEnvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobEnvGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.String**| Environment variable name. | 

### Return type

[**[]EnvVar**](EnvVar.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobEnvSet**
> JobEnvSet(ctx, name, envSetData)


Set new environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Job name. | 
  **envSetData** | [**EnvSetData**](EnvSetData.md)| Environment variables. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobEnvUnset**
> JobEnvUnset(ctx, name, env)


Unset job environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Job name. | 
  **env** | [**[]string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobLog**
> JobLog(ctx, name, optional)


Retrieve logs from a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of job | 
 **optional** | ***JobLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **follow** | **optional.Bool**| attach logs to tty | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJob**
> []Job ListJob(ctx, )


List jobs that runs periodically

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Job**](Job.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerJob**
> TriggerJob(ctx, name)


Trigger a job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of job | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateJob**
> UpdateJob(ctx, name, inputJob)


Update a job that runs periodically

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of job | 
  **inputJob** | [**InputJob**](InputJob.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


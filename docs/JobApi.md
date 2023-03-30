# \JobApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJob**](JobApi.md#CreateJob) | **Post** /1.13/jobs | 
[**GetJob**](JobApi.md#GetJob) | **Get** /1.13/jobs/{name} | 
[**ListJob**](JobApi.md#ListJob) | **Get** /1.13/jobs | 


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

# **GetJob**
> JobInfo GetJob(ctx, )


Get a job that runs periodically

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JobInfo**](JobInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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


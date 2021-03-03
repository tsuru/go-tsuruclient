# \PlanApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlan**](PlanApi.md#DeletePlan) | **Delete** /1.0/plans/{plan} | 
[**PlanCreate**](PlanApi.md#PlanCreate) | **Post** /1.0/plans | 
[**PlanList**](PlanApi.md#PlanList) | **Get** /1.0/plans | 


# **DeletePlan**
> DeletePlan(ctx, plan)


Remove a plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **plan** | **string**| Remove current plan | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlanCreate**
> Plan PlanCreate(ctx, plan)


Create a new plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **plan** | [**Plan**](Plan.md)|  | 

### Return type

[**Plan**](Plan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlanList**
> []Plan PlanList(ctx, )


List plans.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Plan**](Plan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


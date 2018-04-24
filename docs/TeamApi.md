# \TeamApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamCreate**](TeamApi.md#TeamCreate) | **Post** /1.0/teams | 
[**TeamDelete**](TeamApi.md#TeamDelete) | **Delete** /1.0/teams/{team} | 
[**TeamGet**](TeamApi.md#TeamGet) | **Get** /1.4/teams/{team} | 
[**TeamUpdate**](TeamApi.md#TeamUpdate) | **Put** /1.6/teams/{team} | 
[**TeamsList**](TeamApi.md#TeamsList) | **Get** /1.0/teams | 


# **TeamCreate**
> TeamCreate(ctx, teamData)


Create a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamData** | [**TeamData**](TeamData.md)| Team name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamDelete**
> TeamDelete(ctx, team)


Delete a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **team** | **string**| Team name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamGet**
> TeamInfo TeamGet(ctx, team)


Get a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **team** | **string**| Team name. | 

### Return type

[**TeamInfo**](TeamInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamUpdate**
> TeamUpdate(ctx, team, updateData)


Update a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **team** | **string**| Team name. | 
  **updateData** | [**UpdateData**](UpdateData.md)| Team update data. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsList**
> []Team TeamsList(ctx, )


List teams.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Team**](Team.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


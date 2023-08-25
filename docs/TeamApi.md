# \TeamApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamCreate**](TeamApi.md#TeamCreate) | **Post** /1.0/teams | 
[**TeamDelete**](TeamApi.md#TeamDelete) | **Delete** /1.0/teams/{team} | 
[**TeamGet**](TeamApi.md#TeamGet) | **Get** /1.4/teams/{team} | 
[**TeamGroupList**](TeamApi.md#TeamGroupList) | **Get** /1.17/teams/{team}/groups | 
[**TeamQuotaChange**](TeamApi.md#TeamQuotaChange) | **Put** /1.12/teams/{team}/quota | 
[**TeamQuotaGet**](TeamApi.md#TeamQuotaGet) | **Get** /1.12/teams/{team}/quota | 
[**TeamUpdate**](TeamApi.md#TeamUpdate) | **Put** /1.6/teams/{team} | 
[**TeamUserList**](TeamApi.md#TeamUserList) | **Get** /1.17/teams/{team}/users | 
[**TeamsList**](TeamApi.md#TeamsList) | **Get** /1.0/teams | 


# **TeamCreate**
> TeamCreate(ctx, teamCreateArgs)


Create a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamCreateArgs** | [**TeamCreateArgs**](TeamCreateArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamDelete**
> TeamDelete(ctx, team)


Delete a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamGet**
> TeamInfo TeamGet(ctx, team)


Get a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 

### Return type

[**TeamInfo**](TeamInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamGroupList**
> []TeamGroup TeamGroupList(ctx, team)


Get groups of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 

### Return type

[**[]TeamGroup**](TeamGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamQuotaChange**
> TeamQuotaChange(ctx, team, limit)


Changes the team's apps limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 
  **limit** | **float32**| New limit of apps. Negative number indicates unlimited. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamQuotaGet**
> UserQuotaViewResponse TeamQuotaGet(ctx, team)


Get quota of a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 

### Return type

[**UserQuotaViewResponse**](UserQuotaViewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamUpdate**
> TeamUpdate(ctx, team, teamUpdateArgs)


Update a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 
  **teamUpdateArgs** | [**TeamUpdateArgs**](TeamUpdateArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamUserList**
> []TeamUser TeamUserList(ctx, team)


Get users of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team** | **string**| Team name. | 

### Return type

[**[]TeamUser**](TeamUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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


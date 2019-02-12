# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToToken**](AuthApi.md#AssignRoleToToken) | **Post** /1.6/roles/{role_name}/token | 
[**DissociateRoleFromToken**](AuthApi.md#DissociateRoleFromToken) | **Delete** /1.6/roles/{role_name}/token/{token_id} | 
[**TeamTokenCreate**](AuthApi.md#TeamTokenCreate) | **Post** /1.6/tokens | 
[**TeamTokenDelete**](AuthApi.md#TeamTokenDelete) | **Delete** /1.6/tokens/{token_id} | 
[**TeamTokenInfo**](AuthApi.md#TeamTokenInfo) | **Get** /1.6/tokens/{token_id} | 
[**TeamTokenUpdate**](AuthApi.md#TeamTokenUpdate) | **Put** /1.6/tokens/{token_id} | 
[**TeamTokensList**](AuthApi.md#TeamTokensList) | **Get** /1.6/tokens | 


# **AssignRoleToToken**
> AssignRoleToToken(ctx, roleName, assignTokenArgs)


Assigns a role to a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **assignTokenArgs** | [**AssignTokenArgs**](AssignTokenArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DissociateRoleFromToken**
> DissociateRoleFromToken(ctx, roleName, tokenId, context)


Dissociates a role from a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **tokenId** | **string**|  | 
  **context** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokenCreate**
> TeamToken TeamTokenCreate(ctx, teamTokenCreateArgs)


Creates a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamTokenCreateArgs** | [**TeamTokenCreateArgs**](TeamTokenCreateArgs.md)|  | 

### Return type

[**TeamToken**](TeamToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokenDelete**
> TeamTokenDelete(ctx, tokenId)


Deletes a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Token ID. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokenInfo**
> TeamToken TeamTokenInfo(ctx, tokenId)


Shows information about a specific token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Token ID. | 

### Return type

[**TeamToken**](TeamToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokenUpdate**
> TeamToken TeamTokenUpdate(ctx, tokenId, teamTokenUpdateArgs)


Updates a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Token ID. | 
  **teamTokenUpdateArgs** | [**TeamTokenUpdateArgs**](TeamTokenUpdateArgs.md)|  | 

### Return type

[**TeamToken**](TeamToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokensList**
> []TeamToken TeamTokensList(ctx, )


List team tokens.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TeamToken**](TeamToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


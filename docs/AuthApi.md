# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToToken**](AuthApi.md#AssignRoleToToken) | **Post** /1.6/roles/{role_name}/token | 
[**DissociateRoleFromToken**](AuthApi.md#DissociateRoleFromToken) | **Delete** /1.6/roles/{role_name}/token/{token_id} | 
[**TeamTokenCreate**](AuthApi.md#TeamTokenCreate) | **Post** /1.6/tokens | 
[**TeamTokenDelete**](AuthApi.md#TeamTokenDelete) | **Delete** /1.6/tokens/{token_id} | 
[**TeamTokenUpdate**](AuthApi.md#TeamTokenUpdate) | **Put** /1.6/tokens/{token_id} | 
[**TeamTokensList**](AuthApi.md#TeamTokensList) | **Get** /1.6/tokens | 


# **AssignRoleToToken**
> AssignRoleToToken(ctx, roleName, token)


Assigns a role to a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleName** | **string**|  | 
  **token** | [**AssignTokenArgs**](AssignTokenArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DissociateRoleFromToken**
> DissociateRoleFromToken(ctx, roleName, tokenId, context)


Dissociates a role from a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleName** | **string**|  | 
  **tokenId** | **string**|  | 
  **context** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokenCreate**
> TeamToken TeamTokenCreate(ctx, token)


Creates a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **token** | [**TeamTokenCreateArgs**](TeamTokenCreateArgs.md)|  | 

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
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **tokenId** | **string**| Token ID. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTokenUpdate**
> TeamToken TeamTokenUpdate(ctx, tokenId, token)


Updates a team token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **tokenId** | **string**| Token ID. | 
  **token** | [**TeamTokenUpdateArgs**](TeamTokenUpdateArgs.md)|  | 

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


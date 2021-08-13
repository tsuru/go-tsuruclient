# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToGroup**](AuthApi.md#AssignRoleToGroup) | **Post** /1.9/roles/{role_name}/group | 
[**AssignRoleToToken**](AuthApi.md#AssignRoleToToken) | **Post** /1.6/roles/{role_name}/token | 
[**CreateRole**](AuthApi.md#CreateRole) | **Post** /1.0/roles | 
[**DefaultRoleAdd**](AuthApi.md#DefaultRoleAdd) | **Post** /1.0/role/default | 
[**DeleteRole**](AuthApi.md#DeleteRole) | **Delete** /1.0/roles/{role_name} | 
[**DissociateRole**](AuthApi.md#DissociateRole) | **Delete** /1.0/roles/{role_name}/user/{email} | 
[**DissociateRoleFromGroup**](AuthApi.md#DissociateRoleFromGroup) | **Delete** /1.6/roles/{role_name}/group/{group_name} | 
[**DissociateRoleFromToken**](AuthApi.md#DissociateRoleFromToken) | **Delete** /1.6/roles/{role_name}/token/{token_id} | 
[**PermissionAdd**](AuthApi.md#PermissionAdd) | **Post** /1.0/roles/{role_name}/permissions | 
[**RemovePermission**](AuthApi.md#RemovePermission) | **Delete** /1.0/roles{role_name}/permissions/{permission} | 
[**RoleAssign**](AuthApi.md#RoleAssign) | **Post** /1,0/roles/{role_name}/user | 
[**RoleDefaultDelete**](AuthApi.md#RoleDefaultDelete) | **Delete** /1.0/role/default | 
[**TeamTokenCreate**](AuthApi.md#TeamTokenCreate) | **Post** /1.6/tokens | 
[**TeamTokenDelete**](AuthApi.md#TeamTokenDelete) | **Delete** /1.6/tokens/{token_id} | 
[**TeamTokenInfo**](AuthApi.md#TeamTokenInfo) | **Get** /1.7/tokens/{token_id} | 
[**TeamTokenUpdate**](AuthApi.md#TeamTokenUpdate) | **Put** /1.6/tokens/{token_id} | 
[**TeamTokensList**](AuthApi.md#TeamTokensList) | **Get** /1.6/tokens | 
[**UpdateRole**](AuthApi.md#UpdateRole) | **Put** /1.0/roles | 


# **AssignRoleToGroup**
> AssignRoleToGroup(ctx, roleName, assignGroupArgs)


Assigns a role to a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **assignGroupArgs** | [**AssignGroupArgs**](AssignGroupArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **CreateRole**
> CreateRole(ctx, roleAddData)


create a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleAddData** | [**RoleAddData**](RoleAddData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DefaultRoleAdd**
> DefaultRoleAdd(ctx, roleDefaultData)


add a default role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleDefaultData** | [**RoleDefaultData**](RoleDefaultData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole(ctx, roleName)


delete a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DissociateRole**
> DissociateRole(ctx, roleName, email)


Dissociate a role from user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **email** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DissociateRoleFromGroup**
> DissociateRoleFromGroup(ctx, roleName, groupName, context)


Dissociates a role from a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **groupName** | **string**|  | 
  **context** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
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

# **PermissionAdd**
> PermissionAdd(ctx, roleName, permissionData)


add a permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **permissionData** | [**PermissionData**](PermissionData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePermission**
> RemovePermission(ctx, roleName, permission)


remove a permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **permission** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoleAssign**
> RoleAssign(ctx, roleName, roleAssignData)


assign a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**|  | 
  **roleAssignData** | [**RoleAssignData**](RoleAssignData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoleDefaultDelete**
> RoleDefaultDelete(ctx, )


Delete a Default role

### Required Parameters
This endpoint does not need any parameter.

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

# **UpdateRole**
> UpdateRole(ctx, roleUpdateData)


update a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleUpdateData** | [**RoleUpdateData**](RoleUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


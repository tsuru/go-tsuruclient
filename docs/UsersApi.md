# \UsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**APITokenGet**](UsersApi.md#APITokenGet) | **Get** /1.0/users/api-key | 
[**APITokenRegenerate**](UsersApi.md#APITokenRegenerate) | **Post** /1.0/users/api-key | 
[**ChangePassword**](UsersApi.md#ChangePassword) | **Put** /1.0/users/password | 
[**ResetPassword**](UsersApi.md#ResetPassword) | **Post** /1.0/users/{email}/password | 
[**SSHKeyAdd**](UsersApi.md#SSHKeyAdd) | **Post** /1.0/users/keys | 
[**SSHKeyList**](UsersApi.md#SSHKeyList) | **Get** /1.0/users/keys | 
[**SSHKeyRemove**](UsersApi.md#SSHKeyRemove) | **Delete** /1.0/users/keys/{key} | 
[**UserCreate**](UsersApi.md#UserCreate) | **Post** /1.0/users | 
[**UserDelete**](UsersApi.md#UserDelete) | **Delete** /1.0/users | 
[**UserGet**](UsersApi.md#UserGet) | **Get** /1.0/users/info | 
[**UserQuotaChange**](UsersApi.md#UserQuotaChange) | **Put** /1.0/users/{email}/quota | 
[**UserQuotaGet**](UsersApi.md#UserQuotaGet) | **Get** /1.0/users/{email}/quota | 
[**UserTokenDelete**](UsersApi.md#UserTokenDelete) | **Delete** /1.0/users/tokens | 
[**UsersList**](UsersApi.md#UsersList) | **Get** /1.0/users | 


# **APITokenGet**
> string APITokenGet(ctx, email)


Show the API token of an user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**|  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **APITokenRegenerate**
> string APITokenRegenerate(ctx, email)


Regenerate the API Token of an user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**|  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePassword**
> ChangePassword(ctx, changePasswordData)


Change password of logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changePasswordData** | [**ChangePasswordData**](ChangePasswordData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> ResetPassword(ctx, email, token)


Reset password of an user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**|  | 
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SSHKeyAdd**
> SSHKeyAdd(ctx, sSHKeyAddData)


Add SSH key to logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sSHKeyAddData** | [**SshKeyAddData**](SshKeyAddData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SSHKeyList**
> SshKeyListResponse SSHKeyList(ctx, )


Show the list of the ssh keys of logged user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SshKeyListResponse**](SSHKeyListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SSHKeyRemove**
> SSHKeyRemove(ctx, key)


Delete one ssh key of logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserCreate**
> UserCreate(ctx, userData)


Create a new user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userData** | [**UserData**](UserData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserDelete**
> UserDelete(ctx, email)


Delete an user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**| User e-mail. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGet**
> User UserGet(ctx, )


Get information on logged user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserQuotaChange**
> UserQuotaChange(ctx, email, limit)


Change quota of an user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**| User e-mail. | 
  **limit** | **int32**| User new quota. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserQuotaGet**
> UserQuotaViewResponse UserQuotaGet(ctx, email)


Get quota of an user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**| User e-mail. | 

### Return type

[**UserQuotaViewResponse**](UserQuotaViewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserTokenDelete**
> UserTokenDelete(ctx, )


Logout.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersList**
> []User UsersList(ctx, email, optional)


List users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **email** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**|  | 
 **role** | **string**|  | 
 **context** | **string**|  | 

### Return type

[**[]User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


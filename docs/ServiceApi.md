# \ServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceCreate**](ServiceApi.md#InstanceCreate) | **Post** /1.0/services/{service}/instances | 
[**InstanceDelete**](ServiceApi.md#InstanceDelete) | **Delete** /1.0/services/{service}/instances/{instance} | 
[**InstanceGet**](ServiceApi.md#InstanceGet) | **Get** /1.0/services/{service}/instances/{instance} | 
[**InstanceUpdate**](ServiceApi.md#InstanceUpdate) | **Put** /1.0/services/{service}/instances/{instance} | 
[**InstancesList**](ServiceApi.md#InstancesList) | **Get** /1.0/services/instances | 
[**JobServiceInstanceBind**](ServiceApi.md#JobServiceInstanceBind) | **Put** /1.13/services/{service}/instances/{instance}/jobs/{job} | 
[**JobServiceInstanceUnbind**](ServiceApi.md#JobServiceInstanceUnbind) | **Delete** /1.13/services/{service}/instances/{instance}/jobs/{job} | 
[**ServiceAddDoc**](ServiceApi.md#ServiceAddDoc) | **Put** /1.0/services/{name}/doc | 
[**ServiceCreate**](ServiceApi.md#ServiceCreate) | **Post** /1.0/services | 
[**ServiceDelete**](ServiceApi.md#ServiceDelete) | **Delete** /1.0/services/{name} | 
[**ServiceDoc**](ServiceApi.md#ServiceDoc) | **Get** /1.0/services/{name}/doc | 
[**ServiceGrantTeam**](ServiceApi.md#ServiceGrantTeam) | **Put** /1.0/services/{service}/team/{team} | 
[**ServiceInfo**](ServiceApi.md#ServiceInfo) | **Get** /1.0/services/{name} | 
[**ServiceInstanceBind**](ServiceApi.md#ServiceInstanceBind) | **Put** /1.13/services/{service}/instances/{instance}/apps/{app} | 
[**ServiceInstanceBind10**](ServiceApi.md#ServiceInstanceBind10) | **Put** /1.0/services/{service}/instances/{instance}/{app} | 
[**ServiceInstanceGrant**](ServiceApi.md#ServiceInstanceGrant) | **Put** /1.0/services/{service}/instances/permission/{instance}/{team} | 
[**ServiceInstanceRevoke**](ServiceApi.md#ServiceInstanceRevoke) | **Delete** /1.0/services/{service}/instances/permission/{instance}/{team} | 
[**ServiceInstanceStatus**](ServiceApi.md#ServiceInstanceStatus) | **Get** /1.0/services/{service}/instances/{instance}/status | 
[**ServiceInstanceUnbind**](ServiceApi.md#ServiceInstanceUnbind) | **Delete** /1.13/services/{service}/instances/{instance}/apps/{app} | 
[**ServiceInstanceUnbind10**](ServiceApi.md#ServiceInstanceUnbind10) | **Delete** /1.0/services/{service}/instances/{instance}/{app} | 
[**ServicePlans**](ServiceApi.md#ServicePlans) | **Get** /1.0/services/{name}/plans | 
[**ServiceRevokeTeam**](ServiceApi.md#ServiceRevokeTeam) | **Delete** /1.0/services/{service}/team/{team} | 
[**ServiceUpdate**](ServiceApi.md#ServiceUpdate) | **Put** /1.0/services/{name} | 
[**ServicesList**](ServiceApi.md#ServicesList) | **Get** /1.0/services | 


# **InstanceCreate**
> InstanceCreate(ctx, service, serviceInstance)


Create a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **serviceInstance** | [**ServiceInstance**](ServiceInstance.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceDelete**
> InstanceDelete(ctx, service, instance, unbindall)


Remove service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **unbindall** | **bool**| Remove current binds to this instance | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceGet**
> ServiceInstanceInfo InstanceGet(ctx, service, instance)


Get service instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 

### Return type

[**ServiceInstanceInfo**](ServiceInstanceInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceUpdate**
> InstanceUpdate(ctx, service, instance, serviceInstanceUpdateData)


Update a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **serviceInstanceUpdateData** | [**ServiceInstanceUpdateData**](ServiceInstanceUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesList**
> []ServiceList InstancesList(ctx, optional)


List service instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstancesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstancesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **optional.String**| Filter instances by app name | 

### Return type

[**[]ServiceList**](ServiceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobServiceInstanceBind**
> JobServiceInstanceBind(ctx, service, instance, job, jobServiceInstanceBind)


Bind the service instance to job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **job** | **string**| Job name. | 
  **jobServiceInstanceBind** | [**JobServiceInstanceBind**](JobServiceInstanceBind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobServiceInstanceUnbind**
> JobServiceInstanceUnbind(ctx, service, instance, job, jobServiceInstanceUnbind)


Unbind the service instance from job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **job** | **string**| Job name. | 
  **jobServiceInstanceUnbind** | [**JobServiceInstanceUnbind**](JobServiceInstanceUnbind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceAddDoc**
> ServiceAddDoc(ctx, name, optional)


Documentation on a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service name. | 
 **optional** | ***ServiceAddDocOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAddDocOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **doc** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceCreate**
> ServiceCreate(ctx, optional)


Creates a new service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **endpoint** | **optional.String**|  | 
 **multiCluster** | **optional.String**|  | 
 **team** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceDelete**
> ServiceDelete(ctx, name)


Delete a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceDoc**
> ServiceDoc(ctx, name)


Documentation on a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceGrantTeam**
> ServiceGrantTeam(ctx, service, team)


Grant access to team for the service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **team** | **string**| Team name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInfo**
> []ServiceInfo ServiceInfo(ctx, name)


Information on a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service name. | 

### Return type

[**[]ServiceInfo**](ServiceInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceBind**
> ServiceInstanceBind(ctx, service, instance, app, serviceInstanceBind)


Bind the service instance to app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **app** | **string**| App name. | 
  **serviceInstanceBind** | [**ServiceInstanceBind**](ServiceInstanceBind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceBind10**
> ServiceInstanceBind10(ctx, service, instance, app, serviceInstanceBind)


Bind the service instance to app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **app** | **string**| App name. | 
  **serviceInstanceBind** | [**ServiceInstanceBind**](ServiceInstanceBind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceGrant**
> ServiceInstanceGrant(ctx, service, instance, team)


Grant access to team for this service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **team** | **string**| Team name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceRevoke**
> ServiceInstanceRevoke(ctx, service, instance, team)


Revoke access to team for this service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **team** | **string**| Team name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceStatus**
> ServiceInstanceStatus(ctx, service, instance)


Status for service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceUnbind**
> ServiceInstanceUnbind(ctx, service, instance, app, serviceInstanceUnbind)


Unbind the service instance from app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **app** | **string**| App name. | 
  **serviceInstanceUnbind** | [**ServiceInstanceUnbind**](ServiceInstanceUnbind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceUnbind10**
> ServiceInstanceUnbind10(ctx, service, instance, app, serviceInstanceUnbind)


Unbind the service instance from app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **instance** | **string**| Instance name. | 
  **app** | **string**| App name. | 
  **serviceInstanceUnbind** | [**ServiceInstanceUnbind**](ServiceInstanceUnbind.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicePlans**
> []ServicePlan ServicePlans(ctx, name, optional)


Plans for a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service name. | 
 **optional** | ***ServicePlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServicePlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pool** | **optional.String**|  | 

### Return type

[**[]ServicePlan**](ServicePlan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceRevokeTeam**
> ServiceRevokeTeam(ctx, service, team)


Revoke access to team for the service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Service name. | 
  **team** | **string**| Team name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceUpdate**
> ServiceUpdate(ctx, name, optional)


Updates a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Service name. | 
 **optional** | ***ServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **endpoint** | **optional.String**|  | 
 **multiCluster** | **optional.String**|  | 
 **team** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesList**
> []ServiceList ServicesList(ctx, )


List services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ServiceList**](ServiceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


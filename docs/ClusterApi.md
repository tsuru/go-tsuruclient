# \ClusterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterCreate**](ClusterApi.md#ClusterCreate) | **Post** /1.3/provisioner/clusters | 
[**ClusterDelete**](ClusterApi.md#ClusterDelete) | **Delete** /1.3/provisioner/clusters/{cluster_name} | 
[**ClusterList**](ClusterApi.md#ClusterList) | **Get** /1.3/provisioner/clusters | 
[**ClusterUpdate**](ClusterApi.md#ClusterUpdate) | **Post** /1.4/provisioner/clusters/{cluster_name} | 
[**ProvisionerList**](ClusterApi.md#ProvisionerList) | **Get** /1.7/provisioner | 


# **ClusterCreate**
> ClusterCreate(ctx, cluster)


Create cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | [**Cluster**](Cluster.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterDelete**
> ClusterDelete(ctx, clusterName)


Delete cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| Cluster name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterList**
> []Cluster ClusterList(ctx, )


List cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterUpdate**
> ClusterUpdate(ctx, clusterName, cluster)


Update cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| Cluster name. | 
  **cluster** | [**Cluster**](Cluster.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionerList**
> []Provisioner ProvisionerList(ctx, )


List provisioners

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Provisioner**](Provisioner.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


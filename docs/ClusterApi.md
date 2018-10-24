# \ClusterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterCreate**](ClusterApi.md#ClusterCreate) | **Post** /1.3/provisioner/clusters | 
[**ClusterDelete**](ClusterApi.md#ClusterDelete) | **Delete** /1.3/provisioner/clusters/{cluster} | 
[**ClusterList**](ClusterApi.md#ClusterList) | **Get** /1.3/provisioner/clusters | 
[**ClusterUpdate**](ClusterApi.md#ClusterUpdate) | **Post** /1.4/provisioner/clusters/{cluster} | 


# **ClusterCreate**
> ClusterCreate(ctx, clusterCreateData)


Create cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterCreateData** | [**Cluster**](Cluster.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterDelete**
> ClusterDelete(ctx, cluster)


Delete cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cluster** | **string**| Cluster name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

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
> ClusterUpdate(ctx, cluster, clusterUpdateData)


Update cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cluster** | **string**| Cluster name. | 
  **clusterUpdateData** | [**Cluster**](Cluster.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


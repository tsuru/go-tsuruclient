# \VolumeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeBind**](VolumeApi.md#VolumeBind) | **Post** /1.4/volumes/{volume}/bind | 
[**VolumeCreate**](VolumeApi.md#VolumeCreate) | **Post** /1.4/volumes | 
[**VolumeDelete**](VolumeApi.md#VolumeDelete) | **Delete** /1.4/volumes/{volume} | 
[**VolumeGet**](VolumeApi.md#VolumeGet) | **Get** /1.4/volumes/{volume} | 
[**VolumeList**](VolumeApi.md#VolumeList) | **Get** /1.4/volumes | 
[**VolumePlansList**](VolumeApi.md#VolumePlansList) | **Get** /1.4/volumeplans | 
[**VolumeUnbind**](VolumeApi.md#VolumeUnbind) | **Delete** /1.4/volumes/{volume}/bind | 
[**VolumeUpdate**](VolumeApi.md#VolumeUpdate) | **Put** /1.4/volumes/{volume} | 


# **VolumeBind**
> VolumeBind(ctx, volume, volumeBindData)


Bind volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volume** | **string**| Volume name. | 
  **volumeBindData** | [**VolumeBindData**](VolumeBindData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeCreate**
> VolumeCreate(ctx, volume)


Create volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volume** | [**Volume**](Volume.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeDelete**
> VolumeDelete(ctx, volume)


Delete volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volume** | **string**| Volume name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeGet**
> Volume VolumeGet(ctx, volume)


Get a volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volume** | **string**| Volume name. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeList**
> []Volume VolumeList(ctx, )


List volumes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Volume**](Volume.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumePlansList**
> map[string]VolumePlan VolumePlansList(ctx, )


List volume plans.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]VolumePlan**](VolumePlan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeUnbind**
> VolumeUnbind(ctx, volume, volumeBindData)


Unbind volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volume** | **string**| Volume name. | 
  **volumeBindData** | [**VolumeBindData**](VolumeBindData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeUpdate**
> VolumeUpdate(ctx, volume, volumeUpdateData)


Update volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volume** | **string**| Volume name. | 
  **volumeUpdateData** | [**VolumeUpdateData**](VolumeUpdateData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


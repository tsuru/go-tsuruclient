# App

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | App name. | 
**Cluster** | **string** | Cluster name | [optional] 
**Cname** | **[]string** | CNames of App | [optional] 
**Deploys** | **int64** | Number of Deploys | [optional] 
**UnitsMetrics** | [**[]UnitMetrics**](UnitMetrics.md) | Unit metrics. | [optional] 
**AutoscaleRecommendation** | [**[]RecommendedResources**](RecommendedResources.md) | Autoscale Recommendations | [optional] 
**Error** | **string** | Errors during AppGet | [optional] 
**Quota** | [**Quota**](Quota.md) |  | [optional] 
**ServiceInstanceBinds** | [**[]AppServiceInstanceBinds**](App_serviceInstanceBinds.md) | Service instance binds on the app | [optional] 
**Routers** | [**[]AppRouters**](App_routers.md) |  | [optional] 
**InternalAddresses** | [**[]AppInternalAddresses**](App_internalAddresses.md) |  | [optional] 
**VolumeBinds** | [**[]AppVolumeBinds**](App_volumeBinds.md) |  | [optional] 
**Tags** | **[]string** | App tags. | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | [optional] 
**ProcessesTweak** | [**[]AppProcessesTweak**](App_processesTweak.md) |  | [optional] 
**Router** | **string** | App router. | [optional] 
**Routeropts** | **map[string]string** | Custom router options. | [optional] 
**Plan** | [**Plan**](Plan.md) |  | [optional] 
**Lock** | [**AppLock**](AppLock.md) |  | [optional] 
**Pool** | **string** | App pool. | [optional] 
**Provisioner** | **string** | App provisioner. | [optional] 
**Platform** | **string** | App platform. | [optional] 
**Description** | **string** | App description. | [optional] 
**TeamOwner** | **string** | Team that owns the app. | [optional] 
**Teams** | **[]string** |  | [optional] 
**Units** | [**[]Unit**](Unit.md) |  | [optional] 
**Ip** | **string** |  | [optional] 
**Owner** | **string** |  | [optional] 
**Autoscale** | [**[]AutoScaleSpec**](AutoScaleSpec.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



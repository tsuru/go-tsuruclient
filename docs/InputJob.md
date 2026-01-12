# InputJob

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Job name. | 
**Tags** | **[]string** | Job tags. | [optional] 
**Schedule** | **string** | how often this job will run. | [optional] 
**Plan** | **string** | job plan name. | [optional] 
**Pool** | **string** | job pool name. | [optional] 
**Description** | **string** | job description. | [optional] 
**TeamOwner** | **string** | Team that owns the job. | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | [optional] 
**Manual** | **bool** | create a manual job. | [optional] 
**ActiveDeadlineSeconds** | **int64** | job active deadline seconds. | [optional] 
**ConcurrencyPolicy** | **string** | concurrency policy. | [optional] 
**Container** | [**JobSpecContainer**](JobSpecContainer.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



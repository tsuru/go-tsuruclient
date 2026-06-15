# Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Username** | **string** |  | [optional] 
**Password** | **string** |  | [optional] 
**Endpoint** | **string** | Single service endpoint URL. When set without \&quot;endpoints\&quot;, it is stored as endpoints[\&quot;production\&quot;]. Mutually exclusive in practice with \&quot;endpoints\&quot; (if both are sent, \&quot;endpoints\&quot; wins). | [optional] 
**Endpoints** | **map[string]string** | Map of endpoint name to URL, allowing multiple endpoints per service (e.g. one per cluster). Sent as form fields with keys \&quot;endpoints.&lt;name&gt;\&quot;. When provided, it overrides \&quot;endpoint\&quot;. | [optional] 
**MultiCluster** | **string** |  | [optional] 
**Team** | **string** |  | [optional] 
**Encoding** | **string** | Encoding format used to communicate with the service API backend. Valid options are \&quot;form\&quot; (default) and \&quot;json\&quot;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueID** | **string** |  | [optional] 
**Lock** | [**EventTarget**](EventTarget.md) |  | [optional] 
**StartTime** | [**time.Time**](time.Time.md) |  | [optional] 
**EndTime** | [**time.Time**](time.Time.md) |  | [optional] 
**ExpireAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Target** | [**EventTarget**](EventTarget.md) |  | [optional] 
**ExtraTargets** | [**[]EventExtraTarget**](EventExtraTarget.md) |  | [optional] 
**Kind** | [**EventKind**](EventKind.md) |  | [optional] 
**Owner** | [**EventOwner**](EventOwner.md) |  | [optional] 
**SourceIP** | **string** |  | [optional] 
**LockUpdateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**Error** | **string** |  | [optional] 
**Log** | **string** |  | [optional] 
**StructuredLog** | [**[]EventLogEntry**](EventLogEntry.md) |  | [optional] 
**CancelInfo** | [**EventCancelInfo**](EventCancelInfo.md) |  | [optional] 
**Cancelable** | **bool** |  | [optional] 
**Running** | **bool** |  | [optional] 
**Allowed** | [**EventAllowedPermission**](EventAllowedPermission.md) |  | [optional] 
**AllowedCancel** | [**EventAllowedPermission**](EventAllowedPermission.md) |  | [optional] 
**Instance** | [**EventTrackedInstance**](EventTrackedInstance.md) |  | [optional] 
**StartCustomData** | [**EventCustomDataRaw**](EventCustomDataRaw.md) |  | [optional] 
**EndCustomData** | [**EventCustomDataRaw**](EventCustomDataRaw.md) |  | [optional] 
**OtherCustomData** | [**EventCustomDataRaw**](EventCustomDataRaw.md) |  | [optional] 
**CustomData** | [**EventCustomData**](EventCustomData.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



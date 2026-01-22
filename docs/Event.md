# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueID** | **string** |  | [optional] 
**Lock** | [**EventLock**](Event_Lock.md) |  | [optional] 
**StartTime** | [**time.Time**](time.Time.md) |  | [optional] 
**EndTime** | [**time.Time**](time.Time.md) |  | [optional] 
**ExpireAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Target** | [**EventLock**](Event_Lock.md) |  | [optional] 
**ExtraTargets** | [**[]EventExtraTargets**](Event_ExtraTargets.md) |  | [optional] 
**Kind** | [**EventKind**](Event_Kind.md) |  | [optional] 
**Owner** | [**EventKind**](Event_Kind.md) |  | [optional] 
**SourceIP** | **string** |  | [optional] 
**LockUpdateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**Error** | **string** |  | [optional] 
**Log** | **string** |  | [optional] 
**StructuredLog** | [**[]EventStructuredLog**](Event_StructuredLog.md) |  | [optional] 
**CancelInfo** | [**EventCancelInfo**](Event_CancelInfo.md) |  | [optional] 
**Cancelable** | **bool** |  | [optional] 
**Running** | **bool** |  | [optional] 
**Allowed** | [**EventAllowed**](Event_Allowed.md) |  | [optional] 
**AllowedCancel** | [**EventAllowed**](Event_Allowed.md) |  | [optional] 
**Instance** | [**EventInstance**](Event_Instance.md) |  | [optional] 
**StartCustomData** | [**EventStartCustomData**](Event_StartCustomData.md) |  | [optional] 
**EndCustomData** | [**EventStartCustomData**](Event_StartCustomData.md) |  | [optional] 
**OtherCustomData** | [**EventStartCustomData**](Event_StartCustomData.md) |  | [optional] 
**CustomData** | [**EventCustomData**](Event_CustomData.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



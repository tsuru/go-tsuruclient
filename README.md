# Go API client for tsuru

Open source, extensible and Docker-based Platform as a Service (PaaS)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.6
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./tsuru"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppApi* | [**AppCreate**](docs/AppApi.md#appcreate) | **Post** /1.0/apps | 
*AppApi* | [**AppDelete**](docs/AppApi.md#appdelete) | **Delete** /1.0/apps/{app} | 
*AppApi* | [**AppGet**](docs/AppApi.md#appget) | **Get** /1.0/apps/{app} | 
*AppApi* | [**AppList**](docs/AppApi.md#applist) | **Get** /1.0/apps | 
*AppApi* | [**AppQuotaChange**](docs/AppApi.md#appquotachange) | **Put** /1.0/apps/{app}/quota | 
*AppApi* | [**AppQuotaGet**](docs/AppApi.md#appquotaget) | **Get** /1.0/apps/{app}/quota | 
*AppApi* | [**AppRestart**](docs/AppApi.md#apprestart) | **Post** /1.0/apps/{app}/restart | 
*AppApi* | [**AppSetRoutable**](docs/AppApi.md#appsetroutable) | **Post** /1.8/apps/{app}/routable | 
*AppApi* | [**AppUpdate**](docs/AppApi.md#appupdate) | **Put** /1.0/apps/{app} | 
*AppApi* | [**AutoScaleAdd**](docs/AppApi.md#autoscaleadd) | **Post** /1.9/apps/{app}/units/autoscale | 
*AppApi* | [**AutoScaleRemove**](docs/AppApi.md#autoscaleremove) | **Delete** /1.9/apps/{app}/units/autoscale | 
*AppApi* | [**EnvGet**](docs/AppApi.md#envget) | **Get** /1.0/apps/{app}/env | 
*AppApi* | [**EnvSet**](docs/AppApi.md#envset) | **Post** /1.0/apps/{app}/env | 
*AppApi* | [**EnvUnset**](docs/AppApi.md#envunset) | **Delete** /1.0/apps/{app}/env | 
*AuthApi* | [**AssignRoleToGroup**](docs/AuthApi.md#assignroletogroup) | **Post** /1.9/roles/{role_name}/group | 
*AuthApi* | [**AssignRoleToToken**](docs/AuthApi.md#assignroletotoken) | **Post** /1.6/roles/{role_name}/token | 
*AuthApi* | [**DissociateRoleFromGroup**](docs/AuthApi.md#dissociaterolefromgroup) | **Delete** /1.6/roles/{role_name}/group/{group_name} | 
*AuthApi* | [**DissociateRoleFromToken**](docs/AuthApi.md#dissociaterolefromtoken) | **Delete** /1.6/roles/{role_name}/token/{token_id} | 
*AuthApi* | [**TeamTokenCreate**](docs/AuthApi.md#teamtokencreate) | **Post** /1.6/tokens | 
*AuthApi* | [**TeamTokenDelete**](docs/AuthApi.md#teamtokendelete) | **Delete** /1.6/tokens/{token_id} | 
*AuthApi* | [**TeamTokenInfo**](docs/AuthApi.md#teamtokeninfo) | **Get** /1.7/tokens/{token_id} | 
*AuthApi* | [**TeamTokenUpdate**](docs/AuthApi.md#teamtokenupdate) | **Put** /1.6/tokens/{token_id} | 
*AuthApi* | [**TeamTokensList**](docs/AuthApi.md#teamtokenslist) | **Get** /1.6/tokens | 
*ClusterApi* | [**ClusterCreate**](docs/ClusterApi.md#clustercreate) | **Post** /1.3/provisioner/clusters | 
*ClusterApi* | [**ClusterDelete**](docs/ClusterApi.md#clusterdelete) | **Delete** /1.3/provisioner/clusters/{cluster_name} | 
*ClusterApi* | [**ClusterInfo**](docs/ClusterApi.md#clusterinfo) | **Get** /1.8/provisioner/clusters/{cluster_name} | 
*ClusterApi* | [**ClusterList**](docs/ClusterApi.md#clusterlist) | **Get** /1.3/provisioner/clusters | 
*ClusterApi* | [**ClusterUpdate**](docs/ClusterApi.md#clusterupdate) | **Post** /1.4/provisioner/clusters/{cluster_name} | 
*ClusterApi* | [**ProvisionerList**](docs/ClusterApi.md#provisionerlist) | **Get** /1.7/provisioner | 
*EventApi* | [**EventCancel**](docs/EventApi.md#eventcancel) | **Post** /1.1/events/{eventid}/cancel | 
*EventApi* | [**WebhookCreate**](docs/EventApi.md#webhookcreate) | **Post** /1.6/events/webhooks | 
*EventApi* | [**WebhookDelete**](docs/EventApi.md#webhookdelete) | **Delete** /1.6/events/webhooks/{name} | 
*EventApi* | [**WebhookGet**](docs/EventApi.md#webhookget) | **Get** /1.6/events/webhooks/{name} | 
*EventApi* | [**WebhookList**](docs/EventApi.md#webhooklist) | **Get** /1.6/events/webhooks | 
*EventApi* | [**WebhookUpdate**](docs/EventApi.md#webhookupdate) | **Put** /1.6/events/webhooks/{name} | 
*NodeApi* | [**NodeAdd**](docs/NodeApi.md#nodeadd) | **Post** /1.2/node | 
*NodeApi* | [**NodeDelete**](docs/NodeApi.md#nodedelete) | **Delete** /1.2/node/{address} | 
*NodeApi* | [**NodeGet**](docs/NodeApi.md#nodeget) | **Get** /1.2/node/{address} | 
*NodeApi* | [**NodeList**](docs/NodeApi.md#nodelist) | **Get** /1.2/node | 
*NodeApi* | [**NodeUpdate**](docs/NodeApi.md#nodeupdate) | **Put** /1.2/node | 
*NodecontainerApi* | [**NodeContainerCreate**](docs/NodecontainerApi.md#nodecontainercreate) | **Post** /1.2/nodecontainers | 
*PlanApi* | [**DeletePlan**](docs/PlanApi.md#deleteplan) | **Delete** /1.0/plans/{name} | 
*PlanApi* | [**PlanCreate**](docs/PlanApi.md#plancreate) | **Post** /1.0/plans | 
*PlanApi* | [**PlanList**](docs/PlanApi.md#planlist) | **Get** /1.0/plans | 
*PlatformApi* | [**PlatformAdd**](docs/PlatformApi.md#platformadd) | **Post** /1.0/platforms | 
*PlatformApi* | [**PlatformDelete**](docs/PlatformApi.md#platformdelete) | **Delete** /1.0/platforms/{platform} | 
*PlatformApi* | [**PlatformInfo**](docs/PlatformApi.md#platforminfo) | **Get** /1.6/platforms/{platform} | 
*PlatformApi* | [**PlatformList**](docs/PlatformApi.md#platformlist) | **Get** /1.0/platforms | 
*PlatformApi* | [**PlatformRollback**](docs/PlatformApi.md#platformrollback) | **Post** /1.6/platforms/{platform}/rollback | 
*PlatformApi* | [**PlatformUpdate**](docs/PlatformApi.md#platformupdate) | **Put** /1.0/platforms/{platform} | 
*PoolApi* | [**ConstraintList**](docs/PoolApi.md#constraintlist) | **Get** /1.3/constraints | 
*PoolApi* | [**ConstraintSet**](docs/PoolApi.md#constraintset) | **Put** /1.3/constraints | 
*PoolApi* | [**PoolCreate**](docs/PoolApi.md#poolcreate) | **Post** /1.0/pools | 
*PoolApi* | [**PoolDelete**](docs/PoolApi.md#pooldelete) | **Delete** /pools/{pool} | 
*PoolApi* | [**PoolGet**](docs/PoolApi.md#poolget) | **Get** /pools/{pool} | 
*PoolApi* | [**PoolList**](docs/PoolApi.md#poollist) | **Get** /1.0/pools | 
*PoolApi* | [**PoolUpdate**](docs/PoolApi.md#poolupdate) | **Put** /pools/{pool} | 
*RouterApi* | [**RouterCreate**](docs/RouterApi.md#routercreate) | **Post** /1.8/routers | 
*RouterApi* | [**RouterDelete**](docs/RouterApi.md#routerdelete) | **Delete** /1.8/routers/{name} | 
*RouterApi* | [**RouterList**](docs/RouterApi.md#routerlist) | **Get** /1.3/routers | 
*RouterApi* | [**RouterUpdate**](docs/RouterApi.md#routerupdate) | **Put** /1.8/routers/{name} | 
*ServiceApi* | [**InstanceCreate**](docs/ServiceApi.md#instancecreate) | **Post** /1.0/services/{service}/instances | 
*ServiceApi* | [**InstanceDelete**](docs/ServiceApi.md#instancedelete) | **Delete** /1.0/services/{service}/instances/{instance} | 
*ServiceApi* | [**InstanceGet**](docs/ServiceApi.md#instanceget) | **Get** /1.0/services/{service}/instances/{instance} | 
*ServiceApi* | [**InstanceUpdate**](docs/ServiceApi.md#instanceupdate) | **Put** /1.0/services/{service}/instances/{instance} | 
*ServiceApi* | [**InstancesList**](docs/ServiceApi.md#instanceslist) | **Get** /1.0/services/instances | 
*ServiceApi* | [**ServiceBrokerCreate**](docs/ServiceApi.md#servicebrokercreate) | **Post** /1.7/brokers | 
*ServiceApi* | [**ServiceBrokerDelete**](docs/ServiceApi.md#servicebrokerdelete) | **Delete** /1.7/brokers/{name} | 
*ServiceApi* | [**ServiceBrokerList**](docs/ServiceApi.md#servicebrokerlist) | **Get** /1.7/brokers | 
*ServiceApi* | [**ServiceBrokerUpdate**](docs/ServiceApi.md#servicebrokerupdate) | **Put** /1.7/brokers/{name} | 
*ServiceApi* | [**ServiceInstanceBind**](docs/ServiceApi.md#serviceinstancebind) | **Put** /1.0/services/{service}/instances/{instance}/{app} | 
*ServiceApi* | [**ServiceInstanceUnbind**](docs/ServiceApi.md#serviceinstanceunbind) | **Delete** /1.0/services/{service}/instances/{instance}/{app} | 
*ServiceApi* | [**ServicesList**](docs/ServiceApi.md#serviceslist) | **Get** /1.0/services | 
*TeamApi* | [**TeamCreate**](docs/TeamApi.md#teamcreate) | **Post** /1.0/teams | 
*TeamApi* | [**TeamDelete**](docs/TeamApi.md#teamdelete) | **Delete** /1.0/teams/{team} | 
*TeamApi* | [**TeamGet**](docs/TeamApi.md#teamget) | **Get** /1.4/teams/{team} | 
*TeamApi* | [**TeamUpdate**](docs/TeamApi.md#teamupdate) | **Put** /1.6/teams/{team} | 
*TeamApi* | [**TeamsList**](docs/TeamApi.md#teamslist) | **Get** /1.0/teams | 
*UserApi* | [**APITokenGet**](docs/UserApi.md#apitokenget) | **Get** /1.0/users/api-key | 
*UserApi* | [**APITokenRegenerate**](docs/UserApi.md#apitokenregenerate) | **Post** /1.0/users/api-key | 
*UserApi* | [**ChangePassword**](docs/UserApi.md#changepassword) | **Put** /1.0/users/password | 
*UserApi* | [**ResetPassword**](docs/UserApi.md#resetpassword) | **Post** /1.0/users/{email}/password | 
*UserApi* | [**SSHKeyAdd**](docs/UserApi.md#sshkeyadd) | **Post** /1.0/users/keys | 
*UserApi* | [**SSHKeyList**](docs/UserApi.md#sshkeylist) | **Get** /1.0/users/keys | 
*UserApi* | [**SSHKeyRemove**](docs/UserApi.md#sshkeyremove) | **Delete** /1.0/users/keys/{key} | 
*UserApi* | [**UserCreate**](docs/UserApi.md#usercreate) | **Post** /1.0/users | 
*UserApi* | [**UserDelete**](docs/UserApi.md#userdelete) | **Delete** /1.0/users | 
*UserApi* | [**UserGet**](docs/UserApi.md#userget) | **Get** /1.0/users/info | 
*UserApi* | [**UserQuotaChange**](docs/UserApi.md#userquotachange) | **Put** /1.0/users/{email}/quota | 
*UserApi* | [**UserQuotaGet**](docs/UserApi.md#userquotaget) | **Get** /1.0/users/{email}/quota | 
*UserApi* | [**UserTokenDelete**](docs/UserApi.md#usertokendelete) | **Delete** /1.0/users/tokens | 
*UserApi* | [**UsersList**](docs/UserApi.md#userslist) | **Get** /1.0/users | 
*VolumeApi* | [**VolumeBind**](docs/VolumeApi.md#volumebind) | **Post** /1.4/volumes/{volume}/bind | 
*VolumeApi* | [**VolumeCreate**](docs/VolumeApi.md#volumecreate) | **Post** /1.4/volumes | 
*VolumeApi* | [**VolumeDelete**](docs/VolumeApi.md#volumedelete) | **Delete** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeGet**](docs/VolumeApi.md#volumeget) | **Get** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeList**](docs/VolumeApi.md#volumelist) | **Get** /1.4/volumes | 
*VolumeApi* | [**VolumePlansList**](docs/VolumeApi.md#volumeplanslist) | **Get** /1.4/volumeplans | 
*VolumeApi* | [**VolumeUnbind**](docs/VolumeApi.md#volumeunbind) | **Delete** /1.4/volumes/{volume}/bind | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppCreateResponse](docs/AppCreateResponse.md)
 - [AssignGroupArgs](docs/AssignGroupArgs.md)
 - [AssignTokenArgs](docs/AssignTokenArgs.md)
 - [AutoScaleSpec](docs/AutoScaleSpec.md)
 - [ChangePasswordData](docs/ChangePasswordData.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterHelp](docs/ClusterHelp.md)
 - [DynamicRouter](docs/DynamicRouter.md)
 - [Env](docs/Env.md)
 - [EnvSetData](docs/EnvSetData.md)
 - [EventCancelArgs](docs/EventCancelArgs.md)
 - [InputApp](docs/InputApp.md)
 - [Lock](docs/Lock.md)
 - [Machine](docs/Machine.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataItem](docs/MetadataItem.md)
 - [MiniApp](docs/MiniApp.md)
 - [Node](docs/Node.md)
 - [NodeAddData](docs/NodeAddData.md)
 - [NodeCheck](docs/NodeCheck.md)
 - [NodeCheckResult](docs/NodeCheckResult.md)
 - [NodeContainer](docs/NodeContainer.md)
 - [NodeContainerConfig](docs/NodeContainerConfig.md)
 - [NodeGetResponse](docs/NodeGetResponse.md)
 - [NodeListResponse](docs/NodeListResponse.md)
 - [NodeStatus](docs/NodeStatus.md)
 - [NodeUpdateData](docs/NodeUpdateData.md)
 - [PermissionUser](docs/PermissionUser.md)
 - [Plan](docs/Plan.md)
 - [PlanOverride](docs/PlanOverride.md)
 - [PlanRouter](docs/PlanRouter.md)
 - [Platform](docs/Platform.md)
 - [PlatformInfo](docs/PlatformInfo.md)
 - [Pool](docs/Pool.md)
 - [PoolConstraint](docs/PoolConstraint.md)
 - [PoolConstraintSet](docs/PoolConstraintSet.md)
 - [PoolCreateData](docs/PoolCreateData.md)
 - [PoolUpdateData](docs/PoolUpdateData.md)
 - [Provisioner](docs/Provisioner.md)
 - [Quota](docs/Quota.md)
 - [RoleInstance](docs/RoleInstance.md)
 - [RoleUser](docs/RoleUser.md)
 - [Router](docs/Router.md)
 - [Service](docs/Service.md)
 - [ServiceBroker](docs/ServiceBroker.md)
 - [ServiceBrokerConfig](docs/ServiceBrokerConfig.md)
 - [ServiceBrokerConfigAuthConfig](docs/ServiceBrokerConfigAuthConfig.md)
 - [ServiceBrokerConfigAuthConfigBasicAuthConfig](docs/ServiceBrokerConfigAuthConfigBasicAuthConfig.md)
 - [ServiceBrokerConfigAuthConfigBearerConfig](docs/ServiceBrokerConfigAuthConfigBearerConfig.md)
 - [ServiceBrokerList](docs/ServiceBrokerList.md)
 - [ServiceInstance](docs/ServiceInstance.md)
 - [ServiceInstanceBind](docs/ServiceInstanceBind.md)
 - [ServiceInstanceBoundUnit](docs/ServiceInstanceBoundUnit.md)
 - [ServiceInstanceInfo](docs/ServiceInstanceInfo.md)
 - [ServiceInstanceUnbind](docs/ServiceInstanceUnbind.md)
 - [ServiceInstanceUpdateData](docs/ServiceInstanceUpdateData.md)
 - [SetRoutableArgs](docs/SetRoutableArgs.md)
 - [SshKeyAddData](docs/SshKeyAddData.md)
 - [SshKeyListResponse](docs/SshKeyListResponse.md)
 - [Team](docs/Team.md)
 - [TeamCreateArgs](docs/TeamCreateArgs.md)
 - [TeamInfo](docs/TeamInfo.md)
 - [TeamToken](docs/TeamToken.md)
 - [TeamTokenCreateArgs](docs/TeamTokenCreateArgs.md)
 - [TeamTokenUpdateArgs](docs/TeamTokenUpdateArgs.md)
 - [TeamUpdateArgs](docs/TeamUpdateArgs.md)
 - [Unit](docs/Unit.md)
 - [UpdateApp](docs/UpdateApp.md)
 - [Url](docs/Url.md)
 - [User](docs/User.md)
 - [UserData](docs/UserData.md)
 - [UserQuotaViewResponse](docs/UserQuotaViewResponse.md)
 - [Volume](docs/Volume.md)
 - [VolumeBind](docs/VolumeBind.md)
 - [VolumeBindData](docs/VolumeBindData.md)
 - [VolumeBindId](docs/VolumeBindId.md)
 - [VolumePlan](docs/VolumePlan.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookEventFilter](docs/WebhookEventFilter.md)


## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author




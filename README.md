# Go API client for tsuru

Open source, extensible and Docker-based Platform as a Service (PaaS)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.24
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
*AppApi* | [**AppBuild**](docs/AppApi.md#appbuild) | **Post** /1.5/apps/{app}/build | 
*AppApi* | [**AppCnameAdd**](docs/AppApi.md#appcnameadd) | **Post** /1.0/apps/{app}/cname | 
*AppApi* | [**AppCnameDelete**](docs/AppApi.md#appcnamedelete) | **Delete** /1.0/apps/{app}/cname | 
*AppApi* | [**AppCreate**](docs/AppApi.md#appcreate) | **Post** /1.0/apps | 
*AppApi* | [**AppDelete**](docs/AppApi.md#appdelete) | **Delete** /1.0/apps/{app} | 
*AppApi* | [**AppDeploy**](docs/AppApi.md#appdeploy) | **Post** /1.0/apps/{app}/deploy | 
*AppApi* | [**AppGet**](docs/AppApi.md#appget) | **Get** /1.0/apps/{app} | 
*AppApi* | [**AppGetCertificates**](docs/AppApi.md#appgetcertificates) | **Get** /1.24/apps/{app}/certificate | 
*AppApi* | [**AppList**](docs/AppApi.md#applist) | **Get** /1.0/apps | 
*AppApi* | [**AppQuotaChange**](docs/AppApi.md#appquotachange) | **Put** /1.0/apps/{app}/quota | 
*AppApi* | [**AppQuotaGet**](docs/AppApi.md#appquotaget) | **Get** /1.0/apps/{app}/quota | 
*AppApi* | [**AppRestart**](docs/AppApi.md#apprestart) | **Post** /1.0/apps/{app}/restart | 
*AppApi* | [**AppRouterAdd**](docs/AppApi.md#approuteradd) | **Post** /1.5/apps/{app}/routers | 
*AppApi* | [**AppRouterDelete**](docs/AppApi.md#approuterdelete) | **Delete** /1.5/apps/{app}/routers/{router} | 
*AppApi* | [**AppRouterList**](docs/AppApi.md#approuterlist) | **Get** /1.5/apps/{app}/routers | 
*AppApi* | [**AppRouterUpdate**](docs/AppApi.md#approuterupdate) | **Put** /1.5/apps/{app}/routers/{router} | 
*AppApi* | [**AppRun**](docs/AppApi.md#apprun) | **Post** /1.0/apps/{app}/run | 
*AppApi* | [**AppSetCertIssuer**](docs/AppApi.md#appsetcertissuer) | **Put** /1.24/apps/{app}/certissuer | 
*AppApi* | [**AppSetRoutable**](docs/AppApi.md#appsetroutable) | **Post** /1.8/apps/{app}/routable | 
*AppApi* | [**AppStart**](docs/AppApi.md#appstart) | **Post** /1.0/apps/{app}/start | 
*AppApi* | [**AppStop**](docs/AppApi.md#appstop) | **Post** /1.0/apps/{app}/stop | 
*AppApi* | [**AppTeamGrant**](docs/AppApi.md#appteamgrant) | **Put** /1.0/apps/{app}/teams/{team} | 
*AppApi* | [**AppTeamRevoke**](docs/AppApi.md#appteamrevoke) | **Delete** /1.0/apps/{app}/teams/{team} | 
*AppApi* | [**AppUnsetCertIssuer**](docs/AppApi.md#appunsetcertissuer) | **Delete** /1.24/apps/{app}/certissuer | 
*AppApi* | [**AppUpdate**](docs/AppApi.md#appupdate) | **Put** /1.0/apps/{app} | 
*AppApi* | [**AutoScaleAdd**](docs/AppApi.md#autoscaleadd) | **Post** /1.9/apps/{app}/units/autoscale | 
*AppApi* | [**AutoScaleInfo**](docs/AppApi.md#autoscaleinfo) | **Get** /1.9/apps/{app}/units/autoscale | 
*AppApi* | [**AutoScaleRemove**](docs/AppApi.md#autoscaleremove) | **Delete** /1.9/apps/{app}/units/autoscale | 
*AppApi* | [**CertificatUnset**](docs/AppApi.md#certificatunset) | **Delete** /1.0/apps/{app}/certificate | 
*AppApi* | [**CertificateSet**](docs/AppApi.md#certificateset) | **Put** /1.0/apps/{app}/certificate | 
*AppApi* | [**EnvGet**](docs/AppApi.md#envget) | **Get** /1.0/apps/{app}/env | 
*AppApi* | [**EnvSet**](docs/AppApi.md#envset) | **Post** /1.0/apps/{app}/env | 
*AppApi* | [**EnvUnset**](docs/AppApi.md#envunset) | **Delete** /1.0/apps/{app}/env | 
*AppApi* | [**UnitsAdd**](docs/AppApi.md#unitsadd) | **Put** /1.0/apps/{app}/units | 
*AppApi* | [**UnitsRemove**](docs/AppApi.md#unitsremove) | **Delete** /1.0/apps/{app}/units | 
*AuthApi* | [**AssignRoleToGroup**](docs/AuthApi.md#assignroletogroup) | **Post** /1.9/roles/{role_name}/group | 
*AuthApi* | [**AssignRoleToToken**](docs/AuthApi.md#assignroletotoken) | **Post** /1.6/roles/{role_name}/token | 
*AuthApi* | [**CreateRole**](docs/AuthApi.md#createrole) | **Post** /1.0/roles | 
*AuthApi* | [**DefaultRoleAdd**](docs/AuthApi.md#defaultroleadd) | **Post** /1.0/role/default | 
*AuthApi* | [**DeleteRole**](docs/AuthApi.md#deleterole) | **Delete** /1.0/roles/{role_name} | 
*AuthApi* | [**DissociateRole**](docs/AuthApi.md#dissociaterole) | **Delete** /1.0/roles/{role_name}/user/{email} | 
*AuthApi* | [**DissociateRoleFromGroup**](docs/AuthApi.md#dissociaterolefromgroup) | **Delete** /1.6/roles/{role_name}/group/{group_name} | 
*AuthApi* | [**DissociateRoleFromToken**](docs/AuthApi.md#dissociaterolefromtoken) | **Delete** /1.6/roles/{role_name}/token/{token_id} | 
*AuthApi* | [**PermissionAdd**](docs/AuthApi.md#permissionadd) | **Post** /1.0/roles/{role_name}/permissions | 
*AuthApi* | [**RemovePermission**](docs/AuthApi.md#removepermission) | **Delete** /1.0/roles{role_name}/permissions/{permission} | 
*AuthApi* | [**RoleAssign**](docs/AuthApi.md#roleassign) | **Post** /1,0/roles/{role_name}/user | 
*AuthApi* | [**RoleDefaultDelete**](docs/AuthApi.md#roledefaultdelete) | **Delete** /1.0/role/default | 
*AuthApi* | [**TeamTokenCreate**](docs/AuthApi.md#teamtokencreate) | **Post** /1.6/tokens | 
*AuthApi* | [**TeamTokenDelete**](docs/AuthApi.md#teamtokendelete) | **Delete** /1.6/tokens/{token_id} | 
*AuthApi* | [**TeamTokenInfo**](docs/AuthApi.md#teamtokeninfo) | **Get** /1.7/tokens/{token_id} | 
*AuthApi* | [**TeamTokenUpdate**](docs/AuthApi.md#teamtokenupdate) | **Put** /1.6/tokens/{token_id} | 
*AuthApi* | [**TeamTokensList**](docs/AuthApi.md#teamtokenslist) | **Get** /1.6/tokens | 
*AuthApi* | [**UpdateRole**](docs/AuthApi.md#updaterole) | **Put** /1.0/roles | 
*ClusterApi* | [**ClusterCreate**](docs/ClusterApi.md#clustercreate) | **Post** /1.3/provisioner/clusters | 
*ClusterApi* | [**ClusterDelete**](docs/ClusterApi.md#clusterdelete) | **Delete** /1.3/provisioner/clusters/{cluster_name} | 
*ClusterApi* | [**ClusterInfo**](docs/ClusterApi.md#clusterinfo) | **Get** /1.8/provisioner/clusters/{cluster_name} | 
*ClusterApi* | [**ClusterList**](docs/ClusterApi.md#clusterlist) | **Get** /1.3/provisioner/clusters | 
*ClusterApi* | [**ClusterUpdate**](docs/ClusterApi.md#clusterupdate) | **Post** /1.4/provisioner/clusters/{cluster_name} | 
*ClusterApi* | [**ProvisionerList**](docs/ClusterApi.md#provisionerlist) | **Get** /1.7/provisioner | 
*EventApi* | [**EventCancel**](docs/EventApi.md#eventcancel) | **Post** /1.1/events/{eventid}/cancel | 
*EventApi* | [**EventInfo**](docs/EventApi.md#eventinfo) | **Get** /1.1/events/{eventid} | 
*EventApi* | [**WebhookCreate**](docs/EventApi.md#webhookcreate) | **Post** /1.6/events/webhooks | 
*EventApi* | [**WebhookDelete**](docs/EventApi.md#webhookdelete) | **Delete** /1.6/events/webhooks/{name} | 
*EventApi* | [**WebhookGet**](docs/EventApi.md#webhookget) | **Get** /1.6/events/webhooks/{name} | 
*EventApi* | [**WebhookList**](docs/EventApi.md#webhooklist) | **Get** /1.6/events/webhooks | 
*EventApi* | [**WebhookUpdate**](docs/EventApi.md#webhookupdate) | **Put** /1.6/events/webhooks/{name} | 
*JobApi* | [**CreateJob**](docs/JobApi.md#createjob) | **Post** /1.13/jobs | 
*JobApi* | [**DeleteJob**](docs/JobApi.md#deletejob) | **Delete** /1.13/jobs/{name} | 
*JobApi* | [**GetJob**](docs/JobApi.md#getjob) | **Get** /1.13/jobs/{name} | 
*JobApi* | [**JobEnvGet**](docs/JobApi.md#jobenvget) | **Get** /1.16/jobs/{name}/env | 
*JobApi* | [**JobEnvSet**](docs/JobApi.md#jobenvset) | **Post** /1.13/jobs/{name}/env | 
*JobApi* | [**JobEnvUnset**](docs/JobApi.md#jobenvunset) | **Delete** /1.13/jobs/{name}/env | 
*JobApi* | [**JobLog**](docs/JobApi.md#joblog) | **Get** /1.13/jobs/{name}/log | 
*JobApi* | [**ListJob**](docs/JobApi.md#listjob) | **Get** /1.13/jobs | 
*JobApi* | [**TriggerJob**](docs/JobApi.md#triggerjob) | **Post** /1.13/jobs/{name}/trigger | 
*JobApi* | [**UpdateJob**](docs/JobApi.md#updatejob) | **Put** /1.13/jobs/{name} | 
*PlanApi* | [**DeletePlan**](docs/PlanApi.md#deleteplan) | **Delete** /1.0/plans/{plan} | 
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
*ServiceApi* | [**JobServiceInstanceBind**](docs/ServiceApi.md#jobserviceinstancebind) | **Put** /1.13/services/{service}/instances/{instance}/jobs/{job} | 
*ServiceApi* | [**JobServiceInstanceUnbind**](docs/ServiceApi.md#jobserviceinstanceunbind) | **Delete** /1.13/services/{service}/instances/{instance}/jobs/{job} | 
*ServiceApi* | [**ServiceAddDoc**](docs/ServiceApi.md#serviceadddoc) | **Put** /1.0/services/{name}/doc | 
*ServiceApi* | [**ServiceBrokerCreate**](docs/ServiceApi.md#servicebrokercreate) | **Post** /1.7/brokers | 
*ServiceApi* | [**ServiceBrokerDelete**](docs/ServiceApi.md#servicebrokerdelete) | **Delete** /1.7/brokers/{name} | 
*ServiceApi* | [**ServiceBrokerList**](docs/ServiceApi.md#servicebrokerlist) | **Get** /1.7/brokers | 
*ServiceApi* | [**ServiceBrokerUpdate**](docs/ServiceApi.md#servicebrokerupdate) | **Put** /1.7/brokers/{name} | 
*ServiceApi* | [**ServiceCreate**](docs/ServiceApi.md#servicecreate) | **Post** /1.0/services | 
*ServiceApi* | [**ServiceDelete**](docs/ServiceApi.md#servicedelete) | **Delete** /1.0/services/{name} | 
*ServiceApi* | [**ServiceDoc**](docs/ServiceApi.md#servicedoc) | **Get** /1.0/services/{name}/doc | 
*ServiceApi* | [**ServiceGrantTeam**](docs/ServiceApi.md#servicegrantteam) | **Put** /1.0/services/{service}/team/{team} | 
*ServiceApi* | [**ServiceInfo**](docs/ServiceApi.md#serviceinfo) | **Get** /1.0/services/{name} | 
*ServiceApi* | [**ServiceInstanceBind**](docs/ServiceApi.md#serviceinstancebind) | **Put** /1.13/services/{service}/instances/{instance}/apps/{app} | 
*ServiceApi* | [**ServiceInstanceBind10**](docs/ServiceApi.md#serviceinstancebind10) | **Put** /1.0/services/{service}/instances/{instance}/{app} | 
*ServiceApi* | [**ServiceInstanceGrant**](docs/ServiceApi.md#serviceinstancegrant) | **Put** /1.0/services/{service}/instances/permission/{instance}/{team} | 
*ServiceApi* | [**ServiceInstanceRevoke**](docs/ServiceApi.md#serviceinstancerevoke) | **Delete** /1.0/services/{service}/instances/permission/{instance}/{team} | 
*ServiceApi* | [**ServiceInstanceStatus**](docs/ServiceApi.md#serviceinstancestatus) | **Get** /1.0/services/{service}/instances/{instance}/status | 
*ServiceApi* | [**ServiceInstanceUnbind**](docs/ServiceApi.md#serviceinstanceunbind) | **Delete** /1.13/services/{service}/instances/{instance}/apps/{app} | 
*ServiceApi* | [**ServiceInstanceUnbind10**](docs/ServiceApi.md#serviceinstanceunbind10) | **Delete** /1.0/services/{service}/instances/{instance}/{app} | 
*ServiceApi* | [**ServicePlans**](docs/ServiceApi.md#serviceplans) | **Get** /1.0/services/{name}/plans | 
*ServiceApi* | [**ServiceRevokeTeam**](docs/ServiceApi.md#servicerevoketeam) | **Delete** /1.0/services/{service}/team/{team} | 
*ServiceApi* | [**ServiceUpdate**](docs/ServiceApi.md#serviceupdate) | **Put** /1.0/services/{name} | 
*ServiceApi* | [**ServicesList**](docs/ServiceApi.md#serviceslist) | **Get** /1.0/services | 
*TeamApi* | [**TeamCreate**](docs/TeamApi.md#teamcreate) | **Post** /1.0/teams | 
*TeamApi* | [**TeamDelete**](docs/TeamApi.md#teamdelete) | **Delete** /1.0/teams/{team} | 
*TeamApi* | [**TeamGet**](docs/TeamApi.md#teamget) | **Get** /1.4/teams/{team} | 
*TeamApi* | [**TeamGroupList**](docs/TeamApi.md#teamgrouplist) | **Get** /1.17/teams/{team}/groups | 
*TeamApi* | [**TeamQuotaChange**](docs/TeamApi.md#teamquotachange) | **Put** /1.12/teams/{team}/quota | 
*TeamApi* | [**TeamQuotaGet**](docs/TeamApi.md#teamquotaget) | **Get** /1.12/teams/{team}/quota | 
*TeamApi* | [**TeamUpdate**](docs/TeamApi.md#teamupdate) | **Put** /1.6/teams/{team} | 
*TeamApi* | [**TeamUserList**](docs/TeamApi.md#teamuserlist) | **Get** /1.17/teams/{team}/users | 
*TeamApi* | [**TeamsList**](docs/TeamApi.md#teamslist) | **Get** /1.0/teams | 
*UserApi* | [**APITokenGet**](docs/UserApi.md#apitokenget) | **Get** /1.0/users/api-key | 
*UserApi* | [**APITokenRegenerate**](docs/UserApi.md#apitokenregenerate) | **Post** /1.0/users/api-key | 
*UserApi* | [**ChangePassword**](docs/UserApi.md#changepassword) | **Put** /1.0/users/password | 
*UserApi* | [**ResetPassword**](docs/UserApi.md#resetpassword) | **Post** /1.0/users/{email}/password | 
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
*VolumeApi* | [**VolumeUpdate**](docs/VolumeApi.md#volumeupdate) | **Put** /1.4/volumes/{volume} | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppCName](docs/AppCName.md)
 - [AppCertificates](docs/AppCertificates.md)
 - [AppCertificatesCnames](docs/AppCertificatesCnames.md)
 - [AppCertificatesRouters](docs/AppCertificatesRouters.md)
 - [AppCreateResponse](docs/AppCreateResponse.md)
 - [AppId](docs/AppId.md)
 - [AppInternalAddresses](docs/AppInternalAddresses.md)
 - [AppLock](docs/AppLock.md)
 - [AppProcess](docs/AppProcess.md)
 - [AppRouter](docs/AppRouter.md)
 - [AppRouters](docs/AppRouters.md)
 - [AppRunOpts](docs/AppRunOpts.md)
 - [AppServiceInstanceBinds](docs/AppServiceInstanceBinds.md)
 - [AppStartStop](docs/AppStartStop.md)
 - [AppVolumeBinds](docs/AppVolumeBinds.md)
 - [AssignGroupArgs](docs/AssignGroupArgs.md)
 - [AssignTokenArgs](docs/AssignTokenArgs.md)
 - [AutoScalePrometheus](docs/AutoScalePrometheus.md)
 - [AutoScaleSchedule](docs/AutoScaleSchedule.md)
 - [AutoScaleSpec](docs/AutoScaleSpec.md)
 - [AutoScaleSpecBehavior](docs/AutoScaleSpecBehavior.md)
 - [AutoScaleSpecBehaviorScaleDown](docs/AutoScaleSpecBehaviorScaleDown.md)
 - [CertIssuerSetData](docs/CertIssuerSetData.md)
 - [CertificateSetData](docs/CertificateSetData.md)
 - [ChangePasswordData](docs/ChangePasswordData.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterHelp](docs/ClusterHelp.md)
 - [ClusterKubeConfig](docs/ClusterKubeConfig.md)
 - [ClusterKubeConfigCluster](docs/ClusterKubeConfigCluster.md)
 - [ClusterKubeConfigUser](docs/ClusterKubeConfigUser.md)
 - [ClusterKubeConfigUserAuthprovider](docs/ClusterKubeConfigUserAuthprovider.md)
 - [ClusterKubeConfigUserExec](docs/ClusterKubeConfigUserExec.md)
 - [ClusterKubeConfigUserExecEnv](docs/ClusterKubeConfigUserExecEnv.md)
 - [DynamicRouter](docs/DynamicRouter.md)
 - [Env](docs/Env.md)
 - [EnvSetData](docs/EnvSetData.md)
 - [EnvVar](docs/EnvVar.md)
 - [Event](docs/Event.md)
 - [EventCancelArgs](docs/EventCancelArgs.md)
 - [EventStartCustomData](docs/EventStartCustomData.md)
 - [EventTarget](docs/EventTarget.md)
 - [InputApp](docs/InputApp.md)
 - [InputJob](docs/InputJob.md)
 - [InputJobContainer](docs/InputJobContainer.md)
 - [Job](docs/Job.md)
 - [JobInfo](docs/JobInfo.md)
 - [JobServiceInstanceBind](docs/JobServiceInstanceBind.md)
 - [JobServiceInstanceUnbind](docs/JobServiceInstanceUnbind.md)
 - [JobSpec](docs/JobSpec.md)
 - [Lock](docs/Lock.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataItem](docs/MetadataItem.md)
 - [MiniApp](docs/MiniApp.md)
 - [PermissionData](docs/PermissionData.md)
 - [PermissionUser](docs/PermissionUser.md)
 - [Plan](docs/Plan.md)
 - [PlanCpuBurst](docs/PlanCpuBurst.md)
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
 - [RecommendedResources](docs/RecommendedResources.md)
 - [RecommendedResourcesRecommendations](docs/RecommendedResourcesRecommendations.md)
 - [RoleAddData](docs/RoleAddData.md)
 - [RoleAssignData](docs/RoleAssignData.md)
 - [RoleDefaultData](docs/RoleDefaultData.md)
 - [RoleInstance](docs/RoleInstance.md)
 - [RoleUpdateData](docs/RoleUpdateData.md)
 - [RoleUser](docs/RoleUser.md)
 - [Router](docs/Router.md)
 - [Service](docs/Service.md)
 - [ServiceBroker](docs/ServiceBroker.md)
 - [ServiceBrokerConfig](docs/ServiceBrokerConfig.md)
 - [ServiceBrokerConfigAuthConfig](docs/ServiceBrokerConfigAuthConfig.md)
 - [ServiceBrokerConfigAuthConfigBasicAuthConfig](docs/ServiceBrokerConfigAuthConfigBasicAuthConfig.md)
 - [ServiceBrokerConfigAuthConfigBearerConfig](docs/ServiceBrokerConfigAuthConfigBearerConfig.md)
 - [ServiceBrokerList](docs/ServiceBrokerList.md)
 - [ServiceDoc](docs/ServiceDoc.md)
 - [ServiceInfo](docs/ServiceInfo.md)
 - [ServiceInstance](docs/ServiceInstance.md)
 - [ServiceInstanceBind](docs/ServiceInstanceBind.md)
 - [ServiceInstanceBoundUnit](docs/ServiceInstanceBoundUnit.md)
 - [ServiceInstanceInfo](docs/ServiceInstanceInfo.md)
 - [ServiceInstanceUnbind](docs/ServiceInstanceUnbind.md)
 - [ServiceInstanceUpdateData](docs/ServiceInstanceUpdateData.md)
 - [ServiceList](docs/ServiceList.md)
 - [ServicePlan](docs/ServicePlan.md)
 - [SetRoutableArgs](docs/SetRoutableArgs.md)
 - [Team](docs/Team.md)
 - [TeamCreateArgs](docs/TeamCreateArgs.md)
 - [TeamGroup](docs/TeamGroup.md)
 - [TeamInfo](docs/TeamInfo.md)
 - [TeamToken](docs/TeamToken.md)
 - [TeamTokenCreateArgs](docs/TeamTokenCreateArgs.md)
 - [TeamTokenUpdateArgs](docs/TeamTokenUpdateArgs.md)
 - [TeamUpdateArgs](docs/TeamUpdateArgs.md)
 - [TeamUser](docs/TeamUser.md)
 - [Unit](docs/Unit.md)
 - [UnitMetrics](docs/UnitMetrics.md)
 - [UnitsDelta](docs/UnitsDelta.md)
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
 - [VolumeUpdateData](docs/VolumeUpdateData.md)
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




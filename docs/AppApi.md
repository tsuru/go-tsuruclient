# \AppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppBuild**](AppApi.md#AppBuild) | **Post** /1.5/apps/{app}/build | 
[**AppCnameAdd**](AppApi.md#AppCnameAdd) | **Post** /1.0/apps/{app}/cname | 
[**AppCnameDelete**](AppApi.md#AppCnameDelete) | **Delete** /1.0/apps/{app}/cname | 
[**AppCreate**](AppApi.md#AppCreate) | **Post** /1.0/apps | 
[**AppDelete**](AppApi.md#AppDelete) | **Delete** /1.0/apps/{app} | 
[**AppDeploy**](AppApi.md#AppDeploy) | **Post** /1.0/apps/{app}/deploy | 
[**AppGet**](AppApi.md#AppGet) | **Get** /1.0/apps/{app} | 
[**AppList**](AppApi.md#AppList) | **Get** /1.0/apps | 
[**AppQuotaChange**](AppApi.md#AppQuotaChange) | **Put** /1.0/apps/{app}/quota | 
[**AppQuotaGet**](AppApi.md#AppQuotaGet) | **Get** /1.0/apps/{app}/quota | 
[**AppRestart**](AppApi.md#AppRestart) | **Post** /1.0/apps/{app}/restart | 
[**AppRouterAdd**](AppApi.md#AppRouterAdd) | **Post** /1.5/apps/{app}/routers | 
[**AppRouterDelete**](AppApi.md#AppRouterDelete) | **Delete** /1.5/apps/{app}/routers/{router} | 
[**AppRouterList**](AppApi.md#AppRouterList) | **Get** /1.5/apps/{app}/routers | 
[**AppRouterUpdate**](AppApi.md#AppRouterUpdate) | **Put** /1.5/apps/{app}/routers/{router} | 
[**AppRun**](AppApi.md#AppRun) | **Post** /1.0/apps/{app}/run | 
[**AppSetCertIssuer**](AppApi.md#AppSetCertIssuer) | **Put** /1.24/apps/{app}/certissuer | 
[**AppSetRoutable**](AppApi.md#AppSetRoutable) | **Post** /1.8/apps/{app}/routable | 
[**AppStart**](AppApi.md#AppStart) | **Post** /1.0/apps/{app}/start | 
[**AppStop**](AppApi.md#AppStop) | **Post** /1.0/apps/{app}/stop | 
[**AppTeamGrant**](AppApi.md#AppTeamGrant) | **Put** /1.0/apps/{app}/teams/{team} | 
[**AppTeamRevoke**](AppApi.md#AppTeamRevoke) | **Delete** /1.0/apps/{app}/teams/{team} | 
[**AppUnsetCertIssuer**](AppApi.md#AppUnsetCertIssuer) | **Delete** /1.24/apps/{app}/certissuer | 
[**AppUpdate**](AppApi.md#AppUpdate) | **Put** /1.0/apps/{app} | 
[**AutoScaleAdd**](AppApi.md#AutoScaleAdd) | **Post** /1.9/apps/{app}/units/autoscale | 
[**AutoScaleInfo**](AppApi.md#AutoScaleInfo) | **Get** /1.9/apps/{app}/units/autoscale | 
[**AutoScaleRemove**](AppApi.md#AutoScaleRemove) | **Delete** /1.9/apps/{app}/units/autoscale | 
[**CertificatUnset**](AppApi.md#CertificatUnset) | **Delete** /1.0/apps/{app}/certificate | 
[**CertificateSet**](AppApi.md#CertificateSet) | **Put** /1.0/apps/{app}/certificate | 
[**EnvGet**](AppApi.md#EnvGet) | **Get** /1.0/apps/{app}/env | 
[**EnvSet**](AppApi.md#EnvSet) | **Post** /1.0/apps/{app}/env | 
[**EnvUnset**](AppApi.md#EnvUnset) | **Delete** /1.0/apps/{app}/env | 
[**UnitsAdd**](AppApi.md#UnitsAdd) | **Put** /1.0/apps/{app}/units | 
[**UnitsRemove**](AppApi.md#UnitsRemove) | **Delete** /1.0/apps/{app}/units | 


# **AppBuild**
> AppBuild(ctx, app, tag, optional)


Build a Tsuru app image (a regular container image) following the deploy's workflow but don't roll it out to the provisioner. That ends up with a container image that can be pulled by the user or used to deploy the app later.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| Application name | 
  **tag** | **string**| Container image&#39;s tag reference. It must be a valid tag reference according to [container image&#39;s name specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md#pulling-manifests).  Examples: &#x60;staging&#x60;, &#x60;feature-abc&#x60;, &#x60;42.1.0&#x60; | 
 **optional** | ***AppBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.*os.File**| App&#39;s source files (tarball compressed with gzip).  NOTE 1: When &#x60;dockerfile&#x60; field is set, this field contains the container build context files.  NOTE 2: Cannot be presented mutually with &#x60;archive-url&#x60; or &#x60;image&#x60;. | 
 **archiveUrl** | **optional.String**| HTTP URL to app&#39;s source files.  NOTE 1: Tsuru API must be able to download the file over HTTP and the downloaded file must be a tarball compressed with gzip.  NOTE 2: Cannot be presented mutually with &#x60;dockerfile&#x60;, &#x60;file&#x60; or &#x60;image&#x60;.  Example: &#x60;https://my-org.example.com/my-app/v42.tar.gz&#x60;. | 
 **image** | **optional.String**| Container image name. It must be a valid container image name according to [container image&#39;s name specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md#pulling-manifests).  NOTE: Cannot be presented mutually with &#x60;archive-url&#x60;, &#x60;dockerfile&#x60; or &#x60;file&#x60;.  Example: &#x60;registry.example.com/my-org/my-app:v42&#x60; | 
 **dockerfile** | **optional.String**| Content of container file (Dockerfile). It must be a valid file according to [Dockerfile reference](https://docs.docker.com/engine/reference/builder/).  NOTE: Cannot be presented mutually with &#x60;archive-url&#x60; or &#x60;image&#x60;.  Example: &#x60;&#x60;&#x60; FROM alpine:3.16 RUN set -x \\\\     apk add --update --no-cache curl ca-certificates COPY ./app.sh ./tsuru.yaml /var/my-app/ WORKDIR /var/my-app EXPOSE 8888/tcp ENTRYPOINT [\\\&quot;/var/my-app/app.sh\\\&quot;] CMD [\\\&quot;--port\\\&quot;, \\\&quot;8888\\\&quot;] &#x60;&#x60;&#x60; | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppCnameAdd**
> AppCnameAdd(ctx, app, appCName)


adds a cname to app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appCName** | [**AppCName**](AppCName.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppCnameDelete**
> AppCnameDelete(ctx, app, appCName)


remove cname from app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appCName** | [**AppCName**](AppCName.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppCreate**
> AppCreateResponse AppCreate(ctx, inputApp)


Create a new app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputApp** | [**InputApp**](InputApp.md)|  | 

### Return type

[**AppCreateResponse**](AppCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppDelete**
> AppDelete(ctx, app)


Delete a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppDeploy**
> AppDeploy(ctx, app, optional)


Build a new Tsuru app image and roll it out to the provisioner.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| Application name | 
 **optional** | ***AppDeployOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppDeployOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newVersion** | **optional.Bool**| Whether should create a new version while preserving the old ones in the provisioner. | [default to false]
 **overrideVersions** | **optional.Bool**| Whether should replace all versions in the provisioner by this new one. | [default to false]
 **message** | **optional.String**| Message to describe the deploy.  Example: &#x60;Updating the driver connector of MongoDB (vX.Y.Z)&#x60; | 
 **origin** | **optional.String**| The type of client who originates the deploy. | 
 **file** | **optional.Interface of *os.File****optional.*os.File**| App&#39;s source files (tarball compressed with gzip).  NOTE 1: When &#x60;dockerfile&#x60; field is set, this field contains the container build context files.  NOTE 2: Cannot be presented mutually with &#x60;archive-url&#x60; or &#x60;image&#x60;. | 
 **archiveUrl** | **optional.String**| HTTP URL to app&#39;s source files.  NOTE 1: Tsuru API must be able to download the file over HTTP and the downloaded file must be a tarball compressed with gzip.  NOTE 2: Cannot be presented mutually with &#x60;dockerfile&#x60;, &#x60;file&#x60; or &#x60;image&#x60;.  Example: &#x60;https://my-org.example.com/my-app/v42.tar.gz&#x60;. | 
 **image** | **optional.String**| Container image name. It must be a valid container image name according to [container image&#39;s name specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md#pulling-manifests).  NOTE: Cannot be presented mutually with &#x60;archive-url&#x60;, &#x60;dockerfile&#x60; or &#x60;file&#x60;.  Example: &#x60;registry.example.com/my-org/my-app:v42&#x60; | 
 **dockerfile** | **optional.String**| Content of container file (Dockerfile). It must be a valid file according to [Dockerfile reference](https://docs.docker.com/engine/reference/builder/).  NOTE: Cannot be presented mutually with &#x60;archive-url&#x60; or &#x60;image&#x60;.  Example: &#x60;&#x60;&#x60; FROM alpine:3.16 RUN set -x \\\\     apk add --update --no-cache curl ca-certificates COPY ./app.sh ./tsuru.yaml /var/my-app/ WORKDIR /var/my-app EXPOSE 8888/tcp ENTRYPOINT [\\\&quot;/var/my-app/app.sh\\\&quot;] CMD [\\\&quot;--port\\\&quot;, \\\&quot;8888\\\&quot;] &#x60;&#x60;&#x60; | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppGet**
> App AppGet(ctx, app)


Get info about a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

[**App**](App.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppList**
> []MiniApp AppList(ctx, optional)


List apps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locked** | **optional.Bool**| Filter applications by lock status. | 
 **name** | **optional.String**| Filter applications by name. | 
 **owner** | **optional.String**| Filter applications by owner. | 
 **platform** | **optional.String**| Filter applications by platform. | 
 **pool** | **optional.String**| Filter applications by pool. | 
 **status** | **optional.String**| Filter applications by unit status. | 
 **tag** | [**optional.Interface of []string**](string.md)| Filter applications by tag. | 
 **teamOwner** | **optional.String**| Filter applications by team owner. | 
 **simplified** | **optional.Bool**| Returns applications without units list. | 

### Return type

[**[]MiniApp**](MiniApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppQuotaChange**
> AppQuotaChange(ctx, app, limit)


Changes the maximum limit of units allowed for use.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **limit** | **float32**| Number of units allowed for use by the current app. Negative number indicates unlimited. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppQuotaGet**
> Quota AppQuotaGet(ctx, app)


Shows app usage info and its quota limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

[**Quota**](Quota.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRestart**
> AppRestart(ctx, app, appStartStop)


Restart App.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appStartStop** | [**AppStartStop**](AppStartStop.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRouterAdd**
> AppRouterAdd(ctx, app, appRouter)


adds a router to app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appRouter** | [**AppRouter**](AppRouter.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRouterDelete**
> AppRouterDelete(ctx, app, router)


Delete a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **router** | **string**| Router name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRouterList**
> []AppRouter AppRouterList(ctx, app)


list routers from an app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

[**[]AppRouter**](AppRouter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRouterUpdate**
> AppRouterUpdate(ctx, app, router, appRouter)


update a router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **router** | **string**| Router name | 
  **appRouter** | [**AppRouter**](AppRouter.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppRun**
> AppRun(ctx, app, appRunOpts)


run commands inside an app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appRunOpts** | [**AppRunOpts**](AppRunOpts.md)| Options to run commands inside an app | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppSetCertIssuer**
> AppSetCertIssuer(ctx, app, certIssuerSetData)


Set the certificate issuer for the app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| Application name | 
  **certIssuerSetData** | [**CertIssuerSetData**](CertIssuerSetData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppSetRoutable**
> AppSetRoutable(ctx, app, setRoutableArgs)


Sets a version as routable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **setRoutableArgs** | [**SetRoutableArgs**](SetRoutableArgs.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppStart**
> AppStart(ctx, app, appStartStop)


Start App.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appStartStop** | [**AppStartStop**](AppStartStop.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppStop**
> AppStop(ctx, app, appStartStop)


Stop App.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **appStartStop** | [**AppStartStop**](AppStartStop.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppTeamGrant**
> AppTeamGrant(ctx, app, team)


grant access to a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **team** | **string**| Team name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppTeamRevoke**
> AppTeamRevoke(ctx, app, team)


grant access to a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **team** | **string**| Team name | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppUnsetCertIssuer**
> AppUnsetCertIssuer(ctx, app, cname)


Unset the certificate issuer for the app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| Application name | 
  **cname** | **string**| Certificate CNAME | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppUpdate**
> AppUpdate(ctx, app, updateApp)


Update a tsuru app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **updateApp** | [**UpdateApp**](UpdateApp.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoScaleAdd**
> AutoScaleAdd(ctx, app, autoScaleSpec)


Add new unit autoscale spec.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **autoScaleSpec** | [**AutoScaleSpec**](AutoScaleSpec.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoScaleInfo**
> []AutoScaleSpec AutoScaleInfo(ctx, app)


List autoscales for app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

[**[]AutoScaleSpec**](AutoScaleSpec.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoScaleRemove**
> AutoScaleRemove(ctx, app, process)


Remove unit autoscale spec.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **process** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CertificatUnset**
> CertificatUnset(ctx, app)


Unset app certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CertificateSet**
> CertificateSet(ctx, app, certificateSetData)


Create a certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **certificateSetData** | [**CertificateSetData**](CertificateSetData.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvGet**
> []EnvVar EnvGet(ctx, app, optional)


Get app environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
 **optional** | ***EnvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.String**| Environment variable name. | 

### Return type

[**[]EnvVar**](EnvVar.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvSet**
> EnvSet(ctx, app, envSetData)


Set new environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **envSetData** | [**EnvSetData**](EnvSetData.md)| Environment variables. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnvUnset**
> EnvUnset(ctx, app, env, norestart)


Unset app environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **env** | [**[]string**](string.md)|  | 
  **norestart** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnitsAdd**
> UnitsAdd(ctx, app, unitsDelta)


Add units to app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **unitsDelta** | [**UnitsDelta**](UnitsDelta.md)| number of units to add | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnitsRemove**
> UnitsRemove(ctx, app, unitsDelta)


Remove units from app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| App name. | 
  **unitsDelta** | [**UnitsDelta**](UnitsDelta.md)| number of units to remove | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-json-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


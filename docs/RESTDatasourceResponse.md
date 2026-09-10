# RESTDatasourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Attributes** | Pointer to [**[]DataSourceOutput**](DataSourceOutput.md) | List of attributes to fetch for this datasource | [optional] 
**AuthenticationType** | **string** | The authentication type to use while making the REST call. Is linked to &#x60;credentials&#x60;. | 
**Credentials** | Pointer to **string** | Name of the [credentials](#tag/security.credentials) to use for authentication | [optional] 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**ExpectedHttpCodes** | **[]int64** | The success HTTP codes for the request. If the return code is not in this list, the request will be considered failed. | 
**Headers** | Pointer to [**[]Header**](Header.md) | The headers of the request | [optional] 
**Method** | **string** | The HTTP method to use for the request | 
**Name** | **string** | Name of the datasource | 
**NotFoundHttpCodes** | Pointer to **[]int64** | HTTP response codes that indicate a \&quot;not found\&quot; result. Must not overlap with &#x60;expectedHttpCodes&#x60;. When received, the datasource will return a &#x60;not_found&#x60; status with no results. Combined with the &#x60;mandatory&#x60; parameter on a datasource flow entry, this can be used to stop a flow when a required resource is not found. | [optional] 
**Payload** | Pointer to **NullableString** | The body of the request | [optional] 
**PayloadType** | Pointer to **NullableString** | For UI purposes in order to format the body correctly | [optional] 
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while making the request | [optional] 
**Timeout** | **string** | Timeout for the HTTP request. | 
**Type** | **string** | Type of datasource | 
**Url** | **string** | The URL to request | 

## Methods

### NewRESTDatasourceResponse

`func NewRESTDatasourceResponse(id string, authenticationType string, expectedHttpCodes []int64, method string, name string, timeout string, type_ string, url string, ) *RESTDatasourceResponse`

NewRESTDatasourceResponse instantiates a new RESTDatasourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRESTDatasourceResponseWithDefaults

`func NewRESTDatasourceResponseWithDefaults() *RESTDatasourceResponse`

NewRESTDatasourceResponseWithDefaults instantiates a new RESTDatasourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RESTDatasourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RESTDatasourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RESTDatasourceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *RESTDatasourceResponse) GetAttributes() []DataSourceOutput`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RESTDatasourceResponse) GetAttributesOk() (*[]DataSourceOutput, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RESTDatasourceResponse) SetAttributes(v []DataSourceOutput)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RESTDatasourceResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *RESTDatasourceResponse) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *RESTDatasourceResponse) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetAuthenticationType

`func (o *RESTDatasourceResponse) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *RESTDatasourceResponse) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *RESTDatasourceResponse) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetCredentials

`func (o *RESTDatasourceResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RESTDatasourceResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RESTDatasourceResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RESTDatasourceResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDescription

`func (o *RESTDatasourceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RESTDatasourceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RESTDatasourceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RESTDatasourceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *RESTDatasourceResponse) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RESTDatasourceResponse) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RESTDatasourceResponse) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RESTDatasourceResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RESTDatasourceResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RESTDatasourceResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExpectedHttpCodes

`func (o *RESTDatasourceResponse) GetExpectedHttpCodes() []int64`

GetExpectedHttpCodes returns the ExpectedHttpCodes field if non-nil, zero value otherwise.

### GetExpectedHttpCodesOk

`func (o *RESTDatasourceResponse) GetExpectedHttpCodesOk() (*[]int64, bool)`

GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedHttpCodes

`func (o *RESTDatasourceResponse) SetExpectedHttpCodes(v []int64)`

SetExpectedHttpCodes sets ExpectedHttpCodes field to given value.


### GetHeaders

`func (o *RESTDatasourceResponse) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RESTDatasourceResponse) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RESTDatasourceResponse) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RESTDatasourceResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *RESTDatasourceResponse) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *RESTDatasourceResponse) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *RESTDatasourceResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RESTDatasourceResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RESTDatasourceResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetName

`func (o *RESTDatasourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RESTDatasourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RESTDatasourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNotFoundHttpCodes

`func (o *RESTDatasourceResponse) GetNotFoundHttpCodes() []int64`

GetNotFoundHttpCodes returns the NotFoundHttpCodes field if non-nil, zero value otherwise.

### GetNotFoundHttpCodesOk

`func (o *RESTDatasourceResponse) GetNotFoundHttpCodesOk() (*[]int64, bool)`

GetNotFoundHttpCodesOk returns a tuple with the NotFoundHttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundHttpCodes

`func (o *RESTDatasourceResponse) SetNotFoundHttpCodes(v []int64)`

SetNotFoundHttpCodes sets NotFoundHttpCodes field to given value.

### HasNotFoundHttpCodes

`func (o *RESTDatasourceResponse) HasNotFoundHttpCodes() bool`

HasNotFoundHttpCodes returns a boolean if a field has been set.

### GetPayload

`func (o *RESTDatasourceResponse) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *RESTDatasourceResponse) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *RESTDatasourceResponse) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *RESTDatasourceResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *RESTDatasourceResponse) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *RESTDatasourceResponse) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetPayloadType

`func (o *RESTDatasourceResponse) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *RESTDatasourceResponse) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *RESTDatasourceResponse) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *RESTDatasourceResponse) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### SetPayloadTypeNil

`func (o *RESTDatasourceResponse) SetPayloadTypeNil(b bool)`

 SetPayloadTypeNil sets the value for PayloadType to be an explicit nil

### UnsetPayloadType
`func (o *RESTDatasourceResponse) UnsetPayloadType()`

UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
### GetProxy

`func (o *RESTDatasourceResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *RESTDatasourceResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *RESTDatasourceResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *RESTDatasourceResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *RESTDatasourceResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *RESTDatasourceResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *RESTDatasourceResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RESTDatasourceResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RESTDatasourceResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetType

`func (o *RESTDatasourceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RESTDatasourceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RESTDatasourceResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *RESTDatasourceResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RESTDatasourceResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RESTDatasourceResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



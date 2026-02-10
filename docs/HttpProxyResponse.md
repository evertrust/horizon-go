# HttpProxyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** | Name of the proxy | 
**Host** | **string** | Hostname of the proxy | 
**Port** | **int64** | Port of the proxy | 
**Credentials** | Pointer to **NullableString** | Name of the &#x60;password&#x60; [credentials](#tag/security.credentials) to use for Proxy Basic Authentication | [optional] 

## Methods

### NewHttpProxyResponse

`func NewHttpProxyResponse(id string, name string, host string, port int64, ) *HttpProxyResponse`

NewHttpProxyResponse instantiates a new HttpProxyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpProxyResponseWithDefaults

`func NewHttpProxyResponseWithDefaults() *HttpProxyResponse`

NewHttpProxyResponseWithDefaults instantiates a new HttpProxyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HttpProxyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpProxyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpProxyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HttpProxyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpProxyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpProxyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *HttpProxyResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HttpProxyResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HttpProxyResponse) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *HttpProxyResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpProxyResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpProxyResponse) SetPort(v int64)`

SetPort sets Port field to given value.


### GetCredentials

`func (o *HttpProxyResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HttpProxyResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HttpProxyResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *HttpProxyResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *HttpProxyResponse) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *HttpProxyResponse) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



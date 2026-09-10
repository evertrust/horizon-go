# DCVDomainsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | [**[]DCVPolicyDomainStatus**](DCVPolicyDomainStatus.md) | Validation status for each domain managed by this policy | 
**Error** | Pointer to **NullableString** | Error message when the domain list could not be retrieved from the provider | [optional] 

## Methods

### NewDCVDomainsStatus

`func NewDCVDomainsStatus(domains []DCVPolicyDomainStatus, ) *DCVDomainsStatus`

NewDCVDomainsStatus instantiates a new DCVDomainsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCVDomainsStatusWithDefaults

`func NewDCVDomainsStatusWithDefaults() *DCVDomainsStatus`

NewDCVDomainsStatusWithDefaults instantiates a new DCVDomainsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *DCVDomainsStatus) GetDomains() []DCVPolicyDomainStatus`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *DCVDomainsStatus) GetDomainsOk() (*[]DCVPolicyDomainStatus, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *DCVDomainsStatus) SetDomains(v []DCVPolicyDomainStatus)`

SetDomains sets Domains field to given value.


### GetError

`func (o *DCVDomainsStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DCVDomainsStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DCVDomainsStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DCVDomainsStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DCVDomainsStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DCVDomainsStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



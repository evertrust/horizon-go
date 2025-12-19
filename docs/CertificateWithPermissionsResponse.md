# CertificateWithPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CertificateResponse**](CertificateResponse.md) | The certificate object | 
**Permissions** | [**CertificatePermissions**](CertificatePermissions.md) | The permissions of the currently authenticated principal on the certificate | 

## Methods

### NewCertificateWithPermissionsResponse

`func NewCertificateWithPermissionsResponse(certificate CertificateResponse, permissions CertificatePermissions, ) *CertificateWithPermissionsResponse`

NewCertificateWithPermissionsResponse instantiates a new CertificateWithPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithPermissionsResponseWithDefaults

`func NewCertificateWithPermissionsResponseWithDefaults() *CertificateWithPermissionsResponse`

NewCertificateWithPermissionsResponseWithDefaults instantiates a new CertificateWithPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CertificateWithPermissionsResponse) GetCertificate() CertificateResponse`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateWithPermissionsResponse) GetCertificateOk() (*CertificateResponse, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateWithPermissionsResponse) SetCertificate(v CertificateResponse)`

SetCertificate sets Certificate field to given value.


### GetPermissions

`func (o *CertificateWithPermissionsResponse) GetPermissions() CertificatePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CertificateWithPermissionsResponse) GetPermissionsOk() (*CertificatePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CertificateWithPermissionsResponse) SetPermissions(v CertificatePermissions)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



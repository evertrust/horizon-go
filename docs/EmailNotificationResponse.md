# EmailNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**EmailTemplate** | [**EmailTemplate**](EmailTemplate.md) |  | 
**IfPkcs12** | Pointer to **NullableBool** | On events triggering an enrollment, select if mail is sent: - **Always**: set the value to &#x60;null&#x60;  - **Only when a PKCS#12 is available in the request**: set the value to &#x60;true&#x60;  - **Only when a PKCS#12 is not in the request**: set the value to &#x60;false&#x60;  | [optional] 
**AttachPemCertificate** | Pointer to **NullableBool** | Attach the certificate in PEM format if available | [optional] 
**AttachPemBundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PEM format if available | [optional] 
**AttachDerCertificate** | Pointer to **NullableBool** | Attach the certificate in DER format if available | [optional] 
**AttachPkcs7** | Pointer to **NullableBool** | Attach the certificate in PKCS7 format if available | [optional] 
**AttachPkcs7Bundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PKCS7 format if available | [optional] 
**AttachPkcs12** | Pointer to **NullableBool** | Attach the certificate in PKCS#12 format if available | [optional] 
**Name** | **string** | Name of the notification | 
**Retries** | Pointer to **NullableInt64** | Number of retries when the notification fails | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 
**LicenseUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 

## Methods

### NewEmailNotificationResponse

`func NewEmailNotificationResponse(id string, type_ string, emailTemplate EmailTemplate, name string, events []string, ) *EmailNotificationResponse`

NewEmailNotificationResponse instantiates a new EmailNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationResponseWithDefaults

`func NewEmailNotificationResponseWithDefaults() *EmailNotificationResponse`

NewEmailNotificationResponseWithDefaults instantiates a new EmailNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailNotificationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailNotificationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailNotificationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *EmailNotificationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailNotificationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailNotificationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetEmailTemplate

`func (o *EmailNotificationResponse) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *EmailNotificationResponse) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *EmailNotificationResponse) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetIfPkcs12

`func (o *EmailNotificationResponse) GetIfPkcs12() bool`

GetIfPkcs12 returns the IfPkcs12 field if non-nil, zero value otherwise.

### GetIfPkcs12Ok

`func (o *EmailNotificationResponse) GetIfPkcs12Ok() (*bool, bool)`

GetIfPkcs12Ok returns a tuple with the IfPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfPkcs12

`func (o *EmailNotificationResponse) SetIfPkcs12(v bool)`

SetIfPkcs12 sets IfPkcs12 field to given value.

### HasIfPkcs12

`func (o *EmailNotificationResponse) HasIfPkcs12() bool`

HasIfPkcs12 returns a boolean if a field has been set.

### SetIfPkcs12Nil

`func (o *EmailNotificationResponse) SetIfPkcs12Nil(b bool)`

 SetIfPkcs12Nil sets the value for IfPkcs12 to be an explicit nil

### UnsetIfPkcs12
`func (o *EmailNotificationResponse) UnsetIfPkcs12()`

UnsetIfPkcs12 ensures that no value is present for IfPkcs12, not even an explicit nil
### GetAttachPemCertificate

`func (o *EmailNotificationResponse) GetAttachPemCertificate() bool`

GetAttachPemCertificate returns the AttachPemCertificate field if non-nil, zero value otherwise.

### GetAttachPemCertificateOk

`func (o *EmailNotificationResponse) GetAttachPemCertificateOk() (*bool, bool)`

GetAttachPemCertificateOk returns a tuple with the AttachPemCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemCertificate

`func (o *EmailNotificationResponse) SetAttachPemCertificate(v bool)`

SetAttachPemCertificate sets AttachPemCertificate field to given value.

### HasAttachPemCertificate

`func (o *EmailNotificationResponse) HasAttachPemCertificate() bool`

HasAttachPemCertificate returns a boolean if a field has been set.

### SetAttachPemCertificateNil

`func (o *EmailNotificationResponse) SetAttachPemCertificateNil(b bool)`

 SetAttachPemCertificateNil sets the value for AttachPemCertificate to be an explicit nil

### UnsetAttachPemCertificate
`func (o *EmailNotificationResponse) UnsetAttachPemCertificate()`

UnsetAttachPemCertificate ensures that no value is present for AttachPemCertificate, not even an explicit nil
### GetAttachPemBundle

`func (o *EmailNotificationResponse) GetAttachPemBundle() bool`

GetAttachPemBundle returns the AttachPemBundle field if non-nil, zero value otherwise.

### GetAttachPemBundleOk

`func (o *EmailNotificationResponse) GetAttachPemBundleOk() (*bool, bool)`

GetAttachPemBundleOk returns a tuple with the AttachPemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemBundle

`func (o *EmailNotificationResponse) SetAttachPemBundle(v bool)`

SetAttachPemBundle sets AttachPemBundle field to given value.

### HasAttachPemBundle

`func (o *EmailNotificationResponse) HasAttachPemBundle() bool`

HasAttachPemBundle returns a boolean if a field has been set.

### SetAttachPemBundleNil

`func (o *EmailNotificationResponse) SetAttachPemBundleNil(b bool)`

 SetAttachPemBundleNil sets the value for AttachPemBundle to be an explicit nil

### UnsetAttachPemBundle
`func (o *EmailNotificationResponse) UnsetAttachPemBundle()`

UnsetAttachPemBundle ensures that no value is present for AttachPemBundle, not even an explicit nil
### GetAttachDerCertificate

`func (o *EmailNotificationResponse) GetAttachDerCertificate() bool`

GetAttachDerCertificate returns the AttachDerCertificate field if non-nil, zero value otherwise.

### GetAttachDerCertificateOk

`func (o *EmailNotificationResponse) GetAttachDerCertificateOk() (*bool, bool)`

GetAttachDerCertificateOk returns a tuple with the AttachDerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDerCertificate

`func (o *EmailNotificationResponse) SetAttachDerCertificate(v bool)`

SetAttachDerCertificate sets AttachDerCertificate field to given value.

### HasAttachDerCertificate

`func (o *EmailNotificationResponse) HasAttachDerCertificate() bool`

HasAttachDerCertificate returns a boolean if a field has been set.

### SetAttachDerCertificateNil

`func (o *EmailNotificationResponse) SetAttachDerCertificateNil(b bool)`

 SetAttachDerCertificateNil sets the value for AttachDerCertificate to be an explicit nil

### UnsetAttachDerCertificate
`func (o *EmailNotificationResponse) UnsetAttachDerCertificate()`

UnsetAttachDerCertificate ensures that no value is present for AttachDerCertificate, not even an explicit nil
### GetAttachPkcs7

`func (o *EmailNotificationResponse) GetAttachPkcs7() bool`

GetAttachPkcs7 returns the AttachPkcs7 field if non-nil, zero value otherwise.

### GetAttachPkcs7Ok

`func (o *EmailNotificationResponse) GetAttachPkcs7Ok() (*bool, bool)`

GetAttachPkcs7Ok returns a tuple with the AttachPkcs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7

`func (o *EmailNotificationResponse) SetAttachPkcs7(v bool)`

SetAttachPkcs7 sets AttachPkcs7 field to given value.

### HasAttachPkcs7

`func (o *EmailNotificationResponse) HasAttachPkcs7() bool`

HasAttachPkcs7 returns a boolean if a field has been set.

### SetAttachPkcs7Nil

`func (o *EmailNotificationResponse) SetAttachPkcs7Nil(b bool)`

 SetAttachPkcs7Nil sets the value for AttachPkcs7 to be an explicit nil

### UnsetAttachPkcs7
`func (o *EmailNotificationResponse) UnsetAttachPkcs7()`

UnsetAttachPkcs7 ensures that no value is present for AttachPkcs7, not even an explicit nil
### GetAttachPkcs7Bundle

`func (o *EmailNotificationResponse) GetAttachPkcs7Bundle() bool`

GetAttachPkcs7Bundle returns the AttachPkcs7Bundle field if non-nil, zero value otherwise.

### GetAttachPkcs7BundleOk

`func (o *EmailNotificationResponse) GetAttachPkcs7BundleOk() (*bool, bool)`

GetAttachPkcs7BundleOk returns a tuple with the AttachPkcs7Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7Bundle

`func (o *EmailNotificationResponse) SetAttachPkcs7Bundle(v bool)`

SetAttachPkcs7Bundle sets AttachPkcs7Bundle field to given value.

### HasAttachPkcs7Bundle

`func (o *EmailNotificationResponse) HasAttachPkcs7Bundle() bool`

HasAttachPkcs7Bundle returns a boolean if a field has been set.

### SetAttachPkcs7BundleNil

`func (o *EmailNotificationResponse) SetAttachPkcs7BundleNil(b bool)`

 SetAttachPkcs7BundleNil sets the value for AttachPkcs7Bundle to be an explicit nil

### UnsetAttachPkcs7Bundle
`func (o *EmailNotificationResponse) UnsetAttachPkcs7Bundle()`

UnsetAttachPkcs7Bundle ensures that no value is present for AttachPkcs7Bundle, not even an explicit nil
### GetAttachPkcs12

`func (o *EmailNotificationResponse) GetAttachPkcs12() bool`

GetAttachPkcs12 returns the AttachPkcs12 field if non-nil, zero value otherwise.

### GetAttachPkcs12Ok

`func (o *EmailNotificationResponse) GetAttachPkcs12Ok() (*bool, bool)`

GetAttachPkcs12Ok returns a tuple with the AttachPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs12

`func (o *EmailNotificationResponse) SetAttachPkcs12(v bool)`

SetAttachPkcs12 sets AttachPkcs12 field to given value.

### HasAttachPkcs12

`func (o *EmailNotificationResponse) HasAttachPkcs12() bool`

HasAttachPkcs12 returns a boolean if a field has been set.

### SetAttachPkcs12Nil

`func (o *EmailNotificationResponse) SetAttachPkcs12Nil(b bool)`

 SetAttachPkcs12Nil sets the value for AttachPkcs12 to be an explicit nil

### UnsetAttachPkcs12
`func (o *EmailNotificationResponse) UnsetAttachPkcs12()`

UnsetAttachPkcs12 ensures that no value is present for AttachPkcs12, not even an explicit nil
### GetName

`func (o *EmailNotificationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailNotificationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailNotificationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *EmailNotificationResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *EmailNotificationResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *EmailNotificationResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *EmailNotificationResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *EmailNotificationResponse) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *EmailNotificationResponse) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRunPeriod

`func (o *EmailNotificationResponse) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *EmailNotificationResponse) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *EmailNotificationResponse) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *EmailNotificationResponse) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *EmailNotificationResponse) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *EmailNotificationResponse) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
### GetLicenseUsagePercent

`func (o *EmailNotificationResponse) GetLicenseUsagePercent() int64`

GetLicenseUsagePercent returns the LicenseUsagePercent field if non-nil, zero value otherwise.

### GetLicenseUsagePercentOk

`func (o *EmailNotificationResponse) GetLicenseUsagePercentOk() (*int64, bool)`

GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUsagePercent

`func (o *EmailNotificationResponse) SetLicenseUsagePercent(v int64)`

SetLicenseUsagePercent sets LicenseUsagePercent field to given value.

### HasLicenseUsagePercent

`func (o *EmailNotificationResponse) HasLicenseUsagePercent() bool`

HasLicenseUsagePercent returns a boolean if a field has been set.

### SetLicenseUsagePercentNil

`func (o *EmailNotificationResponse) SetLicenseUsagePercentNil(b bool)`

 SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil

### UnsetLicenseUsagePercent
`func (o *EmailNotificationResponse) UnsetLicenseUsagePercent()`

UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
### GetEvents

`func (o *EmailNotificationResponse) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EmailNotificationResponse) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EmailNotificationResponse) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetRunOnRenewed

`func (o *EmailNotificationResponse) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *EmailNotificationResponse) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *EmailNotificationResponse) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *EmailNotificationResponse) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *EmailNotificationResponse) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *EmailNotificationResponse) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



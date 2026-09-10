# EmailNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachDerCertificate** | Pointer to **NullableBool** | Attach the certificate in DER format if available | [optional] 
**AttachPemBundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PEM format if available | [optional] 
**AttachPemCertificate** | Pointer to **NullableBool** | Attach the certificate in PEM format if available | [optional] 
**AttachPkcs7** | Pointer to **NullableBool** | Attach the certificate in PKCS7 format if available | [optional] 
**AttachPkcs7Bundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PKCS7 format if available | [optional] 
**AttachPkcs12** | Pointer to **NullableBool** | Attach the certificate in PKCS#12 format if available | [optional] 
**EmailTemplate** | [**EmailTemplate**](EmailTemplate.md) |  | 
**IfPkcs12** | Pointer to **NullableBool** | On events triggering an enrollment, select if mail is sent: - **Always**: set the value to &#x60;null&#x60;  - **Only when a PKCS#12 is available in the request**: set the value to &#x60;true&#x60;  - **Only when a PKCS#12 is not in the request**: set the value to &#x60;false&#x60;  | [optional] 
**Type** | **string** |  | 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**LicenceUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Name** | **string** | Name of the notification | 
**Retries** | Pointer to **NullableInt64** | Number of retries when the notification fails | [optional] 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 

## Methods

### NewEmailNotification

`func NewEmailNotification(emailTemplate EmailTemplate, type_ string, events []string, name string, ) *EmailNotification`

NewEmailNotification instantiates a new EmailNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationWithDefaults

`func NewEmailNotificationWithDefaults() *EmailNotification`

NewEmailNotificationWithDefaults instantiates a new EmailNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachDerCertificate

`func (o *EmailNotification) GetAttachDerCertificate() bool`

GetAttachDerCertificate returns the AttachDerCertificate field if non-nil, zero value otherwise.

### GetAttachDerCertificateOk

`func (o *EmailNotification) GetAttachDerCertificateOk() (*bool, bool)`

GetAttachDerCertificateOk returns a tuple with the AttachDerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDerCertificate

`func (o *EmailNotification) SetAttachDerCertificate(v bool)`

SetAttachDerCertificate sets AttachDerCertificate field to given value.

### HasAttachDerCertificate

`func (o *EmailNotification) HasAttachDerCertificate() bool`

HasAttachDerCertificate returns a boolean if a field has been set.

### SetAttachDerCertificateNil

`func (o *EmailNotification) SetAttachDerCertificateNil(b bool)`

 SetAttachDerCertificateNil sets the value for AttachDerCertificate to be an explicit nil

### UnsetAttachDerCertificate
`func (o *EmailNotification) UnsetAttachDerCertificate()`

UnsetAttachDerCertificate ensures that no value is present for AttachDerCertificate, not even an explicit nil
### GetAttachPemBundle

`func (o *EmailNotification) GetAttachPemBundle() bool`

GetAttachPemBundle returns the AttachPemBundle field if non-nil, zero value otherwise.

### GetAttachPemBundleOk

`func (o *EmailNotification) GetAttachPemBundleOk() (*bool, bool)`

GetAttachPemBundleOk returns a tuple with the AttachPemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemBundle

`func (o *EmailNotification) SetAttachPemBundle(v bool)`

SetAttachPemBundle sets AttachPemBundle field to given value.

### HasAttachPemBundle

`func (o *EmailNotification) HasAttachPemBundle() bool`

HasAttachPemBundle returns a boolean if a field has been set.

### SetAttachPemBundleNil

`func (o *EmailNotification) SetAttachPemBundleNil(b bool)`

 SetAttachPemBundleNil sets the value for AttachPemBundle to be an explicit nil

### UnsetAttachPemBundle
`func (o *EmailNotification) UnsetAttachPemBundle()`

UnsetAttachPemBundle ensures that no value is present for AttachPemBundle, not even an explicit nil
### GetAttachPemCertificate

`func (o *EmailNotification) GetAttachPemCertificate() bool`

GetAttachPemCertificate returns the AttachPemCertificate field if non-nil, zero value otherwise.

### GetAttachPemCertificateOk

`func (o *EmailNotification) GetAttachPemCertificateOk() (*bool, bool)`

GetAttachPemCertificateOk returns a tuple with the AttachPemCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemCertificate

`func (o *EmailNotification) SetAttachPemCertificate(v bool)`

SetAttachPemCertificate sets AttachPemCertificate field to given value.

### HasAttachPemCertificate

`func (o *EmailNotification) HasAttachPemCertificate() bool`

HasAttachPemCertificate returns a boolean if a field has been set.

### SetAttachPemCertificateNil

`func (o *EmailNotification) SetAttachPemCertificateNil(b bool)`

 SetAttachPemCertificateNil sets the value for AttachPemCertificate to be an explicit nil

### UnsetAttachPemCertificate
`func (o *EmailNotification) UnsetAttachPemCertificate()`

UnsetAttachPemCertificate ensures that no value is present for AttachPemCertificate, not even an explicit nil
### GetAttachPkcs7

`func (o *EmailNotification) GetAttachPkcs7() bool`

GetAttachPkcs7 returns the AttachPkcs7 field if non-nil, zero value otherwise.

### GetAttachPkcs7Ok

`func (o *EmailNotification) GetAttachPkcs7Ok() (*bool, bool)`

GetAttachPkcs7Ok returns a tuple with the AttachPkcs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7

`func (o *EmailNotification) SetAttachPkcs7(v bool)`

SetAttachPkcs7 sets AttachPkcs7 field to given value.

### HasAttachPkcs7

`func (o *EmailNotification) HasAttachPkcs7() bool`

HasAttachPkcs7 returns a boolean if a field has been set.

### SetAttachPkcs7Nil

`func (o *EmailNotification) SetAttachPkcs7Nil(b bool)`

 SetAttachPkcs7Nil sets the value for AttachPkcs7 to be an explicit nil

### UnsetAttachPkcs7
`func (o *EmailNotification) UnsetAttachPkcs7()`

UnsetAttachPkcs7 ensures that no value is present for AttachPkcs7, not even an explicit nil
### GetAttachPkcs7Bundle

`func (o *EmailNotification) GetAttachPkcs7Bundle() bool`

GetAttachPkcs7Bundle returns the AttachPkcs7Bundle field if non-nil, zero value otherwise.

### GetAttachPkcs7BundleOk

`func (o *EmailNotification) GetAttachPkcs7BundleOk() (*bool, bool)`

GetAttachPkcs7BundleOk returns a tuple with the AttachPkcs7Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7Bundle

`func (o *EmailNotification) SetAttachPkcs7Bundle(v bool)`

SetAttachPkcs7Bundle sets AttachPkcs7Bundle field to given value.

### HasAttachPkcs7Bundle

`func (o *EmailNotification) HasAttachPkcs7Bundle() bool`

HasAttachPkcs7Bundle returns a boolean if a field has been set.

### SetAttachPkcs7BundleNil

`func (o *EmailNotification) SetAttachPkcs7BundleNil(b bool)`

 SetAttachPkcs7BundleNil sets the value for AttachPkcs7Bundle to be an explicit nil

### UnsetAttachPkcs7Bundle
`func (o *EmailNotification) UnsetAttachPkcs7Bundle()`

UnsetAttachPkcs7Bundle ensures that no value is present for AttachPkcs7Bundle, not even an explicit nil
### GetAttachPkcs12

`func (o *EmailNotification) GetAttachPkcs12() bool`

GetAttachPkcs12 returns the AttachPkcs12 field if non-nil, zero value otherwise.

### GetAttachPkcs12Ok

`func (o *EmailNotification) GetAttachPkcs12Ok() (*bool, bool)`

GetAttachPkcs12Ok returns a tuple with the AttachPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs12

`func (o *EmailNotification) SetAttachPkcs12(v bool)`

SetAttachPkcs12 sets AttachPkcs12 field to given value.

### HasAttachPkcs12

`func (o *EmailNotification) HasAttachPkcs12() bool`

HasAttachPkcs12 returns a boolean if a field has been set.

### SetAttachPkcs12Nil

`func (o *EmailNotification) SetAttachPkcs12Nil(b bool)`

 SetAttachPkcs12Nil sets the value for AttachPkcs12 to be an explicit nil

### UnsetAttachPkcs12
`func (o *EmailNotification) UnsetAttachPkcs12()`

UnsetAttachPkcs12 ensures that no value is present for AttachPkcs12, not even an explicit nil
### GetEmailTemplate

`func (o *EmailNotification) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *EmailNotification) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *EmailNotification) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetIfPkcs12

`func (o *EmailNotification) GetIfPkcs12() bool`

GetIfPkcs12 returns the IfPkcs12 field if non-nil, zero value otherwise.

### GetIfPkcs12Ok

`func (o *EmailNotification) GetIfPkcs12Ok() (*bool, bool)`

GetIfPkcs12Ok returns a tuple with the IfPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfPkcs12

`func (o *EmailNotification) SetIfPkcs12(v bool)`

SetIfPkcs12 sets IfPkcs12 field to given value.

### HasIfPkcs12

`func (o *EmailNotification) HasIfPkcs12() bool`

HasIfPkcs12 returns a boolean if a field has been set.

### SetIfPkcs12Nil

`func (o *EmailNotification) SetIfPkcs12Nil(b bool)`

 SetIfPkcs12Nil sets the value for IfPkcs12 to be an explicit nil

### UnsetIfPkcs12
`func (o *EmailNotification) UnsetIfPkcs12()`

UnsetIfPkcs12 ensures that no value is present for IfPkcs12, not even an explicit nil
### GetType

`func (o *EmailNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailNotification) SetType(v string)`

SetType sets Type field to given value.


### GetEvents

`func (o *EmailNotification) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EmailNotification) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EmailNotification) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetLicenceUsagePercent

`func (o *EmailNotification) GetLicenceUsagePercent() int64`

GetLicenceUsagePercent returns the LicenceUsagePercent field if non-nil, zero value otherwise.

### GetLicenceUsagePercentOk

`func (o *EmailNotification) GetLicenceUsagePercentOk() (*int64, bool)`

GetLicenceUsagePercentOk returns a tuple with the LicenceUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceUsagePercent

`func (o *EmailNotification) SetLicenceUsagePercent(v int64)`

SetLicenceUsagePercent sets LicenceUsagePercent field to given value.

### HasLicenceUsagePercent

`func (o *EmailNotification) HasLicenceUsagePercent() bool`

HasLicenceUsagePercent returns a boolean if a field has been set.

### SetLicenceUsagePercentNil

`func (o *EmailNotification) SetLicenceUsagePercentNil(b bool)`

 SetLicenceUsagePercentNil sets the value for LicenceUsagePercent to be an explicit nil

### UnsetLicenceUsagePercent
`func (o *EmailNotification) UnsetLicenceUsagePercent()`

UnsetLicenceUsagePercent ensures that no value is present for LicenceUsagePercent, not even an explicit nil
### GetName

`func (o *EmailNotification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailNotification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailNotification) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *EmailNotification) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *EmailNotification) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *EmailNotification) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *EmailNotification) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *EmailNotification) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *EmailNotification) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRunOnRenewed

`func (o *EmailNotification) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *EmailNotification) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *EmailNotification) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *EmailNotification) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *EmailNotification) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *EmailNotification) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
### GetRunPeriod

`func (o *EmailNotification) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *EmailNotification) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *EmailNotification) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *EmailNotification) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *EmailNotification) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *EmailNotification) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



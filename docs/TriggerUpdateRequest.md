# TriggerUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of notification | 
**EmailTemplate** | [**EmailTemplate**](EmailTemplate.md) |  | 
**IfPkcs12** | Pointer to **NullableBool** | On events triggering an enrollment, select if mail is sent: - **Always**: set the value to &#x60;null&#x60;  - **Only when a PKCS#12 is available in the request**: set the value to &#x60;true&#x60;  - **Only when a PKCS#12 is not in the request**: set the value to &#x60;false&#x60;  | [optional] 
**AttachPemCertificate** | Pointer to **NullableBool** | Attach the certificate in PEM format if available | [optional] 
**AttachPemBundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PEM format if available | [optional] 
**AttachDerCertificate** | Pointer to **NullableBool** | Attach the certificate in DER format if available | [optional] 
**AttachPkcs7** | Pointer to **NullableBool** | Attach the certificate in PKCS7 format if available | [optional] 
**AttachPkcs7Bundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PKCS7 format if available | [optional] 
**AttachPkcs12** | Pointer to **NullableBool** | Attach the certificate in PKCS#12 format if available | [optional] 
**Name** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 
**LicenseUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 
**WebhookTemplate** | [**WebhookTemplate**](WebhookTemplate.md) |  | 
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while sending the webhook | [optional] 
**Timeout** | Pointer to **string** | Timeout for the webhook request | [optional] 
**Sequence** | **[]map[string]interface{}** | The REST requests to execute, in execution order. Each request enriches the dictionary with its response for the next one | 
**Connector** | **string** |  | 

## Methods

### NewTriggerUpdateRequest

`func NewTriggerUpdateRequest(type_ string, emailTemplate EmailTemplate, name string, events []string, webhookTemplate WebhookTemplate, sequence []map[string]interface{}, connector string, ) *TriggerUpdateRequest`

NewTriggerUpdateRequest instantiates a new TriggerUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerUpdateRequestWithDefaults

`func NewTriggerUpdateRequestWithDefaults() *TriggerUpdateRequest`

NewTriggerUpdateRequestWithDefaults instantiates a new TriggerUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TriggerUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerUpdateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetEmailTemplate

`func (o *TriggerUpdateRequest) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *TriggerUpdateRequest) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *TriggerUpdateRequest) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetIfPkcs12

`func (o *TriggerUpdateRequest) GetIfPkcs12() bool`

GetIfPkcs12 returns the IfPkcs12 field if non-nil, zero value otherwise.

### GetIfPkcs12Ok

`func (o *TriggerUpdateRequest) GetIfPkcs12Ok() (*bool, bool)`

GetIfPkcs12Ok returns a tuple with the IfPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfPkcs12

`func (o *TriggerUpdateRequest) SetIfPkcs12(v bool)`

SetIfPkcs12 sets IfPkcs12 field to given value.

### HasIfPkcs12

`func (o *TriggerUpdateRequest) HasIfPkcs12() bool`

HasIfPkcs12 returns a boolean if a field has been set.

### SetIfPkcs12Nil

`func (o *TriggerUpdateRequest) SetIfPkcs12Nil(b bool)`

 SetIfPkcs12Nil sets the value for IfPkcs12 to be an explicit nil

### UnsetIfPkcs12
`func (o *TriggerUpdateRequest) UnsetIfPkcs12()`

UnsetIfPkcs12 ensures that no value is present for IfPkcs12, not even an explicit nil
### GetAttachPemCertificate

`func (o *TriggerUpdateRequest) GetAttachPemCertificate() bool`

GetAttachPemCertificate returns the AttachPemCertificate field if non-nil, zero value otherwise.

### GetAttachPemCertificateOk

`func (o *TriggerUpdateRequest) GetAttachPemCertificateOk() (*bool, bool)`

GetAttachPemCertificateOk returns a tuple with the AttachPemCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemCertificate

`func (o *TriggerUpdateRequest) SetAttachPemCertificate(v bool)`

SetAttachPemCertificate sets AttachPemCertificate field to given value.

### HasAttachPemCertificate

`func (o *TriggerUpdateRequest) HasAttachPemCertificate() bool`

HasAttachPemCertificate returns a boolean if a field has been set.

### SetAttachPemCertificateNil

`func (o *TriggerUpdateRequest) SetAttachPemCertificateNil(b bool)`

 SetAttachPemCertificateNil sets the value for AttachPemCertificate to be an explicit nil

### UnsetAttachPemCertificate
`func (o *TriggerUpdateRequest) UnsetAttachPemCertificate()`

UnsetAttachPemCertificate ensures that no value is present for AttachPemCertificate, not even an explicit nil
### GetAttachPemBundle

`func (o *TriggerUpdateRequest) GetAttachPemBundle() bool`

GetAttachPemBundle returns the AttachPemBundle field if non-nil, zero value otherwise.

### GetAttachPemBundleOk

`func (o *TriggerUpdateRequest) GetAttachPemBundleOk() (*bool, bool)`

GetAttachPemBundleOk returns a tuple with the AttachPemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemBundle

`func (o *TriggerUpdateRequest) SetAttachPemBundle(v bool)`

SetAttachPemBundle sets AttachPemBundle field to given value.

### HasAttachPemBundle

`func (o *TriggerUpdateRequest) HasAttachPemBundle() bool`

HasAttachPemBundle returns a boolean if a field has been set.

### SetAttachPemBundleNil

`func (o *TriggerUpdateRequest) SetAttachPemBundleNil(b bool)`

 SetAttachPemBundleNil sets the value for AttachPemBundle to be an explicit nil

### UnsetAttachPemBundle
`func (o *TriggerUpdateRequest) UnsetAttachPemBundle()`

UnsetAttachPemBundle ensures that no value is present for AttachPemBundle, not even an explicit nil
### GetAttachDerCertificate

`func (o *TriggerUpdateRequest) GetAttachDerCertificate() bool`

GetAttachDerCertificate returns the AttachDerCertificate field if non-nil, zero value otherwise.

### GetAttachDerCertificateOk

`func (o *TriggerUpdateRequest) GetAttachDerCertificateOk() (*bool, bool)`

GetAttachDerCertificateOk returns a tuple with the AttachDerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDerCertificate

`func (o *TriggerUpdateRequest) SetAttachDerCertificate(v bool)`

SetAttachDerCertificate sets AttachDerCertificate field to given value.

### HasAttachDerCertificate

`func (o *TriggerUpdateRequest) HasAttachDerCertificate() bool`

HasAttachDerCertificate returns a boolean if a field has been set.

### SetAttachDerCertificateNil

`func (o *TriggerUpdateRequest) SetAttachDerCertificateNil(b bool)`

 SetAttachDerCertificateNil sets the value for AttachDerCertificate to be an explicit nil

### UnsetAttachDerCertificate
`func (o *TriggerUpdateRequest) UnsetAttachDerCertificate()`

UnsetAttachDerCertificate ensures that no value is present for AttachDerCertificate, not even an explicit nil
### GetAttachPkcs7

`func (o *TriggerUpdateRequest) GetAttachPkcs7() bool`

GetAttachPkcs7 returns the AttachPkcs7 field if non-nil, zero value otherwise.

### GetAttachPkcs7Ok

`func (o *TriggerUpdateRequest) GetAttachPkcs7Ok() (*bool, bool)`

GetAttachPkcs7Ok returns a tuple with the AttachPkcs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7

`func (o *TriggerUpdateRequest) SetAttachPkcs7(v bool)`

SetAttachPkcs7 sets AttachPkcs7 field to given value.

### HasAttachPkcs7

`func (o *TriggerUpdateRequest) HasAttachPkcs7() bool`

HasAttachPkcs7 returns a boolean if a field has been set.

### SetAttachPkcs7Nil

`func (o *TriggerUpdateRequest) SetAttachPkcs7Nil(b bool)`

 SetAttachPkcs7Nil sets the value for AttachPkcs7 to be an explicit nil

### UnsetAttachPkcs7
`func (o *TriggerUpdateRequest) UnsetAttachPkcs7()`

UnsetAttachPkcs7 ensures that no value is present for AttachPkcs7, not even an explicit nil
### GetAttachPkcs7Bundle

`func (o *TriggerUpdateRequest) GetAttachPkcs7Bundle() bool`

GetAttachPkcs7Bundle returns the AttachPkcs7Bundle field if non-nil, zero value otherwise.

### GetAttachPkcs7BundleOk

`func (o *TriggerUpdateRequest) GetAttachPkcs7BundleOk() (*bool, bool)`

GetAttachPkcs7BundleOk returns a tuple with the AttachPkcs7Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7Bundle

`func (o *TriggerUpdateRequest) SetAttachPkcs7Bundle(v bool)`

SetAttachPkcs7Bundle sets AttachPkcs7Bundle field to given value.

### HasAttachPkcs7Bundle

`func (o *TriggerUpdateRequest) HasAttachPkcs7Bundle() bool`

HasAttachPkcs7Bundle returns a boolean if a field has been set.

### SetAttachPkcs7BundleNil

`func (o *TriggerUpdateRequest) SetAttachPkcs7BundleNil(b bool)`

 SetAttachPkcs7BundleNil sets the value for AttachPkcs7Bundle to be an explicit nil

### UnsetAttachPkcs7Bundle
`func (o *TriggerUpdateRequest) UnsetAttachPkcs7Bundle()`

UnsetAttachPkcs7Bundle ensures that no value is present for AttachPkcs7Bundle, not even an explicit nil
### GetAttachPkcs12

`func (o *TriggerUpdateRequest) GetAttachPkcs12() bool`

GetAttachPkcs12 returns the AttachPkcs12 field if non-nil, zero value otherwise.

### GetAttachPkcs12Ok

`func (o *TriggerUpdateRequest) GetAttachPkcs12Ok() (*bool, bool)`

GetAttachPkcs12Ok returns a tuple with the AttachPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs12

`func (o *TriggerUpdateRequest) SetAttachPkcs12(v bool)`

SetAttachPkcs12 sets AttachPkcs12 field to given value.

### HasAttachPkcs12

`func (o *TriggerUpdateRequest) HasAttachPkcs12() bool`

HasAttachPkcs12 returns a boolean if a field has been set.

### SetAttachPkcs12Nil

`func (o *TriggerUpdateRequest) SetAttachPkcs12Nil(b bool)`

 SetAttachPkcs12Nil sets the value for AttachPkcs12 to be an explicit nil

### UnsetAttachPkcs12
`func (o *TriggerUpdateRequest) UnsetAttachPkcs12()`

UnsetAttachPkcs12 ensures that no value is present for AttachPkcs12, not even an explicit nil
### GetName

`func (o *TriggerUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *TriggerUpdateRequest) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *TriggerUpdateRequest) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *TriggerUpdateRequest) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *TriggerUpdateRequest) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *TriggerUpdateRequest) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *TriggerUpdateRequest) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRunPeriod

`func (o *TriggerUpdateRequest) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *TriggerUpdateRequest) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *TriggerUpdateRequest) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *TriggerUpdateRequest) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *TriggerUpdateRequest) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *TriggerUpdateRequest) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
### GetLicenseUsagePercent

`func (o *TriggerUpdateRequest) GetLicenseUsagePercent() int64`

GetLicenseUsagePercent returns the LicenseUsagePercent field if non-nil, zero value otherwise.

### GetLicenseUsagePercentOk

`func (o *TriggerUpdateRequest) GetLicenseUsagePercentOk() (*int64, bool)`

GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUsagePercent

`func (o *TriggerUpdateRequest) SetLicenseUsagePercent(v int64)`

SetLicenseUsagePercent sets LicenseUsagePercent field to given value.

### HasLicenseUsagePercent

`func (o *TriggerUpdateRequest) HasLicenseUsagePercent() bool`

HasLicenseUsagePercent returns a boolean if a field has been set.

### SetLicenseUsagePercentNil

`func (o *TriggerUpdateRequest) SetLicenseUsagePercentNil(b bool)`

 SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil

### UnsetLicenseUsagePercent
`func (o *TriggerUpdateRequest) UnsetLicenseUsagePercent()`

UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
### GetEvents

`func (o *TriggerUpdateRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TriggerUpdateRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TriggerUpdateRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetRunOnRenewed

`func (o *TriggerUpdateRequest) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *TriggerUpdateRequest) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *TriggerUpdateRequest) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *TriggerUpdateRequest) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *TriggerUpdateRequest) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *TriggerUpdateRequest) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
### GetWebhookTemplate

`func (o *TriggerUpdateRequest) GetWebhookTemplate() WebhookTemplate`

GetWebhookTemplate returns the WebhookTemplate field if non-nil, zero value otherwise.

### GetWebhookTemplateOk

`func (o *TriggerUpdateRequest) GetWebhookTemplateOk() (*WebhookTemplate, bool)`

GetWebhookTemplateOk returns a tuple with the WebhookTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTemplate

`func (o *TriggerUpdateRequest) SetWebhookTemplate(v WebhookTemplate)`

SetWebhookTemplate sets WebhookTemplate field to given value.


### GetProxy

`func (o *TriggerUpdateRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *TriggerUpdateRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *TriggerUpdateRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *TriggerUpdateRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *TriggerUpdateRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *TriggerUpdateRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *TriggerUpdateRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TriggerUpdateRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TriggerUpdateRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TriggerUpdateRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetSequence

`func (o *TriggerUpdateRequest) GetSequence() []map[string]interface{}`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *TriggerUpdateRequest) GetSequenceOk() (*[]map[string]interface{}, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *TriggerUpdateRequest) SetSequence(v []map[string]interface{})`

SetSequence sets Sequence field to given value.


### GetConnector

`func (o *TriggerUpdateRequest) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *TriggerUpdateRequest) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *TriggerUpdateRequest) SetConnector(v string)`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



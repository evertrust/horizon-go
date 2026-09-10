# TriggerList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AttachDerCertificate** | Pointer to **NullableBool** | Attach the certificate in DER format if available | [optional] 
**AttachPemBundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PEM format if available | [optional] 
**AttachPemCertificate** | Pointer to **NullableBool** | Attach the certificate in PEM format if available | [optional] 
**AttachPkcs7** | Pointer to **NullableBool** | Attach the certificate in PKCS7 format if available | [optional] 
**AttachPkcs7Bundle** | Pointer to **NullableBool** | Attach the certificate and its trust chain (bundle) in PKCS7 format if available | [optional] 
**AttachPkcs12** | Pointer to **NullableBool** | Attach the certificate in PKCS#12 format if available | [optional] 
**EmailTemplate** | [**EmailTemplate**](EmailTemplate.md) |  | 
**IfPkcs12** | Pointer to **NullableBool** | On events triggering an enrollment, select if mail is sent: - **Always**: set the value to &#x60;null&#x60;  - **Only when a PKCS#12 is available in the request**: set the value to &#x60;true&#x60;  - **Only when a PKCS#12 is not in the request**: set the value to &#x60;false&#x60;  | [optional] 
**Type** | **string** | The type of notification | 
**Events** | **[]string** | Event on which the notification runs. This MUST contain only one value. | 
**LicenceUsagePercent** | Pointer to **NullableInt64** | License usage at which the notification needs to run (between 0 and 100). Must be defined on &#x60;on_license_usage&#x60; event and must NOT be defined otherwise. | [optional] 
**Name** | **string** |  | 
**Retries** | Pointer to **NullableInt64** |  | [optional] 
**RunOnRenewed** | Pointer to **NullableBool** | Must be defined on &#x60;on_expire&#x60; event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed. | [optional] 
**RunPeriod** | Pointer to **NullableString** | Time period at which the notification needs to run. Can only be defined on expiration and pending events. | [optional] 
**Proxy** | Pointer to **NullableString** | Name of a Proxy to use while sending the webhook | [optional] 
**Timeout** | Pointer to **string** | Timeout for the webhook request | [optional] 
**WebhookTemplate** | [**WebhookTemplate**](WebhookTemplate.md) |  | 
**Sequence** | **[]map[string]interface{}** | The REST requests to execute, in execution order. Each request enriches the dictionary with its response for the next one | 
**Connector** | **string** |  | 

## Methods

### NewTriggerList200ResponseInner

`func NewTriggerList200ResponseInner(id string, emailTemplate EmailTemplate, type_ string, events []string, name string, webhookTemplate WebhookTemplate, sequence []map[string]interface{}, connector string, ) *TriggerList200ResponseInner`

NewTriggerList200ResponseInner instantiates a new TriggerList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerList200ResponseInnerWithDefaults

`func NewTriggerList200ResponseInnerWithDefaults() *TriggerList200ResponseInner`

NewTriggerList200ResponseInnerWithDefaults instantiates a new TriggerList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttachDerCertificate

`func (o *TriggerList200ResponseInner) GetAttachDerCertificate() bool`

GetAttachDerCertificate returns the AttachDerCertificate field if non-nil, zero value otherwise.

### GetAttachDerCertificateOk

`func (o *TriggerList200ResponseInner) GetAttachDerCertificateOk() (*bool, bool)`

GetAttachDerCertificateOk returns a tuple with the AttachDerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDerCertificate

`func (o *TriggerList200ResponseInner) SetAttachDerCertificate(v bool)`

SetAttachDerCertificate sets AttachDerCertificate field to given value.

### HasAttachDerCertificate

`func (o *TriggerList200ResponseInner) HasAttachDerCertificate() bool`

HasAttachDerCertificate returns a boolean if a field has been set.

### SetAttachDerCertificateNil

`func (o *TriggerList200ResponseInner) SetAttachDerCertificateNil(b bool)`

 SetAttachDerCertificateNil sets the value for AttachDerCertificate to be an explicit nil

### UnsetAttachDerCertificate
`func (o *TriggerList200ResponseInner) UnsetAttachDerCertificate()`

UnsetAttachDerCertificate ensures that no value is present for AttachDerCertificate, not even an explicit nil
### GetAttachPemBundle

`func (o *TriggerList200ResponseInner) GetAttachPemBundle() bool`

GetAttachPemBundle returns the AttachPemBundle field if non-nil, zero value otherwise.

### GetAttachPemBundleOk

`func (o *TriggerList200ResponseInner) GetAttachPemBundleOk() (*bool, bool)`

GetAttachPemBundleOk returns a tuple with the AttachPemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemBundle

`func (o *TriggerList200ResponseInner) SetAttachPemBundle(v bool)`

SetAttachPemBundle sets AttachPemBundle field to given value.

### HasAttachPemBundle

`func (o *TriggerList200ResponseInner) HasAttachPemBundle() bool`

HasAttachPemBundle returns a boolean if a field has been set.

### SetAttachPemBundleNil

`func (o *TriggerList200ResponseInner) SetAttachPemBundleNil(b bool)`

 SetAttachPemBundleNil sets the value for AttachPemBundle to be an explicit nil

### UnsetAttachPemBundle
`func (o *TriggerList200ResponseInner) UnsetAttachPemBundle()`

UnsetAttachPemBundle ensures that no value is present for AttachPemBundle, not even an explicit nil
### GetAttachPemCertificate

`func (o *TriggerList200ResponseInner) GetAttachPemCertificate() bool`

GetAttachPemCertificate returns the AttachPemCertificate field if non-nil, zero value otherwise.

### GetAttachPemCertificateOk

`func (o *TriggerList200ResponseInner) GetAttachPemCertificateOk() (*bool, bool)`

GetAttachPemCertificateOk returns a tuple with the AttachPemCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPemCertificate

`func (o *TriggerList200ResponseInner) SetAttachPemCertificate(v bool)`

SetAttachPemCertificate sets AttachPemCertificate field to given value.

### HasAttachPemCertificate

`func (o *TriggerList200ResponseInner) HasAttachPemCertificate() bool`

HasAttachPemCertificate returns a boolean if a field has been set.

### SetAttachPemCertificateNil

`func (o *TriggerList200ResponseInner) SetAttachPemCertificateNil(b bool)`

 SetAttachPemCertificateNil sets the value for AttachPemCertificate to be an explicit nil

### UnsetAttachPemCertificate
`func (o *TriggerList200ResponseInner) UnsetAttachPemCertificate()`

UnsetAttachPemCertificate ensures that no value is present for AttachPemCertificate, not even an explicit nil
### GetAttachPkcs7

`func (o *TriggerList200ResponseInner) GetAttachPkcs7() bool`

GetAttachPkcs7 returns the AttachPkcs7 field if non-nil, zero value otherwise.

### GetAttachPkcs7Ok

`func (o *TriggerList200ResponseInner) GetAttachPkcs7Ok() (*bool, bool)`

GetAttachPkcs7Ok returns a tuple with the AttachPkcs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7

`func (o *TriggerList200ResponseInner) SetAttachPkcs7(v bool)`

SetAttachPkcs7 sets AttachPkcs7 field to given value.

### HasAttachPkcs7

`func (o *TriggerList200ResponseInner) HasAttachPkcs7() bool`

HasAttachPkcs7 returns a boolean if a field has been set.

### SetAttachPkcs7Nil

`func (o *TriggerList200ResponseInner) SetAttachPkcs7Nil(b bool)`

 SetAttachPkcs7Nil sets the value for AttachPkcs7 to be an explicit nil

### UnsetAttachPkcs7
`func (o *TriggerList200ResponseInner) UnsetAttachPkcs7()`

UnsetAttachPkcs7 ensures that no value is present for AttachPkcs7, not even an explicit nil
### GetAttachPkcs7Bundle

`func (o *TriggerList200ResponseInner) GetAttachPkcs7Bundle() bool`

GetAttachPkcs7Bundle returns the AttachPkcs7Bundle field if non-nil, zero value otherwise.

### GetAttachPkcs7BundleOk

`func (o *TriggerList200ResponseInner) GetAttachPkcs7BundleOk() (*bool, bool)`

GetAttachPkcs7BundleOk returns a tuple with the AttachPkcs7Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs7Bundle

`func (o *TriggerList200ResponseInner) SetAttachPkcs7Bundle(v bool)`

SetAttachPkcs7Bundle sets AttachPkcs7Bundle field to given value.

### HasAttachPkcs7Bundle

`func (o *TriggerList200ResponseInner) HasAttachPkcs7Bundle() bool`

HasAttachPkcs7Bundle returns a boolean if a field has been set.

### SetAttachPkcs7BundleNil

`func (o *TriggerList200ResponseInner) SetAttachPkcs7BundleNil(b bool)`

 SetAttachPkcs7BundleNil sets the value for AttachPkcs7Bundle to be an explicit nil

### UnsetAttachPkcs7Bundle
`func (o *TriggerList200ResponseInner) UnsetAttachPkcs7Bundle()`

UnsetAttachPkcs7Bundle ensures that no value is present for AttachPkcs7Bundle, not even an explicit nil
### GetAttachPkcs12

`func (o *TriggerList200ResponseInner) GetAttachPkcs12() bool`

GetAttachPkcs12 returns the AttachPkcs12 field if non-nil, zero value otherwise.

### GetAttachPkcs12Ok

`func (o *TriggerList200ResponseInner) GetAttachPkcs12Ok() (*bool, bool)`

GetAttachPkcs12Ok returns a tuple with the AttachPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPkcs12

`func (o *TriggerList200ResponseInner) SetAttachPkcs12(v bool)`

SetAttachPkcs12 sets AttachPkcs12 field to given value.

### HasAttachPkcs12

`func (o *TriggerList200ResponseInner) HasAttachPkcs12() bool`

HasAttachPkcs12 returns a boolean if a field has been set.

### SetAttachPkcs12Nil

`func (o *TriggerList200ResponseInner) SetAttachPkcs12Nil(b bool)`

 SetAttachPkcs12Nil sets the value for AttachPkcs12 to be an explicit nil

### UnsetAttachPkcs12
`func (o *TriggerList200ResponseInner) UnsetAttachPkcs12()`

UnsetAttachPkcs12 ensures that no value is present for AttachPkcs12, not even an explicit nil
### GetEmailTemplate

`func (o *TriggerList200ResponseInner) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *TriggerList200ResponseInner) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *TriggerList200ResponseInner) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetIfPkcs12

`func (o *TriggerList200ResponseInner) GetIfPkcs12() bool`

GetIfPkcs12 returns the IfPkcs12 field if non-nil, zero value otherwise.

### GetIfPkcs12Ok

`func (o *TriggerList200ResponseInner) GetIfPkcs12Ok() (*bool, bool)`

GetIfPkcs12Ok returns a tuple with the IfPkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfPkcs12

`func (o *TriggerList200ResponseInner) SetIfPkcs12(v bool)`

SetIfPkcs12 sets IfPkcs12 field to given value.

### HasIfPkcs12

`func (o *TriggerList200ResponseInner) HasIfPkcs12() bool`

HasIfPkcs12 returns a boolean if a field has been set.

### SetIfPkcs12Nil

`func (o *TriggerList200ResponseInner) SetIfPkcs12Nil(b bool)`

 SetIfPkcs12Nil sets the value for IfPkcs12 to be an explicit nil

### UnsetIfPkcs12
`func (o *TriggerList200ResponseInner) UnsetIfPkcs12()`

UnsetIfPkcs12 ensures that no value is present for IfPkcs12, not even an explicit nil
### GetType

`func (o *TriggerList200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerList200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerList200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetEvents

`func (o *TriggerList200ResponseInner) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TriggerList200ResponseInner) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TriggerList200ResponseInner) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetLicenceUsagePercent

`func (o *TriggerList200ResponseInner) GetLicenceUsagePercent() int64`

GetLicenceUsagePercent returns the LicenceUsagePercent field if non-nil, zero value otherwise.

### GetLicenceUsagePercentOk

`func (o *TriggerList200ResponseInner) GetLicenceUsagePercentOk() (*int64, bool)`

GetLicenceUsagePercentOk returns a tuple with the LicenceUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceUsagePercent

`func (o *TriggerList200ResponseInner) SetLicenceUsagePercent(v int64)`

SetLicenceUsagePercent sets LicenceUsagePercent field to given value.

### HasLicenceUsagePercent

`func (o *TriggerList200ResponseInner) HasLicenceUsagePercent() bool`

HasLicenceUsagePercent returns a boolean if a field has been set.

### SetLicenceUsagePercentNil

`func (o *TriggerList200ResponseInner) SetLicenceUsagePercentNil(b bool)`

 SetLicenceUsagePercentNil sets the value for LicenceUsagePercent to be an explicit nil

### UnsetLicenceUsagePercent
`func (o *TriggerList200ResponseInner) UnsetLicenceUsagePercent()`

UnsetLicenceUsagePercent ensures that no value is present for LicenceUsagePercent, not even an explicit nil
### GetName

`func (o *TriggerList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetRetries

`func (o *TriggerList200ResponseInner) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *TriggerList200ResponseInner) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *TriggerList200ResponseInner) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *TriggerList200ResponseInner) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *TriggerList200ResponseInner) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *TriggerList200ResponseInner) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRunOnRenewed

`func (o *TriggerList200ResponseInner) GetRunOnRenewed() bool`

GetRunOnRenewed returns the RunOnRenewed field if non-nil, zero value otherwise.

### GetRunOnRenewedOk

`func (o *TriggerList200ResponseInner) GetRunOnRenewedOk() (*bool, bool)`

GetRunOnRenewedOk returns a tuple with the RunOnRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnRenewed

`func (o *TriggerList200ResponseInner) SetRunOnRenewed(v bool)`

SetRunOnRenewed sets RunOnRenewed field to given value.

### HasRunOnRenewed

`func (o *TriggerList200ResponseInner) HasRunOnRenewed() bool`

HasRunOnRenewed returns a boolean if a field has been set.

### SetRunOnRenewedNil

`func (o *TriggerList200ResponseInner) SetRunOnRenewedNil(b bool)`

 SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil

### UnsetRunOnRenewed
`func (o *TriggerList200ResponseInner) UnsetRunOnRenewed()`

UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
### GetRunPeriod

`func (o *TriggerList200ResponseInner) GetRunPeriod() string`

GetRunPeriod returns the RunPeriod field if non-nil, zero value otherwise.

### GetRunPeriodOk

`func (o *TriggerList200ResponseInner) GetRunPeriodOk() (*string, bool)`

GetRunPeriodOk returns a tuple with the RunPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPeriod

`func (o *TriggerList200ResponseInner) SetRunPeriod(v string)`

SetRunPeriod sets RunPeriod field to given value.

### HasRunPeriod

`func (o *TriggerList200ResponseInner) HasRunPeriod() bool`

HasRunPeriod returns a boolean if a field has been set.

### SetRunPeriodNil

`func (o *TriggerList200ResponseInner) SetRunPeriodNil(b bool)`

 SetRunPeriodNil sets the value for RunPeriod to be an explicit nil

### UnsetRunPeriod
`func (o *TriggerList200ResponseInner) UnsetRunPeriod()`

UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
### GetProxy

`func (o *TriggerList200ResponseInner) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *TriggerList200ResponseInner) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *TriggerList200ResponseInner) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *TriggerList200ResponseInner) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *TriggerList200ResponseInner) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *TriggerList200ResponseInner) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetTimeout

`func (o *TriggerList200ResponseInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TriggerList200ResponseInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TriggerList200ResponseInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TriggerList200ResponseInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetWebhookTemplate

`func (o *TriggerList200ResponseInner) GetWebhookTemplate() WebhookTemplate`

GetWebhookTemplate returns the WebhookTemplate field if non-nil, zero value otherwise.

### GetWebhookTemplateOk

`func (o *TriggerList200ResponseInner) GetWebhookTemplateOk() (*WebhookTemplate, bool)`

GetWebhookTemplateOk returns a tuple with the WebhookTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTemplate

`func (o *TriggerList200ResponseInner) SetWebhookTemplate(v WebhookTemplate)`

SetWebhookTemplate sets WebhookTemplate field to given value.


### GetSequence

`func (o *TriggerList200ResponseInner) GetSequence() []map[string]interface{}`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *TriggerList200ResponseInner) GetSequenceOk() (*[]map[string]interface{}, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *TriggerList200ResponseInner) SetSequence(v []map[string]interface{})`

SetSequence sets Sequence field to given value.


### GetConnector

`func (o *TriggerList200ResponseInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *TriggerList200ResponseInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *TriggerList200ResponseInner) SetConnector(v string)`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



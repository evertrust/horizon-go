# TenantWithAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**BootstrapAccount**](BootstrapAccount.md) |  | 
**Tenant** | [**TenantResponse**](TenantResponse.md) |  | 

## Methods

### NewTenantWithAccountResponse

`func NewTenantWithAccountResponse(account BootstrapAccount, tenant TenantResponse, ) *TenantWithAccountResponse`

NewTenantWithAccountResponse instantiates a new TenantWithAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithAccountResponseWithDefaults

`func NewTenantWithAccountResponseWithDefaults() *TenantWithAccountResponse`

NewTenantWithAccountResponseWithDefaults instantiates a new TenantWithAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TenantWithAccountResponse) GetAccount() BootstrapAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TenantWithAccountResponse) GetAccountOk() (*BootstrapAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TenantWithAccountResponse) SetAccount(v BootstrapAccount)`

SetAccount sets Account field to given value.


### GetTenant

`func (o *TenantWithAccountResponse) GetTenant() TenantResponse`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TenantWithAccountResponse) GetTenantOk() (*TenantResponse, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TenantWithAccountResponse) SetTenant(v TenantResponse)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



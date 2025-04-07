# EntraIDConnectionAttributesConnectionTimeoutConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryWait** | Pointer to **int32** |  | [optional] 
**TokenRefreshMaxTryCount** | Pointer to **int32** |  | [optional] 
**RetryFailureStatusCode** | Pointer to **string** |  | [optional] 
**RetryWaitMaxValue** | Pointer to **int32** |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**ReadTimeout** | Pointer to **int32** |  | [optional] 
**ConnectionTimeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewEntraIDConnectionAttributesConnectionTimeoutConfig

`func NewEntraIDConnectionAttributesConnectionTimeoutConfig() *EntraIDConnectionAttributesConnectionTimeoutConfig`

NewEntraIDConnectionAttributesConnectionTimeoutConfig instantiates a new EntraIDConnectionAttributesConnectionTimeoutConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntraIDConnectionAttributesConnectionTimeoutConfigWithDefaults

`func NewEntraIDConnectionAttributesConnectionTimeoutConfigWithDefaults() *EntraIDConnectionAttributesConnectionTimeoutConfig`

NewEntraIDConnectionAttributesConnectionTimeoutConfigWithDefaults instantiates a new EntraIDConnectionAttributesConnectionTimeoutConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryWait

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryWait() int32`

GetRetryWait returns the RetryWait field if non-nil, zero value otherwise.

### GetRetryWaitOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryWaitOk() (*int32, bool)`

GetRetryWaitOk returns a tuple with the RetryWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWait

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetRetryWait(v int32)`

SetRetryWait sets RetryWait field to given value.

### HasRetryWait

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasRetryWait() bool`

HasRetryWait returns a boolean if a field has been set.

### GetTokenRefreshMaxTryCount

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetTokenRefreshMaxTryCount() int32`

GetTokenRefreshMaxTryCount returns the TokenRefreshMaxTryCount field if non-nil, zero value otherwise.

### GetTokenRefreshMaxTryCountOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetTokenRefreshMaxTryCountOk() (*int32, bool)`

GetTokenRefreshMaxTryCountOk returns a tuple with the TokenRefreshMaxTryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRefreshMaxTryCount

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetTokenRefreshMaxTryCount(v int32)`

SetTokenRefreshMaxTryCount sets TokenRefreshMaxTryCount field to given value.

### HasTokenRefreshMaxTryCount

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasTokenRefreshMaxTryCount() bool`

HasTokenRefreshMaxTryCount returns a boolean if a field has been set.

### GetRetryFailureStatusCode

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryFailureStatusCode() string`

GetRetryFailureStatusCode returns the RetryFailureStatusCode field if non-nil, zero value otherwise.

### GetRetryFailureStatusCodeOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryFailureStatusCodeOk() (*string, bool)`

GetRetryFailureStatusCodeOk returns a tuple with the RetryFailureStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailureStatusCode

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetRetryFailureStatusCode(v string)`

SetRetryFailureStatusCode sets RetryFailureStatusCode field to given value.

### HasRetryFailureStatusCode

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasRetryFailureStatusCode() bool`

HasRetryFailureStatusCode returns a boolean if a field has been set.

### GetRetryWaitMaxValue

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryWaitMaxValue() int32`

GetRetryWaitMaxValue returns the RetryWaitMaxValue field if non-nil, zero value otherwise.

### GetRetryWaitMaxValueOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryWaitMaxValueOk() (*int32, bool)`

GetRetryWaitMaxValueOk returns a tuple with the RetryWaitMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWaitMaxValue

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetRetryWaitMaxValue(v int32)`

SetRetryWaitMaxValue sets RetryWaitMaxValue field to given value.

### HasRetryWaitMaxValue

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasRetryWaitMaxValue() bool`

HasRetryWaitMaxValue returns a boolean if a field has been set.

### GetRetryCount

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetReadTimeout

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *EntraIDConnectionAttributesConnectionTimeoutConfig) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UploadSchemaFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | the file to upload | 
**PathLocation** | **string** | Should be set to &#x60;Datafiles&#x60; to upload to &#x60;job.ecm.imp.file.path&#x60; in &#x60;InternalConfig.groovy&#x60;, or &#x60;SAV&#x60; to upload to &#x60;job.ecm.savfile.path&#x60; in &#x60;InternalConfig.groovy&#x60;.  | 

## Methods

### NewUploadSchemaFileRequest

`func NewUploadSchemaFileRequest(file *os.File, pathLocation string, ) *UploadSchemaFileRequest`

NewUploadSchemaFileRequest instantiates a new UploadSchemaFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSchemaFileRequestWithDefaults

`func NewUploadSchemaFileRequestWithDefaults() *UploadSchemaFileRequest`

NewUploadSchemaFileRequestWithDefaults instantiates a new UploadSchemaFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *UploadSchemaFileRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadSchemaFileRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadSchemaFileRequest) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetPathLocation

`func (o *UploadSchemaFileRequest) GetPathLocation() string`

GetPathLocation returns the PathLocation field if non-nil, zero value otherwise.

### GetPathLocationOk

`func (o *UploadSchemaFileRequest) GetPathLocationOk() (*string, bool)`

GetPathLocationOk returns a tuple with the PathLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathLocation

`func (o *UploadSchemaFileRequest) SetPathLocation(v string)`

SetPathLocation sets PathLocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



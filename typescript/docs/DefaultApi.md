# DefaultApi

All URIs are relative to *https://api.buena.ai/api/v2*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createApiKey**](#createapikey) | **POST** /keys | Create API Key|
|[**createLead**](#createlead) | **POST** /leads | Create Lead|
|[**createVoiceClone**](#createvoiceclone) | **POST** /voice-clones | Create Voice Clone|
|[**deleteVoiceClone**](#deletevoiceclone) | **DELETE** /voice-clones/{voiceId} | Delete Voice Clone|
|[**generateVoicePreview**](#generatevoicepreview) | **POST** /voice-clones/preview | Generate Voice Preview|
|[**healthCheck**](#healthcheck) | **GET** /health | Health Check|
|[**listApiKeys**](#listapikeys) | **GET** /keys | List API Keys|
|[**listLeads**](#listleads) | **GET** /leads | List Leads|
|[**listVoiceClones**](#listvoiceclones) | **GET** /voice-clones | List Voice Clones|
|[**updateVoiceClone**](#updatevoiceclone) | **PUT** /voice-clones/{voiceId} | Update Voice Clone|

# **createApiKey**
> ApiKey createApiKey(createApiKeyRequest)

Create a new API key

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    CreateApiKeyRequest
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let createApiKeyRequest: CreateApiKeyRequest; //

const { status, data } = await apiInstance.createApiKey(
    createApiKeyRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createApiKeyRequest** | **CreateApiKeyRequest**|  | |


### Return type

**ApiKey**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | API key created successfully |  -  |
|**400** | Bad request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createLead**
> Lead createLead(createLeadRequest)

Create a new lead

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    CreateLeadRequest
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let createLeadRequest: CreateLeadRequest; //

const { status, data } = await apiInstance.createLead(
    createLeadRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createLeadRequest** | **CreateLeadRequest**|  | |


### Return type

**Lead**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Lead created successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createVoiceClone**
> VoiceCloneResponse createVoiceClone()

Create a new voice clone from audio file uploads

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let name: string; //Descriptive name for the voice clone (default to undefined)
let files: Array<File>; //Audio files for voice training (1-10 files, max 50MB each) (default to undefined)
let description: string; //Optional description of the voice clone (optional) (default to undefined)

const { status, data } = await apiInstance.createVoiceClone(
    name,
    files,
    description
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | [**string**] | Descriptive name for the voice clone | defaults to undefined|
| **files** | **Array&lt;File&gt;** | Audio files for voice training (1-10 files, max 50MB each) | defaults to undefined|
| **description** | [**string**] | Optional description of the voice clone | (optional) defaults to undefined|


### Return type

**VoiceCloneResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Voice clone created successfully |  -  |
|**400** | Bad request |  -  |
|**413** | File too large |  -  |
|**429** | Voice clone limit exceeded |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteVoiceClone**
> DeleteVoiceCloneResponse deleteVoiceClone()

Permanently delete a voice clone

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let voiceId: string; //The unique identifier of the voice clone (default to undefined)

const { status, data } = await apiInstance.deleteVoiceClone(
    voiceId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **voiceId** | [**string**] | The unique identifier of the voice clone | defaults to undefined|


### Return type

**DeleteVoiceCloneResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Voice clone deleted successfully |  -  |
|**404** | Voice clone not found |  -  |
|**409** | Voice clone is in use |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **generateVoicePreview**
> File generateVoicePreview(voicePreviewRequest)

Generate a voice preview using a voice clone to test quality

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    VoicePreviewRequest
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let voicePreviewRequest: VoicePreviewRequest; //

const { status, data } = await apiInstance.generateVoicePreview(
    voicePreviewRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **voicePreviewRequest** | **VoicePreviewRequest**|  | |


### Return type

**File**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: audio/mpeg, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Voice preview generated successfully |  -  |
|**400** | Bad request |  -  |
|**404** | Voice clone not found |  -  |
|**500** | Voice generation failed |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **healthCheck**
> HealthCheck200Response healthCheck()

Check the API status and health

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.healthCheck();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**HealthCheck200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | API is healthy |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listApiKeys**
> Array<ApiKey> listApiKeys()

List all API keys for the user

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.listApiKeys();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<ApiKey>**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of API keys |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listLeads**
> ListLeads200Response listLeads()

Get a list of leads with optional filtering

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let limit: number; //Maximum number of leads to return (optional) (default to 50)
let offset: number; //Number of leads to skip (optional) (default to 0)

const { status, data } = await apiInstance.listLeads(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | Maximum number of leads to return | (optional) defaults to 50|
| **offset** | [**number**] | Number of leads to skip | (optional) defaults to 0|


### Return type

**ListLeads200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of leads |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listVoiceClones**
> VoiceCloneListResponse listVoiceClones()

Retrieve all voice clones for the authenticated user

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let limit: number; //Number of voice clones to return (1-100) (optional) (default to 50)
let offset: number; //Pagination offset (optional) (default to 0)

const { status, data } = await apiInstance.listVoiceClones(
    limit,
    offset
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | Number of voice clones to return (1-100) | (optional) defaults to 50|
| **offset** | [**number**] | Pagination offset | (optional) defaults to 0|


### Return type

**VoiceCloneListResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of voice clones |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateVoiceClone**
> UpdateVoiceCloneResponse updateVoiceClone(updateVoiceCloneRequest)

Update voice clone settings and metadata

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    UpdateVoiceCloneRequest
} from '@buena/sdk';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let voiceId: string; //The unique identifier of the voice clone (default to undefined)
let updateVoiceCloneRequest: UpdateVoiceCloneRequest; //

const { status, data } = await apiInstance.updateVoiceClone(
    voiceId,
    updateVoiceCloneRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateVoiceCloneRequest** | **UpdateVoiceCloneRequest**|  | |
| **voiceId** | [**string**] | The unique identifier of the voice clone | defaults to undefined|


### Return type

**UpdateVoiceCloneResponse**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Voice clone updated successfully |  -  |
|**404** | Voice clone not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


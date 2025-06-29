# buena_sdk.DefaultApi

All URIs are relative to *https://api.buena.ai/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_api_key**](DefaultApi.md#create_api_key) | **POST** /keys | Create API Key
[**create_lead**](DefaultApi.md#create_lead) | **POST** /leads | Create Lead
[**health_check**](DefaultApi.md#health_check) | **GET** /health | Health Check
[**list_api_keys**](DefaultApi.md#list_api_keys) | **GET** /keys | List API Keys
[**list_leads**](DefaultApi.md#list_leads) | **GET** /leads | List Leads


# **create_api_key**
> ApiKey create_api_key(create_api_key_request)

Create API Key

Create a new API key

### Example

* Api Key Authentication (ApiKeyAuth):

```python
import buena_sdk
from buena_sdk.models.api_key import ApiKey
from buena_sdk.models.create_api_key_request import CreateApiKeyRequest
from buena_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.buena.ai/api/v2
# See configuration.py for a list of all supported configuration parameters.
configuration = buena_sdk.Configuration(
    host = "https://api.buena.ai/api/v2"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuth
configuration.api_key['ApiKeyAuth'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuth'] = 'Bearer'

# Enter a context with an instance of the API client
with buena_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = buena_sdk.DefaultApi(api_client)
    create_api_key_request = buena_sdk.CreateApiKeyRequest() # CreateApiKeyRequest | 

    try:
        # Create API Key
        api_response = api_instance.create_api_key(create_api_key_request)
        print("The response of DefaultApi->create_api_key:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->create_api_key: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_api_key_request** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md)|  | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | API key created successfully |  -  |
**400** | Bad request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_lead**
> Lead create_lead(create_lead_request)

Create Lead

Create a new lead

### Example

* Api Key Authentication (ApiKeyAuth):

```python
import buena_sdk
from buena_sdk.models.create_lead_request import CreateLeadRequest
from buena_sdk.models.lead import Lead
from buena_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.buena.ai/api/v2
# See configuration.py for a list of all supported configuration parameters.
configuration = buena_sdk.Configuration(
    host = "https://api.buena.ai/api/v2"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuth
configuration.api_key['ApiKeyAuth'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuth'] = 'Bearer'

# Enter a context with an instance of the API client
with buena_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = buena_sdk.DefaultApi(api_client)
    create_lead_request = buena_sdk.CreateLeadRequest() # CreateLeadRequest | 

    try:
        # Create Lead
        api_response = api_instance.create_lead(create_lead_request)
        print("The response of DefaultApi->create_lead:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->create_lead: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_lead_request** | [**CreateLeadRequest**](CreateLeadRequest.md)|  | 

### Return type

[**Lead**](Lead.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Lead created successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **health_check**
> HealthCheck200Response health_check()

Health Check

Check the API status and health

### Example

* Api Key Authentication (ApiKeyAuth):

```python
import buena_sdk
from buena_sdk.models.health_check200_response import HealthCheck200Response
from buena_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.buena.ai/api/v2
# See configuration.py for a list of all supported configuration parameters.
configuration = buena_sdk.Configuration(
    host = "https://api.buena.ai/api/v2"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuth
configuration.api_key['ApiKeyAuth'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuth'] = 'Bearer'

# Enter a context with an instance of the API client
with buena_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = buena_sdk.DefaultApi(api_client)

    try:
        # Health Check
        api_response = api_instance.health_check()
        print("The response of DefaultApi->health_check:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->health_check: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**HealthCheck200Response**](HealthCheck200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | API is healthy |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_api_keys**
> List[ApiKey] list_api_keys()

List API Keys

List all API keys for the user

### Example

* Api Key Authentication (ApiKeyAuth):

```python
import buena_sdk
from buena_sdk.models.api_key import ApiKey
from buena_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.buena.ai/api/v2
# See configuration.py for a list of all supported configuration parameters.
configuration = buena_sdk.Configuration(
    host = "https://api.buena.ai/api/v2"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuth
configuration.api_key['ApiKeyAuth'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuth'] = 'Bearer'

# Enter a context with an instance of the API client
with buena_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = buena_sdk.DefaultApi(api_client)

    try:
        # List API Keys
        api_response = api_instance.list_api_keys()
        print("The response of DefaultApi->list_api_keys:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->list_api_keys: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**List[ApiKey]**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of API keys |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_leads**
> ListLeads200Response list_leads(limit=limit, offset=offset)

List Leads

Get a list of leads with optional filtering

### Example

* Api Key Authentication (ApiKeyAuth):

```python
import buena_sdk
from buena_sdk.models.list_leads200_response import ListLeads200Response
from buena_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.buena.ai/api/v2
# See configuration.py for a list of all supported configuration parameters.
configuration = buena_sdk.Configuration(
    host = "https://api.buena.ai/api/v2"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuth
configuration.api_key['ApiKeyAuth'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuth'] = 'Bearer'

# Enter a context with an instance of the API client
with buena_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = buena_sdk.DefaultApi(api_client)
    limit = 50 # int | Maximum number of leads to return (optional) (default to 50)
    offset = 0 # int | Number of leads to skip (optional) (default to 0)

    try:
        # List Leads
        api_response = api_instance.list_leads(limit=limit, offset=offset)
        print("The response of DefaultApi->list_leads:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->list_leads: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int**| Maximum number of leads to return | [optional] [default to 50]
 **offset** | **int**| Number of leads to skip | [optional] [default to 0]

### Return type

[**ListLeads200Response**](ListLeads200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of leads |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


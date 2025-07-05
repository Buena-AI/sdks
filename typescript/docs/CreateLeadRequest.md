# CreateLeadRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**first_name** | **string** | First name of the lead | [default to undefined]
**last_name** | **string** | Last name of the lead | [default to undefined]
**email** | **string** | Email address of the lead | [optional] [default to undefined]
**company** | **string** | Company name | [optional] [default to undefined]
**title** | **string** | Job title | [optional] [default to undefined]
**linkedin_url** | **string** | LinkedIn profile URL | [optional] [default to undefined]

## Example

```typescript
import { CreateLeadRequest } from '@buena/sdk';

const instance: CreateLeadRequest = {
    first_name,
    last_name,
    email,
    company,
    title,
    linkedin_url,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# ApiKey


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | Unique identifier for the API key | [optional] [default to undefined]
**name** | **string** | Name of the API key | [optional] [default to undefined]
**key** | **string** | The actual API key (only shown once during creation) | [optional] [default to undefined]
**permissions** | **Array&lt;string&gt;** | List of permissions | [optional] [default to undefined]
**created_at** | **string** | When the key was created | [optional] [default to undefined]
**last_used** | **string** | When the key was last used | [optional] [default to undefined]

## Example

```typescript
import { ApiKey } from '@buena/sdk';

const instance: ApiKey = {
    id,
    name,
    key,
    permissions,
    created_at,
    last_used,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

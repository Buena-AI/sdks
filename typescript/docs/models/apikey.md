# ApiKey

## Example Usage

```typescript
import { ApiKey } from "openapi";

let value: ApiKey = {};
```

## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `id`                                                                                          | *string*                                                                                      | :heavy_minus_sign:                                                                            | Unique identifier for the API key                                                             |
| `name`                                                                                        | *string*                                                                                      | :heavy_minus_sign:                                                                            | Name of the API key                                                                           |
| `key`                                                                                         | *string*                                                                                      | :heavy_minus_sign:                                                                            | The actual API key (only shown once during creation)                                          |
| `permissions`                                                                                 | *string*[]                                                                                    | :heavy_minus_sign:                                                                            | List of permissions                                                                           |
| `createdAt`                                                                                   | [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date) | :heavy_minus_sign:                                                                            | When the key was created                                                                      |
| `lastUsed`                                                                                    | [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date) | :heavy_minus_sign:                                                                            | When the key was last used                                                                    |
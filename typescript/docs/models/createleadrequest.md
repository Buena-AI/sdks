# CreateLeadRequest

## Example Usage

```typescript
import { CreateLeadRequest } from "openapi";

let value: CreateLeadRequest = {
  firstName: "Remington",
  lastName: "Hagenes",
};
```

## Fields

| Field                     | Type                      | Required                  | Description               |
| ------------------------- | ------------------------- | ------------------------- | ------------------------- |
| `firstName`               | *string*                  | :heavy_check_mark:        | First name of the lead    |
| `lastName`                | *string*                  | :heavy_check_mark:        | Last name of the lead     |
| `email`                   | *string*                  | :heavy_minus_sign:        | Email address of the lead |
| `company`                 | *string*                  | :heavy_minus_sign:        | Company name              |
| `title`                   | *string*                  | :heavy_minus_sign:        | Job title                 |
| `linkedinUrl`             | *string*                  | :heavy_minus_sign:        | LinkedIn profile URL      |
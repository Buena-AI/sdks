<!-- Start SDK Example Usage [usage] -->

```typescript
import { SDK } from "@buena/sdk";

const sdk = new SDK({
  apiKeyAuth: "your-api-key-here",
  serverURL: "https://api.buena.ai/api/v2",
});

async function run() {
  // Health check
  const health = await sdk.healthCheck();
  console.log("API Status:", health);

  // List leads
  const leads = await sdk.listLeads();
  console.log("Leads:", leads);

  // Create a new lead
  const newLead = await sdk.createLead({
    firstName: "John",
    lastName: "Doe",
    email: "john@example.com",
    company: "Example Corp",
  });
  console.log("Created lead:", newLead);
}

run();
```

<!-- End SDK Example Usage [usage] -->

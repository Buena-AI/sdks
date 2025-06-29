import { ClientSDK, RequestOptions } from "../lib/sdks.js";
import * as models from "../models/index.js";
import * as operations from "../models/operations/index.js";
export declare class SDK extends ClientSDK {
    /**
     * Health Check
     *
     * @remarks
     * Check the API status and health
     */
    healthCheck(options?: RequestOptions): Promise<operations.HealthCheckResponse>;
    /**
     * Create API Key
     *
     * @remarks
     * Create a new API key
     */
    createApiKey(request: operations.CreateApiKeyRequest, options?: RequestOptions): Promise<models.ApiKey>;
    /**
     * List API Keys
     *
     * @remarks
     * List all API keys for the user
     */
    listApiKeys(options?: RequestOptions): Promise<Array<models.ApiKey>>;
    /**
     * List Leads
     *
     * @remarks
     * Get a list of leads with optional filtering
     */
    listLeads(request?: operations.ListLeadsRequest | undefined, options?: RequestOptions): Promise<operations.ListLeadsResponse>;
    /**
     * Create Lead
     *
     * @remarks
     * Create a new lead
     */
    createLead(request: models.CreateLeadRequest, options?: RequestOptions): Promise<models.Lead>;
}
//# sourceMappingURL=sdk.d.ts.map
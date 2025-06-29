import * as z from "zod";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./errors/sdkvalidationerror.js";
export type ApiKey = {
    /**
     * Unique identifier for the API key
     */
    id?: string | undefined;
    /**
     * Name of the API key
     */
    name?: string | undefined;
    /**
     * The actual API key (only shown once during creation)
     */
    key?: string | undefined;
    /**
     * List of permissions
     */
    permissions?: Array<string> | undefined;
    /**
     * When the key was created
     */
    createdAt?: Date | undefined;
    /**
     * When the key was last used
     */
    lastUsed?: Date | undefined;
};
/** @internal */
export declare const ApiKey$inboundSchema: z.ZodType<ApiKey, z.ZodTypeDef, unknown>;
/** @internal */
export type ApiKey$Outbound = {
    id?: string | undefined;
    name?: string | undefined;
    key?: string | undefined;
    permissions?: Array<string> | undefined;
    created_at?: string | undefined;
    last_used?: string | undefined;
};
/** @internal */
export declare const ApiKey$outboundSchema: z.ZodType<ApiKey$Outbound, z.ZodTypeDef, ApiKey>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace ApiKey$ {
    /** @deprecated use `ApiKey$inboundSchema` instead. */
    const inboundSchema: z.ZodType<ApiKey, z.ZodTypeDef, unknown>;
    /** @deprecated use `ApiKey$outboundSchema` instead. */
    const outboundSchema: z.ZodType<ApiKey$Outbound, z.ZodTypeDef, ApiKey>;
    /** @deprecated use `ApiKey$Outbound` instead. */
    type Outbound = ApiKey$Outbound;
}
export declare function apiKeyToJSON(apiKey: ApiKey): string;
export declare function apiKeyFromJSON(jsonString: string): SafeParseResult<ApiKey, SDKValidationError>;
//# sourceMappingURL=apikey.d.ts.map
import * as z from "zod";
import { Result as SafeParseResult } from "../../types/fp.js";
import { SDKValidationError } from "../errors/sdkvalidationerror.js";
export type CreateApiKeyRequest = {
    /**
     * Name for the API key
     */
    name?: string | undefined;
    /**
     * List of permissions for the key
     */
    permissions?: Array<string> | undefined;
};
/** @internal */
export declare const CreateApiKeyRequest$inboundSchema: z.ZodType<CreateApiKeyRequest, z.ZodTypeDef, unknown>;
/** @internal */
export type CreateApiKeyRequest$Outbound = {
    name?: string | undefined;
    permissions?: Array<string> | undefined;
};
/** @internal */
export declare const CreateApiKeyRequest$outboundSchema: z.ZodType<CreateApiKeyRequest$Outbound, z.ZodTypeDef, CreateApiKeyRequest>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace CreateApiKeyRequest$ {
    /** @deprecated use `CreateApiKeyRequest$inboundSchema` instead. */
    const inboundSchema: z.ZodType<CreateApiKeyRequest, z.ZodTypeDef, unknown>;
    /** @deprecated use `CreateApiKeyRequest$outboundSchema` instead. */
    const outboundSchema: z.ZodType<CreateApiKeyRequest$Outbound, z.ZodTypeDef, CreateApiKeyRequest>;
    /** @deprecated use `CreateApiKeyRequest$Outbound` instead. */
    type Outbound = CreateApiKeyRequest$Outbound;
}
export declare function createApiKeyRequestToJSON(createApiKeyRequest: CreateApiKeyRequest): string;
export declare function createApiKeyRequestFromJSON(jsonString: string): SafeParseResult<CreateApiKeyRequest, SDKValidationError>;
//# sourceMappingURL=createapikey.d.ts.map
import * as z from "zod";
import { Result as SafeParseResult } from "../../types/fp.js";
import { SDKValidationError } from "../errors/sdkvalidationerror.js";
import * as models from "../index.js";
export type ListLeadsRequest = {
    /**
     * Maximum number of leads to return
     */
    limit?: number | undefined;
    /**
     * Number of leads to skip
     */
    offset?: number | undefined;
};
/**
 * List of leads
 */
export type ListLeadsResponse = {
    leads?: Array<models.Lead> | undefined;
    total?: number | undefined;
};
/** @internal */
export declare const ListLeadsRequest$inboundSchema: z.ZodType<ListLeadsRequest, z.ZodTypeDef, unknown>;
/** @internal */
export type ListLeadsRequest$Outbound = {
    limit: number;
    offset: number;
};
/** @internal */
export declare const ListLeadsRequest$outboundSchema: z.ZodType<ListLeadsRequest$Outbound, z.ZodTypeDef, ListLeadsRequest>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace ListLeadsRequest$ {
    /** @deprecated use `ListLeadsRequest$inboundSchema` instead. */
    const inboundSchema: z.ZodType<ListLeadsRequest, z.ZodTypeDef, unknown>;
    /** @deprecated use `ListLeadsRequest$outboundSchema` instead. */
    const outboundSchema: z.ZodType<ListLeadsRequest$Outbound, z.ZodTypeDef, ListLeadsRequest>;
    /** @deprecated use `ListLeadsRequest$Outbound` instead. */
    type Outbound = ListLeadsRequest$Outbound;
}
export declare function listLeadsRequestToJSON(listLeadsRequest: ListLeadsRequest): string;
export declare function listLeadsRequestFromJSON(jsonString: string): SafeParseResult<ListLeadsRequest, SDKValidationError>;
/** @internal */
export declare const ListLeadsResponse$inboundSchema: z.ZodType<ListLeadsResponse, z.ZodTypeDef, unknown>;
/** @internal */
export type ListLeadsResponse$Outbound = {
    leads?: Array<models.Lead$Outbound> | undefined;
    total?: number | undefined;
};
/** @internal */
export declare const ListLeadsResponse$outboundSchema: z.ZodType<ListLeadsResponse$Outbound, z.ZodTypeDef, ListLeadsResponse>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace ListLeadsResponse$ {
    /** @deprecated use `ListLeadsResponse$inboundSchema` instead. */
    const inboundSchema: z.ZodType<ListLeadsResponse, z.ZodTypeDef, unknown>;
    /** @deprecated use `ListLeadsResponse$outboundSchema` instead. */
    const outboundSchema: z.ZodType<ListLeadsResponse$Outbound, z.ZodTypeDef, ListLeadsResponse>;
    /** @deprecated use `ListLeadsResponse$Outbound` instead. */
    type Outbound = ListLeadsResponse$Outbound;
}
export declare function listLeadsResponseToJSON(listLeadsResponse: ListLeadsResponse): string;
export declare function listLeadsResponseFromJSON(jsonString: string): SafeParseResult<ListLeadsResponse, SDKValidationError>;
//# sourceMappingURL=listleads.d.ts.map
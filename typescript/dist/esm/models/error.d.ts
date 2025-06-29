import * as z from "zod";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./errors/sdkvalidationerror.js";
/**
 * Additional error details
 */
export type Details = {};
/** @internal */
export declare const Details$inboundSchema: z.ZodType<Details, z.ZodTypeDef, unknown>;
/** @internal */
export type Details$Outbound = {};
/** @internal */
export declare const Details$outboundSchema: z.ZodType<Details$Outbound, z.ZodTypeDef, Details>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace Details$ {
    /** @deprecated use `Details$inboundSchema` instead. */
    const inboundSchema: z.ZodType<Details, z.ZodTypeDef, unknown>;
    /** @deprecated use `Details$outboundSchema` instead. */
    const outboundSchema: z.ZodType<Details$Outbound, z.ZodTypeDef, Details>;
    /** @deprecated use `Details$Outbound` instead. */
    type Outbound = Details$Outbound;
}
export declare function detailsToJSON(details: Details): string;
export declare function detailsFromJSON(jsonString: string): SafeParseResult<Details, SDKValidationError>;
//# sourceMappingURL=error.d.ts.map
import * as z from "zod";
import * as models from "../index.js";
import { SDKError } from "./sdkerror.js";
export type ErrorTData = {
    /**
     * Error code
     */
    error: string;
    /**
     * Human-readable error message
     */
    message: string;
    /**
     * Additional error details
     */
    details?: models.Details | undefined;
};
export declare class ErrorT extends SDKError {
    /**
     * Error code
     */
    error: string;
    /**
     * Additional error details
     */
    details?: models.Details | undefined;
    /** The original data that was passed to this error instance. */
    data$: ErrorTData;
    constructor(err: ErrorTData, httpMeta: {
        response: Response;
        request: Request;
        body: string;
    });
}
/** @internal */
export declare const ErrorT$inboundSchema: z.ZodType<ErrorT, z.ZodTypeDef, unknown>;
/** @internal */
export type ErrorT$Outbound = {
    error: string;
    message: string;
    details?: models.Details$Outbound | undefined;
};
/** @internal */
export declare const ErrorT$outboundSchema: z.ZodType<ErrorT$Outbound, z.ZodTypeDef, ErrorT>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace ErrorT$ {
    /** @deprecated use `ErrorT$inboundSchema` instead. */
    const inboundSchema: z.ZodType<ErrorT, z.ZodTypeDef, unknown>;
    /** @deprecated use `ErrorT$outboundSchema` instead. */
    const outboundSchema: z.ZodType<ErrorT$Outbound, z.ZodTypeDef, ErrorT>;
    /** @deprecated use `ErrorT$Outbound` instead. */
    type Outbound = ErrorT$Outbound;
}
//# sourceMappingURL=error.d.ts.map
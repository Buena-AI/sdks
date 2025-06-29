import * as z from "zod";
import { Result as SafeParseResult } from "../../types/fp.js";
import { SDKValidationError } from "../errors/sdkvalidationerror.js";
/**
 * API is healthy
 */
export type HealthCheckResponse = {
    status?: string | undefined;
    timestamp?: Date | undefined;
};
/** @internal */
export declare const HealthCheckResponse$inboundSchema: z.ZodType<HealthCheckResponse, z.ZodTypeDef, unknown>;
/** @internal */
export type HealthCheckResponse$Outbound = {
    status?: string | undefined;
    timestamp?: string | undefined;
};
/** @internal */
export declare const HealthCheckResponse$outboundSchema: z.ZodType<HealthCheckResponse$Outbound, z.ZodTypeDef, HealthCheckResponse>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace HealthCheckResponse$ {
    /** @deprecated use `HealthCheckResponse$inboundSchema` instead. */
    const inboundSchema: z.ZodType<HealthCheckResponse, z.ZodTypeDef, unknown>;
    /** @deprecated use `HealthCheckResponse$outboundSchema` instead. */
    const outboundSchema: z.ZodType<HealthCheckResponse$Outbound, z.ZodTypeDef, HealthCheckResponse>;
    /** @deprecated use `HealthCheckResponse$Outbound` instead. */
    type Outbound = HealthCheckResponse$Outbound;
}
export declare function healthCheckResponseToJSON(healthCheckResponse: HealthCheckResponse): string;
export declare function healthCheckResponseFromJSON(jsonString: string): SafeParseResult<HealthCheckResponse, SDKValidationError>;
//# sourceMappingURL=healthcheck.d.ts.map
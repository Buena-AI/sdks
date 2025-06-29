import * as z from "zod";
import { ClosedEnum } from "../types/enums.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./errors/sdkvalidationerror.js";
/**
 * Current status of the lead
 */
export declare const Status: {
    readonly New: "new";
    readonly Contacted: "contacted";
    readonly Responded: "responded";
    readonly Converted: "converted";
    readonly Unqualified: "unqualified";
};
/**
 * Current status of the lead
 */
export type Status = ClosedEnum<typeof Status>;
export type Lead = {
    /**
     * Unique identifier for the lead
     */
    id?: string | undefined;
    /**
     * First name of the lead
     */
    firstName?: string | undefined;
    /**
     * Last name of the lead
     */
    lastName?: string | undefined;
    /**
     * Email address of the lead
     */
    email?: string | undefined;
    /**
     * Company name
     */
    company?: string | undefined;
    /**
     * Job title
     */
    title?: string | undefined;
    /**
     * LinkedIn profile URL
     */
    linkedinUrl?: string | undefined;
    /**
     * Current status of the lead
     */
    status?: Status | undefined;
    /**
     * When the lead was created
     */
    createdAt?: Date | undefined;
    /**
     * When the lead was last updated
     */
    updatedAt?: Date | undefined;
};
/** @internal */
export declare const Status$inboundSchema: z.ZodNativeEnum<typeof Status>;
/** @internal */
export declare const Status$outboundSchema: z.ZodNativeEnum<typeof Status>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace Status$ {
    /** @deprecated use `Status$inboundSchema` instead. */
    const inboundSchema: z.ZodNativeEnum<{
        readonly New: "new";
        readonly Contacted: "contacted";
        readonly Responded: "responded";
        readonly Converted: "converted";
        readonly Unqualified: "unqualified";
    }>;
    /** @deprecated use `Status$outboundSchema` instead. */
    const outboundSchema: z.ZodNativeEnum<{
        readonly New: "new";
        readonly Contacted: "contacted";
        readonly Responded: "responded";
        readonly Converted: "converted";
        readonly Unqualified: "unqualified";
    }>;
}
/** @internal */
export declare const Lead$inboundSchema: z.ZodType<Lead, z.ZodTypeDef, unknown>;
/** @internal */
export type Lead$Outbound = {
    id?: string | undefined;
    first_name?: string | undefined;
    last_name?: string | undefined;
    email?: string | undefined;
    company?: string | undefined;
    title?: string | undefined;
    linkedin_url?: string | undefined;
    status?: string | undefined;
    created_at?: string | undefined;
    updated_at?: string | undefined;
};
/** @internal */
export declare const Lead$outboundSchema: z.ZodType<Lead$Outbound, z.ZodTypeDef, Lead>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace Lead$ {
    /** @deprecated use `Lead$inboundSchema` instead. */
    const inboundSchema: z.ZodType<Lead, z.ZodTypeDef, unknown>;
    /** @deprecated use `Lead$outboundSchema` instead. */
    const outboundSchema: z.ZodType<Lead$Outbound, z.ZodTypeDef, Lead>;
    /** @deprecated use `Lead$Outbound` instead. */
    type Outbound = Lead$Outbound;
}
export declare function leadToJSON(lead: Lead): string;
export declare function leadFromJSON(jsonString: string): SafeParseResult<Lead, SDKValidationError>;
//# sourceMappingURL=lead.d.ts.map
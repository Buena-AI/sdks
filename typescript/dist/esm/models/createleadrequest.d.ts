import * as z from "zod";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./errors/sdkvalidationerror.js";
export type CreateLeadRequest = {
    /**
     * First name of the lead
     */
    firstName: string;
    /**
     * Last name of the lead
     */
    lastName: string;
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
};
/** @internal */
export declare const CreateLeadRequest$inboundSchema: z.ZodType<CreateLeadRequest, z.ZodTypeDef, unknown>;
/** @internal */
export type CreateLeadRequest$Outbound = {
    first_name: string;
    last_name: string;
    email?: string | undefined;
    company?: string | undefined;
    title?: string | undefined;
    linkedin_url?: string | undefined;
};
/** @internal */
export declare const CreateLeadRequest$outboundSchema: z.ZodType<CreateLeadRequest$Outbound, z.ZodTypeDef, CreateLeadRequest>;
/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export declare namespace CreateLeadRequest$ {
    /** @deprecated use `CreateLeadRequest$inboundSchema` instead. */
    const inboundSchema: z.ZodType<CreateLeadRequest, z.ZodTypeDef, unknown>;
    /** @deprecated use `CreateLeadRequest$outboundSchema` instead. */
    const outboundSchema: z.ZodType<CreateLeadRequest$Outbound, z.ZodTypeDef, CreateLeadRequest>;
    /** @deprecated use `CreateLeadRequest$Outbound` instead. */
    type Outbound = CreateLeadRequest$Outbound;
}
export declare function createLeadRequestToJSON(createLeadRequest: CreateLeadRequest): string;
export declare function createLeadRequestFromJSON(jsonString: string): SafeParseResult<CreateLeadRequest, SDKValidationError>;
//# sourceMappingURL=createleadrequest.d.ts.map
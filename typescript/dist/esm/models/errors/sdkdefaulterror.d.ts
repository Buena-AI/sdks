import { SDKError } from "./sdkerror.js";
/** The fallback error class if no more specific error class is matched */
export declare class SDKDefaultError extends SDKError {
    constructor(message: string, httpMeta: {
        response: Response;
        request: Request;
        body: string;
    });
}
//# sourceMappingURL=sdkdefaulterror.d.ts.map
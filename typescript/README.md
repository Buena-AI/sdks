## @buena/sdk@1.0.0

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install @buena/sdk@1.0.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *https://api.buena.ai/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**createApiKey**](docs/DefaultApi.md#createapikey) | **POST** /keys | Create API Key
*DefaultApi* | [**createLead**](docs/DefaultApi.md#createlead) | **POST** /leads | Create Lead
*DefaultApi* | [**createVoiceClone**](docs/DefaultApi.md#createvoiceclone) | **POST** /voice-clones | Create Voice Clone
*DefaultApi* | [**deleteVoiceClone**](docs/DefaultApi.md#deletevoiceclone) | **DELETE** /voice-clones/{voiceId} | Delete Voice Clone
*DefaultApi* | [**generateVoicePreview**](docs/DefaultApi.md#generatevoicepreview) | **POST** /voice-clones/preview | Generate Voice Preview
*DefaultApi* | [**healthCheck**](docs/DefaultApi.md#healthcheck) | **GET** /health | Health Check
*DefaultApi* | [**listApiKeys**](docs/DefaultApi.md#listapikeys) | **GET** /keys | List API Keys
*DefaultApi* | [**listLeads**](docs/DefaultApi.md#listleads) | **GET** /leads | List Leads
*DefaultApi* | [**listVoiceClones**](docs/DefaultApi.md#listvoiceclones) | **GET** /voice-clones | List Voice Clones
*DefaultApi* | [**updateVoiceClone**](docs/DefaultApi.md#updatevoiceclone) | **PUT** /voice-clones/{voiceId} | Update Voice Clone


### Documentation For Models

 - [ApiKey](docs/ApiKey.md)
 - [CreateApiKeyRequest](docs/CreateApiKeyRequest.md)
 - [CreateLeadRequest](docs/CreateLeadRequest.md)
 - [DeleteVoiceCloneResponse](docs/DeleteVoiceCloneResponse.md)
 - [HealthCheck200Response](docs/HealthCheck200Response.md)
 - [Lead](docs/Lead.md)
 - [ListLeads200Response](docs/ListLeads200Response.md)
 - [ModelError](docs/ModelError.md)
 - [UpdateVoiceCloneRequest](docs/UpdateVoiceCloneRequest.md)
 - [UpdateVoiceCloneResponse](docs/UpdateVoiceCloneResponse.md)
 - [UpdateVoiceCloneResponseData](docs/UpdateVoiceCloneResponseData.md)
 - [VoiceClone](docs/VoiceClone.md)
 - [VoiceCloneListResponse](docs/VoiceCloneListResponse.md)
 - [VoiceCloneListResponseData](docs/VoiceCloneListResponseData.md)
 - [VoiceCloneResponse](docs/VoiceCloneResponse.md)
 - [VoiceCloneResponseData](docs/VoiceCloneResponseData.md)
 - [VoicePreviewRequest](docs/VoicePreviewRequest.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization


Authentication schemes defined for the API:
<a id="BearerAuth"></a>
### BearerAuth

- **Type**: Bearer authentication (JWT)


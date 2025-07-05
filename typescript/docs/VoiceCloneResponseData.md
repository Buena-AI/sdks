# VoiceCloneResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**voiceId** | **string** | Unique voice clone identifier | [optional] [default to undefined]
**name** | **string** | User-defined name for the voice clone | [optional] [default to undefined]
**description** | **string** | Description of the voice clone | [optional] [default to undefined]
**sampleCount** | **number** | Number of audio samples uploaded for training | [optional] [default to undefined]
**createdAt** | **string** | When the voice clone was created | [optional] [default to undefined]

## Example

```typescript
import { VoiceCloneResponseData } from '@buena/sdk';

const instance: VoiceCloneResponseData = {
    voiceId,
    name,
    description,
    sampleCount,
    createdAt,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# VoiceClone


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**voiceId** | **string** | Unique voice clone identifier | [optional] [default to undefined]
**name** | **string** | User-defined name for the voice clone | [optional] [default to undefined]
**description** | **string** | Description of the voice clone | [optional] [default to undefined]
**isActive** | **boolean** | Whether the voice clone is active and available for use | [optional] [default to undefined]
**sampleCount** | **number** | Number of audio samples used to train this voice | [optional] [default to undefined]
**createdAt** | **string** | When the voice clone was created | [optional] [default to undefined]
**updatedAt** | **string** | When the voice clone was last updated | [optional] [default to undefined]
**labels** | **object** | Metadata labels associated with the voice clone | [optional] [default to undefined]

## Example

```typescript
import { VoiceClone } from '@buena/sdk';

const instance: VoiceClone = {
    voiceId,
    name,
    description,
    isActive,
    sampleCount,
    createdAt,
    updatedAt,
    labels,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# coding: utf-8

"""
    Buena.ai API v2

    The most powerful LinkedIn automation and lead management API for modern sales teams and developers.

    The version of the OpenAPI document: 2.0.0
    Contact: support@buena.ai
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from buena_sdk.models.voice_clone_list_response_data import VoiceCloneListResponseData

class TestVoiceCloneListResponseData(unittest.TestCase):
    """VoiceCloneListResponseData unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> VoiceCloneListResponseData:
        """Test VoiceCloneListResponseData
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `VoiceCloneListResponseData`
        """
        model = VoiceCloneListResponseData()
        if include_optional:
            return VoiceCloneListResponseData(
                voice_clones = [
                    buena_sdk.models.voice_clone.VoiceClone(
                        voice_id = '', 
                        name = '', 
                        description = '', 
                        is_active = True, 
                        sample_count = 56, 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        labels = buena_sdk.models.labels.labels(), )
                    ],
                total = 56
            )
        else:
            return VoiceCloneListResponseData(
        )
        """

    def testVoiceCloneListResponseData(self):
        """Test VoiceCloneListResponseData"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

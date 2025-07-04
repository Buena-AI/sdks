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

from buena_sdk.models.create_lead_request import CreateLeadRequest

class TestCreateLeadRequest(unittest.TestCase):
    """CreateLeadRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CreateLeadRequest:
        """Test CreateLeadRequest
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CreateLeadRequest`
        """
        model = CreateLeadRequest()
        if include_optional:
            return CreateLeadRequest(
                first_name = '',
                last_name = '',
                email = '',
                company = '',
                title = '',
                linkedin_url = ''
            )
        else:
            return CreateLeadRequest(
                first_name = '',
                last_name = '',
        )
        """

    def testCreateLeadRequest(self):
        """Test CreateLeadRequest"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

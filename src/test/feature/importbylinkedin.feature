#Feature: Import Lead by Linkedin URL
#  In order to import the Lead using linkedin
#  As an API user
#  I need to be able to request importbylinkedin
#
#  Scenario: does not allow GET method
#    When I send "GET" request to "http://localhost:50000/lead/linkedin/importer"
#    Then the response code should be 405
#    And the response should match json:
#      """
#      {"errors":["Method Not Allowed"],"success":false}
#      """
#
#  Scenario: requires LINKEDIN parameters
#    When I send "POST" request to "http://localhost:50000/lead/linkedin/importer"
#    Then the response code should be 400
#    And the response should match json:
#      """
#      {"errors":["Bad request: Key: 'Linkedin' Error:Field validation for 'Linkedin' failed on the 'required' tag"],"success":false}
#      """
#
#  Scenario: import with linkedin happy flow
#    When I send "POST" request to "http://localhost:50000/lead/linkedin/importer" with linkedin "http://www.linkedin.com/in/adrian-princep-bonser-b9443430"
#    Then the response code should be 200
#    And the response should match json:
#      """
#      {"errors":null,"success":true}
#      """

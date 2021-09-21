#Feature: Import Lead by Email
#  In order to import the Lead using email
#  As an API user
#  I need to be able to request importbyemail
#
#  Scenario: does not allow GET method
#    When I send "GET" request to "http://localhost:50030/lead/email/importer"
#    Then the response code should be 405
#    And the response should match json:
#      """
#      {"errors":["Method Not Allowed"],"success":false}
#      """
#
#  Scenario: requires EMAIL parameters
#    When I send "POST" request to "http://localhost:50030/lead/email/importer"
#    Then the response code should be 400
#    And the response should match json:
#      """
#      {"errors":["Bad request: Key: 'Email' Error:Field validation for 'Email' failed on the 'required' tag"],"success":false}
#      """
#
#  Scenario: import with email happy flow
#    When I send "POST" request to "http://localhost:50030/lead/email/importer" with email "test@test.com"
#    Then the response code should be 200
#    And the response should match json:
#      """
#      {"errors":null,"success":true}
#      """

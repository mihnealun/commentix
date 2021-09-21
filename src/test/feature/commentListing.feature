Feature: List Comments
  In order to list comments for a target
  As an API user
  I need to be able to request comment/list

  Scenario: Receive List of comments
    When I send "GET" request to "http://localhost:50000/comment/list/c62f74df-a965-4468-b0ed-d18919db2403"
    Then the response code should be 200
    And the response should contain:
      """
      {"comments":[{"id":
      """

  Scenario: Receive Empty List of comments
    When I send "GET" request to "http://localhost:50000/comment/list/c62f74df-a965-4468-b0ed-d18919db2404"
    Then the response code should be 200
    And the response should match json:
      """
      {"comments":[]}
      """

Feature: Client Management
  As a system administrator
  I want to manage clients
  So that I can track client information and status

  Background:
    Given the system has a test database

  Scenario: Create a new client successfully
    Given I send a request to create a client with the following details:
      | FirstName    | LastName | Email              | Phone       | Address      | BusinessName | Status | ValidTill  |
      | John         | Doe      | john.doe@test.com  | 1234567890  | 123 Main St  | Test Corp    | Active | 2025-12-31 |
    When the response status should be 201
    Then the response should contain client ID
    And the response should contain FirstName "John"
    And the response should contain Email "john.doe@test.com"

  Scenario: Get client by ID
    Given I send a request to create a client with the following details:
      | FirstName | LastName | Email             | Phone      | Address     | BusinessName | Status | ValidTill  |
      | Jane      | Smith    | jane@test.com    | 9876543210 | 456 Oak Ave | Acme Inc     | Active | 2025-12-31 |
    And I store the client ID from the response
    When I send a request to get client with ID from the response
    Then the response status should be 200
    And the response should contain FirstName "Jane"
    And the response should contain Email "jane@test.com"

  Scenario: Get all clients
    Given a client exists with the following details:
      | FirstName | LastName | Email          | Phone      | Address     | BusinessName | Status | ValidTill  |
      | Bob       | Wilson   | bob@test.com  | 5551234567 | 789 Elm St  | Widget Co    | Active | 2025-12-31 |
    When I send a request to get all clients
    Then the response status should be 200

  Scenario: Update client successfully
    Given I send a request to create a client with the following details:
      | FirstName | LastName | Email            | Phone      | Address     | BusinessName | Status | ValidTill  |
      | Alice     | Brown    | alice@test.com   | 1112223333 | 321 Pine Rd | Tech Solutions | Active | 2025-12-31 |
    And I store the client ID from the response
    When I send a request to update client with ID from stored ID with the following details:
      | FirstName | LastName |
      | Updated   | Name     |
    Then the response status should be 200
    And the response should contain FirstName "Updated"
    And the response should contain LastName "Name"


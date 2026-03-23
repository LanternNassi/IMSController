Feature: User Login
  As a user
  I want to login to the system
  So that I can access protected resources

  Background:
    Given the system has a test database
    And a user exists with email "test@example.com" and password "password123"

  Scenario: Successful login with valid credentials
    Given I have login credentials with email "test@example.com" and password "password123"
    When I send a login request
    Then the response status should be 202
    And the response should contain a token
    And the response should contain username
    And the response should contain email "test@example.com"

  Scenario: Failed login with invalid password
    Given I have login credentials with email "test@example.com" and password "wrongpassword"
    When I send a login request
    Then the response status should be 401
    And the response should contain "Invalid password or username"

  Scenario: Failed login with non-existent email
    Given I have login credentials with email "nonexistent@example.com" and password "password123"
    When I send a login request
    Then the response status should be 401
    And the response should contain "Invalid password or username"


Feature: F1 Greeting

  @story-F1.1
  Rule: A caller who gives a name is greeted by it

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API caller sends "GET /hello?name=Alice"
      Then the response is a 200 JSON object with message "Hello, Alice!"

  @story-F1.2
  Rule: A caller who gives no name gets a generic greeting instead of an error

    Scenario: Greeting a caller with no name
      Given the greeter service is running
      When an API caller sends "GET /hello" with no name parameter
      Then the response is a 200 JSON object with message "Hello, World!"

  @story-F1.3
  Rule: The greeting is always a small JSON object with a single message field

    Scenario: The response shape is a single message field
      Given the greeter service is running
      When an API caller sends "GET /hello?name=Bob"
      Then the response body is a JSON object whose only field is "message"

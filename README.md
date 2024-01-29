# KED Lexer Documentation

## Introduction

The KED lexer is a lexical analyzer for the KED domain-specific language (DSL), which utilizes Cork slang as syntax. It is implemented in Golang and utilizes a test-driven development (TDD) approach to ensure correctness and reliability.

## Lexer Architecture

The KED lexer is composed of the following components:

**Tokenization:** The lexer breaks the input text into individual tokens, each representing a meaningful element of the language, such as identifiers, keywords, operators, and punctuation.

**State Machine:** A finite state machine guides the tokenization process, transitioning between states based on the input characters. Each state corresponds to a specific token type.

**State Transition Table:** The state transition table defines the transitions between states, allowing the lexer to identify the correct token type for each sequence of input characters.

## Tokenization Process

The KED lexer tokenizes the input text according to the following steps:

1. **Initial State:** The lexer starts in the initial state, ready to process the input characters.

2. **Character Examination:** For each input character:

   * **Character Mapping:** Check if the character matches any of the patterns associated with a token type.

   * **State Transition:** If a match is found, transition to the corresponding state. If no match is found, remain in the current state.

   * **Token Generation:** If the state is a terminal state, indicating the end of a token, generate a token object with the corresponding type and lexeme (the string of characters comprising the token).

3. **Output:** Collect and return all generated tokens.

## TDD Approach

A test-driven development (TDD) approach was employed to ensure the correctness and reliability of the lexer. This involved the following steps:

1. **Write a failing test:** Define a test case that describes the expected behavior of the lexer for a specific input sequence.

2. **Write minimum code:** Implement the minimum amount of code necessary to pass the test case.

3. **Refactor code:** Improve the readability and maintainability of the code without changing its functionality.

4. **Refactor tests:** Ensure that the tests remain valid and effective after each refactoring step.

## GitHub Actions Workflow Pipeline

A robust workflow pipeline was developed using GitHub Actions to automatically check the functionality of the lexer before each commit. This pipeline consists of the following steps:

1. **Linting:** Lint the Golang source code to identify potential syntax errors and stylistic issues. # TODO

2. **Unit Testing:** Execute all unit tests to ensure that the lexer correctly identifies and categorizes tokens. 

3. **Code Coverage:** Analyze the code coverage to measure the percentage of code that is being tested by the unit tests. #TODO

4. **Static Analysis:** Perform static analysis to identify potential security vulnerabilities and code quality issues. #TODO

This workflow pipeline helps to maintain code quality, prevent regressions, and ensure that the lexer meets the required standards.

## Conclusion

The KED lexer is a robust and reliable lexical analyzer for the KED DSL, developed using a test-driven development approach and a comprehensive workflow pipeline. It plays a crucial role in the language processing pipeline, enabling the interpretation and manipulation of KED code.

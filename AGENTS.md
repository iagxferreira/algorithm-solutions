# Repo Instructions

This repository is a Kotlin/JVM data structures and algorithms laboratory.

The goal is mastery through reasoning, not a dump of accepted solutions.

## Core Goal

- build deep understanding of algorithms, invariants, complexity, and Kotlin/JVM behavior
- treat each problem as a reusable lesson
- preserve learning history in the repository, not just in chat

## How To Work In This Repo

When starting a new problem or resuming an old one:

1. Read the root [`README.md`](./README.md).
2. Read [`problems/problem.md`](./problems/problem.md).
3. Read the problem-specific README in `problems/<difficulty>/<slug>/README.md`.
4. Read the matching Kotlin implementation in `src/main/kotlin/algorithms/problems/...`.
5. Read the matching test file in `src/test/kotlin/algorithms/problems/...`.
6. Check `git status` before changing anything.

## Problem-Solving Style

Do not jump straight to the optimal solution.

Work through this sequence first:

- Understand
- Constraints
- Examples
- Brute Force
- Bottleneck
- Pattern Recognition
- Data Structure
- Invariant
- Algorithm
- Correctness
- Kotlin Implementation
- Tests
- Complexity
- Reflection

For active study problems, it is often better to guide the reasoning step by step than to dump the final answer immediately.

## Problem File Layout

- Study pages live under `problems/`
- Shared worksheet lives at [`problems/problem.md`](./problems/problem.md)
- Problem writeups live in `problems/<difficulty>/<slug>/README.md`
- Executable Kotlin lives in `src/main/kotlin/algorithms/problems/...`
- Tests live in `src/test/kotlin/algorithms/problems/...`

Keep the study notes and executable code separate.

## Kotlin / JVM Expectations

- Prefer `IntArray`, `LongArray`, `ArrayDeque`, `HashMap`, and `HashSet` when they make sense for DSA
- Be aware of boxing and allocation when choosing collections
- Keep implementations readable before trying to make them clever
- Do not add unnecessary abstractions, frameworks, or application layers

## When Solving Problems

- Use nested loops only when they help explain the brute-force baseline
- Use maps/sets when they eliminate repeated work
- Explain the invariant in plain language before polishing code
- Write the reason the optimization works, not just the final code
- Add or update tests for meaningful edge cases when the implementation changes

## README Expectations

Each problem README should usually contain:

- problem metadata
- a link to [`problems/problem.md`](./problems/problem.md)
- a plain-language understanding of the task
- constraints and examples
- brute-force reasoning
- bottleneck analysis
- pattern, data structure, invariant, correctness
- implementation notes
- test scaffold notes
- complexity
- reflection

## Commit Discipline

Prefer atomic commits with small, reviewable purpose.

Good commit shapes:

- `docs: record two sum study lesson`
- `feat: move problem scaffolds into kotlin source tree`
- `build: add mise toolchain and runnable gradle wrapper`

If a change mixes unrelated concerns, split it before committing.

## Tooling

Use the existing local tooling:

- `mise run test`
- `./gradlew test`

Keep the project lightweight. No Spring Boot, REST APIs, databases, or unnecessary modules.

## Legacy Material

Do not delete the historical archive under [`legacy/`](./legacy) unless explicitly asked.

The old multi-language repository is part of the learning record.

## New Session Shortcut

If this repo is reopened in a fresh session, the first useful questions are:

- Which problem are we working on?
- What does the current README already say?
- What is the brute-force baseline?
- What is the bottleneck?
- Which data structure removes the repeated work?

That is the workflow to preserve.

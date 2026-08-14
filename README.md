# algorithm-solutions

This repository is being refocused into a Kotlin/JVM laboratory for mastering data structures, algorithms, competitive programming, and JVM performance.

The old multi-language solution archive now lives under [`legacy/`](./legacy), preserving the Git history while making room for a cleaner long-term structure.

## Current Focus

- Kotlin/JVM as the primary implementation language
- deep understanding of invariants, complexity, and trade-offs
- reusable explanations for patterns, not just isolated solutions
- lightweight tooling: Gradle, Kotlin, JUnit, and later JMH

## What This Milestone Introduced

- a Kotlin/JVM Gradle build
- a `mise` project config for the local toolchain
- a new documentation structure
- an initial curriculum roadmap
- a small metadata model for future problem tracking
- test infrastructure for JUnit 5

## Run It Easily

From the repository root:

```bash
mise install
mise run test
```

If you prefer `mise` task shortcuts:

```bash
mise run setup
mise run test
```

## Near-Term Scope

The first learning milestone is to prepare the repository for:

- complexity
- arrays
- strings
- hashing
- two pointers
- sliding window

That is the current boundary. The next step is to add the first real Kotlin implementations and pattern notes in those areas.

## Repository Layout

- `docs/` learning roadmap and reference notes
- `notes/` pattern and mistake journals
- `problems/` organized problem library
- `src/main/kotlin/algorithms/problems/` executable Kotlin implementations
- `src/test/kotlin/algorithms/problems/` executable Kotlin tests
- `src/main/kotlin/` reusable Kotlin implementation space
- `src/test/kotlin/` tests and invariants
- `benchmarks/` future JMH work
- `templates/` reusable problem and note templates
- `legacy/` archived solutions from the previous multi-language structure

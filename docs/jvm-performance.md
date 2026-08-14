# JVM Performance

This repository treats the JVM as part of the algorithmic environment, not a black box.

## Concepts That Matter

- primitives
- boxing and unboxing
- heap allocation
- object allocation
- garbage collection
- cache locality
- recursion and stack usage
- string allocation
- mutable versus immutable structures

## Why It Matters

An algorithm can be asymptotically correct and still perform poorly if it allocates too much or forces unnecessary boxing.

## Practical Questions

- Does this data structure allocate per element?
- Is this loop creating temporary objects?
- Is the recursion depth safe?
- Would an array-backed structure be better than a linked one?


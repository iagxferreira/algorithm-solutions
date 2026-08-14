# Kotlin for DSA

Kotlin is expressive, but DSA work needs attention to allocation and boxing.

## Core Types To Learn Well

- `IntArray`
- `LongArray`
- `Array<T>`
- `MutableList<T>`
- `ArrayDeque<T>`
- `HashMap<K, V>`
- `HashSet<T>`
- `PriorityQueue<T>`
- `StringBuilder`

## Important Trade-Offs

- primitive arrays avoid boxing
- generic collections often allocate wrapper objects
- mutable structures are often better for performance-critical code
- idiomatic Kotlin should still be readable under contest pressure

## Interop Notes

- Java collections are available directly
- primitive arrays map well to JVM data structures
- pay attention to nullability and type inference

## DSA Habit

Prefer the simplest Kotlin expression that still makes performance behavior obvious.


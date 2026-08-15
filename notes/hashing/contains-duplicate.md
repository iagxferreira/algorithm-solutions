# Contains Duplicate

This problem is the cleanest example of the difference between:

- remembering values you have already seen
- rearranging the input so duplicates become adjacent

## Main Idea

For this problem, the useful question is:

- Have I seen this value before?

That leads naturally to a `HashSet<Int>`.

## Why the Set Approach Fits

- scan once
- check membership in constant time on average
- stop immediately when a repeated value is found

This gives an average `O(n)` solution.

## Why Sorting Is Different

Sorting can also help because duplicates become neighbors.

Then the reasoning becomes:

- sort the array
- compare adjacent values
- if two neighbors are equal, return `true`

This is valid, but it is a different tradeoff:

- time: `O(n log n)`
- extra memory: depends on the sort implementation
- the original ordering is no longer preserved

## Trade-Off Summary

- `HashSet`: faster average runtime, more memory
- sorting: simpler adjacency check after rearrangement, slower than hashing

## Reusable Lesson

When the answer is a boolean like “does anything repeat?”, ask whether you need:

- the full history
- only membership
- or a sorted view

That choice usually determines the data structure.

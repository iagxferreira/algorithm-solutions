# Complexity

Complexity is a way of describing how an algorithm scales as input grows.

## The Core Questions

- How does runtime grow?
- How much extra memory is needed?
- What happens in the worst case?
- What happens on average?
- Can the cost be spread out across many operations?

## The Main Notations

- Big O describes an upper bound.
- Big Omega describes a lower bound.
- Big Theta describes a tight bound.

## Why Complexity Matters

Two solutions with the same asymptotic class can still behave very differently on the JVM because of:

- boxing
- allocation
- cache locality
- recursion depth
- collection choice

## Amortized Analysis

Some operations are occasionally expensive but cheap on average across a sequence of operations.

The classic example is a dynamic array resize:

- most appends are constant time
- occasional resizes copy the full buffer
- the average append cost is still amortized constant time

## Recurrence Thinking

Recursive algorithms often lead to recurrences.

Useful examples:

- merge sort
- quick sort
- divide and conquer search
- tree traversal

## Practical Rule

Always ask:

1. What is the dominant operation?
2. How many times does it happen?
3. What is the hidden constant on the JVM?


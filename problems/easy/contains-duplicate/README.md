# 217. Contains Duplicate

Platform: LeetCode
Problem URL: [https://leetcode.com/problems/contains-duplicate/](https://leetcode.com/problems/contains-duplicate/)
Difficulty: Easy
Pattern:
Status: Solved

Started: August 15, 2026
Solved: August 15, 2026

Worksheet: [`../problem.md`](/home/iago/workspace/algorithm-solutions/problems/problem.md)

## Session Notes

This problem is about detection, not pairing.

The key question is:

- Have I seen this value before?

That is why a set is a natural fit.

## 1. Understanding

The problem asks whether the array contains any repeated value.

My own wording:

- Scan the list.
- If any number appears more than once, return `true`.
- If every number is unique, return `false`.

What the output represents:

- `true` means at least one duplicate exists.
- `false` means all values are distinct.

What I need to keep in mind:

- I am not looking for the duplicate itself.
- I am only answering whether one exists.

## 2. Constraints

From the prompt:

- Input size: `1 <= nums.length <= 10^5`
- Value range: large integer range
- Duplicates allowed: yes
- Sorted: no guarantee
- Negative values: yes
- Empty input: not required by the problem, but useful to think about defensively
- Time constraints: strongly suggest better than pairwise comparison
- Memory constraints: extra memory is acceptable if it avoids `O(n^2)`

What do these constraints tell me about the possible complexity?

- A nested scan would be too slow on large arrays.
- I should try to remember what I have already seen.
- A set can give fast membership checks.

## 3. Examples

Example 1:

- Input: `nums = [1, 2, 3, 1]`
- Output: `true`
- Why? `1` appears twice.

Example 2:

- Input: `nums = [1, 2, 3, 4]`
- Output: `false`
- Why? Every value is distinct.

Example 3:

- Input: `nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]`
- Output: `true`
- Why? Several values repeat.

My own example:

- Input: `nums = [-5, 10, 7, -5]`
- Output: `true`
- Why? `-5` appears twice.

Relevant edge cases:

- Single element
- All equal
- No duplicates
- Duplicate at the start
- Duplicate at the end
- Negative values

## 4. Brute Force

Approach:

- Compare every element with every later element.
- If any pair matches, return `true`.
- If no pair matches, return `false`.

Why does it work?

- Every possible pair is checked exactly once.

Time Complexity:

- `O(n^2)`

Space Complexity:

- `O(1)`

## 5. Bottleneck

What makes the brute-force solution slow?

- It repeats comparisons many times.

What operation happens repeatedly?

- Checking whether a value has appeared before.

What information am I recomputing?

- Whether a value is already in the previous part of the array.

Can I store something?

- Yes, the values I have already seen.

Can I eliminate unnecessary work?

- Yes, by checking membership in a set instead of rescanning earlier items.

Can ordering help?

- The array is unsorted, so ordering does not solve the problem directly.

Can I process the input only once?

- Yes, and that is the best direction here.

Can I maintain some state?

- Yes, a set of seen values.

## 6. Pattern Recognition

Pattern:

- Hashing
- Set membership

Why:

- I only need to know whether a value has been seen before.

Recognition signals:

- “Contains duplicate”
- “Have I seen this before?”
- “Return a boolean”
- “One pass is enough”

## 7. Data Structure

What data structure am I using?

- `HashSet<Int>`

Why?

- It stores the unique values seen so far.

What operation does it optimize?

- Membership check.

What is its complexity?

- Average `O(1)` insert and lookup.

What would happen if I used another data structure?

- A list would make membership checks slow.
- A map would work too, but it would store more information than needed.

Alternative approach:

- Sort the array and compare adjacent values.
- This is valid because duplicates become neighbors after sorting.
- The tradeoff is `O(n log n)` time instead of the average `O(n)` set-based scan.
- It also changes the input order, so it is a different style of solution.

## 8. Invariant

What must always remain true while my algorithm runs?

- The set contains exactly the values seen so far.
- If the current value is already in the set, then a duplicate exists.

## 9. Algorithm

1. Create an empty set.
2. Scan the array from left to right.
3. For each number, check whether it already exists in the set.
4. If it does, return `true`.
5. Otherwise, add it to the set and continue.
6. If the scan ends, return `false`.

## 10. Correctness

Why does this algorithm always produce the correct result?

- If a duplicate exists, then when the second copy appears, the first copy is already in the set.
- At that moment, the algorithm returns `true`.
- If no duplicate exists, every number is added once and never found twice, so the algorithm returns `false`.

What cases does it handle?

- Duplicate values
- Negative numbers
- Values appearing more than twice
- A duplicate at any position in the array

Why can I safely discard certain information?

- I do not need full history or counts.
- I only need to know whether a value has appeared before.

Why can a pointer move without going backwards?

- This solution is a forward scan, so each value is processed once.

Why can I assume this greedy choice is safe?

- As soon as a repeated value is found, the boolean answer is decided.

Why does the state represent the required subproblem?

- The set captures the exact summary of the prefix needed to decide whether the current value is a duplicate.

## 11. Kotlin Implementation

Implementation notes:

- `HashSet<Int>` is enough because the result is boolean.
- The check must happen before adding the current value.

Final Kotlin solution:

- Implemented in [`ContainsDuplicate.kt`](/home/iago/workspace/algorithm-solutions/src/main/kotlin/algorithms/problems/easy/contains_duplicate/ContainsDuplicate.kt)

## 12. Tests

- Empty cases: defensive check
- One element: should return `false`
- Duplicate values: should return `true`
- Large inputs: confirms the one-pass approach
- Boundaries: duplicate at the first or last position
- Invalid operations: not part of the contract
- Invariants: values seen so far are tracked in the set

Test scaffold:

- Kotlin test file: [`ContainsDuplicateTest.kt`](/home/iago/workspace/algorithm-solutions/src/test/kotlin/algorithms/problems/easy/contains_duplicate/ContainsDuplicateTest.kt)
- What should this test prove? That the function returns `true` exactly when a repeated value appears.

## 13. Complexity

- Time: `O(n)` average
- Space: `O(n)`
- Amortized behavior: set insert and lookup are amortized constant time
- Notes on JVM behavior: `HashSet<Int>` uses boxed `Int` values, which is fine for this learning-scale problem
 
Sorting alternative:

- Time: `O(n log n)`
- Space: depends on the sorting implementation
- Tradeoff: usually less direct than the set approach for this problem

## 14. Reflection

- What did I initially misunderstand? I first thought in terms of “checked positions,” but the useful idea is “seen values.”
- What is the reusable lesson? Many detection problems become simple once you store the prefix summary.
- What should I remember next time? Ask whether the answer is boolean, membership, or pairing before choosing the data structure.

## 15. Revisit

- Next review date: after the next hashing or set-based problem
- Did I solve it independently? yes
- Did I need a hint? only to lock onto the right data structure choice
- Did I need the editorial? no
- Did I forget the pattern? no, after the first pass
- Did I make an implementation mistake? not in the final version
- Did I make a complexity mistake? no
- Did I misunderstand the problem? briefly, before separating “seen values” from “checked indices”

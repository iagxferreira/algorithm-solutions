# Two Sum

Platform: LeetCode
Problem URL: [https://leetcode.com/problems/two-sum](https://leetcode.com/problems/two-sum)
Difficulty: Easy
Pattern:
Status: Solved

Started: August 14, 2026
Solved: August 14, 2026

Worksheet: [`../problem.md`](/home/iago/workspace/algorithm-solutions/problems/problem.md)

## Session Notes

This is the first problem in the study cycle, so the goal here is not to jump to code. The goal is to train the workflow:

1. Restate the problem in my own words.
2. Extract the constraints that matter.
3. Work through examples and edge cases.
4. Write the brute-force baseline.
5. Identify the bottleneck.
6. Recognize the pattern.
7. Choose the data structure and invariant.
8. Prove correctness.
9. Implement in Kotlin.
10. Test, reflect, and revisit.

## 1. Understanding

The problem asks for the indices of two distinct elements in an integer array whose values add up to a target.

My own wording:

- Given a list of integers and a target value, find two positions whose values sum to that target.
- Each input is guaranteed to have exactly one valid answer.
- I may return the indices in any order.

What the output represents:

- The result is not the numbers themselves.
- The result is the pair of array indices that identify the matching numbers.

What I need to keep in mind:

- I cannot use the same element twice.
- I do not need to return a sorted pair.
- I should focus on correctness first, then efficiency.

## 2. Constraints

From the prompt:

- Input size: `2 <= nums.length <= 10^4`
- Value range: `-10^9 <= nums[i] <= 10^9`
- Target range: `-10^9 <= target <= 10^9`
- Duplicates allowed: yes
- Sorted: no guarantee
- Negative values: yes
- Empty input: not allowed by constraints
- Time constraints: imply I should think about something better than brute force if possible
- Memory constraints: not explicitly tight, so a small amount of extra memory may be acceptable if it buys a better runtime
- Valid answer: exactly one solution exists

What do these constraints tell me about the possible complexity?

- `n` can be large enough that `O(n^2)` may be too slow in practice.
- Values can be negative, so I cannot rely on positive-only reasoning.
- The array is unsorted, so ordering-based shortcuts are not given for free.
- Since there is exactly one solution, I can stop as soon as I find it.

## 3. Examples

Example 1:

- Input: `nums = [2, 7, 11, 15]`, `target = 9`
- Output: `[0, 1]`
- Why? `2 + 7 = 9`

Example 2:

- Input: `nums = [3, 2, 4]`, `target = 6`
- Output: `[1, 2]`
- Why? `2 + 4 = 6`

Example 3:

- Input: `nums = [3, 3]`, `target = 6`
- Output: `[0, 1]`
- Why? The same value appears twice, and both positions are valid.

My own example:

- Input: `nums = [5, -2, 8, 1]`, `target = 6`
- Output: `[1, 2]`
- Why? `-2 + 8 = 6`

Relevant edge cases:

- Empty input
- Single element
- Duplicates
- Negative values
- Target negative
- Pair at the beginning
- Pair at the end
- Pair involving duplicate values

## 4. Brute Force

Approach:

- Pick one index.
- Compare it with every later index.
- Check whether the pair sums to the target.
- Stop as soon as the valid pair is found.

Why does it work?

- Every valid pair `(i, j)` with `i < j` is eventually checked exactly once.
- The loops never reuse the same element because the inner loop starts at `i + 1`.

Time Complexity:

- `O(n^2)`

Space Complexity:

- `O(1)`

## 5. Bottleneck

What makes the brute-force solution slow?

- It repeats pair comparisons for many combinations.

What operation happens repeatedly?

- The sum of the current pair.

What information am I recomputing?

- Whether a previous number could have completed the current one.

Can I store something?

- Yes, I can store numbers I have already seen.

Can I eliminate unnecessary work?

- Yes, by checking whether the needed complement already exists instead of rescanning older elements.

Can ordering help?

- Not much here, because the array is unsorted and the answer needs original indices.

Can I process the input only once?

- Yes, that is the optimized direction.

Can I maintain some state?

- Yes, a map of `number -> index`.

## 6. Pattern Recognition

Pattern:

- Hashing
- One-pass lookup

Why:

- For each number, I want to know immediately whether its complement has already appeared.

Recognition signals:

- “Find two numbers that add to a target”
- “Return indices”
- “Exactly one solution”
- “Unsorted input”
- “Can remember earlier values”

## 7. Data Structure

What data structure am I using?

- `HashMap<Int, Int>`

Why?

- It stores each seen number together with its index.

What operation does it optimize?

- Complement lookup.

What is its complexity?

- Average lookup and insert are `O(1)`.

What would happen if I used another data structure?

- A `HashSet` would tell me whether a number exists, but not which index to return.
- A nested array scan would remain `O(n^2)`.

## 8. Invariant

What must always remain true while my algorithm runs?

- Before processing index `i`, the map contains only values from indices `< i`.
- If a complement exists in the map, then I already know an earlier index that can pair with the current one.

## 9. Algorithm

1. Create an empty map from number to index.
2. Scan the array from left to right.
3. For each number, compute the complement `target - current`.
4. If the complement is already in the map, return the stored index and the current index.
5. Otherwise, store the current number and its index in the map.
6. If the loop ends, return an empty array only as a fallback for the LeetCode contract.

## 10. Correctness

Why does this algorithm always produce the correct result?

- When the current number is processed, every earlier number is already recorded.
- If the needed complement was seen earlier, the map returns the exact earlier index that forms the valid pair.
- Because the problem guarantees exactly one solution, the first match found is sufficient.

What cases does it handle?

- Duplicate values
- Negative numbers
- A pair at the beginning, middle, or end of the array
- Cases where the answer uses the same value twice at different indices

Why can I safely discard certain information?

- I do not need every previous pair.
- I only need to know whether a complement has appeared before and where.

Why can a pointer move without going backwards?

- This solution does not use pointers in the two-pointer sense.
- The left-to-right scan is enough because the map preserves all earlier information needed later.

Why can I assume this greedy choice is safe?

- As soon as the complement is found, the answer is complete.
- There is no need to keep searching because the input guarantees one valid pair.

Why does the state represent the required subproblem?

- The map summarizes everything relevant from the prefix of the array.
- That prefix summary is enough to decide whether the current value completes a valid pair.

## 11. Kotlin Implementation

Implementation notes:

- `HashMap<Int, Int>` is the key structure.
- The complement check must happen before storing the current number.
- That order avoids accidentally pairing a number with itself on the same iteration.

Final Kotlin solution:

- Implemented in [`TwoSum.kt`](/home/iago/workspace/algorithm-solutions/src/main/kotlin/algorithms/problems/easy/two_sum/TwoSum.kt)

## 12. Tests

- Empty cases: not needed under the problem constraints, but useful for defensive thinking
- One element: invalid under constraints, but worth noting
- Duplicate values: important, especially `[3, 3]`
- Large inputs: confirms the need to avoid `O(n^2)`
- Boundaries: large positive and negative values
- Invalid operations: not part of the LeetCode contract
- Invariants: map contains only earlier values

Test scaffold:

- Kotlin test file: [`TwoSumTest.kt`](/home/iago/workspace/algorithm-solutions/src/test/kotlin/algorithms/problems/easy/two_sum/TwoSumTest.kt)
- What should this test prove? That the returned pair of indices refers to two distinct positions whose values sum to the target.

## 13. Complexity

- Time: `O(n)` average
- Space: `O(n)`
- Amortized behavior: `HashMap` insert and lookup are amortized constant time
- Notes on JVM behavior: `HashMap<Int, Int>` still involves boxed `Int` keys and values, which is fine here for clarity and performance at this scale

## 14. Reflection

What did I initially misunderstand?

- I started by thinking in terms of “already checked indices” instead of “previously seen values.”

What is the reusable lesson?

- For pair-finding problems, the important question is often: “What earlier value would complete the current one?”

What should I remember next time?

- Brute force is the baseline.
- The real optimization is usually storing useful history so repeated scans disappear.

## 15. Revisit

- Next review date: after the next 3-5 problems in the study cycle
- Did I solve it independently? partially
- Did I need a hint? yes, to shift from indices to seen values
- Did I need the editorial? no
- Did I forget the pattern? initially yes, then no
- Did I make an implementation mistake? yes, while mixing brute force with the map idea
- Did I make a complexity mistake? yes, before the optimized version
- Did I misunderstand the problem? initially yes, then corrected

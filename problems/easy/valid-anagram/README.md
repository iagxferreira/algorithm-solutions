# 242. Valid Anagram

Platform: LeetCode
Problem URL: [https://leetcode.com/problems/valid-anagram/](https://leetcode.com/problems/valid-anagram/)
Difficulty: Easy
Pattern: Hashing / Frequency Counting
Status: Solved

Started: August 15, 2026
Solved: August 15, 2026

Worksheet: [`../problem.md`](/home/iago/workspace/algorithm-solutions/problems/problem.md)

## Session Notes

This problem is about comparing two strings by their character content, not by their order.

The key question is:

- do `s` and `t` contain the same letters with the same frequencies?

That is the idea I want to keep in mind while working through the rest of the sections.

The study path here was:

- start with a map-based counting idea
- notice that the alphabet is limited to lowercase English letters
- replace the map with a fixed `IntArray(26)`
- keep the same frequency logic, but with less overhead

## 1. Understanding

The problem asks whether one string is an anagram of the other.

My own wording:

- Given two strings, determine whether they are made of the same letters in a different order.
- Return `true` if `t` is an anagram of `s`.
- Return `false` otherwise.

What the output represents:

- `true` means the strings contain the same characters with the same counts.
- `false` means at least one character count differs, or the strings do not have the same characters.

What I need to keep in mind:

- order does not matter
- character frequency does matter
- the strings are compared as whole strings, not as substrings

## 2. Constraints

From the prompt:

- Input size: `1 <= s.length, t.length <= 5 * 10^4`
- Character set: lowercase English letters only

What do these constraints tell me about the possible complexity?

- `n` can be large enough that repeated full rescans may be too expensive.
- I should look for a way to summarize the characters instead of comparing them repeatedly.
- Since the character set is limited, a small fixed-frequency summary may be worth thinking about.

## 3. Examples

Example 1:

- Input: `s = "anagram"`, `t = "nagaram"`
- Output: `true`
- Why? Both strings contain the same letters with the same counts.

Example 2:

- Input: `s = "rat"`, `t = "car"`
- Output: `false`
- Why? The letters do not match.

My own example:

- Input: `s = "listen"`, `t = "silent"`
- Output: `true`
- Why? The same letters appear in a different order.

Relevant edge cases:

- Same length but different letters
- Different lengths
- Repeated characters
- All characters equal
- Single-character strings

Follow-up note:

- If the input can include Unicode characters, I would need to think about a broader character mapping than just lowercase English letters.

## 4. Brute Force

Approach:

- Compare the strings by counting the letters in each one.
- A direct brute-force style comparison would keep rechecking letters and counts.

Why does it work?

- If two strings have the same letters with the same frequencies, they are anagrams.
- The brute-force thinking is useful because it exposes the real requirement: frequency, not order.

Time Complexity:

- `O(n^2)` if I keep rescanning to compare characters repeatedly.

Space Complexity:

- `O(1)` for the pure comparison idea, but it is too slow for large inputs.

## 5. Bottleneck

What makes the brute-force solution slow?

- It repeats the same frequency checks over and over.

What operation happens repeatedly?

- Counting how many times each character appears.

What information am I recomputing?

- Whether the two strings contain the same multiset of letters.

Can I store something?

- Yes, I can store counts as I scan each string.

Can I eliminate unnecessary work?

- Yes, by recording frequencies once instead of comparing letter by letter many times.

Can ordering help?

- Not really. The order is irrelevant, so rearranging is not the main idea.

Can I process the input only once?

- Yes, if I keep counts while scanning.

Can I maintain some state?

- Yes, a frequency table.

## 6. Pattern Recognition

Pattern:

- Hashing / frequency counting
- Fixed-frequency array for lowercase letters

Why:

- The input is restricted to lowercase English letters, so the state can be very small and fast.

Recognition signals:

- “Are these two strings anagrams?”
- “Same letters, same counts”
- “Lowercase English letters only”
- “Compare frequencies instead of order”

## 7. Data Structure

What data structure am I using?

- `IntArray(26)`

Why?

- It stores the frequency of each lowercase letter directly.

What operation does it optimize?

- Counting and comparing character frequencies.

What is its complexity?

- Update and lookup are `O(1)`.

What would happen if I used another data structure?

- `HashMap<Char, Int>` would also work and is easier to generalize.
- The fixed array is simpler and faster for this exact problem because the alphabet size is known.

## 8. Invariant

What must always remain true while my algorithm runs?

- After processing the first `i` characters, the count array reflects the difference between the characters seen in `s` and `t` up to that point.
- If all counts end at zero, the strings contain the same letters with the same frequencies.

## 9. Algorithm

1. Check whether the strings have the same length.
2. Create an `IntArray(26)` for character counts.
3. Scan both strings at the same time.
4. Increment the count for the character from `s`.
5. Decrement the count for the character from `t`.
6. Return whether every count is zero.

## 10. Correctness

Why does this algorithm always produce the correct result?

- Each increment records one character from `s`.
- Each decrement cancels one matching character from `t`.
- If the strings contain the same letters with the same counts, every entry in the table ends at zero.
- If any entry is nonzero, one string has more of that letter than the other.

What cases does it handle?

- Equal strings
- Same letters in different order
- Duplicate letters
- Different lengths

Why can I safely discard certain information?

- The positions of the letters do not matter.
- Only the total frequency of each letter matters.

Why does the state represent the required subproblem?

- The frequency table is exactly the summary needed to compare two strings as multisets of letters.

## 11. Kotlin Implementation

Implementation notes:

- The input is lowercase English letters, so a fixed 26-slot array is enough.
- I can update both strings in one pass.
- This avoids the overhead of a map while keeping the code simple.

Final Kotlin solution:

- Use `IntArray(26)` and compare the final counts.

## 12. Tests

Tests:

- Classic anagram case
- Non-anagram with same length
- Repeated letters
- Different lengths

Test scaffold:

- Kotlin test file: [`ValidAnagramTest.kt`](/home/iago/workspace/algorithm-solutions/src/test/kotlin/algorithms/problems/easy/valid_anagram/ValidAnagramTest.kt)
- What should this test prove?

  - The function returns `true` for a real anagram.
  - The function returns `false` when the letter counts do not match.
  - The function handles duplicates correctly.

## 13. Complexity

Time:

- `O(n + k)` where `n` is the string length and `k = 26`
- In practice, this is `O(n)`

Space:

- `O(1)` extra space

Amortized behavior:

- Not relevant here; there is no hash resizing or dynamic lookup structure.

Notes on JVM behavior:

- `IntArray` stores primitive `Int` values without boxing.
- This keeps the solution lightweight compared with a map-based counting structure.

## 14. Reflection

What did I initially misunderstand?

- I first thought about order, but order is not the thing that matters.

What is the reusable lesson?

- When the input alphabet is small and known, a fixed array can replace a hash map.

What should I remember next time?

- Start with the frequency idea, then check whether the constraints let me compress the structure into a fixed-size array.

## 15. Revisit

- Next review date:
- Did I solve it independently?
- Did I need a hint?
- Did I need the editorial?
- Did I forget the pattern?
- Did I make an implementation mistake?
- Did I make a complexity mistake?
- Did I misunderstand the problem?

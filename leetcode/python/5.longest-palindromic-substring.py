class Solution:
    def longestPalindrome(self, s: str) -> str:
        if not s: return ""
        start = 0
        end = 0
        for i in range(len(s)):
            odd = expand(s, i, i)
            even = expand(s, i , i + 1)
            max_len = max(odd, even)

            if max_len >  end - start:
                start = i - (max_len - 1) // 2
                end = i + max_len // 2
        return s[start:end+1]


def expand(s: str, left:int, right:int):
    while left >= 0 and right < len(s) and s[left] == s[right]:
        left -= 1
        right += 1
    return right - left - 1

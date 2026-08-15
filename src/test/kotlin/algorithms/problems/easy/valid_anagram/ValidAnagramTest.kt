package algorithms.problems.easy.valid_anagram

import org.junit.jupiter.api.Assertions.assertFalse
import org.junit.jupiter.api.Assertions.assertTrue
import org.junit.jupiter.api.Test

class ValidAnagramTest {
    @Test
    fun `returns true for the classic example`() {
        assertTrue(isAnagram("anagram", "nagaram"))
    }

    @Test
    fun `returns false when the letters do not match`() {
        assertFalse(isAnagram("rat", "car"))
    }

    @Test
    fun `handles duplicate letters`() {
        assertTrue(isAnagram("aabbcc", "baccab"))
    }

    @Test
    fun `returns false for strings with different lengths`() {
        assertFalse(isAnagram("ab", "a"))
    }
}

package algorithms.problems.easy.contains_duplicate

import org.junit.jupiter.api.Assertions.assertFalse
import org.junit.jupiter.api.Assertions.assertTrue
import org.junit.jupiter.api.Test

class ContainsDuplicateTest {
    @Test
    fun `returns true when a value repeats`() {
        assertTrue(containsDuplicate(intArrayOf(1, 2, 3, 1)))
    }

    @Test
    fun `returns false when all values are unique`() {
        assertFalse(containsDuplicate(intArrayOf(1, 2, 3, 4)))
    }

    @Test
    fun `handles repeated negative values`() {
        assertTrue(containsDuplicate(intArrayOf(-5, 10, 7, -5)))
    }

    @Test
    fun `returns false for a single element`() {
        assertFalse(containsDuplicate(intArrayOf(42)))
    }
}

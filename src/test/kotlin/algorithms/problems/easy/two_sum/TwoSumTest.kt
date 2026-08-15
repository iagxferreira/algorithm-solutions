package algorithms.problems.easy.two_sum

import org.junit.jupiter.api.Assertions.assertArrayEquals
import org.junit.jupiter.api.Test

class TwoSumTest {
    @Test
    fun `returns indices for the classic example`() {
        assertArrayEquals(intArrayOf(0, 1), twoSum(intArrayOf(2, 7, 11, 15), 9))
    }

    @Test
    fun `returns indices for an answer in the middle of the array`() {
        assertArrayEquals(intArrayOf(1, 2), twoSum(intArrayOf(3, 2, 4), 6))
    }

    @Test
    fun `handles duplicate values`() {
        assertArrayEquals(intArrayOf(0, 1), twoSum(intArrayOf(3, 3), 6))
    }

    @Test
    fun `handles negative values`() {
        assertArrayEquals(intArrayOf(1, 2), twoSum(intArrayOf(5, -2, 8, 1), 6))
    }
}

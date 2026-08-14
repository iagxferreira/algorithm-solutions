package algorithms.problems.easy.two_sum

fun twoSum(nums: IntArray, target: Int): IntArray {
    val seen = HashMap<Int, Int>()
    for (i in nums.indices) {
        // Look for the number that would complete the target with the current value.
        val needed = target - nums[i]
        if (seen.containsKey(needed)) {
            return intArrayOf(seen[needed]!!, i)
        }
        // Store the current value for future complements.
        seen[nums[i]] = i
    }
    return intArrayOf()
}

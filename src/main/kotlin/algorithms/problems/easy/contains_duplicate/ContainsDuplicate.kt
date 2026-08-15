package algorithms.problems.easy.contains_duplicate

fun containsDuplicate(nums: IntArray): Boolean {
    val seen = HashSet<Int>()
    for (num in nums) {
        if (!seen.add(num)) {
            return true
        }
    }
    return false
}

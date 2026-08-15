package algorithms.problems.easy.valid_anagram

fun isAnagram(s: String, t: String): Boolean {
    if (s.length != t.length) return false

    val counts = IntArray(26)

    for (i in s.indices) {
        counts[s[i] - 'a']++
        counts[t[i] - 'a']--
    }

    return counts.all { it == 0 }
}

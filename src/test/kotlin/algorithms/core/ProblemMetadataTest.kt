package algorithms.core

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ProblemMetadataTest {
    @Test
    fun `metadata preserves curriculum fields`() {
        val metadata = ProblemMetadata(
            platform = "LeetCode",
            problem = "Two Sum",
            difficulty = Difficulty.EASY,
            pattern = "Hashing",
            dataStructure = "HashMap",
            algorithm = "One-pass lookup",
            status = ProblemStatus.LEARNING,
        )

        assertEquals("LeetCode", metadata.platform)
        assertEquals("Two Sum", metadata.problem)
        assertEquals(Difficulty.EASY, metadata.difficulty)
        assertEquals("Hashing", metadata.pattern)
        assertEquals("HashMap", metadata.dataStructure)
        assertEquals("One-pass lookup", metadata.algorithm)
        assertEquals(ProblemStatus.LEARNING, metadata.status)
    }
}

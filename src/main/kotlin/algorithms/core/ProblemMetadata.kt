package algorithms.core

data class ProblemMetadata(
    val platform: String,
    val problem: String,
    val difficulty: Difficulty,
    val pattern: String,
    val dataStructure: String,
    val algorithm: String,
    val status: ProblemStatus,
)


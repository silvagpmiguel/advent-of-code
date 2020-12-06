package service

// Args data class
data class Args(
        val cookie: String,
        val action: String,
        val day: String,
        val year: String,
        var level: String = "",
        var answer: String = "",
        var filepath: String = "",
)
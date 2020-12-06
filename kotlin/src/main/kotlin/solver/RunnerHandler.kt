package solver

import java.io.File
import kotlin.system.measureNanoTime
import kotlin.time.ExperimentalTime

// Runner is the structure of a puzzle runner
data class Runner(val input: String, val solver: Solver)

class RunnerHandler(val filepath: String, val solvers: List<Solver>, val runners: MutableList<Runner> = mutableListOf()) {

    init {
        val file = File(filepath)

        if (!file.exists()) {
            throw Exception("Error, file '${filepath}' doesn't exist!")
        }

        val input = file.readText()

        solvers.forEach {
            runners.add(Runner(input = input, solver = it))
        }
    }

    // Run a puzzle runner
    fun run() {
        var part1: String?
        var part2: String?

        runners.forEach {
            val elapsedTimeProcess = measureNanoTime {
                it.solver.processInput(it.input)
            }

            val elapsedTimePart1 = measureNanoTime {
                part1 = it.solver.part1()
            }

            val elapsedTimePart2 = measureNanoTime {
                part2 = it.solver.part2()
            }

            Solution(
                processingTime = elapsedTimeProcess,
                part1 = part1,
                part1Time = elapsedTimePart1,
                part2 = part2,
                part2Time = elapsedTimePart2
            ).print()
        }
    }
}
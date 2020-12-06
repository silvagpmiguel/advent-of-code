package solver

interface Solver{

    // processInput of a puzzle.
    fun processInput(content: String)

    // part1 is the solution for the part 1 of the puzzle
    fun part1(): String?

    // part2 is the solution for the part 2 of the puzzle
    fun part2(): String?

}
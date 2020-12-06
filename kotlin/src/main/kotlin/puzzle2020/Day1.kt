package puzzle2020

import solver.Solver

class Day1(val entries: MutableList<Int> = mutableListOf(), val target: Int = 2020) : Solver {
    override fun processInput(content: String) {
        val lines = content.trim().split("\n")
        lines.forEach { entries.add(Integer.parseInt(it)) }
    }

    override fun part1(): String? {
        val len = entries.size
        var i = 0

        while (i < len - 1) {
            var j = i + 1
            while (j < len) {
                val entry1 = entries[i]
                val entry2 = entries[j]
                if (entry1 + entry2 == target) {
                    return (entry1 * entry2).toString()
                }
                j++
            }
            i++
        }

        return "None"
    }

    override fun part2(): String? {
        val len = entries.size
        var i = 0

        while (i < len - 1) {
            var j = i + 1
            while (j < len) {
                var k = j + 1
                while (k < len) {
                    val entry1 = entries[i]
                    val entry2 = entries[j]
                    val entry3 = entries[k]
                    if (entry1 + entry2 + entry3 == target) {
                        return (entry1 * entry2).toString()
                    }
                    k++
                }
                j++
            }
            i++
        }

        return "None"
    }
}

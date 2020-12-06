package solver

import kotlin.math.round
import kotlin.math.roundToInt

// Solution is the structure that stores the details of a puzzle solution
class Solution (
        val processingTime: Long,
        val part1: String?,
        val part1Time: Long,
        val part2Time: Long,
        val part2: String?,
){
    // format time units
    fun formatTime(t: Long): String{
        var time = t.toDouble()
        val units = listOf("ns", "us", "ms", "s")
        var i = 0
        while((time/1000) > 0.1){
            time /= 1000
            i++
        }
        if (i>3){
            throw Exception("Damn, that's a huge time. You need to optimize it :)")
        }
        return String.format("%.2f%s", time, units.get(i))
    }

    // PrintSolution to System.Out
    fun print(){
        println("Input Process: ${formatTime(processingTime)}")
        println("Part 1: ${part1} (in ${formatTime(part1Time)})")
        println("Part 1: ${part2} (in ${formatTime(part2Time)})")
    }
}
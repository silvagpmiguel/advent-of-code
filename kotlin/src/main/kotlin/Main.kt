import service.*
import solver.RunnerHandler
import solver.Solver
import puzzle2020.*
import java.io.File

val solverMap = mapOf(
    "2020" to mapOf(
        "1" to  Day1(),
        "2" to  Day2(),
        "3" to  Day3(),
        "4" to  Day4(),
        "5" to  Day5(),
        "6" to  Day6(),
        "7" to  Day7(),
        "8" to  Day8(),
        "9" to  Day9(),
        "10" to Day10(),
    )
)

fun getSolver(year: String, day: String): Any {
    val solverList = solverMap.getOrElse(year, {
        throw Exception("No solvers available for year ${year}")
    })
    return solverList.getOrElse(day, {
        throw Exception("No solvers available in year ${year} for day ${day}")
    })
}

fun main(args: Array<String>) {
    val arg = ArgsHandler(args.size, args).getArgs()
    val endpoint = "https://adventofcode.com/${arg.year}/day/${arg.day}/${arg.action}"
    val service = Service()

    when(arg.action){
        "solve" -> {
            val solver = getSolver(arg.year, arg.day)
            RunnerHandler(filepath=arg.filepath, solvers=listOf(solver as Solver)).run()
        }
        "answer" -> {
            val answer = service.POST(endpoint, arg.cookie, arg.level, arg.answer)
            println(answer)
        }
        "input" ->{
            val input = service.GET(endpoint, arg.cookie)
            File(arg.filepath).writeText(input)
            println("Input file created: ${arg.filepath}")
        }
        "default" -> println("Error: Wrong action.\nRequired Arguments: \"input|answer|solve\" \"year\" \"day\" [level] [answer]")
    }
}

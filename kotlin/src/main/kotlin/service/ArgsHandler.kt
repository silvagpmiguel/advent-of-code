package service

import java.nio.file.Paths

class ArgsHandler(val argc: Int, val args: Array<String>){

    // Store arguments to be used in main script
    fun getArgs(): Args{
        if (argc < 3 || argc > 5) {
            throw Exception("wrong arguments.\nUsage: ./main \"input|answer|solve\" \"year\" \"day\" [level] [answer]")
        }

        val env = Env("../.env")

        val cookie = env.getEnvVariable("COOKIE")

        var newArgs = Args(cookie = cookie,
            action = args[0],
            year = args[1],
            day = args[2],
            filepath = "src/main/kotlin/year${args[1]}/input/${args[2]}.in",
        )

        if (argc == 4) {
            newArgs.level = args[3]
        } else if (argc == 5) {
            newArgs.level = args[3]
            newArgs.answer = args[4]
        }

        return newArgs
    }
}

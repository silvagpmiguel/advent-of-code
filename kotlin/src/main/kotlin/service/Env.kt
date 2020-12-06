package service

import java.io.File

// Env class used to get environment variables from an .env file
class Env(val path: String, val envMap: HashMap<String,String> = HashMap()){

    init{
        val file = File(path)

        if (!file.exists()){
            throw Exception("Error, environment file '${path}' doesn't exist!")
        }

        file.forEachLine {
            val splitted = it.split("=")

            if (splitted.isEmpty() || splitted.size != 2) {
                throw Exception("Error, file '${path}' is not valid!")
            }

            envMap[splitted[0]] = splitted[1]
        }
    }

    // Get environment variable from .env file
    fun getEnvVariable(variable: String):String {
        return envMap.getOrElse(variable, {
            throw Exception("Error, that environment variable doesn't exist!")
        })
    }
}
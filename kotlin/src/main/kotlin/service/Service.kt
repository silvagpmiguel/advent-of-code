package service

import java.net.URI
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.http.HttpResponse

// Service to get and post to AoC
class Service(val client: HttpClient = HttpClient.newHttpClient()) {

    // Get returns the input of a puzzle
    fun GET(endpoint: String, cookie: String) : String {
        val request: HttpRequest = HttpRequest.newBuilder()
                .uri(URI.create(endpoint))
                .header("cookie", "session=${cookie}")
                .GET()
                .build()
        val response = client.send(request, HttpResponse.BodyHandlers.ofString())
        println(response.body())
        return response.body()
    }

    // Post sends the answer to AoC
    fun POST(endpoint: String, cookie: String, level: String, answer: String) : String? {
        val form = "level=${level}&answer=${answer}"
        val request: HttpRequest = HttpRequest.newBuilder()
        .uri(URI.create(endpoint))
        .header("cookie", "session=${cookie}")
        .header("content-type", "application/x-www-form-urlencoded")
        .POST(HttpRequest.BodyPublishers.ofString(form))
        .build()
        val response = client.send(request, HttpResponse.BodyHandlers.ofString())
        val regex = Regex("<article><p>[A-Za-z0-9'?!.:]*")
        val answer = regex.find(response.body())
        if (answer != null){
            return answer.value.split("<article><p>")[1]
        }
        return null
    }
}
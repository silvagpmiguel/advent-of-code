# Advent of Code

My take for [Advent of Code](https://adventofcode.com/)

## Golang

### Move to Project
```
cd go
``` 

### Build Project
```
go build
```

### Run Program
```
./aoc "input|answer|solve" "year" "day" [level] [answer]
```

## Kotlin

### Move to Project
```
cd kotlin
``` 

### Build Project
```
./mvnw clean package
```

### Run Program
```
java11 -jar target/aoc.jar "input|answer|solve" "year" "day" [level] [answer]
```

## Important
You can get the **input** of a day in a year, or you can **send** your answer directly to Advent of Code, but in order to do that you need to configure the `.env` file located at the **root** of this whole project. 

When inside the `.env` file, replace the value of the key **COOKIE** with your Advent of Code **cookie**.
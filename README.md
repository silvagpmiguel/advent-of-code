# Advent of Code

My take for [Advent of Code](https://adventofcode.com/)

## Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/silvagpmiguel/advent-of-code/go)](https://goreportcard.com/report/github.com/silvagpmiguel/advent-of-code/go)
[![GoDoc](https://godoc.org/github.com/silvagpmiguel/advent-of-code/go?status.svg)](https://godoc.org/github.com/silvagpmiguel/advent-of-code/go)

### Build Project

```
go build
```

### Run Program

```
./aoc "input|answer|solve" "year" "day" [level] [answer]
```

## Kotlin

### Build Project
```
./mvnw clean package
```

### Run Program
```
java11 -jar target/aoc.jar "input|answer|solve" "year" "day" [level] [answer]
```

## Typescript

### Build & Run
```
npm start "year" "day"
```

## Important
You can get the **input** of a day in a year, or you can **send** your answer directly to Advent of Code, but in order to do that you need to configure the `.env` file located at the **root** of this whole project. 

When inside the `.env` file, replace the value of the key **COOKIE** with your Advent of Code **cookie**.

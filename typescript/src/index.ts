import Day1 from './day/Day1'
import { readDayInputAsNumberArray } from './parse/input'
import { exit } from 'process'
import { Day } from './day/Day'

interface Days2021 {
  day1: Day
}

interface YearDays {
  year2021: Days2021
}

const DAYS: Readonly<YearDays | any> = {
  year2021: {
    day1: Day1(readDayInputAsNumberArray(1), 3),
  },
}

if (process.argv.length != 4) {
  console.error('Error, insert advent of code year and day!')
  console.error('npm start <year> <day>')
  exit()
}

const YEAR: Readonly<string> = `year${process.argv[2]}`
const DAY: Readonly<string> = `day${process.argv[3]}`

if (!(YEAR in DAYS && DAY in DAYS[YEAR])) {
  console.error(`Error, unfortunately there is no implementation for ${YEAR} - ${DAY} :(`)
  exit()
}

console.log(`Advent of Code ${YEAR} - ${DAY}`)
console.log(`part1: ${DAYS[YEAR][DAY].part1()}`)
console.log(`part2: ${DAYS[YEAR][DAY].part2()}`)

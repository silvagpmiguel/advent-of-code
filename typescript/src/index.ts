import { AdventRunner } from '@advent/core'

if (process.argv.length == 4) {
  const year = process.argv[2]
  const day = process.argv[3]
  new AdventRunner(year, day).run()
} else {
  console.error('Error, insert advent of code year and day!\nnpm start <year> <day>')
}

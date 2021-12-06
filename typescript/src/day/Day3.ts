import { StringArrayInputParts } from '../parse/input'
import { Day } from './Day'

export default function Day3(inputs: StringArrayInputParts): Day {
  const p1Input: string[] = inputs.part1
  const p2Input: string[] = inputs.part2
  const computeDecimalRating = (gamma: string, epsilonRate: string) => parseInt(gamma, 2) * parseInt(epsilonRate, 2)
  const convergeInputRatings = (input: any, currentIndex: number): void => {
    const oxygenLen = input.oxygen.length
    const co2Len = input.co2.length
    const counter: any = {
      oxygen: { 0: [], 1: [] },
      co2: { 0: [], 1: [] },
    }

    if (oxygenLen != 1) {
      input.oxygen.forEach((val: number, i: number) => counter.oxygen[input.oxygen[i][currentIndex]].push(val))
      input.oxygen = counter.oxygen[0].length > counter.oxygen[1].length ? counter.oxygen[0] : counter.oxygen[1]
    }

    if (co2Len != 1) {
      input.co2.forEach((val: number, i: number) => counter.co2[input.co2[i][currentIndex]].push(val))
      input.co2 = counter.co2[0].length <= counter.co2[1].length ? counter.co2[0] : counter.co2[1]
    }
  }

  return {
    part1: (): number => {
      let gamma: string = ''
      let epsilon: string = ''
      const lineLen = p1Input[0].length
      const counter: any = {}

      for (let i = 0; i < lineLen; i++) {
        counter[i] = { 0: 0, 1: 0 }
      }

      for (let i = 0; i < lineLen; i++) {
        p1Input.forEach((_, j: number) => counter[i][p1Input[j][i]]++)
        if (counter[i][0] > counter[i][1]) {
          gamma += '0'
          epsilon += '1'
        } else {
          gamma += '1'
          epsilon += '0'
        }
      }

      return computeDecimalRating(gamma, epsilon)
    },
    part2: (): number => {
      const input: any = {
        oxygen: p2Input.slice(),
        co2: p2Input.slice(),
      }

      for (let i = 0; i < p2Input[0].length; i++) {
        convergeInputRatings(input, i)
      }

      return computeDecimalRating(input.oxygen[0], input.co2[0])
    },
  }
}
